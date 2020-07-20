package main

import "fmt"

func main() {

	var value int

	fmt.Println("Введите число")
	if fmt.Scanln(&value); isEven(value) {
		fmt.Println("Вы ввели четное число")
	} else {
		fmt.Println("Вы ввели нечетное число")
	}
}

func isEven(value int) bool { return value % 2 == 0 }