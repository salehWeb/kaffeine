//go:build binarySearch
// +build binarySearch
package main

import (
	"fmt"
)

func main() {
	data := []int{-63, -9, 1, 2, 3, 4, 5, 7, 9, 11, 100}

	fmt.Println(BinarySearch(data, 11))
	fmt.Println(BinarySearch(data, -9))
	fmt.Println(BinarySearch(data, 11888))
}

func BinarySearch(data []int, target int) int {
	left := 0
	right := len(data) - 1

    // (dividend+(divisor-1))/divisor is the standard way to round up
    // golang by default round down integers
    mid := left + ((right-left)+1)/2

	for left <= right {
		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

		mid = left + ((right-left)+1)/2
	}

	return -1
}
