package rbt

// Node is the interface that the items must
// satisfy to store in the red-black tree.
// The Compare function is used by rbt to compare
// two pairs of items in the tree.
type Node interface {
	Compare(Node) bool
}
