package main

import (
	"github.com/davecgh/go-spew/spew"
	"structs"
)

func main() {
	contacts := structs.GetNewContactBook()
	contacts.AddContact("Artem", "SPB", 898989898989)
	contacts.AddContact("Max", "SPB", 898989898989)
	spew.Dump(contacts)
}
