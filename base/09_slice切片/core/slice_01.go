package core

import (
	"fmt"
)

func append(s []string, data ...string) []string {
	m := len(s)
	n := m + len(data)
	t := make([]string, n)
	fmt.Printf("%T %v\n", data, data)

	for i, _ := range s {
		t[i] = s[i]
	}
	copy(t[m:n], data)
	return t
}

func Slice01() {
	s1 := []string{"abc", "edf", "hig", "klm", "nop"}
	s2 := []string{"123", "456"}
	fmt.Printf("s1: %v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v %v %v\n", s2, len(s2), cap(s2))

	s3 := append(s1, s2...)
	fmt.Printf("s3: %v %v %v\n", s3, len(s3), cap(s3))
}
