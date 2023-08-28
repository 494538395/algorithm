package memory

import (
	"container/list"
	"fmt"
	"github.com/dgraph-io/ristretto/z"
	//"github.com/OpenIMSDK/Open-IM-Server/pkg/utils"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// 新增内存缓存，底层基于原生 Map ，性能远高于使用缓存服务，如 Redis 等。
// 目前实现的功能有：
// 1.同步缓存
// 2.TTL
// 3.key 最大数量限制

// Q&A
// Q：为什么要引入这个内存缓存？直接使用 Redis 不行吗？
// A: 为性能敏感场景提供多一种思路。
// A: 直接使用 Redis 可以。不过对于极高要求性能的场景，使用内存缓存可以更好的解决

// Q：什么场景适合使用内存缓存？
// A：适用下面几种场景
// 1.读多写少
// 2.计算密集型的场景
// 3.key 总量较少、value 体积较小

// Q：可以即使用内存缓存同时使用 Redis 等缓存服务吗？
// A：可以

// Q：现在是全部切到内存缓存了吗？
// A：不是。现在只把部分逻辑切到内存缓存，后续会做成开关。

// Q：如何实现 TTL 能力？
// A：定时清除位于最小堆顶的元素

const (
	defaultExpireTime = time.Second * 5 // 12H
	maxKeys           = 2
)

var (
	setBufSize = 32 * 1024 // 32 kb
)

var log = logrus.New()

type itemCallback func(*Item)

type itemFlag byte

const (
	itemNew itemFlag = iota
	itemDelete
	itemUpdate
)

// Item 通用的item.
type Item struct {
	flag       itemFlag
	Key        uint64
	Conflict   uint64
	Value      interface{}
	Expiration time.Time
	wg         *sync.WaitGroup
}

// MemoryCache cache 结构.
type MemoryCache struct {

	// store 借鉴 concurrent hashmap 思想实现的 map
	store store

	// setBuf set 缓冲区，应对于高并发进行批量操作的场景
	setBuf chan *Item

	// isClosed 是否关闭的标识位
	isClosed bool

	// stop 用作终止 processItem 的定时任务
	stop chan struct{}

	// 当前 key 数量
	currentKey int

	// 最大 key 数量限制
	maxKeys int

	// KeyToHash
	keyToHash func(interface{}) (uint64, uint64)

	// cleanupTicker 定时 GC 的时间间隔.
	cleanupTicker *time.Ticker

	recentUsedOrder *list.List                    // 双链表用于记录最近使用的顺序
	keysMap         map[interface{}]*list.Element // 哈希表用于快速查找缓存项

	maxExpireTime time.Duration
}

// NewMemoryCache 创建新的memory cache.
func NewMemoryCache() MemoryCache {
	log.SetLevel(logrus.WarnLevel)
	logrus.SetOutput(os.Stdout) // 日志打印在控制台，不写文件
	luc := list.New()
	km := make(map[interface{}]*list.Element)
	c := MemoryCache{
		store:           newStore(luc, km),
		setBuf:          make(chan *Item, setBufSize),
		keyToHash:       z.KeyToHash,
		cleanupTicker:   time.NewTicker(time.Duration(bucketDurationSecs) * time.Second / 2),
		recentUsedOrder: luc,
		keysMap:         km,
		maxKeys:         maxKeys,
		maxExpireTime:   defaultExpireTime,
	}
	go c.processItems()
	return c
}

// processItems 单独的协程，用于处理 item 以及 GC
func (c *MemoryCache) processItems() {

	onEvict := func(i *Item) {
		fmt.Println(fmt.Sprintf("item -->%v has been droped", i))
	}

	for {
		select {
		case i := <-c.setBuf:
			if i.wg != nil {
				i.wg.Done()
				continue
			}
			switch i.flag {
			case itemNew:
				if c.currentKey+1 > c.maxKeys {
					c.deleteOldestKey(onEvict)
				}
				c.store.Set(i)
				c.currentKey++
			case itemUpdate:
				continue
			case itemDelete:
				continue
			}
		case <-c.cleanupTicker.C:
			c.store.Cleanup(onEvict)
		case <-c.stop:
			return
		}
	}
}

func (c *MemoryCache) Set(key, value interface{}) bool {
	return c.SetWithTTL(key, value, c.maxExpireTime)
}

func (c *MemoryCache) SetWithTTL(key, value interface{}, ttl time.Duration) bool {
	if c == nil || c.isClosed || key == nil {
		return false
	}

	var expiration time.Time
	switch {
	case ttl == 0:
		// 没有过期时间
		break
	case ttl < 0:
		// 无效的过期时间
		return false
	default:
		expiration = time.Now().Add(ttl)
	}

	keyHash, conflictHash := c.keyToHash(key)
	i := &Item{
		flag:       itemNew,
		Key:        keyHash,
		Conflict:   conflictHash,
		Value:      value,
		Expiration: expiration,
	}

	if _, ok := c.store.Update(i); ok {
		i.flag = itemUpdate
		return true
	}
	// Attempt to send item to policy.
	select {
	case c.setBuf <- i:
		return true
	default:
		return false
	}
}

func (c *MemoryCache) Wait() {
	if c == nil || c.isClosed {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.setBuf <- &Item{wg: wg}
	wg.Wait()
}

// Get 根据指定 key 查询其 value。bool 表示是否查找到
func (c *MemoryCache) Get(key interface{}) (interface{}, bool) {
	if c == nil || c.isClosed || key == nil {
		return nil, false
	}
	keyHash, conflictHash := c.keyToHash(key)
	value, ok := c.store.Get(keyHash, conflictHash)

	return value, ok
}

// --------------

// GetPre 获取 memoryCache 的值.
//func (c *MemoryCache) GetPre(key interface{}) (value interface{}, found bool) {
//	c.mutex.RLock()
//	start := time.Now()
//	defer func() {
//		c.mutex.RUnlock()
//		log.Debug(fmt.Sprintf("MemoryCache Get 耗时: %v ms", time.Since(start).Milliseconds()))
//	}()
//
//	item, exists := c.cache[key]
//	if exists && item.Expiration.After(time.Now()) {
//		return item.Value, true
//	}
//	// 由于 memoryCache 是嵌在 DB 里面的属性，做加载的动作放在 调用 Get 的地方去做
//	return nil, false
//}

//SetPrev 改变memory cache 的value.
//func (c *MemoryCache) SetPrev(key interface{}, value interface{}) {
//	c.mutex.Lock()
//	start := time.Now()
//	defer func() {
//		c.mutex.Unlock()
//		log.Debug(fmt.Sprintf("MemoryCache Set 耗时: %v ms", time.Since(start).Milliseconds()))
//	}()
//
//	// 检查 key 是否已经在缓存中
//	if existingItem, found := c.cache[key]; found {
//		// 若存在，刷新 value 和 tll
//		existingItem.Value = value
//		existingItem.Expiration = time.Now().Add(c.maxExpireTime)
//
//		c.recentUsedOrder.MoveToFront(c.keysMap[key]) // 刚访问的元素直接移到队头
//	} else {
//		length := len(c.cache)
//		if length > c.maxKeys {
//			c.deleteOldestKey()
//		}
//		// 创建新的项并添加到缓存中
//		expTime := time.Now().Add(c.maxExpireTime)
//		item := &Item{
//			//Key:        key,
//			Value:      value,
//			Expiration: expTime,
//		}
//		c.cache[key] = item
//		element := c.recentUsedOrder.PushFront(key) // 新的元素直接移到队头
//		c.keysMap[key] = element
//		heap.Push(c.expirationQ, item)
//	}
//}

// 清理 oldestKey 的时候不会清理在 堆 中的数据.
// nolint:gocritic,gosimple
func (c *MemoryCache) deleteOldestKey(onEvict itemCallback) {
	// LRU delete
	element := c.recentUsedOrder.Back()
	if element != nil {
		item := element.Value.(*Item)
		c.store.Del(item.Key, item.Conflict)
		c.recentUsedOrder.Remove(element)
		delete(c.keysMap, item.Key)
		onEvict(element.Value.(*Item))
	}
}
