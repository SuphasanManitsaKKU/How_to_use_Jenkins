package main

import (
	"fmt"
	"test/pkg" // Import the pkg package
)

func main() {
	result := pkg.Add(1, 2)
	fmt.Println("Result:", result)
}
