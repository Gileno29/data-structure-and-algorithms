package tree

import (
	"fmt"

	"github.com/gammazero/deque"
)

type Node struct {
	Data  int
	Left  *Node
	Rigth *Node
}

type BinaryTree struct {
	Root *Node
}

func (b *BinaryTree) Insert(d int) {
	if b.Root == nil {
		b.Root = &Node{Data: d, Left: nil, Rigth: nil}
	} else {
		b.InsertRecursive(d, b.Root)
	}
}

func (b *BinaryTree) InsertRecursive(data int, n *Node) {
	if data < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: data, Left: nil, Rigth: nil}
		} else {
			b.InsertRecursive(data, n.Left)
		}
	} else {
		if n.Rigth == nil {
			n.Rigth = &Node{Data: data, Left: nil, Rigth: nil}
		} else {
			b.InsertRecursive(data, n.Rigth)

		}
	}
}

func (b *BinaryTree) Search(data int) bool {

	return b.SearchRecursive(b.Root, data)

}

func (b *BinaryTree) SearchRecursive(node *Node, data int) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	} else if data < node.Data {
		return b.SearchRecursive(node.Left, data)
	} else {
		return b.SearchRecursive(node.Rigth, data)
	}

}

func (b *BinaryTree) DFS(data int) bool {

	return b.DFSRecursive(b.Root, data)

}

func (b *BinaryTree) DFSRecursive(node *Node, data int) bool {

	if node == nil {
		return false
	}
	fmt.Println(node.Data)
	if node.Data == data {
		return true

	}
	if b.DFSRecursive(node.Left, data) {
		return true
	}
	if b.DFSRecursive(node.Rigth, data) {
		return true
	}
	return false
}

func (b *BinaryTree) BFS(target int) bool {
	if b.Root == nil {
		return false
	}

	var q deque.Deque[*Node]

	q.PushFront(b.Root)

	for q.Len() != 0 {

		node := q.PopFront()
		if node.Data == target {
			return true
		}
		fmt.Println(node.Data)

		if node.Left != nil {
			q.PushFront(node.Left)
		}

		if node.Rigth != nil {
			q.PushFront(node.Rigth)

		}
	}
	return false

}
func (b *BinaryTree) PreorderTraversal() []int {
	var result []int

	b.PreorderRecursive(b.Root, &result)

	return result
}

func (b *BinaryTree) PreorderRecursive(node *Node, result *[]int) {

	if node != nil {
		*result = append(*result, node.Data)
		b.PreorderRecursive(node.Left, result)
		b.PreorderRecursive(node.Rigth, result)
	}
}

func (b *BinaryTree) InOrderTraversal() []int {
	var result []int

	b.InOrderRecursive(b.Root, &result)

	return result
}

func (b *BinaryTree) InOrderRecursive(node *Node, result *[]int) {

	if node != nil {
		b.InOrderRecursive(node.Left, result)
		*result = append(*result, node.Data)
		b.InOrderRecursive(node.Rigth, result)

	}
}

func (b *BinaryTree) PostOrderTraversal() []int {
	var result []int

	b.PostOrderRecursive(b.Root, &result)

	return result
}

func (b *BinaryTree) PostOrderRecursive(node *Node, result *[]int) {

	if node != nil {
		b.PostOrderRecursive(node.Left, result)
		b.PostOrderRecursive(node.Rigth, result)
		*result = append(*result, node.Data)

	}
}
func Show(n *Node) {
	if n == nil {
		return
	}
	Show(n.Left)
	fmt.Printf("%d ->", n.Data)
	Show(n.Rigth)
}
