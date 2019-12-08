package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/suzuito/bacon-number-go/entity"
)

func main() {
	nodeStore := nodes01
	startNodeID := entity.NodeID("01")

	queue := NewQueueImpl()
	tableStore := NewTableStoreImpl()
	dvr := entity.DVRImpl{
		NodeStore:  &nodeStore,
		TableStore: tableStore,
		Queue:      queue,
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(1)
	if err := dvr.Next(ctx, startNodeID); err != nil {
		fmt.Printf("Start error: %+v", err)
		return
	}
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			currentNodeID, fromNodeID, cost := queue.Dequeue()
			if currentNodeID == "" {
				continue
			}
			var updatedTable *entity.Table
			var err error
			updatedTable, err = dvr.Update(
				ctx,
				currentNodeID,
				fromNodeID,
				cost,
			)
			if err != nil {
				fmt.Printf("Update error: %+v", err)
				break
			}
			if updatedTable != nil {
				go func() {
					if err := dvr.Next(ctx, currentNodeID); err != nil {
						fmt.Printf("Next error: %+v", err)
					}
				}()
			}
			body, _ := json.MarshalIndent(tableStore.tables, "", "  ")
			fmt.Println(string(body))
		}
	}()
	wg.Wait()
}
