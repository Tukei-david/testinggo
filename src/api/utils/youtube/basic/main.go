package main

import "fmt"

func addNumber(n int) (result int) {
	result = n + 2
	return result
}

func main() {
	result := addNumber(2)
	fmt.Println(result)
}
