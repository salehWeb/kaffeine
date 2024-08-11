//go:build insertionSort
// +build insertionSort

package main

import "fmt"

func main() {
	data := []int{3, 4, 5, -63, 2, 1, -9, 7, 9}
	fmt.Println(insertionSort(data))
}

func insertionSort(data []int) []int {
	ln := len(data)
	for i := 1; i < ln; i++ {
		// item to insert
		item := data[i]
		// end of sorted sub array
		j := i - 1

		for j > -1 {
			// found place for item
			if data[j] <= item {
				break
			}

			// push sorted array items forward until you find place for item
			data[j+1] = data[j]
			j--
		}

		// found item place and set value
		data[j+1] = item
	}

	return data
}
