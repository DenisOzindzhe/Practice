package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("More than two")
		}

		if i%2 == 0 {
			fmt.Println("Number is ", i, " Even")
		} else {
			fmt.Println("Number is ", i, " Odd")
		}

	}
}
