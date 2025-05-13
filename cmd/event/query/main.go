// Блокирующая очередь.
package main

import (
	query "concurrency/internal/event/query"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	q := query.NewQueue()

	wg.Add(1)
	go func() {
		for i := range 100 {
			q.Put(i)
		}
		wg.Done()
	}()
	wg.Wait()

	total := 0

	wg.Add(1)
	go func() {
		for range 100 {
			total += q.Get()
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Put x100, Get x100, Total:", total)
}
