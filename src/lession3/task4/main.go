package main

//прошу прощения, спешу, не успеваю сдать до закрытия такой возможности )
import "lession3/task4/structs"

func main() {
	contacts := &structs.ContactBook{}
	contacts.AddContact("Artem", "SPB", 898989898989)
	contacts.AddContact("Max", "SPB", 898989898989)
}
