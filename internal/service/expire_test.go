package service

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestExpireHeap_Push(t *testing.T) {
	heap := NewExpireHeap(10)
	assert.Equal(t, 0, len(heap.heap))
	heap.Push("test_1", 100)
	p := heap.heap[0]
	assert.Equal(t, int64(100), p.expireAt)
	assert.Equal(t, "test_1", p.key)
}
func TestExpireHeap_Pop(t *testing.T) {
	heap := NewExpireHeap(10)
	heap.Push("test_1", 100)
	assert.Equal(t, int64(100), heap.heap[0].expireAt)
	heap.Pop()
	assert.Equal(t, 0, len(heap.heap))
}
