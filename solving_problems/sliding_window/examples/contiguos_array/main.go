package main

import "fmt"

// Given an array of integers nums and a positive integer k, find the maximum sum of any contiguous subarray of size k

func maxSum(nums []int, k int) int {
	max := 0
	left := 0
	currentSum := 0

	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		if right >= k {
			currentSum -= nums[left]
			left++
		}

		if (right - left + 1) == k {
			if currentSum > max || left == 0 { // left == 0 handles the first full window
				max = currentSum
			}
		}

	}

	return max

}
func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	nums2 := []int{2, 3, 4, 1, 5}
	k := 3

	fmt.Println(maxSum(nums, k))

	fmt.Println(maxSum(nums2, 2))
}
