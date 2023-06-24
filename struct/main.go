package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// nayeem := person{"Nayeem", "Nishaat"}
	// nayeem := person{firstName: "Nayeem", lastName: "Nishaat"}
	// fmt.Println(nayeem)

	var nayeem person

	nayeem.firstName = "Nayeem"
	nayeem.lastName = "Nishaat"

	fmt.Printf("%+v", nayeem)
}
