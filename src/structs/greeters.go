package structs

import "fmt"

type Stranger struct {
	Name string
}

func (s *Stranger) Strange() {}
func (s *Stranger) SayHello() {
	fmt.Printf("Hi, citizen, I'm stranger. My name is %s\n", s.Name)
}

type Citizen struct {
	Name    string
	Country string
	City    string
}

func (s *Stranger) Work() {}
func (c *Citizen) SayHello() {
	fmt.Printf("Hi, there, I'm citizen. Nice to meet you. My name is %s. I'm from %s, %s\n", c.Name, c.City, c.Country)
}
