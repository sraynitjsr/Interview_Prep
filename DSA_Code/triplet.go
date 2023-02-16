package main

import (
	"fmt"
	"sort"
)

func sixteen() {
	fmt.Println("Inside Sixteen")
	mySlice := []int{1, 4, 6, 10, 12, 15, 25, 30, 45}
	tripletSum := 35
	var start, end int
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	})
	for i := 0; i < len(mySlice); i++ {
		start = i + 1
		end = len(mySlice) - 1
		for start < end {
			if mySlice[i]+mySlice[start]+mySlice[end] == tripletSum {
				fmt.Println("Triplet is", mySlice[start], mySlice[i], mySlice[end])
				return
			} else if mySlice[i]+mySlice[start]+mySlice[end] < tripletSum {
				start++
			} else {
				end--
			}
		}
	}
	fmt.Println("No Triplet Found")
}
