package main

import (
	can "concurrency/internal/racedata/wererace2"
	"fmt"
	"time"
)

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := can.Delay(50*time.Millisecond, work)
	defer cancel()
	time.Sleep(100 * time.Millisecond)
}
