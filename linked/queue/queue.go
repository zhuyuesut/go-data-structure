package queue

import "container/list"

type queue struct {
	data *list.List
}

func New() *queue {
	return &queue{
		data: list.New(),
	}
}

func (q *queue) Len() int {
	return q.data.Len()
}

func (q *queue) L() interface{} {
	return q.data.Front().Value
}

func (q *queue) R() interface{} {
	return q.data.Back().Value
}

func (q *queue) LPop() interface{} {
	e := q.data.Front()
	q.data.Remove(e)
	return e.Value
}

func (q *queue) RPop() interface{} {
	e := q.data.Back()
	q.data.Remove(e)
	return e.Value
}

func (q *queue) RPush(v interface{}) {
	q.data.PushBack(v)
}

func (q *queue) LPush(v interface{}) {
	q.data.PushFront(v)
}
