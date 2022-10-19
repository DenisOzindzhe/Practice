package main

import "fmt"

//Пример передачи переменной в функцию - сама переменная не меняется
//Для того чтобы это исправить и поменять переменную используют специальный тип - указатели
func main() {
	var x int = 5
	some(x)        // 10
	fmt.Println(x) //5
	somemodif(&x)
	fmt.Println(x) //10

}
func some(x int) {
	fmt.Println(x + 5)
}
func somemodif(xPtr *int) {
	*xPtr = *xPtr + 5 //Указываем на место в памяти где хранится передаваемое значение
}
