package sliding

func Sliding(s []string) int {
	characters := s
	l, r := 0, 0

	max := 1

	counter := make(map[string]int)

	// the first letter in the string
	counter[string(characters[0])] = 1

	for r < (len(characters) - 1) {
		r += 1

		_, ok := counter[string(characters[r])]
		if !ok {
			counter[string(characters[r])] = 1
		} else {
			counter[string(characters[r])] += 1
		}

		for counter[string(characters[r])] == 3 {
			counter[string(characters[l])] -= 1
			l += 1

		}

		if (r-l)+1 > max {
			max = (r - l) + 1

		}
	}
	return max
}
