package generic

import (
	"github.com/adverax/containers/lists"
	"sort"
)

type FLOAT interface {
	~float64 | ~float32
}

type INTEGER interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type ORDERED interface {
	INTEGER | FLOAT | ~string
}

type Set[T ORDERED] map[T]struct{}

func NewSet[T ORDERED](values ...T) Set[T] {
	set := make(map[T]struct{})
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func (set Set[T]) Append(values ...T) {
	for _, value := range values {
		set[value] = struct{}{}
	}
}

func (set Set[T]) Len() int {
	return len(set)
}

func (set Set[T]) Add(value T) {
	set[value] = struct{}{}
}

func (set Set[T]) Remove(value T) {
	delete(set, value)
}

func (set Set[T]) Contains(value T) bool {
	_, ok := set[value]
	return ok
}

func (set Set[T]) Values() lists.List[T] {
	values := make(lists.List[T], 0, len(set))
	for value := range set {
		values = append(values, value)
	}
	sort.Sort(values)
	return values
}

func Union[T ORDERED](lists ...[]T) lists.List[T] {
	set := make(Set[T])
	for _, list := range lists {
		set.Append(list...)
	}
	return set.Values()
}
