package quickselect

import (
	"math/rand"
	"sort"
)

// Interface for quick select alg
type Interface interface {
	sort.Interface
	Value(index int) interface{}
	SubSlice(start, end int) Interface
}

// IntSlice Interface wrapper for int slice
type IntSlice []int

// Len of int slice
func (items IntSlice) Len() int { return len(items) }

// Less sorts int slice in increasing order
func (items IntSlice) Less(i, j int) bool { return items[i] <= items[j] }

// Swap two elems
func (items IntSlice) Swap(i, j int) { items[i], items[j] = items[j], items[i] }

// Value of the items[index]
func (items IntSlice) Value(index int) interface{} { return items[index] }

// SubSlice of the origin slice
func (items IntSlice) SubSlice(i, j int) Interface { return items[i:j] }

// Float64Slice Interface wrapper for float64 slice
type Float64Slice []float64

// Len of Float64Slice slice
func (items Float64Slice) Len() int { return len(items) }

// Less sorts Float64Slice slice in increasing order
func (items Float64Slice) Less(i, j int) bool { return items[i] <= items[j] }

// Swap two elems
func (items Float64Slice) Swap(i, j int) { items[i], items[j] = items[j], items[i] }

// Value of the items[index]
func (items Float64Slice) Value(index int) interface{} { return items[index] }

// SubSlice of the origin slice
func (items Float64Slice) SubSlice(i, j int) Interface { return items[i:j] }

// QuickSelect for ints
func QuickSelect(items Interface, target int) interface{} {
	if items.Len() == 1 {
		return items.Value(0)
	} else if items.Len() < 512 {
		sort.Sort(items)
		return items.Value(target - 1)
	}

	randIdx := rand.Int() % items.Len()
	items.Swap(randIdx, items.Len()-1)
	//items[randIdx], items[len(items)-1] = items[len(items)-1], items[randIdx]

	i, j := 0, items.Len()-2
	for i <= j {
		if items.Less(i, items.Len()-1) {
			//if items[i] <= items[len(items)-1] {
			i++
			continue
		}

		if !items.Less(j, items.Len()-1) {
			//if items[j] > items[len(items)-1] {
			j--
			continue
		}
		items.Swap(i, j)
	}
	items.Swap(i, items.Len()-1)

	if target == i+1 {
		return items.Value(i)
	} else if target < i+1 {
		return QuickSelect(items.SubSlice(0, i), target)
	} else {
		return QuickSelect(items.SubSlice(i+1, items.Len()), target-i-1)
	}
}

// Float64s for float64 slice
func Float64s(items []float64, target int) interface{} {
	if len(items) == 1 {
		return items[0]
	} else if len(items) < 512 {
		sort.Float64s(items)
		return items[target-1]
	}

	pivot := rand.Int() % len(items)
	items[pivot], items[len(items)-1] = items[len(items)-1], items[pivot]

	i, j := 0, len(items)-2
	for i <= j {
		if items[i] <= items[len(items)-1] {
			i++
			continue
		}

		if items[j] > items[len(items)-1] {
			j--
			continue
		}
		items[i], items[j] = items[j], items[i]
	}
	items[i], items[len(items)-1] = items[len(items)-1], items[i]

	if target == i+1 {
		return items[i]
	} else if target < i+1 {
		return Float64s(items[:i], target)
	} else {
		return Float64s(items[i+1:], target-i-1)
	}
}
