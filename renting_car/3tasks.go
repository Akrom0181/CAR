package main

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for i, v := range s {
		sum += roman[string(v)]
		if i != 0 {
			if roman[string(s[i-1])] < roman[string(v)] {
				sum -= 2 * roman[string(s[i-1])]
			}
		}
	}
	return sum
}

func isPalindrome(x int) bool {
	var b int
	var t int = x
	if x < 0 {
		return false
	}
	for x > 0 {
		b = b*10 + x%10
		x = x / 10

	}
	if t == b {
		return true
	} else {
		return false
	}

}

func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}
