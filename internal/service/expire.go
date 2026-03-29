package service

type expiringKey struct {
	key      string
	expireAt int64
}

type ExpireHeap struct {
	heap []expiringKey
}

func NewExpireHeap(baseCapacity int) *ExpireHeap {
	return &ExpireHeap{
		heap: make([]expiringKey, 0, baseCapacity),
	}
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}

func (h *ExpireHeap) swap(i int, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}
func (h *ExpireHeap) siftUp(idx int) {
	for idx > 0 {
		p := parent(idx)
		if h.heap[p].expireAt <= h.heap[idx].expireAt {
			break
		}
		h.swap(idx, p)
		idx = p
	}
}

func (h *ExpireHeap) siftDown(idx int) {
	for {
		smallest := idx
		l := left(idx)
		r := right(idx)
		if l < len(h.heap) && h.heap[l].expireAt < h.heap[smallest].expireAt {
			smallest = l
		}
		for r < len(h.heap) && h.heap[l].expireAt < h.heap[smallest].expireAt {
			smallest = r
		}
		if smallest == idx {
			break
		}
		h.swap(idx, smallest)
		idx = smallest
	}
}

func (h *ExpireHeap) Push(key string, expireAt int64) {
	h.heap = append(h.heap, expiringKey{key, expireAt})
	h.siftUp(len(h.heap) - 1)
}

func (h *ExpireHeap) Pop() (string, bool) {
	if len(h.heap) == 0 {
		return "", false
	}
	if len(h.heap) == 1 {
		p := h.heap[0].key
		h.heap = h.heap[:0]
		return p, true
	}
	p := h.heap[0].key
	h.heap[len(h.heap)-1], h.heap[0] = h.heap[0], h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0)
	return p, true
}
