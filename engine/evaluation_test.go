package engine

import (
	"testing"
	gmc "github.com/malbrecht/chess"
)

func TestEvaluateMaterial(t *testing.T) {
	board := gmc.MustParseFen("r4rk1/2pp1ppp/8/8/5P2/8/PPPPP1PP/RNBQKBNR b KQ c3 0 12")
	v := evaluateMaterial(board)
	if v <23.5 || v > 24.5 {
		t.Errorf("Expected material advantage for white to be around 24. Actual value: %v", v)
	}
}

// TODO(tboldt): add more sophisticated tests for material

// TODO(tboldt): add test for square control

// TODO(tboldt): add overall test