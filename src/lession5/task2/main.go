package main

import (
	"io"
	"log"
	"os"
)

//Полагаю, надо изменить неупрощенный вариант
func main() {

	file, err := os.Open("main.go")
	if err != nil {
		// не надо молчать об ошибках
		// return → log.Fatal
		log.Fatal(err)
	}
	defer file.Close()

	// не нужна нам промежуточная переменная для вывода в консольку
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal(err)
	}
}
