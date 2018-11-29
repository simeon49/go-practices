package core

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, v*k)
	}
	return t
}

func insert(t *Tree, value int) *Tree {
	if t == nil {
		t = &Tree{nil, value, nil}
	} else if value < t.Value {
		t.Left = insert(t.Left, value)
	} else {
		t.Right = insert(t.Right, value)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
