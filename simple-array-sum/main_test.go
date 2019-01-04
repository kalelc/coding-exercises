package main

import "testing"

type elements struct {
	values []int32
	sum    int32
}

var okTests = []elements{
	{[]int32{1, 2, 3, 4, 5, 6}, 21},
	{[]int32{10, 21, 32, 43}, 106},
}

func TestSimpleArraySum(t *testing.T) {
	for _, test := range okTests {
		sum := SimpleArraySum(test.values)
		if sum != test.sum {
			t.Error(
				`For`, test.values,
				`Expected`, sum,
				`Got`, test.sum,
			)
		}
	}
}
