// TryAcquire
package main

import (
	traffic "concurrency/internal/trafficlight/tryacquire"
	"fmt"
	"sync"
	"time"
)

// конец решения

func main() {
	const maxConc = 4
	sema := traffic.NewSemaphore(maxConc)

	const nCalls = 12
	var wg sync.WaitGroup
	wg.Add(nCalls)

	var nOK, nBusy int
	for i := 0; i < nCalls; i++ {
		if !sema.TryAcquire() {
			nBusy++
			wg.Done()
			continue
		}
		go func() {
			defer wg.Done()
			defer sema.Release()
			time.Sleep(10 * time.Millisecond)
			fmt.Print(".")
			nOK++
		}()
	}

	wg.Wait()

	fmt.Println()
	fmt.Printf("%d calls: %d OK, %d busy\n", nCalls, nOK, nBusy)
	/*
		....
		12 calls: 4 OK, 8 busy
	*/
}
