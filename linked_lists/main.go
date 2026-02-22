package main

import "fmt"

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

	fmt.Println("--- Teste 1: Adicionando elementos ---")
	list.AddToFront(10) // [10]
	list.AddToFront(20) // [20, 10]
	list.AddToEnd(30)   // [20, 10, 30]
	list.Display()      // Esperado: 20 <-> 10 <-> 30

	fmt.Println("\n--- Teste 2: Removendo da frente ---")
	val1 := list.RemoveFromFront()
	fmt.Printf("Removido: %d | ", val1)
	list.Display() // Esperado: 10 <-> 30

	fmt.Println("\n--- Teste 3: Removendo do fim ---")
	val2 := list.RemoveFromEnd()
	fmt.Printf("Removido: %d | ", val2)
	list.Display() // Esperado: 10

	fmt.Println("\n--- Teste 4: Esvaziando a lista ---")
	list.RemoveFromFront()
	list.Display() // Esperado: NIL

	fmt.Printf("Remover de lista vazia: %d\n", list.RemoveFromFront()) // Esperado: -1

	fmt.Println("\n--- Teste 5: Re-populando e checando ordem ---")
	list.AddToEnd(1)
	list.AddToEnd(2)
	list.AddToFront(0)
	list.Display() // Esperado: 0 <-> 1 <-> 2
}
