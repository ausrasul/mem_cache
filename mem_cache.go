package main

type Cache struct {
	storage map[string]interface{}
	lock    chan bool
}

func NewCache() Cache {
	return Cache{
		storage: make(map[string]interface{}),
		lock:    make(chan bool, 1),
	}
}

func (c Cache) Read(key string) interface{} {
	c.lock <- true
	defer func() { <-c.lock }()
	return c.storage[key]
}
func (c Cache) Write(key string, val interface{}) {
	c.lock <- true
	defer func() { <-c.lock }()
	c.storage[key] = val
}
