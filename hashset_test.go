package goset

import (
	"testing"
)

func TestHashSet_Size(t *testing.T) {
	set := newHashSet[int]()
	Equal(set.Size(), 0, t)
	set.Add(1)
	set.Add(2)
	set.Add(3)
	Equal(set.Size(), 3, t)
}

func TestHashSet_IsEmpty(t *testing.T) {
	set := newHashSet[int]()
	Equal(set.IsEmpty(), true, t)
	set.Add(1)
	Equal(set.IsEmpty(), false, t)
}

func TestHashSet_IsNotEmpty(t *testing.T) {
	set := newHashSet[int]()
	Equal(set.IsNotEmpty(), false, t)
	set.Add(1)
	Equal(set.IsNotEmpty(), true, t)
}

func TestHashSet_Contains(t *testing.T) {
	set := newHashSet[int]()
	Equal(set.Contains(1), false, t)
	set.Add(1)
	Equal(set.Contains(1), true, t)
}

func TestHashSet_Add(t *testing.T) {
	set := newHashSet[int]()
	Equal(set.Add(1), true, t)
	Equal(set.Add(2), true, t)
	Equal(set.Add(3), true, t)
	Equal(set.Add(4), true, t)
	//重复添加
	Equal(set.Add(1), false, t)
	m := map[int]struct{}(set)
	MapContains(m, 1, t)
	MapContains(m, 2, t)
	MapContains(m, 3, t)
	MapContains(m, 4, t)

	MapNotContains(m, 5, t)
	Equal(set.Size(), 4, t)
}

func TestHashSet_Remove(t *testing.T) {
	set := newHashSet[int]()
	Equal(set.Remove(1), false, t)
	set.Add(1)
	Equal(set.Remove(1), true, t)
	Equal(set.Size(), 0, t)
}

func TestHashSet_ForEach(t *testing.T) {
	set := newHashSet[int]()
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		set.Add(v)
	}
	set.ForEach(func(value int) {
		InSlice(value, s, t)
	})
}

func BenchmarkHashSet_Add(b *testing.B) {
	set := newHashSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkHashSet_Contains(b *testing.B) {
	set := newHashSet[int]()
	for i := 0; i < 0xffffff; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i)
	}
}

func BenchmarkHashSet_Remove(b *testing.B) {
	set := newHashSet[int]()
	for i := 0; i < 0xffffff; i++ {
		set.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(i)
	}
}
