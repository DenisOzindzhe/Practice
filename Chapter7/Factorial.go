package main

import "fmt"

/*
	Рекурсивный вызов функции - функция вызывает сама себя до тех пор пока не будет выполнено условия при котором функция
	вернет значение отличное от вызова себя
*/
func main() {
	fmt.Println(factorial(5))
}
func factorial(a uint) uint {
	if a == 0 {
		return 1
	}
	return (a * factorial(a-1))
}
