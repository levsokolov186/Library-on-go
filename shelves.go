package main

import "fmt"

const MaxItemsPerShelf = 5

type ShelfItem interface {
	Book | Newspaper | Comics
}

type Shelf[T ShelfItem] struct {
	Number int
	Items  []T
}

func NewShelf[T ShelfItem](number int) *Shelf[T] {
	return &Shelf[T]{
		Number: number,
		Items:  make([]T, 0, MaxItemsPerShelf),
	}
}

func (s *Shelf[T]) Add(item T) error {
	if len(s.Items) >= MaxItemsPerShelf {
		return fmt.Errorf("полка %d уже заполнена: максимум %d объектов", s.Number, MaxItemsPerShelf)
	}

	s.Items = append(s.Items, item)
	return nil
}

func (s *Shelf[T]) Count() int {
	return len(s.Items)
}

func (s *Shelf[T]) IsFull() bool {
	return len(s.Items) >= MaxItemsPerShelf
}
