package main

func reverseString(s string) string {
	runes := []rune(s)

	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

// Alternative implementation using a builder
func reverseStringWithBuilder(s string) string {
	runes := []rune(s)
	builder := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		builder[len(runes)-1-i] = runes[i]
	}
	return string(builder)
}