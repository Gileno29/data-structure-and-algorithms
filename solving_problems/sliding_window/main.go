package main

import (
	"fmt"
	"slide-window/sliding"
)

func main() {

	array := []string{"b", "c", "b", "b", "b", "c", "b", "a"}

	fmt.Println(sliding.Sliding(array))
}
