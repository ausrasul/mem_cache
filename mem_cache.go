package main

import "sync"

type Cache struct {
	storage map[string]interface{}
	mtx     *sync.RWMutex
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]interface{}),
		mtx:     &sync.RWMutex{},
	}
}

func (c Cache) Read(key string) interface{} {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.storage[key]
}
func (c Cache) Write(key string, val interface{}) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.storage[key] = val
}
