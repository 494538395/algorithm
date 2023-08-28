package memory

import (
	"container/list"
	"sync"
	"time"
)

type storeItem struct {
	key        uint64      // hash key
	conflict   uint64      // 为了应对可能出现的  hash 冲突，使用双层 hash 策略
	value      interface{} //  value
	expiration time.Time   // ttl
}

// store 借鉴 concurrent hashmap 思想实现的 Map
type store interface {
	// Get 根据指定 key 返回对应的 value，bool 表示是否 Get 到
	Get(uint64, uint64) (interface{}, bool)
	// Expiration 返回指定 key 的过期时间
	Expiration(uint64) time.Time
	// Set 新增 / 更新 Map 中的 KV 键值对
	Set(*Item)
	// Del 删除  Map 中的 KV 键值对
	Del(uint64, uint64) (uint64, interface{})
	// Update 根据 KV，如果成功返回 true
	Update(*Item) (interface{}, bool)
	// Cleanup GC 删除过期 item
	Cleanup(onEvict itemCallback)
}

// newStore 返回默认的 store 实现
func newStore(ruo *list.List, km map[interface{}]*list.Element) store {
	return newShardedMap(ruo, km)
}

// 分片数量定义
const numShards uint64 = 3

type shardedMap struct {
	shards          []*lockedMap
	expiryMap       *expirationMap
	recentUsedOrder *list.List // 双链表用于记录最近使用的顺序
	keysMap         map[interface{}]*list.Element
}

func newShardedMap(ruo *list.List, km map[interface{}]*list.Element) *shardedMap {
	sm := &shardedMap{
		shards:          make([]*lockedMap, int(numShards)),
		expiryMap:       newExpirationMap(),
		recentUsedOrder: ruo,
		keysMap:         km,
	}
	for i := range sm.shards {
		sm.shards[i] = newLockedMap(sm.expiryMap, sm.recentUsedOrder, sm.keysMap)
	}
	return sm
}

func (sm *shardedMap) Get(key, conflict uint64) (interface{}, bool) {
	idx := key % numShards
	return sm.shards[idx].get(key, conflict)
}

func (sm *shardedMap) Expiration(key uint64) time.Time {
	return sm.shards[key%numShards].Expiration(key)
}

func (sm *shardedMap) Set(i *Item) {
	if i == nil {
		// If item is nil make this Set a no-op.
		return
	}

	idx := i.Key % numShards
	sm.shards[idx].Set(i)
}

func (sm *shardedMap) Del(key, conflict uint64) (uint64, interface{}) {
	idx := key % numShards

	return sm.shards[idx].Del(key, conflict)
}

func (sm *shardedMap) Update(newItem *Item) (interface{}, bool) {
	idx := newItem.Key % numShards
	return sm.shards[idx].Update(newItem)
}

func (sm *shardedMap) Cleanup(onEvict itemCallback) {
	sm.expiryMap.cleanup(sm, onEvict)
}

type lockedMap struct {
	sync.RWMutex
	data            map[uint64]storeItem
	em              *expirationMap
	recentUsedOrder *list.List
	keysMap         map[interface{}]*list.Element
	maxKeys         int
}

func newLockedMap(em *expirationMap, ruo *list.List, km map[interface{}]*list.Element) *lockedMap {
	return &lockedMap{
		data:            make(map[uint64]storeItem),
		em:              em,
		recentUsedOrder: ruo,
		keysMap:         km,
	}
}

func (m *lockedMap) get(key, conflict uint64) (interface{}, bool) {
	m.RLock()
	item, ok := m.data[key]
	m.RUnlock()
	if !ok {
		return nil, false
	}
	if conflict != 0 && (conflict != item.conflict) {
		return nil, false
	}

	// Handle expired items.
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return nil, false
	}
	return item.value, true
}

func (m *lockedMap) Expiration(key uint64) time.Time {
	m.RLock()
	defer m.RUnlock()
	return m.data[key].expiration
}

func (m *lockedMap) Set(i *Item) {
	if i == nil {
		// If the item is nil make this Set a no-op.
		return
	}

	m.Lock()
	defer m.Unlock()
	item, ok := m.data[i.Key]

	if ok {
		// 更新过期时间
		m.recentUsedOrder.MoveToFront(m.keysMap[i.Key]) // 刚访问的元素直接移到队头
		m.em.update(i.Key, i.Conflict, item.expiration, i.Expiration)
	} else {
		m.em.add(i.Key, i.Conflict, i.Expiration)
		element := m.recentUsedOrder.PushFront(i) // 刚访问的元素直接移到队头
		m.keysMap[i.Key] = element
	}

	m.data[i.Key] = storeItem{
		key:        i.Key,
		conflict:   i.Conflict,
		value:      i.Value,
		expiration: i.Expiration,
	}
}

func (m *lockedMap) Del(key, conflict uint64) (uint64, interface{}) {
	m.Lock()
	item, ok := m.data[key]
	if !ok {
		m.Unlock()
		return 0, nil
	}
	if conflict != 0 && (conflict != item.conflict) {
		m.Unlock()
		return 0, nil
	}

	if !item.expiration.IsZero() {
		m.em.del(key, item.expiration)
	}

	delete(m.data, key)
	m.Unlock()
	return item.conflict, item.value
}

func (m *lockedMap) Update(newItem *Item) (interface{}, bool) {
	m.Lock()
	item, ok := m.data[newItem.Key]
	if !ok {
		m.Unlock()
		return nil, false
	}
	// 双层 hash 判断，防止 hash 冲突
	if newItem.Conflict != 0 && (newItem.Conflict != item.conflict) {
		m.Unlock()
		return nil, false
	}

	m.em.update(newItem.Key, newItem.Conflict, item.expiration, newItem.Expiration)
	// 更新过期时间
	m.recentUsedOrder.MoveToFront(m.keysMap[newItem.Key]) // 刚访问的元素直接移到队头
	m.data[newItem.Key] = storeItem{
		key:        newItem.Key,
		conflict:   newItem.Conflict,
		value:      newItem.Value,
		expiration: newItem.Expiration,
	}

	m.Unlock()
	return item.value, true
}
