//go:build selectionSort
// +build selectionSort

package main

import "fmt"

func main() {
    data := []int{8,92,119,-199,179, 18888, 1078961, 2128}
    fmt.Println(SelectionSort(data))
}


func SelectionSort(data []int) []int {
    ln := len(data)
    var minIndex, temp int
        
    for i := 0; i < ln-1; i++ {
        minIndex = i;
        for j := i+1; j < ln; j++ {
            if data[minIndex] > data[j] {
                minIndex = j
            }
        } 

        temp = data[i]
        data[i] = data[minIndex]
        data[minIndex] = temp
    }

    return data
}
