package main

import (
	"fmt"

	"./common"
)

func init() {
	fmt.Printf("Init in main 1 -> GlobalStr: '%v' localStr: '%v'\n", GlobalStr, localStr)
}

func init() {
	fmt.Println("Init in main 2.")
}

var (
	GlobalStr = "GlobalStr in main"
	localStr  = "localStr in main"
)

func main() {
	fmt.Println("========= main start =========")
	common.Foo()
	fmt.Println("========= main end =========")
}
