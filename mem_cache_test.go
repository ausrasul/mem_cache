package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

// mem cache is initialized with eviction time
func TestCacheInit(t *testing.T) {
	cache := NewCache()
	if reflect.TypeOf(cache).Name() != "Cache" {
		t.Error("Expected type Cache, got ", reflect.TypeOf(cache))
	}
}

// It reads what it writes, unique keys
func TestReadWrite(t *testing.T) {
	cache := NewCache()
	tests := []struct {
		key      string
		value    interface{}
		write_it bool
		expect   interface{}
		errMsg   string
	}{
		{
			"a",
			1,
			true,
			1,
			"Expect a 1, got ",
		},
		{
			"a",
			2,
			true,
			2,
			"Expect a 2, got ",
		},
		{
			"b",
			3,
			true,
			3,
			"Expect b 3, got ",
		},
		{
			"c",
			3,
			false,
			nil,
			"Expect nil, got ",
		},
	}
	for _, tt := range tests {
		if tt.write_it {
			cache.Write(tt.key, tt.value)
		}
		if cache.Read(tt.key) != tt.expect {
			t.Error(tt.errMsg, cache.Read(tt.key))
		}
	}
}

func TestConcurrentRead(t *testing.T) {
	cache := NewCache()
	cache.Write("a", 1)

	var wg sync.WaitGroup

	a := func() {
		for i := 0; i < 1000; i++ {
			cache.Read("a")
		}
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		go a()
		wg.Add(1)
	}
	wg.Wait()
}

func TestConcurrentWrite(t *testing.T) {
	cache := NewCache()
	var wg sync.WaitGroup

	a := func() {
		for i := 0; i < 1000; i++ {
			cache.Write("a", 1)
		}
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		go a()
		wg.Add(1)
	}
	wg.Wait()
}

func TestConcurrentReadWrite(t *testing.T) {
	cache := NewCache()
	var wg sync.WaitGroup

	a := func() {
		for i := 0; i < 1000; i++ {
			cache.Write("a", 1)
		}
		wg.Done()
	}
	b := func() {
		for i := 0; i < 1000; i++ {
			cache.Read("a")
		}
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		go a()
		go b()
		wg.Add(2)
	}

	wg.Wait()
}

//Write ops benchmark

func BenchmarkWrite(b *testing.B) {
	cache := NewCache()
	for i := 0; i < b.N; i++ {
		cache.Write("c"+fmt.Sprint(i), "a")
	}
}
func BenchmarkOverwrite(b *testing.B) {
	cache := NewCache()
	for i := 0; i < b.N; i++ {
		cache.Write("c", "a")
	}
}

func BenchmarkRead(b *testing.B) {
	cache := NewCache()
	cache.Write("c", "a")
	for i := 0; i < b.N; i++ {
		cache.Read("c")
	}
}

func BenchmarkReadMiss(b *testing.B) {
	cache := NewCache()
	for i := 0; i < b.N; i++ {
		cache.Read("c")
	}
}

func BenchmarkWriteGoMap(b *testing.B) {
	cache := make(map[string]string)
	for i := 0; i < b.N; i++ {
		cache["c"+fmt.Sprint(i)] = "a"
	}
}
func BenchmarkOverwriteGoMap(b *testing.B) {
	cache := make(map[string]string)
	for i := 0; i < b.N; i++ {
		cache["c"] = "a"
	}
}

func BenchmarkReadGoMap(b *testing.B) {
	cache := make(map[string]string)
	cache["c"] = "a"
	for i := 0; i < b.N; i++ {
		_ = cache["c"]
	}
}
func BenchmarkReadMissGoMap(b *testing.B) {
	cache := make(map[string]string)
	for i := 0; i < b.N; i++ {
		_ = cache["c"]
	}
}
