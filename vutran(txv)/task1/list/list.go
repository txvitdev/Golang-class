package list

type List[T any] interface {
	Add(item T)
	Get(index int) (T, error)
	Remove(index int) error
	Size() int
	Print()
}
