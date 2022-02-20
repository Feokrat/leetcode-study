package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	var subStr []rune
	res := 0
	for _, char := range s {
		if ifContains, pos := contains(char, subStr); ifContains {
			if len(subStr) > res {
				res = len(subStr)
			}
			subStr = subStr[pos+1:]
			subStr = append(subStr, char)
		} else {
			subStr = append(subStr, char)
		}
	}

	if len(subStr) > res {
		res = len(subStr)
	}
	return res
}

func contains(char rune, s []rune) (bool, int) {
	for pos, c := range s {
		if char == c {
			return true, pos
		}
	}
	return false, 0
}