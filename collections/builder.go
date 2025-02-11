package collections

import "errors"

type Builder[T any] struct {
	collection *Collection[T]
}

func NewBuilder[T any]() *Builder[T] {
	return &Builder[T]{
		collection: &Collection[T]{},
	}
}

func (that *Builder[T]) WithComparator(comparator Comparator[T]) *Builder[T] {
	that.collection.comparator = comparator
	return that
}

func (that *Builder[T]) WithUnique(unique bool) *Builder[T] {
	that.collection.unique = unique
	return that
}

func (that *Builder[T]) WithSorted(sorted bool) *Builder[T] {
	that.collection.sorted = sorted
	return that
}

func (that *Builder[T]) Build() (*Collection[T], error) {
	if err := that.checkRequiredFields(); err != nil {
		return nil, err
	}

	return that.collection, nil
}

func (that *Builder[T]) checkRequiredFields() error {
	if that.collection.comparator == nil && that.collection.sorted {
		return ErrRequiredFieldComparator
	}

	return nil
}

var (
	ErrRequiredFieldComparator = errors.New("required field 'comparator' is not set")
)
