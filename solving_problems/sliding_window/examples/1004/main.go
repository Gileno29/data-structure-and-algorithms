package main

import (
	"fmt"
)

//  Max Consecutive Ones III (LeetCode 1004) - Médio
// Problema: Dado um array binário e um inteiro k, retorne o número máximo de 1s consecutivos se puder inverter no máximo k zeros.

func MaxConsecutiveNums(arr []int, k int) int {
	//rigth := 0
	left := 0
	max := 0
	for rigth := 0; rigth < len(arr); rigth++ {
		//expand window
		if arr[rigth] == 0 {
			k--
		}
		for k < 0 {
			if arr[left] == 0 {
				k++
			}
			left++
		}

		//the calc must be done after the window be valid again
		if dist := rigth - left + 1; dist > max {
			max = dist
		}

	}

	return max
}

func main() {
	arr := []int{1, 0, 0, 0, 1, 0, 1, 1, 1, 1}
	fmt.Println(MaxConsecutiveNums(arr, 3))

}
