package structs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phones  []int  `json:"phones"`
}

func (c *Contact) AddPhoneNumber(num int) { c.Phones = append(c.Phones, num) }

type ContactBook struct {
	contacts map[string]*Contact
}

func (c *ContactBook) AddContact(name, address string, phone int) {
	c.contacts[name] = &Contact{
		Name:    name,
		Address: address,
		Phones:  []int{phone},
	}

	j, err:= json.Marshal(c.contacts)
	if err != nil {
		log.Fatal("Can not marshal the structure")
	}

	err = ioutil.WriteFile("phonebook.json", j, 0644)
	if err != nil {
		log.Fatal("Can't write to file")
	}
}

func GetNewContactBook() *ContactBook {
	c := &ContactBook{}
	c.contacts = make(map[string]*Contact)
	return c
}
