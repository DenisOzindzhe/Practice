package main

import "fmt"

func main() {
	//Массивы в go
	fmt.Println("Массивы")
	score := [5]float64{4.2, 5.3, 3.1, 4.9, 3.4}
	fmt.Println("Оценки пользователя =", score)

	total := 1.0
	/*
		Долгий способ записи цикла
		for i := 0; i < len(score); i++ {
			total += score[i]
		}
	*/
	for _, value := range score {
		total += value
	}
	//Приведение типа в Go, используем название типа как функцию
	total = total / float64(len(score))
	fmt.Println("Средний балл пользователя = ", total)

}
