package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	jim := person{firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email:   "jimparty@gmail.com",
			zipCode: 12345,
		}}
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
