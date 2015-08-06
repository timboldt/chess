package main

import (
	"fmt"
	gmc "github.com/malbrecht/chess"
)

func Evaluate(b *gmc.Board) int {
	return 1
}

func main() {
	board := gmc.MustParseFen("r4rk1/2pp1ppp/8/8/5P2/8/PPPPP1PP/RNBQKBNR b KQ c3 0 12")
	s := Evaluate(board)
	fmt.Printf("%d\n", s)
	moves := board.LegalMoves()
	fmt.Printf(moves[0].San(board))
	fmt.Printf("\n")
}
