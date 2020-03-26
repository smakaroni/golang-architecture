package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from person - this is my name", p.first)
}

//any TYPE that has the methods specified by an interface is also of the interface type

type human interface {
	speak()
}

func main() {
	p1 := person{
		first:"Emma",
	}

	fmt.Printf("%T\n", p1)

	var x human
	// only works because person has the method speak()
	x = p1
	fmt.Printf("%T\n", x)
}
