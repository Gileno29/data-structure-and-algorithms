package hash_map

func FirstUniqChar(s string) int {
	characters := []rune(s)
	d := make(map[string][]int)

	for index, value := range characters {

		if _, ok := d[string(value)]; !ok {
			d[string(value)] = []int{index, 1}

		} else {
			d[string(value)][1]++
		}
	}
	var firstUniqLetter string
	encontrou := false

	// Percorremos a string original para garantir a ordem de "quem apareceu primeiro"
	for _, value := range characters {
		charStr := string(value)

		// Verificamos no seu mapa 'd' se a contagem (posição 1) é igual a 1
		if d[charStr][1] == 1 {
			firstUniqLetter = charStr
			encontrou = true
			break // Já achamos a primeira, podemos parar o loop
		}
	}

	if encontrou {
		return d[firstUniqLetter][0]
	} else {
		return -1
	}

}
