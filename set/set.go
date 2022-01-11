package set

import "fmt"

type Set map[interface{}]struct{}

func NewSet(xx ...interface{}) Set {
	s := make(Set)
	s.Add(xx...)
	return s
}

func (s Set) Contains(xx ...interface{}) bool {
	for _, x := range xx {
		_, found := s[x]
		if !found {
			return false
		}
	}
	return true
}

func (s Set) Pop() interface{} {
	for k := range s {
		delete(s, k)
		return k
	}
	panic(fmt.Errorf("no item in set"))
}

func (s Set) Add(xx ...interface{}) int {
	var c int
	for _, x := range xx {
		_, found := s[x]
		if !found {
			c++
			s[x] = struct{}{}
		}
	}
	return c
}

func (s Set) Remove(xx ...interface{}) int {
	var c int
	for _, x := range xx {
		_, found := s[x]
		if found {
			c++
			delete(s, x)
		}
	}
	return c
}

func (s Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func Equal(a, b Set) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b.Contains(k) {
			return false
		}
	}
	return true
}

func Copy(desc, src Set) {
	for k := range src {
		desc.Add(k)
	}
}

func RemoveIf(s Set, f func(x ...interface{}) bool) int {
	var c int
	for k := range s {
		if f(k) {
			c++
			delete(s, k)
		}
	}
	return c
}

func Slice(s Set) []interface{} {
	r := make([]interface{}, 0, len(s))
	for k := range s {
		r = append(r, k)
	}
	return r
}
