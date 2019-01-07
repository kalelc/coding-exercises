package main

import (
	"math"
)

const Elements = 3

// first element has number of elements into of array

func DiagonalDifference(arr *[][]int32) int32 {
	var result1 int32
	var result2 int32
	elements := *arr

	for i := 0; i < Elements; i++ {
		temp := (Elements - 1) - i
		result1 += elements[i][i]
		result2 += elements[i][temp]
	}

	return int32(math.Abs(float64(result1 - result2)))
}
