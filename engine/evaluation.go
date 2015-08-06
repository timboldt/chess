package engine

import (
	gmc "github.com/malbrecht/chess"
)

const materialMultiplier = 1.0
const squareControlMultiplier = 0.1

func pieceValue(p gmc.Piece) int {
	switch p.Type() {
		case gmc.Pawn: return 1
		case gmc.Knight, gmc.Bishop: return 3
		case gmc.Rook: return 5
		case gmc.Queen: return 9
	}
	return 0
}

func evaluateMaterial(b *gmc.Board) float64 {
	total := 0.0
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			piece := b.Piece[gmc.Square(file, rank)]
			if piece.Color() == gmc.White {
				total += float64(pieceValue(piece))
			} else {
				total -= float64(pieceValue(piece))
			}
		}
	}
	return total
}

func evaluateSquareControl(b *gmc.Board) float64 {
	return 0
}

func evaluate(b *gmc.Board) float64 {
	return evaluateMaterial(b) * materialMultiplier + evaluateSquareControl(b) * squareControlMultiplier
}
