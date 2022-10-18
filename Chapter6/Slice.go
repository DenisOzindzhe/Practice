package main

import "fmt"

func main() {
	var slice1 []float64
	var slice2 []float64

	slice1 = make([]float64, 0, 1)
	slice2 = make([]float64, 0, 2)
	slice1 = append(slice1, 9)
	slice2 = append(slice2, 4, 5, 6)
	copy(slice1, slice2)

	var slice3 []float64
	slice3 = make([]float64, 4, 8)
	slice3 = append(slice3, 1, 2, 3, 4)
	slice3[0] = 3
	fmt.Println(slice3)
	var slice4 []int
	slice4 = make([]int, 3, 9)
	fmt.Println(slice4)
	//	for _, value := range slice1 {
	//		fmt.Println(value)
	//	}

}
