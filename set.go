package goset

type Set interface {
	Add(i interface{}) bool

	Pop() interface{}

	Size() int

	Clear()

	Copy() Set

	Contains(elems ...interface{}) bool

	SymmetricDiff(other Set) Set

	Diff(s Set) Set

	Union(s Set) Set

	Intersect(s Set) Set

	IsSubset(other Set) bool

	IsSuperset(other Set) bool

	Equal(s Set) bool

	Remove(elems ...interface{})

	String() string

	Range(f func(key, value interface{}) bool)

	ToSlice() []interface{}
}

func NewSet(s ...interface{}) Set {
	set := newThreadUnsafeSet()

	for _, v := range s {
		set.Add(v)
	}

	return set
}

func NewSetWith(elems ...interface{}) Set {
	return NewSetFromSlice(elems)
}

func NewSetFromSlice(s []interface{}) Set {
	return NewSet(s...)
}

func NewThreadSafeSet() Set {
	return newThreadSafeSet()
}

func NewThreadSetWith(elems ...interface{}) Set {
	return NewThreadSafeSetFromSlice(elems)
}

func NewThreadSafeSetFromSlice(s []interface{}) Set {
	set := newThreadSafeSet()

	for _, v := range s {
		set.Add(v)
	}

	return set
}
