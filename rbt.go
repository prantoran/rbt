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

// Transplant replace node ref u with node ref v
func (t *Tree) Transplant(u, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

