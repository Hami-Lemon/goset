package goset

import (
	"testing"
)

func makeSet(num int) GoSet[int] {
	set := newHashSet[int]()
	for i := 0; i < num; i++ {
		set.Add(i)
	}
	return &set
}

func setEqual[T comparable](set1 GoSet[T], set2 GoSet[T], t *testing.T) {
	Equal(set1.Size(), set2.Size(), t)
	set1.ForEach(func(v T) {
		Equal(set1.Contains(v), true, t)
	})
}

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
	Equal(set.Add(2), false, t)
	Equal(set.Add(3), false, t)
	Equal(set.Add(4), false, t)
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

func TestHashSet_Copy(t *testing.T) {
	set := makeSet(1)
	setEqual(set, set.Copy(), t)
	set = makeSet(0)
	setEqual(set, set.Copy(), t)
	set = makeSet(10)
	setEqual(set, set.Copy(), t)
	set = makeSet(100)
	setEqual(set, set.Copy(), t)
}

func TestHashSet_IsSub(t *testing.T) {
	set := makeSet(20)
	sub := makeSet(10)
	subSet := makeSet(20)
	noSub := newHashSet[int]()
	for i := 0; i < 22; i++ {
		noSub.Add(i + 10)
	}
	Equal(set.IsSub(sub, false), true, t)
	Equal(set.IsSub(sub, true), true, t)
	Equal(set.IsSub(subSet, false), true, t)
	//需要真子集
	Equal(set.IsSub(subSet, true), false, t)
	Equal(set.IsSub(&noSub, false), false, t)
	Equal(set.IsSub(&noSub, true), false, t)

}

func TestHashSet_Union(t *testing.T) {
	set := makeSet(20)
	set1 := makeSet(10)
	set2 := newHashSet[int]()
	for i := 0; i < 10; i++ {
		set2.Add(10 + i)
	}
	setEqual(set, set1.Union(&set2), t)
}

func TestHashSet_Intersection(t *testing.T) {
	set1 := makeSet(10)
	set2 := makeSet(20)
	setEqual(set1.Intersection(set2), set1, t)
}

func TestHashSet_Complement(t *testing.T) {
	set := makeSet(20)
	set1 := makeSet(10)
	com := newHashSet[int]()
	for i := 0; i < 10; i++ {
		com.Add(10 + i)
	}
	setEqual[int](set1.Complement(set), &com, t)
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
