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
