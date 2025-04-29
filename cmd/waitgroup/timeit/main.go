package main

import (
	timer "concurrency/internal/waitGroup/timeit"
	"fmt"
	"math/rand"
	"time"
)

// начало решения

// timeit выполняет nIter вызовов функции fn
// с помощью nWorkers параллельных обработчиков,
// и возвращает время выполнения в миллисекундах.

func main() {
	rnd := rand.New(rand.NewSource(42))

	fn := func() {
		// "работа" занимает от 10 до 50 мс
		n := 10 + rnd.Intn(40)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	const nIter = 96
	for _, nWorkers := range []int{1, 2, 4, 16} {
		elapsed := timer.Timeit(nIter, nWorkers, fn)
		fmt.Printf("%d iterations, %d workers, took %dms\n", nWorkers*(nIter/nWorkers), nWorkers, elapsed)
	}
	// результаты могут отличаться
	// 96 iterations, 1 workers, took 2998ms
	// 96 iterations, 2 workers, took 1511ms
	// 96 iterations, 4 workers, took 809ms
	// 96 iterations, 16 workers, took 229ms
}
