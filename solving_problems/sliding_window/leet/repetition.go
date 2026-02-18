package lett

func ContainsNearbyDuplicate(nums []int, k int) bool {
	// Usamos um map como um "Set" (apenas para verificar existência)
	window := make(map[int]struct{})

	for i, num := range nums {
		// 1. Se o número já está na janela, achamos o par!
		if _, found := window[num]; found {
			return true
		}

		// 2. Adiciona o número atual à janela
		window[num] = struct{}{}

		// 3. Mantém a janela no tamanho k.
		// Se passamos de k, removemos o elemento mais antigo.
		if i >= k {
			delete(window, nums[i-k])
		}
	}

	return false
}
