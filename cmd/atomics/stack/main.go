// Конкурентно-безопасный стек.
package main

import (
	"fmt"
	"sync"
	"time"
	stack "concurrency/internal/atomics/stack"
)



func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	stack := &stack.Stack{}
	for i := range 1000 {
		go func() {
			time.Sleep(time.Millisecond)
			stack.Push(i)
			wg.Done()
		}()
	}

	wg.Wait()
	count := 0
	for _, ok := stack.Pop(); ok; _, ok = stack.Pop() {
		count++
	}
	fmt.Println(count)
}
