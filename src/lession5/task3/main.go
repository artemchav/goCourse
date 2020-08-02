package main

import (
	"encoding/csv"
	"os"
)

func main() {
	addr := [][]string{
		{"Ira", "SPB"},
		{"Anya", "Surgut"},
		{"Petya", "Moscow"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(addr)

	w.Comma = ';'
	w.WriteAll(addr)
}
