package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	addNonAtomic()
	addAtomic()
	addNonAtomicMutex()
}

func addNonAtomic() {
	var wg sync.WaitGroup
	var val int
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			val += 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Non-atomic add Total: ", val)
}

func addNonAtomicMutex() {
	var wg sync.WaitGroup
	var val int
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			val += 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Non-atomic with mutex add Total: ", val)
}
func addAtomic() {
	var wg sync.WaitGroup
	var val int32
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&val, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Atomic add Total: ", val)
}
