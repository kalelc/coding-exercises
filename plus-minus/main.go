package main

import "fmt"

func main() {
	var values = []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(&values)
}

func plusMinus(arr *[]int32) {
	elements := *arr
	var values = []int32{0, 0, 0}
	var i int32

	for _, j := range elements {
		if j > 0 {
			values[0]++
		} else if j < 0 {
			values[1]++
		} else {
			values[2]++
		}
		i++
	}

	for _, value := range values {
		fmt.Printf("%f\n", float64(value)/float64(i))
	}

}
