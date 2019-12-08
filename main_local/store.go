package main

import (
	"context"

	"github.com/suzuito/bacon-number-go/entity"
)

type NodeStoreImpl struct {
	nodes map[entity.NodeID]entity.Node
}

func (n *NodeStoreImpl) GetNode(
	id entity.NodeID,
	node *entity.Node,
) error {
	ret := entity.Node{}
	exist := true
	if ret, exist = n.nodes[id]; !exist {
		return entity.NotExistErr
	}
	*node = ret
	return nil
}

func (n *NodeStoreImpl) PutNode(
	node *entity.Node,
) error {
	n.nodes[node.ID] = *node
	return nil
}

type TableStoreImpl struct {
	tables map[entity.NodeID]entity.Table
}

func NewTableStoreImpl() *TableStoreImpl {
	return &TableStoreImpl{
		tables: map[entity.NodeID]entity.Table{},
	}
}

func (t *TableStoreImpl) UpdateTable(
	ctx context.Context,
	id entity.NodeID,
	fn func(
		currentTable *entity.Table,
	) (*entity.Table, error),
) (*entity.Table, error) {
	currentTable := entity.Table{}
	exist := false
	if currentTable, exist = t.tables[id]; !exist {
		currentTable = entity.Table{}
	}
	updatedTable, err := fn(&currentTable)
	if err != nil {
		return nil, err
	}
	if updatedTable != nil {
		t.tables[id] = *updatedTable
	}
	return updatedTable, nil
}
