// Пишем семафор
package main

import (
	traffic "concurrency/internal/trafficlight/writetrafficlight"
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConc = 4
	sema := traffic.NewSemaphore(maxConc)
	start := time.Now()

	const nCalls = 12
	var wg sync.WaitGroup
	wg.Add(nCalls)

	for i := 0; i < nCalls; i++ {
		sema.Acquire()
		go func() {
			defer wg.Done()
			defer sema.Release()
			time.Sleep(10 * time.Millisecond)
			fmt.Print(".")
		}()
	}

	wg.Wait()

	fmt.Printf("\n%d calls took %d ms\n", nCalls, time.Since(start).Milliseconds())
}
