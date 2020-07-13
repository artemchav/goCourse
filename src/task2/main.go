package main

import (
	"fmt"
	"math"
)

var a, b, c float64

func main() {

	println("Введите длину катетов треугольника, разделив их пробелом")
	fmt.Scanln(&a, &b)

	c = math.Sqrt(a * a + b * b)
	p := a + b + c
	s := a * b / 2

	println("Стороны треугольника:", a, b, c)
	println("P = ", p)
	println("S = ", s)
}
