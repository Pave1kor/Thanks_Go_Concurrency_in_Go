// Concurrent-группа
package main

import (
	concGroup "concurrency/internal/waitgroup/concurrentgroup"
	"fmt"
	"time"
)

func main() {
	work := func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Print(".")
	}

	cg := concGroup.NewConcGroup()
	for i := 0; i < 4; i++ {
		cg.Run(work)
	}
	cg.Wait()
}
