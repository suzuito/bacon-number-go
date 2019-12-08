package entity

// TableRow ...
type TableRow struct {
	DestinationID NodeID
	NextID NodeID
	Cost int
}

// Table ...
type Table map[NodeID]*TableRow


type DVR struct {
	NodeStore NodeStore
	TableStore TableStore
}

func (d *DVRImpl) Update(
	currentID NodeID,
	fromID NodeID,
	fromTable *Table,
) error {
	currentNode := Node{}
	if err := d.NodeStore.GetNode(currentID, &currentNode); err != nil {
		return err
	}

	
}