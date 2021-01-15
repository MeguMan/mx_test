package cache

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestLRU_Set(t *testing.T) {
	l := TestLru(t)
	l.Set("firstKey", "newValue")
	newVal, _ := l.Get("firstKey")
	assert.Equal(t, "newValue", newVal)
	assert.Equal(t, l.queue.Front(), l.items["firstKey"])

	for i:= 0; i < l.capacity; i++ {
		l.Set(strconv.Itoa(i), "value")
	}
	newVal, _ = l.Get("firstKey")
	assert.Equal(t, nil, newVal)
}

func TestLRU_Get(t *testing.T) {
	l := TestLru(t)
	val, err := l.Get("firstKey")
	assert.NoError(t, err)
	assert.NotNil(t, val)
	val, err = l.Get("nonexistentKey")
	assert.Error(t, err)
	assert.Nil(t, val)
}
