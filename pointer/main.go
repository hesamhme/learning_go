package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p := Person{"Hesam", "Hashemi", 30}
	fmt.Println(p)
	// change type to a pointer to the same type with &
	ChangeName(&p)
	fmt.Println(p)

}

func ChangeName(p *Person) {
	p.FirstName = "ahmad"
}
