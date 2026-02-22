package main

import "fmt"

func Buble(nums []int) {
	size := len(nums)

	for idx, _ := range nums {
		isSorted := true
		fmt.Println(idx, nums)
		for i := 0; i < (size - 1); i++ {
			tmp := nums[i]
			if nums[i] > nums[i+1] {
				isSorted = false
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
		}
		if isSorted {
			return
		}
	}
}

func main() {
	Buble([]int{5, 4, 3, 2, 1})
	Buble([]int{1, 2, 3, 4, 5})
}
