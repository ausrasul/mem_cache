package main

type Cache struct {
	storage map[string]interface{}
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]interface{}),
	}
}

func (c Cache) Read(key string) interface{} {
	return c.storage[key]
}
func (c Cache) Write(key string, val interface{}) {
	c.storage[key] = val
}
