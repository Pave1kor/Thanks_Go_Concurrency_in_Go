package main

import (
	gm "concurrency/internal/event/game"
	"fmt"
	"time"
)

func main() {
	// создаем новую игру
	game := gm.NewGame(3)

	// игроки делают ставки
	go game.Play("Alice", 10)
	go game.Play("Bob", 21)
	go game.Play("Cindy", 30)
	time.Sleep(10 * time.Millisecond)

	// завершаем игру
	go game.Finish()
	go game.Finish()
	time.Sleep(10 * time.Millisecond)
	winner := game.Finish()

	// оглашаем победителя
	time.Sleep(10 * time.Millisecond)
	fmt.Println("winner:", winner)
	// winner: {Bob 21}
}
