package main

import "fmt"

func main() {
	ptrX := new(int)
	some(ptrX)
	fmt.Println(*ptrX) //5

}
func some(ptrX *int) {
	*ptrX = 5
}
