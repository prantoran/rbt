package rbt

type Color int

const (
	Red   Color = 1
	Black Color = 2
)

type node struct {
	left, right, parent *node
	value               Item
	c                   Color
}

func (n *node) Color() Color {
	return n.c
}
