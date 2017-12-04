package main

import "fmt"

func main() {
	var (
		// M = target total
		N, M, i uint64
		// number and value of coins
		V []uint64
	)
	fmt.Scan(&M, &N)
	V = make([]uint64, N)
	for i = 0; i < N; i++ {
		fmt.Scan(&V[i])
	}
	fmt.Println(dpCoinChange(V, M))

}

// exponential running time O(2^N)
func naiveCoinChange(M int, v []int, index int) int {
	if M < 0 {
		return 0
	}
	if M == 0 {
		return 1
	}
	if len(v) == int(index) {
		return 0
	}

	return naiveCoinChange(M-v[index], v, index) + naiveCoinChange(M, v, index+1)

}

func dpCoinChange(v []uint64, M uint64) uint64 {
	dpTable := make([][]uint64, len(v)+1)
	for i := 0; i < len(v)+1; i++ {
		dpTable[i] = make([]uint64, M+1)
	}
	for i := 0; i <= len(v); i++ {
		dpTable[i][0] = 1
	}
	var i, j uint64
	for i = 1; i <= M; i++ {
		dpTable[0][i] = 0
	}

	for i = 1; i <= uint64(len(v)); i++ {
		for j = 1; j <= M; j++ {

			if v[i-1] <= j {
				dpTable[i][j] = dpTable[i-1][j] + dpTable[i][j-v[i-1]]
			} else {
				dpTable[i][j] = dpTable[i-1][j]
			}
		}
	}
	fmt.Println(dpTable)

	return dpTable[len(v)][M]

}
