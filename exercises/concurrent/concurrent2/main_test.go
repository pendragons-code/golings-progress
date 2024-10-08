// concurrent2
// Make the tests pass!

package main_test

import (
	"sync"
	"testing"
)

var (
	mutex sync.Mutex
)

func TestCounter(t *testing.T) {
	counter := updateCounter()
	if counter != 100 {
		t.Errorf("Counter should be 100, but got %d", counter)
	}
}

func updateCounter() int {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++ // Many goroutines trying to update the counter? We need some protection here!
			mutex.Unlock()
			// https://reliasoftware.com/blog/concurrency-in-golang
		}()
	}
	wg.Wait()
	return counter
}
