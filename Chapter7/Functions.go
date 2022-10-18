package main

import "fmt"

func main() {
	Interface("Folk\nHere is averageFunc")
	score := []float64{4, 4, 4, 4, 5, 5}
	fmt.Println(Middle(score))
}

func Interface(value string) {
	fmt.Println("Welcome to my App", value)
}
func Middle(slice []float64) float64 {
	total := 0.0
	//slice = make([]float64, len(slice))
	for _, value := range slice {
		total += value
	}
	//Приведение типа в Go, используем название типа как функцию

	return total / float64(len(slice))
}
