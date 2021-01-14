package cache

import (
	"errors"
)

var (
	errNotFound = errors.New("row with this key wasn't found")
)

func (l *LRU) Set(key string, value interface{}) {
	if element, exists := l.items[key]; exists == true {
		l.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
	}
	if l.queue.Len() == l.capacity {
		l.purge()
	}
	item := NewItem(key, value)
	element := l.queue.PushFront(item)
	l.items[item.Key] = element
}

func (l *LRU) Get(key string) (interface{}, error) {
	element, exists := l.items[key]
	if !exists {
		return nil, errNotFound
	}
	l.queue.MoveToFront(element)
	return element.Value.(*Item).Value, nil
}