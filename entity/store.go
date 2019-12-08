package entity

import "context"

import "fmt"

var (
	NotExistErr = fmt.Errorf("Not exist")
)

type NodeStore interface {
	GetNode(id NodeID, node *Node) error
	PutNode(node *Node) error
}

type TableStore interface {
	UpdateTable(
		ctx context.Context,
		id NodeID,
		fn func(
			currentTable *Table,
		) (*Table, error),
	) (*Table, error)
}
