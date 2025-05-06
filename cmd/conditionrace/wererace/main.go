package main

import (
	wererace "concurrency/internal/conditionrace/wererace"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// запускаем работу через 50 мс
	work := func() {
		fmt.Println("work done")
		wg.Done()
	}
	cancel := wererace.Delay(1*time.Millisecond, work)
	defer cancel()

	// отменяем работу через 20 мс c вероятностью 50%
	time.Sleep(20 * time.Millisecond)
	if rand.Intn(2) == 0 {
		cancel()
		fmt.Println("canceled")
		wg.Done()
	}
	wg.Wait()
}
