package main

import "fmt"

// SearchRotated encontra o índice de um elemento num array rotacionado
func SearchRotated(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2 // Evita overflow de inteiros

		if nums[mid] == target {
			return mid
		}

		// Verifica se a metade esquerda está ordenada necessário apenas para listas rotacionadas
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1 // Alvo está na esquerda
			} else {
				low = mid + 1  // Alvo está na direita
			}
		} else {
			// Caso contrário, a metade direita deve estar ordenada
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1  // Alvo está na direita
			} else {
				high = mid - 1 // Alvo está na esquerda
			}
		}
	}

	return -1
}

func main() {
	// Array ordenado [0,1,2,4,5,6,7] rotacionado no índice 4
	rotatedArray := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	index := SearchRotated(rotatedArray, target)

	if index != -1 {
		fmt.Printf("Elemento %d encontrado no índice: %d\n", target, index)
	} else {
		fmt.Println("Elemento não encontrado.")
	}
}
