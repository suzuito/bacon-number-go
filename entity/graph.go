package entity

type NodeID string

type Node struct {
	ID          NodeID
	Adjacencies []NodeID
}

func NewNode(id NodeID) *Node {
	return &Node{
		ID:          id,
		Adjacencies: []NodeID{},
	}
}

func (n *Node) AddAdjacency(id NodeID) {
	for _, adjID := range n.Adjacencies {
		if adjID == id {
			return
		}
	}
	n.Adjacencies = append(n.Adjacencies, id)
}
