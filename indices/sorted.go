package index

import (
	"github.com/adverax/containers/collections"
)

type Sorted[T any] struct {
	collections.Collection[T]
}

func (that *Sorted[T]) Truncate(
	iterator func(item T) bool,
) {
	items := that.Items()
	for len(items) > 0 {
		item := items[0]
		if !iterator(item) {
			return
		}
		that.SkipHead(1)
		items = that.Items()
	}
}

func NewSorted[T any](comparator collections.Comparator[T]) *Sorted[T] {
	return &Sorted[T]{
		Collection: *collections.NewCollection(comparator, false),
	}
}
