package main

import (
	"interfaces"
	"structs"
)

func main() {

	citizen := &structs.Citizen{
		Name:    "Upyachka",
		Country: "Zimbabwe",
		City:    "Abracadabra",
	}

	stranger := &structs.Stranger{"Bomzh"}
	greet(citizen, stranger)
}

func greet(greeters ...interfaces.Greeter) {
	for _, g := range greeters {
		g.SayHello()
	}
}
