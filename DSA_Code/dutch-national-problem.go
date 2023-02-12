package main

import "fmt"

func dutchNatioanlProbelm() {
	mySlice := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	zeros := 0
	ones := 0
	twos := 0
	for _, data := range mySlice {
		switch data {
		case 0:
			zeros++
		case 1:
			ones++
		default:
			twos++
		}
	}
	for i := 1; i <= zeros; i++ {
		fmt.Print("0 ")
	}
	for i := 1; i <= ones; i++ {
		fmt.Print("1 ")
	}
	for i := 1; i <= twos; i++ {
		fmt.Print("2 ")
	}
}
