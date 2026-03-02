package tree

import "fmt"

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

func Show(n *Node) {
	if n == nil {
		return
	}
	Show(n.Left)
	fmt.Printf("%d ->", n.Data)
	Show(n.Rigth)
}
