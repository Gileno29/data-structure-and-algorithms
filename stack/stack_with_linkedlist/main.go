package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	Top  *Node[T]
	Size int
}

func (s *Stack[T]) push(element T) {
	newNode := Node[T]{Value: element}
	newNode.Next = s.Top
	s.Top = &newNode
	s.Size++

}

func (s *Stack[T]) Pop() (T, error) {

	if s.Top == nil {
		return s.Top.Value, errors.New("the stack is empty")
	}
	item := s.Top
	s.Top = item.Next
	s.Size--

	return item.Value, nil
}
func (s *Stack[T]) Peek() (T, error) {

	if s.Top == nil {
		return s.Top.Value, errors.New("the stack is empty")
	}
	item := s.Top

	return item.Value, nil
}

func (s *Stack[T]) GetSize() int {
	return s.Size

}

type Pessoa struct {
	Nome  string
	Idade int
	CPF   string
}

func main() {
	p := Pessoa{
		Nome:  "gileno",
		Idade: 28,
		CPF:   "11111111111111",
	}

	p2 := Pessoa{
		Nome:  "paulina",
		Idade: 28,
		CPF:   "2222222222",
	}
	var s Stack[Pessoa]

	s.push(p)
	s.push(p2)

	fmt.Println(s.GetSize())
	fmt.Println(s.Peek())
}
