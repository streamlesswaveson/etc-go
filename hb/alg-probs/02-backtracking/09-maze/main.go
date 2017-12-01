package main

import "fmt"

var mazeTable [][]int
var solutionTable [][]int
var mazeSize int

func main() {
	initialize()

	if solveMaze(0, 0) {
		printResults()
	} else {
		fmt.Println("No solution")
	}

}
func initialize() {
	fmt.Scan(&mazeSize)
	mazeTable = make([][]int, mazeSize)
	solutionTable = make([][]int, mazeSize)
	for i := 0; i < mazeSize; i++ {
		mazeTable[i] = make([]int, mazeSize)
		solutionTable[i] = make([]int, mazeSize)
		for j := 0; j < mazeSize; j++ {
			fmt.Scan(&mazeTable[i][j])
		}
	}
}
func solveMaze(x int, y int) bool {

	if isFinished(x, y) {
		return true
	}

	// only searching to the right and down
	if isValid(x, y) {
		solutionTable[x][y] = 1

		// move right
		if solveMaze(x+1, y) {
			return true
		}
		// move down
		if solveMaze(x, y+1) {
			return true
		}
		//otherwise we need to backtrack
		solutionTable[x][y] = 0
	}
	return false

}
func isValid(x int, y int) bool {

	if x < 0 || x >= mazeSize {
		return false
	}
	if y < 0 || y >= mazeSize {
		return false
	}
	if mazeTable[x][y] != 1 {
		return false
	}

	return true

}
// the end is defined as the bottom right corner
func isFinished(x int, y int) bool {
	if x == mazeSize-1 && y == mazeSize-1 {
		solutionTable[x][y] = 1
		return true
	}
	return false

}

func printResults() {
	for i := 0; i < len(solutionTable); i++ {
		for j := 0; j < len(solutionTable); j++ {
			if solutionTable[i][j] == 1 {
				fmt.Print(" S ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}

}
