package main

import (
	"structs"
)

func main() {
	contacts := structs.GetNewContactBook()
	contacts.AddContact("Artem", "SPB", 898989898989)
	contacts.AddContact("Max", "SPB", 898989898989)
	contacts.AddContact("Igor", "SPB", 8982389898989)
}
