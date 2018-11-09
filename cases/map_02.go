package cases

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount is a test
func WordCount(s string) map[string]int {
	m := map[string]int{}
	f := strings.Fields(s)
	// fmt.Printf("%T %v %v\n", f, f, len(f))
	for _, v := range f {
		m[v]++
	}
	return m
}

// Map02 is a map test
func Map02() {
	wc.Test(WordCount)
}
