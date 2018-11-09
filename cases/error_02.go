package cases

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}

func Error02() {
	fmt.Println(ErrNegativeSqrt(-2).Error())
}
