package main

import "fmt"

// Minimum Size Subarray Sum (LeetCode 209) - Médio
// Problema: Dado um array de inteiros positivos e um target, encontre o comprimento
// mínimo de um subarray cuja soma seja maior ou igual ao target.

func maxLength(arr []int, k int) int {
	left := 0
	rigth := 0
	res := len(arr) + 1 // can´t be initiated in 0 because we want's the min value
	total := 0          // the sum of items, it will use to make or window
	for i := 0; i < len(arr); i++ {
		rigth++
		total += arr[i]
		for total >= k {
			total -= arr[left]

			if dist := rigth - left + 1; dist < res { // (rigth - left +1) is the formula to legth of an window
				res = dist
			}

			left++ // will move the window
		}

	}
	if res > len(arr) { // if all elements is less tha target it will have the initial value: len(arr)+1
		return 0
	}

	return res

}
func main() {
	arr := []int{1, 2, 3, 4, 3, 6}
	arr2 := []int{0, 1, 0, 1, 0, 0}
	fmt.Println(maxLength(arr, 3))
	fmt.Println(maxLength(arr2, 3))
}
