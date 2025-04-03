package list

import "fmt"

type ArrayList[T any] struct {
	items []T
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{}
}

func (list *ArrayList[T]) Add(item T) {
	list.items = append(list.items, item)
}

func (list *ArrayList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= len(list.items) {
		return zero, fmt.Errorf("Index is out of range")
	}

	return list.items[index], nil
}

func (list *ArrayList[T]) Size() int {
	return len(list.items)
}

func (list *ArrayList[T]) Remove(index int) error {
	if index < 0 || index >= len(list.items) {
		return fmt.Errorf("Index is out of range")
	}
	list.items = append(list.items[:index], list.items[index+1:]...)
	return nil
}

func (list *ArrayList[T]) Print() {
	for _, v := range list.items {
		fmt.Println(v)
	}
}
