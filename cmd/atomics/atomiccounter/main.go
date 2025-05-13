package main

import (
	"fmt"
	"sync"
	counter "concurrency/internal/atomics/atomiccounter"
)
func main() {
	var wg sync.WaitGroup

	var total counter.Total
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10000 {
				total.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("total", total.Value())
}