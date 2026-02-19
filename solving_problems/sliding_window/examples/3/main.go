package main

import "fmt"

//  Longest Substring Without Repeating Characters (LeetCode 3) - Médio
// Problema: Encontre o comprimento da maior substring sem caracteres repetidos.

func MaxLength(s string) int {
	lo := 0
	max := 0
	// Usamos um mapa para contar as ocorrências na janela atual
	counts := make(map[byte]int)

	for hi := 0; hi < len(s); hi++ {
		charHi := s[hi]
		counts[charHi]++

		// Enquanto houver duplicata (o caractere atual apareceu 2 vezes)
		for counts[charHi] > 1 {
			charLo := s[lo]
			counts[charLo]-- // Removemos o caractere da esquerda do mapa
			lo++             // Encolhemos a janela pela esquerda
		}

		// A distância entre hi e lo + 1 é o tamanho da substring atual
		if current := hi - lo + 1; current > max {
			max = current
		}
	}
	return max
}

func main() {
	fmt.Println(MaxLength("abcabcbb"))
}
