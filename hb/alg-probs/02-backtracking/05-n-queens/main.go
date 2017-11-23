package main

import "fmt"

var numQueens int

var chessBoard [][]int

func main()  {
	fmt.Scan(&numQueens)

	chessBoard = make([][]int,numQueens)
	for i:= range chessBoard {
		chessBoard[i] = make([]int, numQueens)
	}

	solve()
}

func solve() {
	if setQueens(0) {
		printQueens()
	} else {
		fmt.Println("No solution")
	}
}

func setQueens(colIndex int) bool {
	if  colIndex == numQueens {
		return true
	}

	for rowIndex:=0; rowIndex<numQueens; rowIndex++ {

		if isPlaceValid(rowIndex, colIndex) {
			chessBoard[rowIndex][colIndex] = 1

			if setQueens(colIndex+1) {
				return true
			}

			// backtracking!
			chessBoard[rowIndex][colIndex] = 0
		}
	}
	return false
}

func isPlaceValid(rowIndex int, colIndex int) bool  {

	for i:=0; i<colIndex; i++ {
		if chessBoard[rowIndex][i] == 1 {
			return false
		}
	}

	for i, j :=rowIndex, colIndex; i>=0 && j>=0; i,j = i-1, j-1 {
		if chessBoard[i][j] == 1 {
			return false
		}
	}

	for i, j := rowIndex, colIndex; i<len(chessBoard) && j>=0; i,j = i+1, j-1 {
		if chessBoard[i][j] == 1 {
			return false
		}
	}
	return true

}

func printQueens() {
	for i:=0; i<numQueens; i++ {
		for j :=0; j<numQueens; j++ {
			if chessBoard[i][j] == 1 {
				fmt.Print(" * ")
			} else {
				fmt.Print(" - ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}