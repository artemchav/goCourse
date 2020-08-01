package main

import (
	"fmt"
	"sort"
	"structs"
)

func main() {

	phones := &structs.PhoneBook{
		structs.Contact{Name: "Ololo", Address: "Podval", Phones: []int{123123}},
		structs.Contact{Name: "Abcdef", Address: "City", Phones: []int{123123}},
		structs.Contact{Name: "Zcbn", Address: "Dom", Phones: []int{123123}},
		structs.Contact{Name: "Ycbn", Address: "Home", Phones: []int{123123}},
	}

	fmt.Printf("%#v\n", phones)
	sort.Sort(phones)
	fmt.Printf("%#v\n", phones)

}
