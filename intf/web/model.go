package web

import "github.com/suzuito/bacon-number-go/entity"

type Node struct {
	ID          entity.NodeID   `json:"id"`
	Adjacencies []entity.NodeID `json:"adjacencies"`
	Table       *Table          `json:"table"`
}

func NewNode(n *entity.Node, tbl *entity.Table) *Node {
	return &Node{
		ID:          n.ID,
		Adjacencies: n.Adjacencies,
		Table:       NewTable(tbl),
	}
}

type TableRow struct {
	DestinationID entity.NodeID `json:"destinationID"`
	NextID        entity.NodeID `json:"nextID"`
	Cost          int           `json:"cost"`
}

type Table []*TableRow

func NewTable(tbl *entity.Table) *Table {
	table := Table{}
	for _, row := range *tbl {
		table = append(table, &TableRow{
			DestinationID: row.DestinationID,
			NextID:        row.NextID,
			Cost:          row.Cost,
		})
	}
	return &table
}
