package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}
type Human struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Human) Speak() string {
	return "Hola!"
}

func main() {
	dog := Dog{}
	human := Human{}

	fmt.Println(dog.Speak())
	fmt.Println(human.Speak())

	var s Speaker = Dog{}

	if d, ok := s.(Dog); ok {
		fmt.Println("Dog:", d.Speak())
	} else {
		fmt.Println("Not a Dog")
	}

}
