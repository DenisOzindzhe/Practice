package main

import "fmt"

func main() {
	myfunc(23, 52, 25, 25, 15, 25, 24, 24, 23, 21, 3, 2)

	anonvoid := func(a, b int) int { return a + b }
	fmt.Println(anonvoid(23, 23))
}
func myfunc(a ...int) {
	fmt.Println(a)
}
