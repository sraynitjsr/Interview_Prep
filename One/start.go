package One

import "fmt"

func Start() {
	fmt.Println("Inside Stock Buy Sell to Maximize Profit")

	mySlice := []int{100, 180, 260, 310, 40, 535, 695}
	maxProfit := 0
	for i := 1; i < len(mySlice); i++ {
		if mySlice[i] > mySlice[i-1] {
			maxProfit = maxProfit + mySlice[i] - mySlice[i-1]
		}
	}
	fmt.Println(maxProfit)
}
