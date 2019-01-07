package main

const Elements = 3

func CompareTriplets(a, b *[]int32) []int32 {
	var result = []int32{0, 0}

	first := *a
	second := *b

	for i := 0; i < Elements; i++ {
		if first[i] > second[i] {
			result[0]++
		} else if first[i] < second[i] {
			result[1]++
		}
	}

	return result
}
