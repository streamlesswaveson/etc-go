package main

import "fmt"

func main()  {
	var N int
	fmt.Scan(&N)

	fmt.Println(fact(N))

}

func fact(i int) int {
	if i == 1 {
		return 1
	}
	return i * fact(i-1)
}