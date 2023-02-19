package Two

import "fmt"

func Start() {
	fmt.Println("Inside Majority Element")
	mySlice := []int{1, 1, 2, 1, 3, 5, 1}
	myMap := make(map[int]int)
	halfLength := len(mySlice)
	for _, data := range mySlice {
		currentValue, present := myMap[data]
		if present {
			myMap[data] = currentValue + 1
		} else {
			myMap[data] = 1
		}
	}
	for k, v := range myMap {
		if v > halfLength/2 {
			fmt.Println("Majority Element is =>", k)
			return
		}
	}
	fmt.Println("No Majority ELement Found")
}
