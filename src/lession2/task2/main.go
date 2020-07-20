package main

import (
	"fmt"
	"strconv"
)

func main() {

	var value int

	fmt.Println("Введите число")

	_, err := fmt.Scanln(&value)
	if err != nil {
		panic("Неверное число")
	}

	if isDivideBy3(value){
		println("Число делится на 3 без остатка")
	} else {
		println("Число не делится на 3 без остатка")
	}
}

//число делится на 3 без остатка, если сумма его цифр делится на 3 без остатка
func isDivideBy3(value int) bool {

	var charSum int

	valueString := strconv.Itoa(value)
	for _, n := range valueString {
		char, _ := strconv.Atoi(string(n))
		charSum += char
	}

	return charSum % 3 == 0
}
