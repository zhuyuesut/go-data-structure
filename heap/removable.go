package heap

import std "container/heap"

func NewRemovable(less func(interface{}, interface{}) bool) RemovableHeap {
	return &removable{withIdx: newWithIdx(less), cnt: make(map[interface{}]int)}
}

type removable struct {
	*withIdx
	sum int
	cnt map[interface{}]int
}

func (h *removable) Len() int {
	return h.sum
}

func (h *removable) Front() interface{} {
	return h.withIdx.data[0]
}

func (h *removable) PopOne() interface{} {
	f := h.withIdx.data[0]
	h.cnt[f]--
	h.sum--
	if h.cnt[f] == 0 {
		delete(h.cnt, f)
		std.Pop(h.withIdx)
	}
	return f
}

func (h *removable) Pop() (interface{}, int) {
	f := h.withIdx.data[0]
	cnt := h.cnt[f]
	delete(h.cnt, f)
	std.Pop(h.withIdx)
	h.sum -= cnt
	return f, cnt
}

func (h *removable) RemoveOne(x interface{}) {
	if h.cnt[x] > 0 {
		h.cnt[x]--
		h.sum--
		if h.cnt[x] == 0 {
			delete(h.cnt, x)
			idx := h.withIdx.idx[x]
			std.Remove(h.withIdx, idx)
		}
	}
}

func (h *removable) Remove(x interface{}) int {
	cnt := h.cnt[x]
	if cnt > 0 {
		h.sum -= cnt
		delete(h.cnt, x)
		idx := h.withIdx.idx[x]
		std.Remove(h.withIdx, idx)
	}
	return cnt
}

func (h *removable) PushOne(x interface{}) {
	h.sum++

	if h.cnt[x] > 0 {
		h.cnt[x]++
		return
	}

	std.Push(h.withIdx, x)
	h.cnt[x]++
}

func (h *removable) Push(x interface{}, cnt int) {
	if cnt <= 0 {
		return
	}
	h.sum += cnt
	if h.cnt[x] > 0 {
		h.cnt[x] += cnt
		return
	}
	std.Push(h.withIdx, x)
	h.cnt[x] += cnt
}

/*
The following is the standard library heap operation agent
*/

func newWithIdx(less func(interface{}, interface{}) bool) *withIdx {
	return &withIdx{less: less, idx: make(map[interface{}]int)}
}

type withIdx struct {
	data []interface{}
	idx  map[interface{}]int
	less func(interface{}, interface{}) bool
}

func (h *withIdx) Len() int {
	return len(h.data)
}

func (h *withIdx) Push(x interface{}) {
	h.data = append(h.data, x)
	h.idx[x] = len(h.data) - 1
}

func (h *withIdx) Pop() interface{} {
	t := (h.data)[len(h.data)-1]
	h.data = (h.data)[:len(h.data)-1]
	delete(h.idx, t)
	return t
}

func (h *withIdx) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *withIdx) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.idx[h.data[i]] = i
	h.idx[h.data[j]] = j
}
