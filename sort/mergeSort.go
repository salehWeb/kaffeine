//go:build mergeSort
// +build mergeSort

package main

import "fmt"

func main() {
	data := []int{9, -10, 92, 5, 32, 12, 188, -10}
	fmt.Println(Sort(data))
}

func Merge(data []int, write []int, start int, mid int, end int) {
	l := start
	r := mid + 1
	i := start

	for l <= mid && r <= end {
		if data[l] <= data[r] {
			write[i] = data[l]
			l++
		} else {
			write[i] = data[r]
			r++
		}
		i++
	}

	for l <= mid {
		write[i] = data[l]
		l++
		i++
	}

	for r <= end {
		write[i] = data[r]
		r++
		i++
	}

	for i = start; i <= end; i++ {
		data[i] = write[i]
	}
}

func MergeSort(data []int, write []int, start int, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2

	MergeSort(data, write, start, mid)
	MergeSort(data, write, mid+1, end)

	Merge(data, write, start, mid, end)
}

func Sort(data []int) []int {
	ln := len(data)
	write := make([]int, ln)

	MergeSort(data, write, 0, ln-1)

	return data
}
