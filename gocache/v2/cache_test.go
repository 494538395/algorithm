package memory

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	cache := NewMemoryCache()

	cache.Set("name", "jerry")
	cache.Wait()
	cache.Set("name", "tom")
	cache.Set("age", 21)

	cache.Wait()
	cache.Set("hobby", "football")

	value, found := cache.Get("name")
	fmt.Println("value-->", value)
	fmt.Println("found-->", found)

	time.Sleep(10 * time.Minute)

}
