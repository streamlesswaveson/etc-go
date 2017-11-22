package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var N int
	fmt.Scan(&N)
	fmt.Println(reverse(N))
}


func reverse(n int) int {
	str := strconv.Itoa(n)
	rev := ""
	for _,s := range str {
		rev = string(s) + rev
	}
	if rev[len(rev)-1] == '-' {
		rev = string(rev[0:len(rev)-1])
		v, _ := strconv.Atoi(rev)
		return v*-1

	}
	v, _ := strconv.Atoi(rev)
	return v
}