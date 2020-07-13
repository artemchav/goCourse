package main

import "fmt"

const YEARS = 5

var deposit, percent float64

func main(){
	println("Введите сумму вклада")
	fmt.Scanln(&deposit)
	println("Введите годовой процент")
	fmt.Scanln(&percent)

	for i:=0; i<YEARS; i++ {
		deposit += deposit * percent / 100
		fmt.Printf("%d-й год: %d \n", i+1, int(deposit))
	}

}
