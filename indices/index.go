package indicies

type Index[T any] interface {
	Reset()
	Push(item T)
	Pop() (T, error)
	Truncate(iterator func(item T) bool)
}
