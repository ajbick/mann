package main

import (
	"fmt"
	"time"
)

// The board dimensions.
const numRows = 6
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
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
	numCalls++

	if numVisited == numRows*numCols {
		if !requireClosedTour {
			return true
		}
		return true //false
	}

	for _, o := range moveOffsets {
		r := curRow + o.dr
		c := curCol + o.dc

		if r < 0 || r >= numRows {
			continue
		}
		if c < 0 || c >= numCols {
			continue
		}
		if board[r][c] != unvisited {
			continue
		}

		board[r][c] = numVisited

		if findTour(board, numRows, numCols, r, c, numVisited+1) {
			return true
		}

		board[r][c] = unvisited
	}

	return false
}

func main() {
	numCalls = 0

	// Initialize the move offsets.
	initializeOffsets()

	// Create the blank board.
	board := makeBoard(numRows, numCols)

	// Try to find a tour.
	start := time.Now()
	board[0][0] = 0
	if findTour(board, numRows, numCols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dumpBoard(board)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", numCalls)
}
