package local

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/suzuito/bacon-number-go/entity"
)

type Data struct {
	CurrentID entity.NodeID
	FromID    entity.NodeID
	Cost      int
}

type QueueImpl struct {
	lock  *sync.Mutex
	datas []Data
	out   chan *Data
}

func NewQueueImpl(out chan *Data) *QueueImpl {
	return &QueueImpl{
		lock:  &sync.Mutex{},
		datas: []Data{},
		out:   out,
	}
}

func (q *QueueImpl) Enqueue(
	currentID entity.NodeID,
	fromID entity.NodeID,
	cost int,
) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	go func() {
		q.out <- &Data{CurrentID: currentID, FromID: fromID, Cost: cost}
	}()
	return nil
}

func RunQueue(
	ctx context.Context,
	tableStore *TableStoreImpl,
	onQueue chan *Data,
	dvr *entity.DVRImpl,
) {
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
