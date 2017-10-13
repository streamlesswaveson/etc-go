package main

import "fmt"

func main() {

	var size, tmp, search int
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&tmp)
		arr[i] = tmp
	}
	fmt.Scan(&search)

	where := binarySearch(arr, search)
	fmt.Println(where)

}

func binarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		elem := arr[mid]
		if elem == item {
			return mid
		}
		if elem < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1

}
