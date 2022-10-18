package main

import "fmt"

func main() {
	var slice []uint = Evens(10)
	fmt.Println(slice)
}

/*
	Функция не принимает в себя параметров и возвращает функцию которая возврашает положительное четное число
	Данный способ относится к функциональному программированию, и вызов функции в функции называется замыканием
*/
func makeevengenerator() func() uint {
	var i uint = 0
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

/*
	Функция возвращает слайс (срез) из а положительных кратных 2 чисел, использует в себе функцию описанную выше
*/
func Evens(a uint) []uint {
	var ret []uint
	nexteven := makeevengenerator()
	for i := 0; i < int(a); i++ {
		ret = append(ret, nexteven())
	}
	return ret
}
