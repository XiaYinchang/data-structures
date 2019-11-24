package main

import (
	"fmt"
)

// BubbleSort sort the input data ascending
func BubbleSort(input []int) []int {
	// sorted 标志位引入是为了提早结束扫描
	sorted := false
	count := len(input)
	times := 0
	for !sorted {
		sorted = true
		for i := 1; i < count; i++ {
			if input[i-1] > input[i] {
				input[i-1], input[i] = input[i], input[i-1]
				sorted = false
			}
		}
		times++
		fmt.Printf("this is %d scan\n", times)
	}
	return input
}

func main() {
	fmt.Printf("%v\n", BubbleSort([]int{8, 93, 2, 15, 9, 3, 4, 917, 3, 59, 64, 0, 9, 6, 5, 38, 6, 65, 49, 6947, 19, 2, 3}))
}
