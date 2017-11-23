package main

import "fmt"

var numOfVertices int
var hamiltonianPath []int
var adjacencyMatrix [][]int

func main() {
	var N, tmp int
	fmt.Scan(&N)

	matrix := make([][]int, N)

	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&tmp)
			matrix[i][j] = tmp
		}
	}

	solveHamiltonian(matrix)

}

func solveHamiltonian(matrix [][]int) {

	adjacencyMatrix = matrix
	hamiltonianPath = make([]int, len(matrix))
	numOfVertices = len(matrix)

	hamiltonianPath[0] = 0

	if !findFeasibleSolution(1) {
		fmt.Println("No solution")
	} else {
		showHamiltonianPath()
	}
}
func findFeasibleSolution(position int) bool {

	if position == numOfVertices {
		if adjacencyMatrix[hamiltonianPath[position-1]][hamiltonianPath[0]] == 1 {
			return true
		} else {
			return false
		}
	}
	for vertexIndex := 1; vertexIndex < numOfVertices; vertexIndex++ {
		if isFeasible(vertexIndex, position) {
			hamiltonianPath[position] = vertexIndex

			if findFeasibleSolution(position + 1) {
				return true
			}
		}
	}
	return false

}
func isFeasible(vertexIndex int, actualPostion int) bool {
	// first criterion - are the two nodes connected?
	if adjacencyMatrix[actualPostion-1][vertexIndex] == 0 {
		return false
	}

	// second criterion - have we visited it before?
	for i := 0; i < actualPostion; i++ {
		if hamiltonianPath[i] == vertexIndex {
			return false
		}
	}

	return true
}

func showHamiltonianPath() {
	for i := 0; i < len(hamiltonianPath); i++ {
		fmt.Printf("%v ", hamiltonianPath[i])
	}
	fmt.Println(hamiltonianPath[0])
}
