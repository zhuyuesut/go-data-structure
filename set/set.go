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

func (s Set) Add(xx ...interface{}) {
	for _, x := range xx {
		s[x] = struct{}{}
	}
}

func (s Set) Remove(xx ...interface{}) {
	for _, x := range xx {
		delete(s, x)
	}
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

func Remove(desc, src Set) {
	for k := range src {
		desc.Remove(k)
	}
}

func Intersection(a, b Set) Set {
	r := make(Set)
	if len(b) < len(a) {
		a, b = b, a
	}
	for k := range a {
		if b.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func Slice(s Set) []interface{} {
	r := make([]interface{}, 0, len(s))
	for k := range s {
		r = append(r, k)
	}
	return r
}
