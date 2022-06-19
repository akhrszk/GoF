package main

import "fmt"

type Iterable[T any] []T

func (i Iterable[T]) Iterator() Iterator[T] {
	return Iterator[T]{i, 0}
}

type Iterator[T any] struct {
	Iterable[T]
	cursor int
}

func (i Iterator[T]) HasNext() bool {
	return len(i.Iterable) > i.cursor
}

func (i *Iterator[T]) Next() T {
	curr := (i.Iterable)[i.cursor]
	i.cursor++
	return curr
}

type Item struct {
	id   string
	name string
}

func main() {
	list := []Item{{"1234", "Golang"}, {"9999", "Java"}}
	iterator := Iterable[Item](list).Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
