package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

func (p *Student) greeting() {
	fmt.Println("Hello~ Student")
}

type Student struct {
	Person
	school string
	grade  int
}

func main() {
	var s Student
	s.greeting()
	s.Person.greeting()
}
