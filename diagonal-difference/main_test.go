package main

import (
	"testing"
)

type elements struct {
	arr    *[][]int32
	result int32
}

var okTests = []elements{
	{
		&[][]int32{
			{1, 2, 3},
			{4, 5, 6},
			{9, 8, 9},
		}, 2,
	},
}

func TestDiagonalDifference(t *testing.T) {
	for _, test := range okTests {
		result := DiagonalDifference(test.arr)
		if !(test.result == result) {
			t.Error(
				`For`, test.arr,
				`Expected`, test.result,
				`Got`, result,
			)
		}
	}
}
