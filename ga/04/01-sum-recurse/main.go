package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(sum(arr))
}

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sum(arr[1:])
}
