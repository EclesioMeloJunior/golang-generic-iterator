package main

import "fmt"

type iterable[T any] struct {
	idx int
	set []T
}

func (i *iterable[T]) Next() T{
	var t T

	if i.HasNext() {
		t = i.set[i.idx]
		i.idx += 1

		return t
	}

	return t
}

func (i *iterable[T]) HasNext() bool {
	return i.idx < len(i.set)
}

func toIter[T any](sl *[]T) *iterable[T] {
	iter := iterable[T]{
		idx: 0,
	}

	iter.set = make([]T, len(*sl))
	copy(iter.set, (*sl)[:])

	return &iter
}

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5}
	intIter := toIter[int](&intSlice)

	for intIter.HasNext() {
		fmt.Println(intIter.Next())
	}

	strings := []string{"omg", "generics"}
	anotherIter := toIter[string](&strings)

	for anotherIter.HasNext() {
		fmt.Println(anotherIter.Next())
	}
}