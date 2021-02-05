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

func main() {
	// alex := person{"Alex", "Santos"}
	// fmt.Printf("%+v\n", alex)

	anderson := person{
		firstName: "Anderson",
		lastName:  "Silva",
		contact: contactInfo{
			email:   "anderson@gmail.co",
			zipCode: 3800,
		},
	}
	anderson.lastName = "Silveira"
	fmt.Printf("%+v\n", anderson)

	var bruce person
	bruce.firstName = "Bruce"
	bruce.lastName = "Wayne"

	fmt.Println(bruce)
}
