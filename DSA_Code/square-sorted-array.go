package main

func squareSortedArray() {
	mySlice := []int{-6, -3, -1, 2, 4, 5}
	n := len(mySlice)
	var k int
	for k = 0; k < n; k++ {
		if mySlice[k] >= 0 {
			break
		}
	}
	i := k - 1
	j := k
	ind := 0

}
