package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

// my comment
func main() {
	fmt.Println("Я домашка")
	fmt.Println(Add(3, 4))
}
