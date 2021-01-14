package cache

import (
	"container/list"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLru() *LRU {
	return &LRU{
		capacity: 1000,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func NewItem(key string, value interface{}) *Item {
	return &Item{
		Key:   key,
		Value: value,
	}
}

func (l *LRU) purge() {
	if element := l.queue.Back(); element != nil {
		item := l.queue.Remove(element).(*Item)
		delete(l.items, item.Key)
	}
}

