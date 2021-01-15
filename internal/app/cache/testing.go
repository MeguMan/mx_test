package cache

import (
	"container/list"
	"testing"
)

func TestLru(t *testing.T) *LRU {
	t.Helper()
	l := LRU{
		capacity: 3,
		items: make(map[string]*list.Element),
		queue:    list.New(),
	}
	item := &Item{
		Key:   "firstKey",
		Value: "ItemsValue",
	}
	element := l.queue.PushFront(item)
	l.items[item.Key] = element
	return &l
}