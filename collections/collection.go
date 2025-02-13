package collections

import (
	"errors"
	"sort"
)

type Comparator[T any] interface {
	Less(a, b T) bool
	Equal(a, b T) bool
	Greater(a, b T) bool
}

// Collection - sorted collection of items
type Collection[T any] struct {
	items      []T
	comparator Comparator[T]
	unique     bool
	sorted     bool
}

func (that *Collection[T]) Len() int {
	return len(that.items)
}

func (that *Collection[T]) Less(i, j int) bool {
	return that.comparator.Less(that.items[i], that.items[j])
}

func (that *Collection[T]) Swap(i, j int) {
	that.items[i], that.items[j] = that.items[j], that.items[i]
}

// Reset - clear collection
func (that *Collection[T]) Reset() {
	that.items = nil
}

// Clone returns a copy of the collection
func (that *Collection[T]) Clone() *Collection[T] {
	res := &Collection[T]{
		items:      make([]T, len(that.items)),
		comparator: that.comparator,
		unique:     that.unique,
	}
	copy(res.items, that.items)
	return res
}

// Contains checks if item is in collection
func (that *Collection[T]) Contains(item T) bool {
	l := len(that.items)
	if l == 0 {
		return false
	}

	i := that.search(item)
	if i == l {
		return false
	}

	return that.comparator.Equal(that.items[i], item)
}

// Include allows to append item to collection
func (that *Collection[T]) Include(item T) bool {
	l := len(that.items)
	if l == 0 {
		that.items = []T{item}
		return true
	}

	if !that.sorted {
		that.items = append(that.items, item)
		return true
	}

	i := that.search(item)
	if i == l {
		that.items = append(that.items, item)
		return true
	}
	if that.unique {
		if that.comparator.Equal(that.items[i], item) {
			return false
		}
	}

	that.items = append(that.items, item)
	copy(that.items[i+1:], that.items[i:])
	that.items[i] = item
	return true
}

// Exclude allows to delete item from collection
func (that *Collection[T]) Exclude(item T) bool {
	l := len(that.items)
	if l == 0 {
		return false
	}

	i := that.search(item)
	if i == l {
		return false
	}

	if !that.comparator.Equal(that.items[i], item) {
		return false
	}

	if l == 1 {
		that.items = make([]T, 0)
		return true
	}

	that.items = append(that.items[:i], that.items[i+1:]...)
	return true
}

// Add - join two collections
func (that *Collection[T]) Add(bs *Collection[T], unique bool) *Collection[T] {
	if !that.sorted {
		return nil
	}

	la := len(that.items)
	lb := len(bs.items)
	if la == 0 {
		return bs
	}
	if lb == 0 {
		return that
	}

	a := 0
	b := 0
	comparator := that.comparator
	c := &Collection[T]{
		items:      make([]T, 0, la+lb),
		comparator: comparator,
		unique:     unique,
	}
	for a < la && b < lb {
		if comparator.Less(that.items[a], bs.items[b]) {
			c.items = append(c.items, that.items[a])
			a++
			continue
		}
		if comparator.Greater(that.items[a], bs.items[b]) {
			c.items = append(c.items, bs.items[b])
			b++
			continue
		}
		c.items = append(c.items, that.items[a])
		if !unique {
			c.items = append(c.items, bs.items[b])
		}
		a++
		b++
	}
	if a < la {
		c.items = append(c.items, that.items[a:]...)
	}
	if b < lb {
		c.items = append(c.items, bs.items[b:]...)
	}
	return c
}

// Sub - subtract one collection from another
func (that *Collection[T]) Sub(bs *Collection[T]) *Collection[T] {
	if !that.sorted {
		return nil
	}

	la := len(that.items)
	lb := len(bs.items)
	if la == 0 {
		return nil
	}
	if lb == 0 {
		return that
	}

	a := 0
	b := 0
	comparator := that.comparator
	c := &Collection[T]{
		items:      make([]T, 0, la),
		comparator: comparator,
	}
	for a < la && b < lb {
		if comparator.Less(that.items[a], bs.items[b]) {
			c.items = append(c.items, that.items[a])
			a++
			continue
		}
		if comparator.Greater(that.items[a], bs.items[b]) {
			b++
			continue
		}
		a++
		b++
	}
	if a < la {
		c.items = append(c.items, that.items[a:]...)
	}
	return c
}

// IndexOf - returns founded index of item
func (that *Collection[T]) IndexOf(item T) int {
	l := len(that.items)
	if l == 0 {
		return -1
	}

	i := that.search(item)
	if i == l {
		return -1
	}

	if that.comparator.Equal(that.items[i], item) {
		return i
	}

	return -1
}

func (that *Collection[T]) search(item T) int {
	l := len(that.items)

	if that.sorted {
		return sort.Search(l, func(i int) bool { return !that.comparator.Less(that.items[i], item) })
	}

	if that.comparator == nil {
		return l
	}

	for i, v := range that.items {
		if that.comparator.Equal(v, item) {
			return i
		}
	}

	return l
}

func (that *Collection[T]) Items() []T {
	return that.items
}

// Push is alias for include
func (that *Collection[T]) Push(item T) {
	that.Include(item)
}

// Pop returns first item of collection
func (that *Collection[T]) Pop() (item T, err error) {
	if len(that.items) == 0 {
		return item, ErrNoMatch
	}

	item = that.items[0]
	that.items = that.items[1:]

	return item, nil
}

// SkipHead allow skip n first items
func (that *Collection[T]) SkipHead(n int) {
	if n > len(that.items) {
		n = len(that.items)
	}
	that.items = that.items[n:]
}

// SkipTail allow skip n last items
func (that *Collection[T]) SkipTail(n int) {
	if n > len(that.items) {
		n = len(that.items)
	}
	that.items = that.items[:n]
}

// SetSorted allow to set sorted flag
func (that *Collection[T]) SetSorted(sorted bool) {
	if that.sorted == sorted {
		return
	}
	that.sorted = sorted
	if sorted {
		sort.Sort(that)
	}
}

func (that *Collection[T]) GetSorted() bool {
	return that.sorted
}

func (that *Collection[T]) GetUnique() bool {
	return that.unique
}

func (that *Collection[T]) GetComparator() Comparator[T] {
	return that.comparator
}

var (
	ErrNoMatch = errors.New("no match")
)
