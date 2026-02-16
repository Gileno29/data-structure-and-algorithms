package main

import (
	"binary-search-implementation/search"
	"fmt"
)

func main() {
	// Array ordenado [0,1,2,4,5,6,7] rotacionado no índice 4
	rotatedArray := []int{4, 5, 6, 7, 0, 1, 2}
	Array := []int{0, 1, 2, 4, 5, 6, 7}
	target := 0

	index := search.SearchRotated(rotatedArray, target)
	index2 := search.Search(Array, target)

	if index != -1 {
		fmt.Printf("Elemento %d encontrado no índice: %d\n", target, index)
	} else {
		fmt.Println("Elemento não encontrado.")
	}

	if index2 != -1 {
		fmt.Printf("Elemento %d encontrado no índice: %d\n", target, index2)
	} else {
		fmt.Println("Elemento não encontrado.")
	}
}
