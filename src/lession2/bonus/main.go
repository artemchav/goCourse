package main

import (
	"fmt"
)

//Методичка: Ветвление с помощью условий
//Поменяем программу так, чтобы она указывала, является ли число четным.
//Без условий выполнения сделать подобное было бы невозможно (с)

//Даладна ))

func main() {

	var (
		value  int
		prefix = []string{"", "не"}
	)

	println("Введите число")
	fmt.Scanln(&value)
	index := value % 2
	fmt.Printf("Вы ввели %sчетное число", prefix[index])
}
