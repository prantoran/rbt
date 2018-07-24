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
	ret := ""
	ret += " ("
	if u.left != nil {
		ret += print(u.left)
	}
	ret += fmt.Sprintf("%v", u.key)
	if u.right != nil {
		ret += print(u.right)
	}
	ret += ") "
	return ret
}

func TestLeftRotateRoot(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.leftRotate(tree.root)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestLeftRotateRootLeft(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.leftRotate(tree.root.left)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestLeftRotateRootRight(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.leftRotate(tree.root.right)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestRightRotateRoot(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.rightRotate(tree.root)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestRightRotateRootLeft(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.rightRotate(tree.root.left)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestRightRotateRootRight(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.rightRotate(tree.root.right)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestTransplantRoot(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	tree.transplant(tree.root, tree.root.left)
	s = print(tree.root)
	t.Logf("tree: %v", s)

	tree = DummyTree()
	tree.transplant(tree.root, tree.root.right)
	s = print(tree.root)
	t.Logf("tree: %v", s)

}

func TestFind(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	n := tree.find(Int{4})
	if n != tree.root {
		t.Errorf("key 4 not root, searching for: %v", Int{4})

		if n != nil {
			t.Errorf("found key: %v", n.key)
		}
	}

	n = tree.find(Int{1})
	if n != tree.root.left.left {
		t.Errorf("key 1 not root.l.l, searching for: %v", Int{1})

		if n != nil {
			t.Errorf("found key: %v", n.key)
		}
	}

	n = tree.find(Int{7})
	if n != tree.root.right.right {
		t.Errorf("key 1 not root.r.r, searching for: %v", Int{7})

		if n != nil {
			t.Errorf("found key: %v", n.key)
		}
	}

	n = tree.find(Int{3})
	if n != tree.root.left.right {
		t.Errorf("key 1 not root.l.r, searching for: %v", Int{3})

		if n != nil {
			t.Errorf("found key: %v", n.key)
		}
	}
}

func TestMinimum(t *testing.T) {
	tree := DummyTree()
	s := print(tree.root)
	t.Logf("tree: %v", s)

	n := tree.minimum(tree.root.right)
	t.Logf("n.key: %v", n.key)
	expected := Int{5}
	if n.key != expected {
		t.Errorf("min of root.right not 5, found: %v", n.key)
	}
}
