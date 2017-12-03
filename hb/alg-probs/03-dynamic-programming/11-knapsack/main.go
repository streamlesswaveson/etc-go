package main

import (
	"fmt"
	"math"
)

var (
	numOfItems       int
	knapsackCapacity int
	knapsackTable    [][]int
	totalBenefit     int
	weights          []int
	values           []int
)

func main() {
	fmt.Scan(&numOfItems)
	fmt.Scan(&knapsackCapacity)

	weights = make([]int, numOfItems+1)
	values = make([]int, numOfItems+1)
	for i := 1; i <= numOfItems; i++ {
		fmt.Scan(&weights[i])
		fmt.Scan(&values[i])
	}
	//fmt.Println(weights)
	//fmt.Println(values)
	knapsackTable = make([][]int, numOfItems+1)
	for i := 0; i < numOfItems+1; i++ {
		knapsackTable[i] = make([]int, knapsackCapacity+1)
	}
	//fmt.Println(knapsackTable)

	solve()
	showResult()

}

func solve() {

	for i := 1; i < numOfItems+1; i++ {
		for w := 1; w < knapsackCapacity+1; w++ {
			notTakingItem := knapsackTable[i-1][w]

			takingItem := 0

			if weights[i] <= w {
				takingItem = values[i] + knapsackTable[i-1][w-weights[i]]
			}
			myval := math.Max(float64(notTakingItem), float64(takingItem))
			knapsackTable[i][w] = int(myval)
		}
	}
	totalBenefit = knapsackTable[numOfItems][knapsackCapacity]

}

func showResult() {
	fmt.Printf("Total benefit: %v\n", totalBenefit)

	for n, w := numOfItems, knapsackCapacity; n > 0; n-- {
		if knapsackTable[n][w] != 0 && knapsackTable[n][w] != knapsackTable[n-1][w] {
			fmt.Printf("We take item %v\n", n)
			w = w - weights[n]
		}
	}

}
