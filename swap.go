package main

func swap(a, b int) (int, int) {
	return b, a
}

func swapUsingPointers(a, b *int) {
	*a, *b = *b, *a
}