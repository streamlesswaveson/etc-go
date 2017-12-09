package main

import "fmt"

var cache = map[int]int{}
var N int

func main()  {
	fmt.Scan(&N)

	fmt.Println(fib(N))
}


func fib(n int) int  {

	if v,ok := cache[n]; ok {
		return v
	}

	if n < 2 {
		return n
	}

	result := fib(n-1) + fib(n-2)
	cache[n] = result
	return result
}

