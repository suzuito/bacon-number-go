package entity

import "context"

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

// Table ...
type Table map[NodeID]*TableRow

func (t *Table) UpdateCost(
	fromID NodeID,
	fromTable *Table,
	cost int,
) {
	if _, exist := (*t)[fromID]; !exist {
		(*t)[fromID] = NewTableRow(fromID)
	}
	for _, row := range *t {

		row.NextID = fromID
		row.Cost += cost
	}
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
			updatedTable := currentTable.Clone()
			updatedTable.UpdateCost(fromID, fromTable, cost)
			diffTable := currentTable.Merge(updatedTable)
			if len(*diffTable) <= 0 {
				// Stable
				return nil, nil
			}
			return updatedTable, nil
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
