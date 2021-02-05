package main

import "fmt"

type address struct {
	street string
	number int
}

type person struct {
	name string
	age  int
	address
}

func main() {
	felipe := person{
		name: "Felipe",
		age:  36,
		address: address{
			street: "Rua GC",
			number: 72,
		},
	}

	// felipePointer := &felipe
	// felipePointer.updateName("Pantoja")
	felipe.updateName("Pantoja")

	fmt.Printf("%+v\n", felipe)

	n1 := 10
	n2 := &n1
	*n2 = 20
	fmt.Println(n1, *n2)

	name := "Bill"
	fmt.Println(&name)
	fmt.Println(*&name)

	printNamePointer(&name)
}

func (p *person) updateName(newName string) {
	(*p).name = newName
}

func printNamePointer(name *string) {
	fmt.Println(&name)
}
