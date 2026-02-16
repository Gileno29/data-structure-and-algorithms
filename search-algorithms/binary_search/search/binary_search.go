package search

// SearchRotated encontra o índice de um elemento num array rotacionado
func Search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2 // Evita overflow de inteiros

		if nums[mid] == target {
			return mid
		}
		// Caso contrário, a metade direita deve estar ordenada
		if target > nums[mid] && target <= nums[high] {
			low = mid + 1 // Alvo está na direita
		} else {
			high = mid - 1 // Alvo está na esquerda
		}
	}

	return -1
}
