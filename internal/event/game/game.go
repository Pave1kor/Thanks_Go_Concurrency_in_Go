package game

import (
	"math"
	"sync"
)

// начало решения

// Game представляет игру.
type Game struct {
	stakes chan stake
	once   sync.Once
	winner stake
}

// NewGame создает новую игру на nPlayers игроков.
func NewGame(nPlayers int) *Game {
	return &Game{
		stakes: make(chan stake, nPlayers),
	}
}

// Play принимает ставку от игрока.
func (g *Game) Play(player string, num float64) {
	select {
	case g.stakes <- stake{player: player, num: num}:
	default:
	}

}

// Finish завершает игру и определяет победителя.
func (g *Game) Finish() stake {
	g.once.Do(func() {
		g.winner = g.decideWinner()
	})
	return g.winner
}

// конец решения
// type Game struct {
//     stakes  chan stake
//     players int
//     done    chan struct{}
//     once    sync.Once
//     winner  stake
// }

// // NewGame создает новую игру на nPlayers игроков.
// func NewGame(nPlayers int) *Game {
//     return &Game{
//         stakes:  make(chan stake, nPlayers),
//         done:    make(chan struct{}),
//         players: nPlayers,
//     }
// }

// // Play принимает ставку от игрока.
// func (g *Game) Play(player string, num float64) {
//     if g.players <= 0 {
//         return
//     }

//     g.stakes <- stake{player: player, num: num}
//     g.players--

//     if g.players == 0 {
//         close(g.done)
//     }
// }

// // Finish завершает игру и определяет победителя.
// func (g *Game) Finish() stake {
//     <-g.done  // Ждем, пока все игроки сделают ставки

//     g.once.Do(func() {
//         close(g.stakes)
//         g.winner = g.decideWinner()
//     })

//	    return g.winner
//	}
//
// stake представляет ставку игрока.
type stake struct {
	player string
	num    float64
}

// decideWinner определяет победителя игры.
// Победитель - игрок, чья ставка ближе всего к средней.
func (g *Game) decideWinner() stake {
	// собираем все ставки
	var s []stake
	for range len(g.stakes) {
		s = append(s, <-g.stakes)
	}

	// находим среднюю ставку
	total := 0.0
	for _, stake := range s {
		total += stake.num
	}
	avg := total / float64(len(s))

	// побеждает тот, чья ставка ближе всего к средней
	var winner stake
	minDist := math.Inf(1)
	for _, stake := range s {
		if dist := math.Abs(stake.num - avg); dist < minDist {
			minDist = dist
			winner = stake
		}
	}

	return winner
}
