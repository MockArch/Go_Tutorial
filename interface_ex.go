package main

import "fmt"

type Animal interface {
	Speak() string
}

type dog struct {
}

func (d dog) Speak() string {
	return "woof"
}

type cat struct {
}

func (c cat) Speak() string {
	return "meow"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{dog{}, cat{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
