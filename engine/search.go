package engine

import (
	"container/heap"
	"fmt"
	gmc "github.com/malbrecht/chess"
)

type CandidateMoveSequence struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type SearchQueue []*CandidateMoveSequence

func (pq SearchQueue) Len() int { return len(pq) }

func (pq SearchQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater-than here.
	return pq[i].priority > pq[j].priority
}

func (pq *SearchQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*CandidateMoveSequence)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *SearchQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	board := gmc.MustParseFen("r4rk1/2pp1ppp/8/8/5P2/8/PPPPP1PP/RNBQKBNR b KQ c3 0 12")
	moves := board.LegalMoves()
	fmt.Printf(moves[0].San(board))
	fmt.Printf("\n")
}