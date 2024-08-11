//go:build bubbleSort
// +build bubbleSort

package main

import "fmt"

func main() {
	data := []int{-19, -18, 16, 100, 11, 145, -100, 19}
	fmt.Println(BubbleSort(data))
}

func BubbleSort(data []int) []int {
	ln := len(data)

	for i := ln - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}

	return data
}
