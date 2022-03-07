package scene

import (
	"github.com/google/uuid"
)
type Node struct {
	id       string
	parent   *Node
	children []*Node
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) AddNewNode() *Node {
	newChild := Node{
		id: uuid.NewString(),
		parent: n,
		children: make([]*Node, 1),
	}
	return &newChild
}
