package sets

import "sort"

// StringSet operations for string keys
type StringSet interface {
	Insert(items ...string)
	Delete(items ...string)
	Has(item string) bool
	HasAll(items ...string) bool
	HasAny(items ...string) bool
	IsSuperset(right StringSet) bool
	Equal(right StringSet) bool
	Diff(right StringSet) StringSet
	Union(right StringSet) StringSet
	Intersection(right StringSet) StringSet
	List() []string
	SortedList() []string
	PopAny() (string, bool)
	Len() int
}

// NewStringSet create a StringSet from a string slice
func NewStringSet(items ...string) StringSet {
	ss := ssetImpl{}
	if len(items) > 0 {
		ss.Insert(items...)
	}
	return ss
}

type ssetImpl map[string]struct{}

func (ss ssetImpl) Insert(items ...string) {
	for _, item := range items {
		ss[item] = struct{}{}
	}
}
func (ss ssetImpl) Delete(items ...string) {
	for _, item := range items {
		delete(ss, item)
	}
}
func (ss ssetImpl) Has(item string) bool {
	_, exists := ss[item]
	return exists
}
func (ss ssetImpl) HasAll(items ...string) bool {
	for _, item := range items {
		if !ss.Has(item) {
			return false
		}
	}
	return true
}
func (ss ssetImpl) HasAny(items ...string) bool {
	for _, item := range items {
		if ss.Has(item) {
			return true
		}
	}
	return false
}
func (ss ssetImpl) IsSuperset(right StringSet) bool {
	rightImpl := right.(ssetImpl)
	for item := range rightImpl {
		if !ss.Has(item) {
			return false
		}
	}
	return true
}
func (ss ssetImpl) Equal(right StringSet) bool {
	return len(ss) == right.Len() && ss.IsSuperset(right)
}
func (ss ssetImpl) Diff(right StringSet) StringSet {
	result := NewStringSet()
	for item := range ss {
		if !right.Has(item) {
			result.Insert(item)
		}
	}
	return result
}
func (ss ssetImpl) Union(right StringSet) StringSet {
	result := NewStringSet()
	for item := range ss {
		result.Insert(item)
	}
	rightImpl := right.(ssetImpl)
	for item := range rightImpl {
		result.Insert(item)
	}
	return result
}
func (ss ssetImpl) Intersection(right StringSet) StringSet {
	s1 := ss
	s2 := right.(ssetImpl)
	if len(s1) > s2.Len() {
		s1, s2 = s2, s1
	}
	result := NewStringSet()
	for item := range s1 {
		if s2.Has(item) {
			result.Insert(item)
		}
	}
	return result
}
func (ss ssetImpl) List() []string {
	result := make([]string, 0, len(ss))
	for item := range ss {
		result = append(result, item)
	}
	return result
}
func (ss ssetImpl) SortedList() []string {
	list := ss.List()
	sort.Strings(list)
	return list
}
func (ss ssetImpl) PopAny() (string, bool) {
	for item := range ss {
		ss.Delete(item)
		return item, true
	}
	return "", false
}
func (ss ssetImpl) Len() int {
	return len(ss)
}
