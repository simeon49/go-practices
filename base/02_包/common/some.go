package common

import (
	"fmt"
)

func init() {
	fmt.Printf("Init in common -> GlobalStr: '%v' localStr: '%v'\n", GlobalStr, localStr)
}

const (
	GlobalStr = "GlobalStr in common"
	localStr  = "localStr in common"
)

func Foo() {
	fmt.Printf("Foo is called")
}
