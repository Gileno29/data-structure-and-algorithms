package main

import (
	"fmt"
	lett "slide-window/leet"
	"slide-window/sliding"
)

func main() {

	array := []string{"b", "c", "b", "b", "b", "c", "b", "a"}

	fmt.Println(sliding.Sliding(array))

	nums := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 1, 2, 3}

	fmt.Println(lett.ContainsNearbyDuplicate(nums, 3))
	fmt.Println(lett.ContainsNearbyDuplicate(nums2, 2))
}
