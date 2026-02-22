package main

import "fmt"

func Buble(nums []int) {
	size := len(nums)

	for idx, _ := range nums {
		fmt.Println(idx, nums)
		for i := 0; i < (size - 1); i++ {
			tmp := nums[i]
			if nums[i] > nums[i+1] {
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
		}
	}
}

func main() {
	Buble([]int{5, 4, 3, 2, 1})
}
