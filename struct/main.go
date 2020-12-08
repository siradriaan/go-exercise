package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber string
}

func main() {
	alex := person{
		firstName: "Adrian",
		lastName:  "Kurnianto",
		contactInfo: contactInfo{
			email:       "sir.adss.23@gmail.com",
			phoneNumber: "081389051948",
		},
	}
	alexPointer := &alex
	alexPointer.updateFirstName("Adrian")
	alex.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (pointerToPerson *person) updateFirstName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname
}
