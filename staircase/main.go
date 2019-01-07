package main

import "fmt"

func main() {
	Staircase(5)
}

func Staircase(n int32) {
	var i int32

	for ; i < n; i++ {
		content := ""

		var j, k int32

		for ; j <= (n-i)-1; j++ {
			content += " "
		}

		for ; k <= i; k++ {
			content += "#"
		}
		fmt.Println(content)
	}
}
