package main

import (
	"fmt"
)

var fibonacci = []int{1,1}
func main() {

	var value int

	fmt.Println("Введите количество чисел Фибоначчи")

	_, err := fmt.Scanln(&value)
	if err != nil {
		panic("Неверное число")
	}

	fmt.Println(Fibonacci(value))
}
func Fibonacci(n int) []int{
	if n > len(fibonacci) {
		for i := len(fibonacci); i < n; i++ {
			fibonacci = append(fibonacci, fibonacci[i-1] + fibonacci[i-2])
		}
	}
	return fibonacci[:n]
}
