package main

import (
	"fmt"
)

var memoize map[uint64]uint64

func main()  {
	var N uint64
	fmt.Scan(&N)

	memoize = map[uint64]uint64{}
	memoize[0] = 0
	memoize[1] = 1


	//fmt.Println(fibDP(N))
	//fmt.Println(naiveFib(N))
	for i:=uint64(0); i<=N; i++ {
		fmt.Println(fibDP(i))
	}


}

func fibDP(n uint64) uint64 {

	if memoize == nil {
		memoize = map[uint64]uint64{}
		memoize[0] = 0
		memoize[1] = 1
	}

	if val, ok := memoize[n]; ok {
		return val
	}

	memoize[n-1] = fibDP(n-1)
	memoize[n-2] = fibDP(n-2)

	num := memoize[n-1] + memoize[n-2]
	memoize[n] = num
	return num
}

func naiveFib(n uint64) uint64  {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return naiveFib(n -1) + naiveFib(n-2)
}