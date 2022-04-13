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

func (h *hashSet[T]) IsNotEmpty() bool {
	return len(*h) != 0
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
