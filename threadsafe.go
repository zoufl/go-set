package goset

import "sync"

type threadSafeSet struct {
	set threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() *threadSafeSet {
	return &threadSafeSet{
		set: newThreadUnsafeSet(),
	}
}

func (s *threadSafeSet) Add(i interface{}) bool {
	s.Lock()
	ret := s.set.Add(i)
	s.Unlock()

	return ret
}

func (s *threadSafeSet) Pop() interface{} {
	s.RLock()
	defer s.Unlock()

	return s.set.Pop()
}

func (s *threadSafeSet) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.set.Size()
}

func (s *threadSafeSet) Clear() {
	s.Lock()
	s.set = newThreadUnsafeSet()
	s.Unlock()
}

func (s *threadSafeSet) Copy() Set {
	s.RLock()
	copiedSet := s.set.Copy()
	set := &threadSafeSet{set: copiedSet.(threadUnsafeSet)}
	s.RUnlock()

	return set
}

func (s *threadSafeSet) Contains(elems ...interface{}) bool {
	s.RLock()
	ret := s.set.Contains(elems...)
	s.RUnlock()

	return ret
}

func (s *threadSafeSet) SymmetricDiff(other Set) Set {
	o := other.(*threadSafeSet)

	s.Lock()
	o.Lock()
	diffSet := s.set.SymmetricDiff(o.set).(threadUnsafeSet)
	set := &threadSafeSet{set: diffSet}
	s.Unlock()
	o.Unlock()

	return set
}

func (s *threadSafeSet) Diff(other Set) Set {
	o := other.(*threadSafeSet)

	s.Lock()
	o.Lock()
	diffSet := s.set.Diff(o.set).(threadUnsafeSet)
	set := &threadSafeSet{set: diffSet}
	s.Unlock()
	o.Unlock()

	return set
}

func (s *threadSafeSet) Union(other Set) Set {
	o := other.(*threadSafeSet)

	s.Lock()
	o.Lock()
	unionSet := s.set.Union(o.set).(threadUnsafeSet)
	set := &threadSafeSet{set: unionSet}
	s.Unlock()
	o.Unlock()

	return set
}

func (s *threadSafeSet) Intersect(other Set) Set {
	o := other.(*threadSafeSet)

	s.Lock()
	o.Lock()
	intersectSet := s.set.Intersect(o.set).(threadUnsafeSet)
	set := &threadSafeSet{set: intersectSet}
	s.Unlock()
	o.Unlock()

	return set
}

func (s *threadSafeSet) IsSubset(other Set) bool {
	o := other.(*threadSafeSet)

	s.RLock()
	o.RLock()
	ret := s.set.IsSubset(o.set)
	s.RUnlock()
	o.RUnlock()

	return ret
}

func (s *threadSafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(s)
}

func (s *threadSafeSet) Equal(other Set) bool {
	o := other.(*threadSafeSet)

	s.RLock()
	o.RLock()

	ret := s.set.Equal(o.set)

	s.RUnlock()
	o.RUnlock()

	return ret
}

func (s *threadSafeSet) Remove(elems ...interface{}) {
	s.Lock()
	s.set.Remove(elems...)
	s.Unlock()
}

func (s *threadSafeSet) String() string {
	s.RLock()
	ret := s.set.String()
	s.RUnlock()

	return ret
}

func (s *threadSafeSet) Range(f func(key, value interface{}) bool) {
	s.RLock()
	defer s.RUnlock()

	s.set.Range(f)
}

func (s *threadSafeSet) ToSlice() []interface{} {
	s.RLock()
	ret := s.set.ToSlice()
	s.RUnlock()

	return ret
}
