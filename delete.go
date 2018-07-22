package rbt

// Delete finds the node z for item i, deletes the item,
// and fixes the chain of nodes upwards if
// the properties are broken
func (t *Tree) Delete(i Item) {
	z := t.find(i)
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
		// symmetrical
		if x == x.parent.right {
			w := x.parent.left
			if w.c == Red { // case 1
				w.c = Black
				x.parent.c = Red
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.left.c == Black && w.right.c == Black { // case 2
				w.c = Red
				x = x.parent
				continue
			}
			if w.left.c == Black { // case 3
				w.right.c = Black
				w.c = Red
				t.leftRotate(w)
				w = x.parent.left
			}
			// case 4
			w.c = x.parent.c
			x.parent.c = Black
			w.left.c = Black
			t.rightRotate(x.parent)
			x = t.root
			continue
		}
	}
	x.c = Black
}
