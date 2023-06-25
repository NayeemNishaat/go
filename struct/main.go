package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// nayeem := person{"Nayeem", "Nishaat"}
	// nayeem := person{firstName: "Nayeem", lastName: "Nishaat"}
	// fmt.Println(nayeem)

	// var nayeem person
	// nayeem.firstName = "Nayeem"
	// nayeem.lastName = "Nishaat"
	// fmt.Printf("%+v", nayeem)

	saymon := person{
		firstName: "Saymon",
		lastName:  "Ali",
		// contact: contactInfo{
		// 	email:   "saymon@mail.com",
		// 	zipCode: 12451,
		// },
		contactInfo: contactInfo{
			email:   "saymon@mail.com",
			zipCode: 12451,
		},
	}

	// saymonPointer := &saymon
	// saymonPointer.updateName("Sammy")

	saymon.updateName("Sam")
	saymon.print()

	v := true
	updateBool(&v)
	fmt.Println(v)

	b := bl(true)
	b.updtBl()
	fmt.Println(b)
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
	// p.firstName = firstName
}

// func (pointerToP *person) updateName(firstName string) {
// 	(*pointerToP).firstName = firstName
// }

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateBool(b *bool) {
	*b = false
}

type bl bool

func (b *bl) updtBl() {
	*b = false
}
