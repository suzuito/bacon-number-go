package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/suzuito/bacon-number-go/entity"
	"github.com/suzuito/bacon-number-go/local"
)

func main() {
	// nodeStore := nodes01
	// startNodeID := entity.NodeID("01")
	nodeStore := local.Nodes02
	startNodeID := entity.NodeID("02")
	// nodeStore := nodes03
	// startNodeID := entity.NodeID("01")

	onQueue := make(chan *local.Data)
	queue := local.NewQueueImpl(onQueue)
	tableStore := local.NewTableStoreImpl()
	dvr := entity.DVRImpl{
		NodeStore:  &nodeStore,
		TableStore: tableStore,
		Queue:      queue,
	}

	ctx := context.Background()
	if err := dvr.Next(ctx, startNodeID); err != nil {
		fmt.Printf("Start error: %+v", err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	i := 0
	go func() {
		defer wg.Done()
		for {
			i++
			select {
			case dd := <-onQueue:
				if err := dvr.OnNode(ctx, dd.CurrentID, dd.FromID, dd.Cost); err != nil {
					fmt.Printf("Update error: %+v", err)
					break
				}
			default:
				fmt.Println("Queued datas is empty")
				time.Sleep(time.Second)
				continue
			}
			fmt.Printf("== %d ==\n", i)
			fmt.Println(tableStore.String())
		}
	}()
	wg.Wait()
}
