package main

import (
	race "concurrency/internal/conditionrace/limiter"
	"fmt"
)

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := race.Throttle(5, work)
	defer cancel()

	const n = 8
	var nOK, nErr int
	for i := 0; i < n; i++ {
		err := handle()
		if err == nil {
			nOK += 1
		} else {
			nErr += 1
		}
	}
	fmt.Println()
	fmt.Printf("%d calls: %d OK, %d busy\n", n, nOK, nErr)
}
