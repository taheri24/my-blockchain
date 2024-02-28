package msgnet

import (
	"testing"
	"time"
)

func TestSimplePush(t *testing.T) {
	mesh := New[string]()
	bucketID, _ := NewBucketID("200")
	mesh.Push(bucketID, "hello world")
	mesh.Sync(time.Millisecond)
	values := mesh.Pull(bucketID)
	t.Log(values)
}
