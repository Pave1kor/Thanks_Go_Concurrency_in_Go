// Concurrent-группа с паникой
package main

import (
	concgroup "concurrency/internal/waitgroup/concurrentgroup"
	"fmt"
	"math/rand"
)

func main() {
	work := func() {
		if rand.Intn(4) == 1 {
			panic("oopsie")
		}
		// do stuff
	}

	defer func() {
		val := recover()
		if val == nil {
			fmt.Println("work done")
		} else {
			fmt.Printf("panicked: %v!\n", val)
		}
	}()

	p := concgroup.NewConcGroup()

	for i := 0; i < 4; i++ {
		p.Run(work)
	}

	p.Wait()
}
