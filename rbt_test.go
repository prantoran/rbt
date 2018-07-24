
package rbt

import (
	"fmt"
	"testing"
)

func DummyTree() *Tree {
	/*
			                ___7(B)
			               |
			     |--6(R)---|
			     |         |___5(B)
		         |
			4(B)-|
			     |          ___3(B)
			     |         |
			     |--2(R)---|
					       |___1(B)

	*/
	t := Tree{}
	t.root = &node{
		left:   nil,
		right:  nil,
		parent: nil,
		key:    Int{4},
		c:      Black,
	}

	t.root.left = &node{
		left:   nil,
		right:  nil,
		parent: t.root,
		key:    Int{2},
		c:      Red,
	}

	t.root.right = &node{
		left:   nil,
		right:  nil,
		parent: t.root,
		key:    Int{6},
		c:      Red,
	}

	t.root.left.left = &node{
		left:   nil,
		right:  nil,
		parent: t.root.left,
		key:    Int{1},
		c:      Black,
	}

	t.root.left.right = &node{
		left:   nil,
		right:  nil,
		parent: t.root.left,
		key:    Int{3},
		c:      Black,
	}

	t.root.right.left = &node{
		left:   nil,
		right:  nil,
		parent: t.root.right,
		key:    Int{5},
		c:      Black,
	}

	t.root.right.right = &node{
		left:   nil,
		right:  nil,
		parent: t.root.right,
		key:    Int{7},
		c:      Black,
	}
	return &t
}

func print(u *node) string {
	ret := "("
	if u.left != nil {
		ret += print(u.left)
	}
	ret += fmt.Sprintf("%v", u.key)
	if u.right != nil {
		ret += print(u.right)
	}
	ret += ")"
	return ret
}

func TestLeftRotate(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)
}
