package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Println("Конвертор градусов из F в С\nВведите градусы Фаренгейт")
	fmt.Scanf("%f", &fahrenheit)
	c := (fahrenheit - 32.0) * 5.0 / 9.0
	fmt.Println("Конвертирование прошло успешно, F =", fahrenheit, "Будет равно С=", c)
}
