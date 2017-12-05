package main

import (
	"fmt"
	"math"
)

var (
	dpTable [][]int
	lengthOfRod int
	N int
	prices []int
)
func main()  {
	fmt.Scan(&lengthOfRod, &N)
	prices = make([]int, N)
	for i:=0;i<N; i++ {
		fmt.Scan(&prices[i])
	}


	solve()
	//showResult()
}

func solve()  {
	dpTable := make([][]int, len(prices)+1)
	for i := 0; i < len(prices)+1; i++ {
		dpTable[i] = make([]int, lengthOfRod+1)
	}

	for i:=1; i<=len(prices); i++ {
		for j:=1; j<=lengthOfRod; j++ {
			if i <= j {
				dpTable[i][j] = int(math.Max(float64(dpTable[i-1][j]), float64(prices[i] + dpTable[i-1][j-1])))
			} else {
				dpTable[i][j] = dpTable[i-1][j]
			}
		}
	}

	fmt.Println(dpTable)
}

func showResult(){
	fmt.Printf("Optimal profit: %v", dpTable[len(prices)-1][lengthOfRod])
}