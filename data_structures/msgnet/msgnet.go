package msgnet

import (
	"sync"
	"sync/atomic"
	"time"

	"taheri24.ir/blockchain/data_structures/stack"
)

const (
	MAX_CHANNEL      = 4
	MAX_NODE         = 3000
	CHANNEL_BUF_SIZE = 10
)

type msgnetActionMode int8

const (
	pushVal msgnetActionMode = iota
	pullVal
)

type msgnetItem[T any] struct {
	id         BucketID
	wg         *sync.WaitGroup
	pushValue  T
	pullValues []T
	mode       msgnetActionMode
}
type intArray []uint32
type msgnet[T any] struct {
	channels    [MAX_CHANNEL]chan *msgnetItem[T]
	syncCounter [MAX_CHANNEL]int64
}

type nodeHeader struct {
	// Primary Values
	nextIndex, backIndex uint32
	// DEBUG Fields, TODO: remove its after stable level
	isActive     bool
	bucketID     BucketID
	currentIndex uint32
}

func New[T any]() MsgNet[T] {
	var channels [MAX_CHANNEL](chan *msgnetItem[T])
	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan *msgnetItem[T], CHANNEL_BUF_SIZE)
	}

	return &msgnet[T]{channels: channels}
}

func (m msgnet[T]) Pull(id BucketID) []T {
	arrayIdx := id.ArrayIndex()
	chIdx := arrayIdx % MAX_CHANNEL
	wg := new(sync.WaitGroup)
	wg.Add(1)
	item := &msgnetItem[T]{
		wg:         wg, // TODO: add
		pullValues: nil,
		mode:       pullVal,
	}
	m.channels[chIdx] <- item
	wg.Wait()
	return item.pullValues
}

func (m msgnet[T]) Sync(d time.Duration) *time.Ticker {
	t := time.NewTicker(d)
	go func() {
		counter := 0
		opsCounter := m.syncCounter[0:]
		for range t.C {

			for i := range opsCounter {
				if opsCounter[i] != m.syncCounter[i] {
					counter = 0
					opsCounter = m.syncCounter[0:]
					break
				}
			}
			if counter > 10 {
				t.Stop()
				break
			}
		}
	}()
	return t
}

func (m msgnet[T]) Push(id BucketID, val T) {
	arrayIdx := id.ArrayIndex()
	chIdx := arrayIdx % MAX_CHANNEL
	atomic.AddInt64(&m.syncCounter[chIdx], 1)
	m.channels[chIdx] <- &msgnetItem[T]{
		wg:         &sync.WaitGroup{}, // TODO: add
		pullValues: make([]T, 0, 5),
		mode:       pushVal,
	}

}
func (m msgnet[T]) Handle(idx int) {
	nodes, headers := make([]T, MAX_NODE), make([]nodeHeader, MAX_NODE)
	firstLinkedNode, lastLinkedNode := make(intArray, MAX_NODE), make(intArray, MAX_NODE)
	idleNodeStack := stack.New[uint32](MAX_NODE)
	for i := uint32(0); i < MAX_NODE; i++ {
		idleNodeStack.Push(i)
	}
	ch := m.channels[idx]
	for item := range ch {
		atomic.AddInt64(&m.syncCounter[idx], 1)
		switch item.mode {
		case pushVal:
			upcomingIdx, ok := idleNodeStack.Pop()
			if !ok {
				panic("429!SystemTooBusy")
			}
			lastPtr := &lastLinkedNode[item.id.ArrayIndex()]
			headers[*lastPtr].nextIndex = upcomingIdx
			headers[upcomingIdx] = nodeHeader{
				backIndex: *lastPtr, nextIndex: 0,
				// DEBUG Fields, TODO: remove its after stable level
				isActive:     true,
				bucketID:     item.id,
				currentIndex: upcomingIdx,
			}
			nodes[upcomingIdx] = item.pushValue
			*lastPtr = upcomingIdx

		case pullVal:
			arrayIdx := item.id.ArrayIndex()
			idx := firstLinkedNode[arrayIdx]
			pullVals := []T{}
			for {
				if idx == 0 {
					break
				}
				h, n := &headers[idx], nodes[idx]
				pullVals = append(pullVals, n)
				h.isActive = false
				idleNodeStack.Push(idx)
				idx = h.nextIndex
			}

			item.pullValues = pullVals
			item.wg.Done()
		}
	}

}
