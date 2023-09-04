package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.v[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func syncMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1_000; i++ {
		c.Inc("eternal sunshine of the spotless mind")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("eternal sunshine of the spotless mind"))
}
