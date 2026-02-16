package main

import (
	"fmt"
	"trie-implementation/trie"
)

func main() {

	Trie := trie.NewTrie()

	fmt.Println(Trie)

	Trie.Insert("apple")
	Trie.Insert("banana")
	Trie.Insert("app")
	Trie.Insert("orange")

	fmt.Println(Trie.Search("ap"))

	fmt.Println(Trie.Starts_with("ap"))

	fmt.Println(Trie.ShowAll())

}
