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

// transplant replace node ref u with node ref v,
// should only be used when u has less than 2 children
func (t *Tree) transplant(u, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// find returns the node in tree that contains the item i
func (t *Tree) find(i Item) *node {
	z := node{
		key: i,
	}
	x := t.root

	for x != nil {
		if x.key == i {
			return x
		}

		if z.key.Compare(x.key) {
			x = x.left
			continue
		}
		x = x.right
	}

	return nil
}

func (t *Tree) minimum(x *node) *node {
	if x == nil {
		return nil
	}
	var ret = x
	for ret.left != nil {
		ret = ret.left
	}
	return ret
}
