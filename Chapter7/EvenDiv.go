package main

/*
Напишите функцию, которая принимает число,
делит его пополам и возвращает true в случае, если исходное число чётное, и false, если нечетное.
Например, half(1) должна вернуть (0, false),
 в то время как half(2) вернет (1, true).
*/

import "fmt"

func main() {
	defer fmt.Println("Closing")
	fmt.Println(solution(1))
	fmt.Println(solution(2))
}

func solution(x int) (out int, is bool) {
	return x / 2, x%2 == 0
}
