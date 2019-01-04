package main

const Elements = 3

func CompareTriplets(a, b []int32) []int32 {
	var result = []int32{0, 0}

	for i := 0; i < Elements; i++ {
		if a[i] > b[i] {
			result[0]++
		} else if a[i] < b[i] {
			result[1]++
		}
	}

	return result
}
