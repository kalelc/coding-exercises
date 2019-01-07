package main

import (
	"testing"
)

type elements struct {
	a      *[]int32
	b      *[]int32
	result []int32
}

var okTests = []elements{
	{&[]int32{5, 6, 7}, &[]int32{3, 6, 10}, []int32{1, 1}},
	{&[]int32{17, 28, 30}, &[]int32{99, 16, 8}, []int32{2, 1}},
}

func TestCompareTriplets(t *testing.T) {
	for _, test := range okTests {
		result := CompareTriplets(test.a, test.b)
		if !compareResults(result, test.result) {
			t.Error(
				`For`, test.a,
				`Expected`, test.result,
				`Got`, result,
			)
		}
	}
}

func compareResults(a, b []int32) bool {
	status := false
	for i, j := range a {
		if b[i] == j {
			status = true
		} else {
			status = false
		}
	}
	return status
}
