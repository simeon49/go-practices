package structures

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (h *Human) Eat(food string) {
	fmt.Printf("%v is eating %v\n", h.name, food)
}

type Man struct {
	Human
}

func (man *Man) Work() {
	fmt.Printf("%v is working", man.name)
}

// Struct01 :
// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
func Struct01() {
	human := Human{
		name: "Jack",
		age:  26,
	}
	fmt.Printf("%T %v\n", human, human)
	human.Eat("apple")

	man := Man{
		Human: Human{
			name: "Tom",
			age:  21,
		},
	}
	fmt.Printf("%T %v\n", man, man)
	man.Eat("hamburger")
	man.Work()
}
