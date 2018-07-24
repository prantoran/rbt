package rbt

const (
	eps = 1e-9
)

// Item is the interface that the items must
// satisfy to store in the red-black tree.
// The Compare function is used by rbt to compare
// two pairs of items in the tree.
type Item interface {
	//Compare should return true if parent item is less than compared item
	Compare(u interface{}) bool
}

type Int struct {
	key int
}

func (i Int) Compare(u interface{}) bool {
	return i.key < u.(Int).key
}

type Long struct {
	key int64
}

func (l *Long) Compare(u interface{}) bool {
	return l.key < u.(Long).key
}

type Float struct {
	key float64
}

func (f *Float) Compare(u interface{}) bool {
	return f.key+eps < u.(Float).key
}
