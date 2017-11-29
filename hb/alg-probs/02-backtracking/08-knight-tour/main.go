package main

import (
	"fmt"
	"math"
)

const NUM_OF_MOVES = 8

var N int
var board [][]int
var xMoves = []int{2, 1, -1, -2, -2, -1, 1, 2}
var yMoves = []int{1, 2, 2, 1, -1, -2, -2, -1}

func main() {

	fmt.Scan(&N)

	intializeBoard()

	solveTour()

}
func solveTour() {
	board[0][0] = 0

	if solve(1, 0, 0) {
		printSolution()
	} else {
		fmt.Println("No solution")
	}

}
func solve(stepCount int, x int, y int) bool {
	// if we've visited every cell
	if stepCount == N*N {
		return true
	}

	for i := 0; i < NUM_OF_MOVES; i++ {
		nextX := x + xMoves[i]
		nextY := y + yMoves[i]

		if isStepValid(nextX, nextY) {
			board[nextX][nextY] = stepCount

			if solve(stepCount+1, nextX, nextY) {
				return true
			}
			// backtrack!
			board[nextX][nextY] = math.MinInt32
		}
	}
	return false

}
func isStepValid(x int, y int) bool {
	if x < 0 || x >= N {
		return false
	}
	if y < 0 || y >= N {
		return false
	}

	if board[x][y] != math.MinInt32 {
		return false
	}
	return true
}

func printSolution() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%v ", board[i][j])
		}
		fmt.Println()
	}

}

func intializeBoard() {
	board = make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
		for j := 0; j < N; j++ {
			board[i][j] = math.MinInt32
		}
	}

}
