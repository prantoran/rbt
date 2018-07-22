package rbt

// Tree is the Red-Black Tree
type Tree struct {
	root *node
}

// leftRotate rotates left by one node with x's position as pivot
func (t *Tree) leftRotate(x *node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// rightRotate rotates right by one node with x's position as pivot.
// It is symmetrical to leftRotate
func (t *Tree) rightRotate(x *node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.right = x
	x.parent = y

}

// Insert encapsulates Item i with node, inserts node z
// into tree in a similar way to binary search tree,
// then fixes any violations to red black tree properties
func (t *Tree) Insert(i Item) {
	z := node{
		key: i,
	}
	var y *node
	x := t.root
	for x != nil {
		y = x
		if z.key.Compare(x.key) {
			x = x.left
			continue
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = &z
	} else if z.key.Compare(y.key) {
		y.left = &z
	} else {
		y.right = &z
	}
	z.left = nil
	z.right = nil
	z.c = Red
	t.insertFixup(&z)
}

func (t *Tree) insertFixup(z *node) {

}
