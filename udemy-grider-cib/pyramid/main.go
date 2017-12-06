package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		str := ""

		for j := N - i - 1; j > 0; j-- {
			str += " "
		}
		for j := 0; j < (i*2)+1; j++ {
			str += "#"
		}
		for j := N - i - 1; j > 0; j-- {
			str += " "
		}
		fmt.Println(str)
	}

}
