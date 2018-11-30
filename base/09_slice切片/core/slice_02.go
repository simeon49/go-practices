package core

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	t := make([][]uint8, dy)
	for i := range t {
		t[i] = make([]uint8, dx)
		for j := range t[i] {
			t[i][j] = uint8((i * j))
		}
	}
	return t
}

func Slice02() {
	pic.Show(Pic)
}
