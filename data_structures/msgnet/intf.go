package msgnet

import (
	"strconv"
	"strings"
	"time"
)

type BucketID [2]uint64

func NewBucketID(id string) (BucketID, error) {
	segments := strings.SplitN(id, ":", 2)
	var err error
	var bid BucketID
	bid[0], err = strconv.ParseUint(segments[0], 10, 64)
	if err != nil {
		return bid, err
	}

	if len(segments) == 2 {

		bid[1], err = strconv.ParseUint(segments[1], 10, 64)
	}
	return bid, err
}

func (b BucketID) ArrayIndex() uint64 {
	if arrayIndex := b[1]; arrayIndex != 0 {
		return arrayIndex
	}
	return b[0]
}

func (b BucketID) PrivateIndex() uint64 {
	return b[0]
}

type MsgNet[TData any] interface {
	Pull(BucketID) []TData
	Push(BucketID, TData)
	Sync(time.Duration) *time.Ticker
}
