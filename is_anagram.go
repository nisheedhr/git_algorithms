package algorithms

func isAnagram(s string, t string) bool {

	runeMap := make(map[rune]rune)

	for _, runeVal := range s {
		runeMap[runeVal] += 1
	}

	for _, runeVal := range t {
		runeMap[runeVal] -= 1
	}

	for _, v := range runeMap {
		if v != 0 {
			return false
		}
	}
	return true
}
