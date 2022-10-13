package main

import "fmt"

func main() {
	var feet float64
	fmt.Println("Конвертор из футов в метры\nВведите футы")
	fmt.Scanf("%f", &feet)
	metres := feet * 0.3048 //(1 фут = 0.3048 метр) => 1 метр  = 3.28
	fmt.Println("Конвертирование прошло успешно, футы =", feet, "Будут равны м =", metres)
}
