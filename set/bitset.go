package set

import (
	"encoding/base64"
	"fmt"
)

type BitSet struct {
	data []byte
	l    int
}

func NewBitSetFromString(s string) *BitSet {
	var bs BitSet
	var t string
	_, err := fmt.Sscanf(s, "%d %s", &bs.l, &t)
	if err != nil {
		panic(err)
	}
	bs.data, err = base64.StdEncoding.DecodeString(t)
	if err != nil {
		panic(err)
	}
	return &bs
}

func NewBitSet(xx ...int) *BitSet {
	var bs BitSet
	bs.Add(xx...)
	return &bs
}

func (bs BitSet) String() string {
	bs.Compression()
	return fmt.Sprintf("%d %s", bs.l, base64.StdEncoding.EncodeToString(bs.data))
}

func (bs BitSet) Contains(xx ...int) bool {
	for _, x := range xx {
		index, mark := x/8, byte(1<<(7-x%8))
		if index >= len(bs.data) {
			return false
		}
		if bs.data[index]&mark == 0 {
			return false
		}
	}
	return true
}

func (bs *BitSet) Compression() {
	for len(bs.data) > 0 {
		last := bs.data[len(bs.data)-1]
		if last != 0 {
			break
		}
		bs.data = bs.data[:len(bs.data)-1]
	}
}

func (bs *BitSet) Pop() int {
	bs.Compression()
	for {
		last := bs.data[len(bs.data)-1]
		for i := 0; i < 8; i++ {
			mark := byte(1 << (7 - i))
			if last&mark != 0 {
				bs.data[len(bs.data)-1] ^= mark
				bs.l--
				return (len(bs.data)-1)*8 + i
			}
		}
	}
}

func (bs *BitSet) Add(xx ...int) int {
	var c int
	for _, x := range xx {
		index, mark := x/8, byte(1<<(7-x%8))
		for index >= len(bs.data) {
			bs.data = append(bs.data, 0)
		}
		if bs.data[index]&mark == 0 {
			c++
			bs.l++
			bs.data[index] ^= mark
		}
	}
	return c
}

func (bs BitSet) Remove(xx ...int) int {
	var c int
	for _, x := range xx {
		index, mark := x/8, byte(1<<(7-x%8))
		if index >= len(bs.data) {
			continue
		}
		if bs.data[index]&mark != 0 {
			c++
			bs.l--
			bs.data[index] ^= mark
		}
	}
	return c
}

func (bs *BitSet) Clear() {
	bs.data = nil
	bs.l = 0
}

func (bs BitSet) Len() int {
	return bs.l
}

func (bs BitSet) Each(f func(i int) bool) {
	if bs.l == 0 {
		return
	}
	for index, x := range bs.data {
		if x != 0 {
			for i := 0; i < 8; i++ {
				mark := byte(1 << (7 - i))
				if bs.data[index]&mark != 0 {
					if !f(index*8 + i) {
						return
					}
				}
			}
		}
	}
}

func EqualBitSet(a, b BitSet) bool {
	if a.Len() != b.Len() {
		return false
	}
	if len(a.data) > len(b.data) {
		a, b = b, a
	}
	for i, x := range a.data {
		if b.data[i] != x {
			return false
		}
	}
	return true
}

func CopyBitSet(desc, src *BitSet) {
	src.Each(func(i int) bool {
		desc.Add(i)
		return true
	})
}

func SliceBit(bs BitSet) []int {
	r := make([]int, 0, bs.Len())
	bs.Each(func(i int) bool {
		r = append(r, i)
		return true
	})
	return r
}
