package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	fmt.Println(reverseString("नमस्ते"))
	fmt.Println(reverseString("ankur"))
	fmt.Println(reverseString("Go 🚀"))

	fmt.Println(reverseStringWithBuilder("नमस्ते"))
	fmt.Println(reverseStringWithBuilder("ankur"))
	fmt.Println(reverseStringWithBuilder("Go 🚀"))

	fmt.Println(fizzbuzz(15))

	x := 1
	y := 2

	fmt.Println("Before swap: x =", x, ", y =", y)
	// fmt.Print("Address of x:", &x, ", y:", &y)
	x, y = swap(x, y)
	fmt.Println("After swap: x =", x, ", y =", y)
	// fmt.Print("Address of x:", &x, ", y:", &y)
	swapUsingPointers(&x, &y)
	fmt.Println("After swap using pointers: x =", x, ", y =", y)
	// fmt.Print("Address of x:", &x, ", y:", &y)

	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	text := "I am learning go. Go is fun and go is powerful"
	res := uniqueWords(text)

	fmt.Println(res)

	rect := Rectangle{
		Width:  10,
		Height: 5,
	}

	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())

	nums := []int{1, 2, 3, 4, 5}
	a := nums[1:3]
	b := nums[2:4]
	a[1] = 999
	fmt.Println(nums, a, b)

	numbers := []int{1, 2, 2, 3, 1, 4, 3, 5}

	fmt.Println(removeDuplicates(numbers))

}
