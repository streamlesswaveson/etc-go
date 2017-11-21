package main

import "fmt"

/*
Towers of Hanoi
Minimum number of moves to solve is 2^n - 1 (exponential time)

 */
func main() {
	var N int
	fmt.Scan(&N)

	solveHanoi(N, "A","B","C")

}

func solveHanoi(n int, from string, middle string, to string) {

	if n == 1 {
		fmt.Println("plate 1 from " + from + " to " + to)
		return
	}

	solveHanoi(n-1, from, to, middle)
	fmt.Println("plate", n, "from " + from + " to " + to)
	solveHanoi(n-1, middle, from, to)
}
