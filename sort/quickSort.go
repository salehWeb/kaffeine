package main

import "fmt"

func main() {
	data := []int{1, 1, 12, 35, 643, -3298, 230, 282, 11, 23, 4, 5, 6, 7, 8, 9, 3, 2, 11}

	fmt.Println(QuickSort(data))
}

func QuickSort(data []int) []int {
	quickSort(data, 0, len(data)-1)
	return data
}

func quickSort(data []int, start int, end int) {
	if end-start < 1 {
		return
	}

	pivot := partition(data, start, end)
	quickSort(data, start, pivot-1)
	quickSort(data, pivot+1, end)
}

func partition(data []int, start int, end int) int {
	pivot := data[end]
	i := start - 1
	j := start
	var temp int

	for j < end {
		if data[j] < pivot {
			i++
			temp = data[i]
			data[i] = data[j]
			data[j] = temp
		}

        j++
	}

	i++
	temp = data[i]
	data[i] = data[end]
	data[end] = temp

	return i
}
