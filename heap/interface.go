package heap

type RemovableHeap interface {
	Len() int
	PopOne() interface{}
	Pop() (interface{}, int)
	RemoveOne(interface{})
	Remove(interface{}) int
	PushOne(interface{})
	Push(interface{}, int)
	Front() interface{}
}

type Heap interface {
	Len() int
	Push(interface{})
	Pop() interface{}
	Front() interface{}
}
