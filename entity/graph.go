package entity

type NodeID string

type Node struct {
	ID NodeID
	Adjacencies []string
}
