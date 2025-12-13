package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func changeName(person *Person) {
	person.Name = "Mark"
}

func main() {
	person := Person{"Vad"}
	fmt.Println(person.Name)
	changeName(&person)
	fmt.Println(person.Name)
}
