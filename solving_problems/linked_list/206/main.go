package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func newNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

func (l *DoublyLinkedList) newDoblyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *DoublyLinkedList) AddToFront(value int) {
	n := newNode(value)
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	n.Next = l.Head
	l.Head.Prev = n
	l.Head = n

}

func (l *DoublyLinkedList) AddToEnd(value int) {
	n := newNode(value)
	if l.Tail == nil {
		l.Head = n
		l.Tail = n
		return
	}
	n.Prev = l.Tail
	l.Tail.Next = n
	l.Tail = n

}

func (l *DoublyLinkedList) RemoveFromFront() int {
	if l.Head == nil {
		return -1
	}
	val := l.Head.Value
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	return val
}
func (l *DoublyLinkedList) RemoveFromEnd() int {
	if l.Tail == nil {
		return -1
	}
	val := l.Tail.Value
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil // Lista ficou vazia
	}
	return val

}

// Função auxiliar para visualizar a lista
func (l *DoublyLinkedList) Display() {
	curr := l.Head
	fmt.Print("Lista: ")
	for curr != nil {
		fmt.Printf("%d <-> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("NIL")
}
func main() {
	list := &DoublyLinkedList{}
	list.AddToFront(10)
	list.AddToFront(20)
	list.AddToEnd(30)
	list.Display()

	var NewList *Node = nil
	current := list.Head

	for current != nil {
		nextNode := current.Next // Salva o próximo antes de alterar

		current.Next = NewList  // Inverte o ponteiro Next
		current.Prev = nextNode // Inverte o ponteiro Prev (para manter a consistência da Double Linked)

		NewList = current  // Move NewList para o nó atual
		current = nextNode // Avança para o próximo nó original
	}

	fmt.Println("\nPercorrendo a nova lista invertida:")
	// Iteração correta sobre um nó:
	temp := NewList
	for temp != nil {
		fmt.Printf("Valor: %d\n", temp.Value)
		temp = temp.Next // ESSENCIAL: Avançar o ponteiro para evitar loop infinito
	}
}
