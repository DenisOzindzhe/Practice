package main

/*
Напишите функцию с переменным числом параметров,
которая находит наибольшее число в списке.
*/

import (
	"fmt"
)

func main() {
	defer fmt.Println("Closing")
	fmt.Println(solution(1, 2, 32, 5, 16, 1, 3, 14, 9, 2, 3, 1))
}

func solution(x ...int) int {
	var slice []int
	slice = x
	maxvalue := slice[0]
	for _, value := range x {
		if value >= maxvalue {
			maxvalue = value
		}
	}
	return maxvalue
}
