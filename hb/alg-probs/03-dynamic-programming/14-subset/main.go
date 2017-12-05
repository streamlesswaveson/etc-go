package main

import "fmt"

var (
	dpTable [][]bool
	S       []int
	sum, N  int
)

func main() {
	//S = []int{5, 2, 3, 1}
	//sum = 9

	fmt.Scan(&sum, &N)
	S = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&S[i])
	}

	initTable()
	//fmt.Println(dpTable)
	solve()
	showIntegers()

}

func initTable() {
	dpTable = make([][]bool, len(S)+1)
	for i := 0; i < len(S)+1; i++ {
		dpTable[i] = make([]bool, sum+1)
	}
}

func solve() {

	// first column is true
	for i := 0; i < len(S)+1; i++ {
		dpTable[i][0] = true
	}

	for rowIndex := 1; rowIndex < len(S)+1; rowIndex++ {
		for colIndex := 1; colIndex < sum+1; colIndex++ {

			if colIndex < S[rowIndex-1] {
				dpTable[rowIndex][colIndex] = dpTable[rowIndex-1][colIndex]
			} else {

				// copy above if true
				if dpTable[rowIndex-1][colIndex] == true {
					dpTable[rowIndex][colIndex] = dpTable[rowIndex-1][colIndex]
				} else {
					dpTable[rowIndex][colIndex] = dpTable[rowIndex-1][colIndex-S[rowIndex-1]]
				}
			}
		}
	}
	fmt.Printf("Solution: %v\n", dpTable[len(S)][sum])

}

func showIntegers() {
	colIndex := sum
	rowIndex := len(S)

	for colIndex > 0 || rowIndex > 0 {
		if dpTable[rowIndex][colIndex] == dpTable[rowIndex-1][colIndex] {
			rowIndex = rowIndex - 1
		} else {
			fmt.Printf("We take %v\n", S[rowIndex-1])
			colIndex = colIndex - S[rowIndex-1]
			rowIndex = rowIndex - 1
		}
	}

}
