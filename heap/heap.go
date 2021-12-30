package heap

import std "container/heap"

func New(less func(interface{}, interface{}) bool) Heap {
	return &heap{newAgent(less)}
}

type heap struct {
	*agent
}

func (h *heap) Push(x interface{}) {
	std.Push(h.agent, x)
}

func (h *heap) Pop() interface{} {
	return std.Pop(h.agent)
}

func (h *heap) Front() interface{} {
	return h.agent.data[0]
}

/*
The following is the standard library heap operation agent
*/

type agent struct {
	data []interface{}
	less func(interface{}, interface{}) bool
}

func newAgent(less func(interface{}, interface{}) bool) *agent {
	return &agent{less: less}
}

func (h *agent) Len() int {
	return len(h.data)
}

func (h *agent) Push(x interface{}) {
	h.data = append(h.data, x)
}

func (h *agent) Pop() interface{} {
	t := (h.data)[len(h.data)-1]
	h.data = (h.data)[:len(h.data)-1]
	return t
}

func (h *agent) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *agent) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
