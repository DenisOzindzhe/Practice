package main

import "fmt"

/*
Последовательность чисел Фибоначчи определяется
как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Напишите рекурсивную функцию, находящую fib(n).
*/

func main() {
	defer fmt.Println("Closing")
	fmt.Println(fibo(6))
}
func fibo(x int) int {
	if x == 1 {
		return 1
	} else if x == 0 {
		return 0
	} else {
		return fibo(x-1) + fibo(x-2)
	}
}
