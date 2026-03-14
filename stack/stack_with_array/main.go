package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T

		return zeroValue, errors.New("The stack is empty")
	}

	index := len(s.Items) - 1
	item := s.Items[index]
	s.Items = s.Items[:index]

	return item, nil

}
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T

		return zeroValue, errors.New("The stack is empty")
	}

	return s.Items[len(s.Items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack[T]) Size() int {
	if s.IsEmpty() {
		return 0
	}

	return len(s.Items)
}
func main() {
	item := []string{"gileno", "paulina"}

	s1 := Stack[string]{Items: item}

	fmt.Println(s1)
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())

}
