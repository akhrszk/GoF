package main

import "fmt"

type Aggregate[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Item struct {
	name string
}

type ItemAggregate struct {
	items []Item
}

func (c ItemAggregate) Iterator() Iterator[Item] {
	return &ItemIterator{0, c.items}
}

type ItemIterator struct {
	index int
	items []Item
}

func (i ItemIterator) HasNext() bool {
	return len(i.items) > i.index
}

func (i *ItemIterator) Next() Item {
	cur := i.items[i.index]
	i.index++
	return cur
}

type Tester struct {
	a Aggregate[Item]
}

func main() {
	a := ItemAggregate{[]Item{{"Golang"}, {"Java"}}}
	i := a.Iterator()
	for i.HasNext() {
		fmt.Println(i.Next().name)
	}
}
