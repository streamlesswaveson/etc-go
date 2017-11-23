package main

import "fmt"

var (
	numVertices int
	numColors int
	colors []int
	adjacencyMatrix [][]int
)
func main()  {
	var N, tmp, myNumColors int
	fmt.Scan(&myNumColors)
	fmt.Scan(&N)

	matrix := make([][]int, N)

	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&tmp)
			matrix[i][j] = tmp
		}
	}

	solve(matrix, myNumColors)
}

func solve(matrix [][]int, myNumColors int) {
	numColors = myNumColors
	adjacencyMatrix = matrix
	numVertices = len(adjacencyMatrix)
	colors = make([]int, numVertices)

	if solveProblem(0) {
		showResults()
	} else {
		fmt.Println("No solution")
	}
}
func solveProblem(nodeIndex int) bool {
	if nodeIndex == numVertices {
		return true
	}

	for colorIndex :=1; colorIndex <=numColors; colorIndex++ {
		if isColorValid(nodeIndex, colorIndex) {
			colors[nodeIndex] = colorIndex

			if solveProblem(nodeIndex +1) {
				return true
			}

		}
	}
	return false
}

func isColorValid(nodeIndex int, colorIndex int) bool {

	for i:=0;i<numVertices; i++ {
		// adjacent nodes with the same color
		if adjacencyMatrix[nodeIndex][i] == 1 && colorIndex == colors[i] {
			return false
		}
	}
	return true
}

func showResults() {
	for i:=0; i<numVertices; i++ {
		fmt.Printf("Node %v has color index %v\n", i+1, colors[i])
	}
}