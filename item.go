package rbt

// Item is the interface that the items must
// satisfy to store in the red-black tree.
// The Compare function is used by rbt to compare
// two pairs of items in the tree.
type Item interface {
	//Compare should return true if parent item is less than compared item
	Compare(Item) bool
}
