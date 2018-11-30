package main

import (
	"fmt"
)

/**
 *	数据类型
 */
func main() {
	// Bool
	b1 := true
	fmt.Printf("b1: %v %T\n", b1, b1)

	// 整型
	num1 := 1
	num2 := 999999999999999
	fmt.Printf("num1: %v %T\n", num1, num1)
	fmt.Printf("num2: %v %T\n", num2, num2)

	// 浮点类型
	f1 := 12.43
	f2 := 1.243e100
	fmt.Printf("f1: %v %T\n", f1, f1)
	fmt.Printf("f2: %v %T\n", f2, f2)

	// 复数类型
	z1 := 1 + 12i
	fmt.Printf("z1: %v %T\n", z1, z1)

	// 字符串类型
	s1 := "abcdef"
	fmt.Printf("s1: %v %T len=%v\n", s1, s1, len(s1))

	// struct
	type Obj struct {
	}
	var obj Obj
	if obj == (Obj{}) {
		fmt.Println("obj is empty")
	}
	fmt.Printf("obj: %v %T\n", obj, obj)
	// or fmt.Printf("obj: %v %v\n", obj, reflect.TypeOf(obj))
}
