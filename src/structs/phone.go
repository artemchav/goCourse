package structs

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
	//TODO: save JSON to file. Have not enough time now (
}

func GetNewContactBook() *ContactBook {
	c := &ContactBook{}
	c.contacts = make(map[string]*Contact)
	return c
}
