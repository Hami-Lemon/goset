package goset

type GoSet[T comparable] interface {
	// Size 获取集合中的元素个数
	Size() int
	// IsEmpty 获取集合是否为空
	IsEmpty() bool
	// IsNotEmpty 获取集合是否为非空
	IsNotEmpty() bool
	// Contains 判断集合中是否包含某元素
	Contains(value T) bool
	// Add 向集合中添加元素，如果该元素已经存在，返回false
	Add(value T) bool
	// Remove 从集合中移除元素，如果不存在该元素，返回false
	Remove(value T) bool
	// ForEach 遍历集合中的所有元素
	ForEach(consumer func(v T))
}

func New[T comparable](value ...T) GoSet[T] {
	s := newHashSet[T]()
	for _, v := range value {
		s.Add(v)
	}
	return &s
}

func NewConcurrent[T comparable](value ...T) GoSet[T] {
	s := newConcurrentSet[T]()
	for _, v := range value {
		s.Add(v)
	}
	return &s
}
