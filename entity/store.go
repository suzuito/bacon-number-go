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
		currentID NodeID,
		fromID NodeID,
		fn func(
			currentTable, fromTable *Table,
		) (*Table, error),
	) (*Table, error)
}
