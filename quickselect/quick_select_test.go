package quickselect

import (
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	ns := []int{10000, 100000, 500000, 600000, 800000, 1000000, 5000000, 10000000}
	for _, n := range ns {
		items := make([]float64, 0, n)
		for i := 0; i < n; i++ {
			items = append(items, float64(rand.Int()%n))
		}

		randFind := rand.Int()%n + 1
		v := QuickSelect(Float64Slice(items), randFind).(float64)
		v2 := Float64s(items, randFind)

		sort.Float64s(items)
		t.Log(v, v2, items[randFind-1])
		if v != items[randFind-1] || v != v2 {
			t.Fatalf("not equal results: %v, %v, %v", items[randFind-1], v, v2)
		}
	}
}
func BenchmarkQuickSelect(b *testing.B) {
	for n := 1; n < b.N; n++ {
		items := make([]float64, 0, n)
		for i := 0; i < n; i++ {
			items = append(items, float64(rand.Int()%n))
		}
		QuickSelect(Float64Slice(items), rand.Int()%n+1)
	}
}
func BenchmarkQuickSelectFloat64(b *testing.B) {
	for n := 1; n < b.N; n++ {
		items := make([]float64, 0, n)
		for i := 0; i < n; i++ {
			items = append(items, float64(rand.Int()%n))
		}
		Float64s(items, rand.Int()%n+1)
		Float64s(items, rand.Int()%n+1)
		Float64s(items, rand.Int()%n+1)
		Float64s(items, rand.Int()%n+1)
		Float64s(items, rand.Int()%n+1)
		Float64s(items, rand.Int()%n+1)
		Float64s(items, rand.Int()%n+1)
	}
}
func BenchmarkQuickSortSelect(b *testing.B) {
	for n := 1; n < b.N; n++ {
		items := make([]float64, 0, n)
		for i := 0; i < n; i++ {
			items = append(items, float64(rand.Int()%n))
		}
		sort.Float64s(items)
	}
}
