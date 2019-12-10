package main

import (
	"context"
	"github.com/suzuito/bacon-number-go/entity"
	"github.com/suzuito/bacon-number-go/intf/web"
	"github.com/suzuito/bacon-number-go/local"
)

func main() {
	// nodeStore := nodes01
	// startNodeID := entity.NodeID("01")
	// nodeStore := local.Nodes02
	// startNodeID := entity.NodeID("02")
	// nodeStore := nodes03
	// startNodeID := entity.NodeID("01")

	nodeStore := local.NodeStoreImpl{
		Nodes: make(map[entity.NodeID]*entity.Node),
	}
	onQueue := make(chan *local.Data)
	queue := local.NewQueueImpl(onQueue)
	tableStore := local.NewTableStoreImpl()
	dvr := entity.DVRImpl{
		NodeStore:  &nodeStore,
		TableStore: tableStore,
		Queue:      queue,
	}
	ctx := context.Background()
	go local.RunQueue(
		ctx,
		tableStore,
		onQueue,
		&dvr,
	)
	r := web.NewRoute(&nodeStore, tableStore, &dvr)
	if err := r.Run("localhost:8081"); err != nil {
		panic(err)
	}
}
