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
}