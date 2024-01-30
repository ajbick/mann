package main

import (
	"fmt"
	//"time"
)

// The board dimensions.
const numRows = 8
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
    dr, dc  int
}

var moveOffsets []Offset

var numCalls int64



func initializeOffsets() {
	moveOffsets = []Offset{
		Offset{-2, -1},
		Offset{-1, -2},
		Offset{2, 1},
		Offset{1, 2},
		Offset{-2, 1},
		Offset{1, -2},
		Offset{2, -1},
		Offset{-1, 2},
	}
}
		
func makeBoard(r int, c int) [][]int {
	b := make([][]int, r)
	for i := range b {
		b[i] = make([]int, c)
		for s := 0; s < c; s++ {
			b[i][s] = unvisited
		}
	}
	return b
}

func dumpBoard(board [][]int) {
	for x := range board {
		for y := range board[x] {
			fmt.Printf("%02d ", board[x][y])
		}
		fmt.Println()
	}
}

// Try to extend a knight's tour starting at (curRow, curCol).
// Return true or false to indicate whether we have found a solution.
func findTour(board [][]int, numRows, numCols, curRow, curCol, numVisited int) bool {

}



func main() {
	board := makeBoard(8, 8)
	dumpBoard(board)
}
