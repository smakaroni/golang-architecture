package main

import "fmt"

type person struct {
	first string
}

type girl struct {
	person
	girl bool
}

func (g girl) speak() {
	fmt.Println("Im a girl - this is my name", g.first)
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
		first:"Jokke",
	}

	p2 := girl{
		person: person{
			first: "Emma",
		},
		girl:   true,
	}

	fmt.Printf("%T\n", p1)

	var x, y human
	// only works because person has the method speak()
	x = p1
	y = p2
	x.speak()
	y.speak()

}
