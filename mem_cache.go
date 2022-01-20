package main

import "sync"

type Cache struct {
	storage sync.Map
}

func NewCache() Cache {
	return Cache{
		storage: sync.Map{},
	}
}

func (c *Cache) Read(key string) interface{} {
	val, _ := c.storage.Load(key)
	return val
}
func (c *Cache) Write(key string, val interface{}) {
	c.storage.Store(key, val)
}
