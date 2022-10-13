package main

import "fmt"

func main() {
	fmt.Println("Divide by 3 numbers from 1 to 100 includes")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
