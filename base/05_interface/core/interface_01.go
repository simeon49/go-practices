package core

import (
	"fmt"
)

type Human interface {
	Eat(string)
}

type Man struct {
	name string
}

func (man *Man) Eat(food string) {
	fmt.Printf("the man \"%v\" is eating %v\n", man.name, food)
}

func Interface01() {
	var tom Human
	tom = &Man{"Tom"}
	tom.Eat("apple")
}
