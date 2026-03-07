package main

import (
	"binarytree/tree"
	"fmt"
)

func main() {

	btree := tree.BinaryTree{Root: nil}

	btree.Insert(5)
	btree.Insert(3)
	btree.Insert(1)
	btree.Insert(10)
	btree.Insert(15)
	btree.Insert(7)

	fmt.Println(btree.Search(2))

	tree.Show(btree.Root)
	fmt.Println()

	fmt.Println("Pre order Traversal:", btree.PreorderTraversal())

	fmt.Println("In order Traversal:", btree.InOrderTraversal())

	fmt.Println("Post order Traversal:", btree.PostOrderTraversal())

	fmt.Println(btree.DFSRecursive(btree.Root, 22))
	fmt.Println(btree.BFS(20))
}
