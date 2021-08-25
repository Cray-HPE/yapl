package util

import "sync"

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.value++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.value
}

func (c *SafeCounter) Lock() {
	c.mu.Lock()
}

func (c *SafeCounter) Unlock() {
	c.mu.Unlock()
}
