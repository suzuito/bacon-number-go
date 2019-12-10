package entity

import (
	"context"
	"fmt"
	"strings"
)

// TableRow ...
type TableRow struct {
	DestinationID NodeID
	NextID        NodeID
	Cost          int
}

func NewTableRow(destID NodeID) *TableRow {
	return &TableRow{
		DestinationID: destID,
		NextID:        "",
		Cost:          -1,
	}
}

func (t *TableRow) Clone() *TableRow {
	return &TableRow{
		DestinationID: t.DestinationID,
		NextID:        t.NextID,
		Cost:          t.Cost,
	}
}

func (t *TableRow) String() string {
	return fmt.Sprintf("%10s %10s %10d", t.DestinationID, t.NextID, t.Cost)
}

// Table ...
type Table map[NodeID]*TableRow

func (t *Table) UpdateCost(
	fromID NodeID,
	fromTable *Table,
	cost int,
) *Table {
	diffTable := Table{}
	if _, exist := (*t)[fromID]; !exist {
		(*t)[fromID] = NewTableRow(fromID)
		newedRow := &TableRow{
			DestinationID: fromID,
			NextID:        fromID,
			Cost:          cost,
		}
		(*t)[fromID] = newedRow
		diffTable[fromID] = newedRow
	}
	for destinationID, fromRow := range *fromTable {
		if _, exist := (*t)[destinationID]; !exist {
			newedRow := &TableRow{
				DestinationID: fromRow.DestinationID,
				NextID:        fromID,
				Cost:          fromRow.Cost + cost,
			}
			(*t)[destinationID] = newedRow
			diffTable[destinationID] = newedRow
		} else {
			currentRow := (*t)[destinationID]
			if fromRow.Cost+cost < currentRow.Cost {
				currentRow.NextID = fromID
				currentRow.Cost = fromRow.Cost + cost
				diffTable[destinationID] = currentRow
			}
		}
	}
	return &diffTable
}

func (t *Table) Clone() *Table {
	table := Table{}
	for _, row := range *t {
		table[row.DestinationID] = row.Clone()
	}
	return &table
}

func (t *Table) Merge(newed *Table) *Table {
	updated := Table{}
	for destID, newedRow := range *newed {
		currentRow, exist := (*t)[destID]
		if !exist {
			(*t)[destID] = newedRow
			updated[destID] = newedRow
			continue
		}
		if newedRow.Cost < currentRow.Cost {
			(*t)[destID] = newedRow
			updated[destID] = newedRow
		}
	}
	return &updated
}

func (t *Table) String() string {
	ret := []string{}
	for _, row := range *t {
		ret = append(ret, row.String())
	}
	return strings.Join(ret, "\n")
}

type DVRImpl struct {
	NodeStore  NodeStore
	TableStore TableStore
	Queue      Queue
}

func (d *DVRImpl) Update(
	ctx context.Context,
	currentID NodeID,
	fromID NodeID,
	cost int,
) (*Table, error) {
	return d.TableStore.UpdateTable(
		ctx,
		currentID,
		fromID,
		func(currentTable, fromTable *Table) (*Table, error) {
			diffTable := currentTable.UpdateCost(fromID, fromTable, cost)
			if len(*diffTable) <= 0 {
				// Stable
				return nil, nil
			}
			return currentTable, nil
		},
	)
}

func (d *DVRImpl) Next(
	ctx context.Context,
	currentID NodeID,
) error {
	currentNode := Node{}
	if err := d.NodeStore.GetNode(currentID, &currentNode); err != nil {
		return err
	}
	for _, adjID := range currentNode.Adjacencies {
		if err := d.Queue.Enqueue(adjID, currentID, 1); err != nil {
			return err
		}
	}
	return nil
}

func (d *DVRImpl) OnNode(
	ctx context.Context,
	currentID NodeID,
	fromID NodeID,
	cost int,
) error {
	if currentID == "" {
		return nil
	}
	var updatedTable *Table
	var err error
	updatedTable, err = d.Update(
		ctx,
		currentID,
		fromID,
		cost,
	)
	if err != nil {
		fmt.Printf("Update error: %+v", err)
		return err
	}
	if updatedTable != nil {
		if err := d.Next(ctx, currentID); err != nil {
			fmt.Printf("Next error: %+v", err)
			return err
		}
	}
	return nil
}
