package search

// SearchRotated encontra o índice de um elemento num array rotacionado
func BinarySearch(nums []int, target int, l int, h int) int {
	low := l
	high := h

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

func ExponentialSearch(arr []int, target int) int {

	if arr[0] == target {
		return 0
	}

	n := len(arr)

	if n == 0 {
		return -1
	}

	right := 1

	for right < n && arr[right] < target {
		if arr[right] == target {
			return right
		}
		right *= 2

	}
	return BinarySearch(arr, target, right/2, min(right, n-1))
}
