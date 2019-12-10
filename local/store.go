package local

import (
	"context"
	"fmt"
	"strings"

	"github.com/suzuito/bacon-number-go/entity"
)

type NodeStoreImpl struct {
	Nodes map[entity.NodeID]*entity.Node
}

func (n *NodeStoreImpl) GetNode(
	id entity.NodeID,
	node *entity.Node,
) error {
	var ret *entity.Node
	exist := true
	if ret, exist = n.Nodes[id]; !exist {
		return entity.NotExistErr
	}
	*node = *ret
	return nil
}

func (n *NodeStoreImpl) GetNodes(
	nodes *[]*entity.Node,
) error {
	for _, node := range n.Nodes {
		*nodes = append(*nodes, node)
	}
	return nil
}

func (n *NodeStoreImpl) PutNode(
	node *entity.Node,
) error {
	n.Nodes[node.ID] = node
	return nil
}

func (n *NodeStoreImpl) PutEdge(
	tailID, headID entity.NodeID,
	both bool,
) error {
	tailNode, exist := n.Nodes[tailID]
	if !exist {
		return entity.NotExistErr
	}
	headNode, exist := n.Nodes[headID]
	if !exist {
		return entity.NotExistErr
	}
	tailNode.AddAdjacency(headID)
	if both {
		headNode.AddAdjacency(tailID)
	}
	fmt.Println(tailNode)
	fmt.Println(headNode)
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
	currentID entity.NodeID,
	fromID entity.NodeID,
	fn func(
		currentTable, fromTable *entity.Table,
	) (*entity.Table, error),
) (*entity.Table, error) {
	exist := false
	currentTable := entity.Table{}
	if currentTable, exist = t.tables[currentID]; !exist {
		currentTable = entity.Table{}
	}
	fromTable := entity.Table{}
	if fromTable, exist = t.tables[fromID]; !exist {
		fromTable = entity.Table{}
	}
	updatedTable, err := fn(&currentTable, &fromTable)
	if err != nil {
		return nil, err
	}
	if updatedTable != nil {
		t.tables[currentID] = *updatedTable
	}
	return updatedTable, nil
}

func (t *TableStoreImpl) GetTable(
	ctx context.Context,
	id entity.NodeID,
	tbl *entity.Table,
) error {
	if table, exist := t.tables[id]; !exist {
		return nil
	} else {
		*tbl = table
	}
	return nil
}

func (t *TableStoreImpl) String() string {
	ret := []string{}
	for nodeID, table := range t.tables {
		ret = append(ret, string(nodeID)+"\n"+table.String())
	}
	return strings.Join(ret, "\n")
}
