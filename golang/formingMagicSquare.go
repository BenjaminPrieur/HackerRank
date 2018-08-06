package main

import (
	"fmt"
	"sort"
)

func formingMagicSquare(s [][]int32) int32 {
	cost := []int32{0, 0, 0, 0, 0, 0, 0, 0}
	magicMat := [][]int32{
		{4, 9, 2, 3, 5, 7, 8, 1, 6},
		{4, 3, 8, 9, 5, 1, 2, 7, 6},
		{2, 9, 4, 7, 5, 3, 6, 1, 8},
		{2, 7, 6, 9, 5, 1, 4, 3, 8},
		{8, 1, 6, 3, 5, 7, 4, 9, 2},
		{8, 3, 4, 1, 5, 9, 6, 7, 2},
		{6, 7, 2, 1, 5, 9, 8, 3, 4},
		{6, 1, 8, 7, 5, 3, 2, 9, 4},
	}

	for i := 0; i < 8; i++ {
		j := 0
		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				cost[i] = cost[i] + Abs(magicMat[i][j]-s[k][l])
				j++
			}
		}
	}
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] < cost[j]
	})

	return cost[0]
}

// Absolute value for int32
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func formingMagicSquareSample() {
	magicSquare := [][]int32{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}}
	result := formingMagicSquare(magicSquare)
	fmt.Println(result)
}
