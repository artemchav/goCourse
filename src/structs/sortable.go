package structs


type PhoneBook []Contact

func (s PhoneBook) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s PhoneBook) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s PhoneBook) Len() int {
	return len(s)
}
