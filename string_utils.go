package main

import "strconv"

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


func fizzbuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}