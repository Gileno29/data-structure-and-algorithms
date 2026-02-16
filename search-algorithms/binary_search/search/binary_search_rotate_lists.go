package search

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
				low = mid + 1 // Alvo está na direita
			}
		} else {
			// Caso contrário, a metade direita deve estar ordenada
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1 // Alvo está na direita
			} else {
				high = mid - 1 // Alvo está na esquerda
			}
		}
	}

	return -1
}
