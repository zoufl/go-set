package goset

import (
	"fmt"
	"strings"
)

type threadUnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (s threadUnsafeSet) Add(i interface{}) bool {
	if _, ok := s[i]; ok {
		return false
	}

	s[i] = struct{}{}

	return true
}

func (s threadUnsafeSet) Pop() interface{} {
	for v := range s {
		delete(s, v)

		return v
	}

	return nil
}

func (s threadUnsafeSet) Size() int {
	return len(s)
}

func (s threadUnsafeSet) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s threadUnsafeSet) Copy() Set {
	copiedSet := newThreadUnsafeSet()

	for k := range s {
		copiedSet.Add(k)
	}

	return copiedSet
}

func (s threadUnsafeSet) Contains(elems ...interface{}) bool {
	for _, v := range elems {
		if _, ok := s[v]; !ok {
			return false
		}
	}

	return true
}

func (s threadUnsafeSet) SymmetricDiff(other Set) Set {
	aDiff := s.Diff(other)
	bDiff := other.Diff(s)

	return aDiff.Union(bDiff)
}

func (s threadUnsafeSet) Diff(other Set) Set {
	set := newThreadUnsafeSet()

	for k := range s {
		if !other.Contains(k) {
			set.Add(k)
		}
	}

	return s
}

func (s threadUnsafeSet) Union(other Set) Set {
	set := other.Copy()

	for k := range s {
		set.Add(k)
	}

	return set
}

func (s threadUnsafeSet) Intersect(other Set) Set {
	set := newThreadUnsafeSet()

	var s1 = s
	var s2 Set = other
	if other.Size() < s.Size() {
		s1 = other.(threadUnsafeSet)
		s2 = s
	}

	for k := range s1 {
		if s2.Contains(k) {
			set[k] = struct{}{}
		}
	}

	return set
}

func (s threadUnsafeSet) IsSubset(other Set) bool {
	if other.Size() > s.Size() {
		return false
	}

	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}

func (s threadUnsafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(s)
}

func (s threadUnsafeSet) Equal(other Set) bool {
	if s.Size() != other.Size() {
		return false
	}

	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}

	return true
}

func (s threadUnsafeSet) Remove(elems ...interface{}) {
	for _, v := range elems {
		delete(s, v)
	}
}

func (s threadUnsafeSet) String() string {
	strs := make([]string, 0, len(s))

	for k := range s {
		strs = append(strs, fmt.Sprintf("%v", k))
	}

	return "Set{" + strings.Join(strs, ", ") + "}"
}

func (s threadUnsafeSet) Range(f func(key, value interface{}) bool) {
	for k, v := range s {
		if !f(k, v) {
			break
		}
	}
}

func (s threadUnsafeSet) ToSlice() []interface{} {
	slice := make([]interface{}, 0, len(s))

	for k := range s {
		slice = append(slice, k)
	}

	return slice
}
