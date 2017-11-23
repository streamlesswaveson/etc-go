package main

import (
	"fmt"
	"sort"
)


func main()  {
	var N, K int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Scan(&K)

	fmt.Println(binarySearch(arr, 0, N-1, K))


}

func binarySearch(arr []int, start int, end int, item int) int {

	if end < start {
		fmt.Println("not found!")
		return -1
	}

	middle := (start + end) / 2

	if item == arr[middle] {
		return middle
	} else if item < arr[middle] {
		return binarySearch(arr, start, middle -1, item)
	} else {
		return binarySearch(arr, middle+1, end, item)
	}

}