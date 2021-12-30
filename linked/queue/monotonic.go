package queue

type monotonic struct {
	Queue
	less func(interface{}, interface{}) bool
}

func NewMonotonic(less func(interface{}, interface{}) bool) Queue {
	return &monotonic{
		Queue: New(),
		less:  less,
	}
}

func (m *monotonic) RPush(v interface{}) {
	for m.Len() > 0 {
		if m.less(m.R(), v) {
			break
		}
		m.RPop()
	}
	m.Queue.RPush(v)
}

func (m *monotonic) LPush(v interface{}) {
	for m.Len() > 0 {
		if m.less(v, m.L()) {
			break
		}
		m.LPop()
	}
	m.Queue.LPush(v)
}
