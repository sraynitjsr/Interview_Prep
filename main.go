package main

import (
	"fmt"

	"github.com/sraynitjsr/One"
	"github.com/sraynitjsr/Two"
)

func main() {
	fmt.Println("\nInside Main")

	fmt.Println("")
	One.Start()
	fmt.Println("")
	Two.Start()

	fmt.Println("")
}
