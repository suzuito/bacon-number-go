package entity

type Queue interface {
	Enqueue(
		currentID NodeID,
		fromID NodeID,
		cost int,
	) error
}
