package main

import "fmt"
// https://www.udemy.com/algorithmic-problems-in-java/learn/v4/t/lecture/4068848?start=0
func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(calcFact(N))

}

func calcFact(n int) int {
	return fact(1, n)
}

// faster implementation because it does not need to unwind
func fact(accumulator int, n int) int {

	if n == 1 {
		return accumulator
	}

	return fact(accumulator*n, n-1)

}
