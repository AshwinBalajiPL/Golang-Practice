package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type employee struct {
	name    string
	age     int
	role    string
	domain  string
	contact contactInfo
}

func (e employee) print() {
	fmt.Printf("%+v", e)
	fmt.Println()
}

func main() {
	e1 := employee{
		name:   "Ashwin",
		age:    21,
		role:   "GET",
		domain: "Cloud And Backend",
		contact: contactInfo{
			email: "ashwin.1901023@srec.ac.in",
			phone: 7092158285,
		},
	}
	e1.print()
	e1.updateName("Ashwin Balaji")
	e1.print()
}

func (e *employee) updateName(name string) {
	(*e).name = name
}
