package queue

type Queue interface {
	Len() int
	L() interface{}
	R() interface{}
	LPop() interface{}
	RPop() interface{}
	LPush(v interface{})
	RPush(v interface{})
}
