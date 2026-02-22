package main

import (
	"fmt"
)

// 5. Permutation in String (LeetCode 567) - Médio
// Problema: Dadas duas strings s1 e s2, retorne true se s2 contiver uma permutação de s1.

// CheckInclusion verifica se s2 contém uma permutação de s1 utilizando Maps.
// Esta abordagem é mais genérica e funciona com qualquer caractere Unicode.
func CheckInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	// Criamos mapas para armazenar a frequência dos caracteres
	s1Count := make(map[rune]int)
	s2Count := make(map[rune]int)

	// 1. Mapeia a frequência de s1 e da primeira janela de s2
	// Usamos []rune para garantir que lidamos corretamente com caracteres Unicode
	r1 := []rune(s1)
	r2 := []rune(s2)

	for i := 0; i < n1; i++ {
		s1Count[r1[i]]++
		s2Count[r2[i]]++
	}

	// 2. Desliza a janela por s2
	for i := 0; i < n2-n1; i++ {
		// No Go, reflect.DeepEqual pode comparar mapas, mas é lento.
		// Para performance, o ideal é uma função de comparação manual.
		if mapsAreEqual(s1Count, s2Count) {
			return true
		}

		// Remove o caractere que sai da janela (à esquerda)
		charSaindo := r2[i]
		s2Count[charSaindo]--
		if s2Count[charSaindo] == 0 {
			delete(s2Count, charSaindo) // Limpa o mapa para manter a comparação correta
		}

		// Adiciona o caractere que entra na janela (à direita)
		charEntrando := r2[i+n1]
		s2Count[charEntrando]++
	}

	// Verifica a última janela
	return mapsAreEqual(s1Count, s2Count)
}

// mapsAreEqual verifica se dois mapas de frequência são idênticos.
func mapsAreEqual(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	// Testes
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Printf("Usando Mapas - s1: %s em s2: %s -> %v\n", s1, s2, CheckInclusion(s1, s2))

	s1_2 := "ab"
	s2_2 := "eidboaoo"
	fmt.Printf("Usando Mapas - s1: %s em s2: %s -> %v\n", s1_2, s2_2, CheckInclusion(s1_2, s2_2))
}
