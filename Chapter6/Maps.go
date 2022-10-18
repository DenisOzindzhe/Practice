package main

import "fmt"

func main() {
	myMap := map[string]map[string]string{
		"D": {
			"Name":   "Denis",
			"Gender": "Male",
		},
		"V": {
			"Name":   "Vas",
			"Gender": "Male",
		},
	}
	//myMap = make(map[string]map[string]string)
	//Поиск элемента из мапы
	if name, ok := myMap["D"]; ok {
		//если есть такой эллемент, как в запросе, вернуть его и true, если такого в списке нету - то возвращает false
		fmt.Println(name, ok)
	} else {
		fmt.Println("not found")
	}
}
