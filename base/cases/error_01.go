package cases

import (
	"fmt"
)

type MyError struct {
	code string
	msg  string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("Error code: %v msg: %v", err.code, err.msg)
}

func run() error {
	return &MyError{"404", "page not found!"}
}

func Error01() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
