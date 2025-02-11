package collections

import (
	"fmt"
)

type MyType struct {
	value int
}

type MyTypeComparator struct{}

func (c MyTypeComparator) Less(a, b *MyType) bool {
	return a.value < b.value
}

func (c MyTypeComparator) Equal(a, b *MyType) bool {
	return a.value == b.value
}

func (c MyTypeComparator) Greater(a, b *MyType) bool {
	return a.value > b.value
}

func ExampleNewCollection() {
	collection := NewCollection[*MyType](&MyTypeComparator{})

	collection.Include(&MyType{1})
	collection.Include(&MyType{3})
	collection.Include(&MyType{2})
	collection.Include(&MyType{4})

	for _, item := range collection.Items() {
		fmt.Println("Item:", item.value)
	}

	// Output:
	// Item: 1
	// Item: 2
	// Item: 3
	// Item: 4
}
