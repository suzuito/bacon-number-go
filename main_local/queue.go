package main

import (
	"sync"

	"github.com/suzuito/bacon-number-go/entity"
)

type data struct {
	currentID entity.NodeID
	fromID    entity.NodeID
	cost      int
}

type QueueImpl struct {
	lock  *sync.Mutex
	datas []data
}

func NewQueueImpl() *QueueImpl {
	return &QueueImpl{
		lock:  &sync.Mutex{},
		datas: []data{},
	}
}

func (q *QueueImpl) Enqueue(
	currentID entity.NodeID,
	fromID entity.NodeID,
	cost int,
) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.datas = append(q.datas, data{
		currentID: currentID,
		fromID:    fromID,
		cost:      cost,
	})
	return nil
}

func (q *QueueImpl) Dequeue() (entity.NodeID, entity.NodeID, int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	datas := q.datas
	q.datas = []data{}
	currentID := entity.NodeID("")
	fromID := entity.NodeID("")
	cost := -1
	for i, data := range datas {
		if i == 0 {
			currentID = data.currentID
			fromID = data.fromID
			cost = data.cost
			continue
		}
		q.datas = append(q.datas, data)
	}
	return currentID, fromID, cost
}
