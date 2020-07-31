package main

import "fmt"

const COURSE = 70.4

var rubles float32

func main(){
	println("Введите сумму в рублях")
	fmt.Scanln(&rubles)
	dollars := rubles / COURSE
	println(dollars)
}
