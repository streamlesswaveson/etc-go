package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(quicksort(arr))

}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]

	var left, right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] >= pivot {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}
