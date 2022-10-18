package main

import "fmt"

func main() {
	x := []int{
		4, 96, 86, 68,
		57, 82, 63, 70,
		37, 1, 83, 27,
		19, 97, 3, 17, 23,
	}

	var maxvalue int = x[0]
	for _, value := range x {

		if value < maxvalue {
			maxvalue = value
		}
	}
	fmt.Println(maxvalue)
}
