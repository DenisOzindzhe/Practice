package main

import "fmt"

/*
	defer часто используется в случаях, когда нужно освободить
	 ресурсы после завершения.
	 Например, открывая файл необходимо убедиться, что позже он должен быть закрыт. C defer это выглядит так:

	f, _ := os.Open(filename)
	defer f.Close()
*/

func main() {
	first()
	defer Second()
}
func first() {
	fmt.Println("First")
}
func Second() {
	fmt.Println("Second")
}
