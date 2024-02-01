package genericCache

import (
	"sync"
)

type Cache[T any] struct {
	m sync.Map
}

type item[T any] struct {
	value T
}

func NewCache[T any]() *Cache[T] {
	return &Cache[T]{}
}

func (c *Cache[T]) Set(key string, value T) {
	c.m.Store(key, item[T]{
		value: value,
	})
}

func (c *Cache[T]) Get(key string) T {
	var zero T

	if i, ok := c.m.Load(key); ok {
		if it, exist := i.(item[T]); exist {
			return it.value
		}
	}

	return zero
}

func (c *Cache[T]) Delete(key string) {
	c.m.Delete(key)
}
