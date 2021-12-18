package main

import (
	"sync"
)

// struct is collection of fields whereas interface type is a set of method signature
// Go doesn't have class but you can define methods on types
/**
method vs function: methods are functions with a special receiver argument. For example:
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
ref: https://go.dev/tour
*/

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewURLStore() *URLStore {
	// maps are reference type
	// new just allocates memory, not initializes memory; make allocates and initializes memory
	return &URLStore{urls: make(map[string]string)}
}

/**
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
ref: https://go.dev/tour
*/
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// if _, ok pattern
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Put(url string) string {
	// for {} is infinite loop so there must be something inside the loop to break out of the loop
	// this for loop keeps trying until a unique key is generated by genbKey
	// for existing keys, our Set method will return false, so the loop will continue
	for {
		key := genKey()
		if ok := s.Set(key, url); ok {
			return key
		}

	}
}