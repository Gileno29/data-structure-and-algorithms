package main

import (
	"binarytree/tree"
	"fmt"
)

func main() {

	btree := tree.BinaryTree{Root: nil}

	btree.Insert(10)
	btree.Insert(20)
	btree.Insert(30)
	btree.Insert(40)
	btree.Insert(550)
	btree.Insert(320)
	btree.Insert(1000)
	btree.Insert(120)
	btree.Insert(110)
	btree.Insert(330)
	btree.Insert(220)
	btree.Insert(11110)

	fmt.Println(btree.Search(2))

	tree.Show(btree.Root)
}
