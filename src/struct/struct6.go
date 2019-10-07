package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

func (p *Person) babo() {
	fmt.Println("Babo~")
}

type Student struct {
	p      Person
	school string
	grade  int
}

func main() {
	var s Student
	var b Student
	s.p.greeting()
	b.p.babo()
}
