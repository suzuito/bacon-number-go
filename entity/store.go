package entity

type NodeStore interface {
	GetNode(id NodeID, node *Node) error
	PutNode(node *Node) error
}

type TableStore interface {
	GetTable(id NodeID, *Table) error
	PutTable(id NodeID, table *Table) error
}