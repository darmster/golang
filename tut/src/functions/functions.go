package main

// A function can take zero or more arguments.

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 12))
}

