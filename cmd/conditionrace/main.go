// Конкурентно-безопасная карта.
package main

import (
	conc "concurrency/internal/conditionrace/safemap"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getSet() {
	var wg sync.WaitGroup
	wg.Add(2)

	m := conc.NewConcMap[string, int]()

	go func() {
		defer wg.Done()
		m.Set("hello", rand.Intn(100))
	}()

	go func() {
		defer wg.Done()
		m.Set("hello", rand.Intn(100))
	}()

	wg.Wait()
	fmt.Println("hello =", m.Get("hello"))
	// hello = 71 (случайное)
}

func setIfAbsent() {
	var wg sync.WaitGroup
	wg.Add(2)

	m := conc.NewConcMap[string, int]()

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond)
		m.SetIfAbsent("hello", 42)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		m.SetIfAbsent("hello", 84)
	}()

	wg.Wait()
	fmt.Println("hello =", m.Get("hello"))
	// hello = 42 (от первой горутины)
}

func compute() {
	var wg sync.WaitGroup
	wg.Add(2)

	m := conc.NewConcMap[string, int]()

	go func() {
		defer wg.Done()
		for range 100 {
			m.Compute("hello", func(v int) int {
				return v + 1
			})
		}
	}()

	go func() {
		defer wg.Done()
		for range 100 {
			m.Compute("hello", func(v int) int {
				return v + 1
			})
		}
	}()

	wg.Wait()
	fmt.Println("hello =", m.Get("hello"))
	// hello = 200 (каждая горутина увеличила hello на 100)
}

func main() {
	getSet()
	fmt.Println("---")
	setIfAbsent()
	fmt.Println("---")
	compute()
}
