package goset

type hashSet[T comparable] map[T]struct{}

func newHashSet[T comparable]() hashSet[T] {
	return make(hashSet[T])
}

func (h *hashSet[T]) Size() int {
	return len(*h)
}

func (h *hashSet[T]) IsEmpty() bool {
	return len(*h) == 0
}

func (h *hashSet[T]) Contains(value T) bool {
	_, ok := (*h)[value]
	return ok
}

func (h *hashSet[T]) Add(value T) bool {
	_, ok := (*h)[value]
	if ok {
		return false
	}
	(*h)[value] = struct{}{}
	return true
}

func (h *hashSet[T]) Remove(value T) bool {
	_, ok := (*h)[value]
	if !ok {
		return false
	}
	delete(*h, value)
	return true
}

func (h *hashSet[T]) ForEach(consumer func(v T)) {
	for value, _ := range *h {
		consumer(value)
	}
}

// Copy 复制当前集合
func (h *hashSet[T]) Copy() GoSet[T] {
	set := newHashSet[T]()
	h.ForEach(func(v T) {
		set.Add(v)
	})
	return &set
}

//IsSub 判断slice是否是当前集合的子集，如果subSet为true，则只有slice为当前集合的真子集时才会返回true
func (h *hashSet[T]) IsSub(set GoSet[T], subSet bool) bool {
	is := true
	set.ForEach(func(v T) {
		if !h.Contains(v) {
			is = false
			return
		}
	})
	if is && subSet {
		is = h.Size() != set.Size()
	}
	return is
}

// Union 返回当前集合和set的并集
func (h *hashSet[T]) Union(set GoSet[T]) GoSet[T] {
	res := newHashSet[T]()
	h.ForEach(func(v T) {
		res.Add(v)
	})
	set.ForEach(func(v T) {
		res.Add(v)
	})
	return &res
}

// Intersection 返回当前集合和set的交集
func (h *hashSet[T]) Intersection(set GoSet[T]) GoSet[T] {
	res := newHashSet[T]()
	h.ForEach(func(v T) {
		if set.Contains(v) {
			res.Add(v)
		}
	})
	return &res
}

// Complement 返回当前集合在set中的补集，即把set作为全集
func (h *hashSet[T]) Complement(set GoSet[T]) GoSet[T] {
	res := newHashSet[T]()
	set.ForEach(func(v T) {
		if !h.Contains(v) {
			res.Add(v)
		}
	})
	return &res
}
