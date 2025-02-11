package index

import "github.com/adverax/containers/collections"

type FIFO[T any] struct {
	items []T
}

func (that *FIFO[T]) Reset() {
	that.items = nil
}

func (that *FIFO[T]) PushMultiple(item ...T) error {
	that.items = append(that.items, item...)
	return nil
}

func (that *FIFO[T]) Push(item T) {
	that.items = append(that.items, item)
}

func (that *FIFO[T]) Pop() (item T, err error) {
	if len(that.items) == 0 {
		return item, collections.ErrNoMatch
	}

	item = that.items[0]
	that.items = that.items[1:]

	return item, nil
}

func (that *FIFO[T]) Truncate(
	iterator func(item T) bool,
) {
	for len(that.items) > 0 {
		item := that.items[0]
		if !iterator(item) {
			return
		}
		that.items = that.items[1:]
	}
}

func NewFIFO[T any]() *FIFO[T] {
	return &FIFO[T]{}
}
