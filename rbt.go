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
	var y *node
	for z.parent != nil && z.parent.c == Red {
		// root must always be black, hence if z.parent is red,
		// then z.parent.parent exists
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right // uncle
			if y.c == Red {
				z.parent.c = Black
				y.c = Black
				z.parent.parent.c = Red
				z = z.parent.parent
				continue
			}
			if z == z.parent.right {
				z = z.parent
				t.leftRotate(z)
			}
			z.parent.c = Black
			z.parent.parent.c = Red
			t.rightRotate(z.parent.parent)
			continue
		}
		// symmetrical to the above left case
		if z.parent == z.parent.parent.right {
			y = z.parent.parent.left
			if y.c == Red {
				z.parent.c = Black
				y.c = Black
				z.parent.parent.c = Red
				z = z.parent.parent
				continue
			}
			if z == z.parent.left {
				z = z.parent
				t.rightRotate(z)
			}
			z.parent.c = Black
			z.parent.parent.c = Red
			t.leftRotate(z.parent.parent)
			continue
		}
	}
	t.root.c = Black
}

// transplant replace node ref u with node ref v
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

func (t *Tree) Find(i Item) *node {
	z := node{
		key: i,
	}
	x := t.root
	var y *node
	for x != nil && z.key.Compare(x.key) {
		y = x.left
	}
	if y.key == i {
		return y
	}
	return nil
}

// Delete finds the node z for item i, deletes the item,
// and fixes the chain of nodes upwards if
// the properties are broken
func (t *Tree) Delete(i Item) {
	z := t.Find(i)
	y := z
	yc := y.c // y-original-color
	var x *node
	if z.left == nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimum(z.right)
		yc = y.c
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.c = z.c
	}
	if yc == Black {
		t.deleteFixup(x)
	}

}

func (t *Tree) minimum(n *node) *node {
	return nil
}

func (t *Tree) deleteFixup(x *node) {
	for x != t.root && x.c == Black {
		if x == x.parent.left {
			w := x.parent.right
			if w.c == Red { // case 1
				w.c = Black
				x.parent.c = Red
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.c == Black && w.right.c == Black { // case 2
				w.c = Red
				x = x.parent
				continue
			}
			if w.right.c == Black { // case 3
				w.left.c = Black
				w.c = Red
				t.rightRotate(w)
				w = x.parent.right
			}
			// case 4
			w.c = x.parent.c
			x.parent.c = Black
			w.right.c = Black
			t.leftRotate(x.parent)
			x = t.root
			continue
		}
	}
	x.c = Black
}
