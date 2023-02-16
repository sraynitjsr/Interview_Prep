package main

import "fmt"

func majority() {
	mySlice := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	myMap := make(map[int]int)
	halfOfSlice := len(mySlice) / 2
	for _, val := range mySlice {
		v, present := myMap[val]
		if present {
			myMap[val] = v + 1
		} else {
			myMap[val] = 1
		}
	}
	for k, v := range myMap {
		if v > halfOfSlice {
			fmt.Println("Majority Element Is =>", k)
			return
		}
	}
	fmt.Println("No Majority Element Found", myMap)
}
