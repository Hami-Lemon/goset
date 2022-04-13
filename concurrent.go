package goset

import "sync"

type concurrentSet[T comparable] struct {
	lock sync.RWMutex
	data map[T]struct{}
}

func newConcurrentSet[T comparable]() concurrentSet[T] {
	return concurrentSet[T]{
		data: make(map[T]struct{}),
	}
}

func (c *concurrentSet[T]) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.data)
}

func (c *concurrentSet[T]) IsEmpty() bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.data) == 0
}

func (c *concurrentSet[T]) IsNotEmpty() bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.data) != 0
}

func (c *concurrentSet[T]) Contains(value T) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.data[value]
	return ok
}
func (c *concurrentSet[T]) Add(value T) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, ok := c.data[value]
	if ok {
		return false
	}
	c.data[value] = struct{}{}
	return true
}
func (c *concurrentSet[T]) Remove(value T) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ok := c.data[value]
	if !ok {
		return false
	}
	delete(c.data, value)
	return true
}
func (c *concurrentSet[T]) ForEach(consumer func(v T)) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	for value, _ := range c.data {
		consumer(value)
	}
}
