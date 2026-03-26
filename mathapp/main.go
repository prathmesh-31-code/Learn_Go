package main

// Packages
// A package is a collection of related Go files in the same directory that are compiled together

import (
	"fmt"
	"mathapp/mathutils"
)

func main() {
	a, b := 10, 5

	fmt.Println("Add: ", mathutils.Add(a, b))
	fmt.Println("Subtract: ", mathutils.Subtract(a, b))
	fmt.Println("Multiply: ", mathutils.Multiply(a, b))

	result, err := mathutils.Divide(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Divide: ", result)
	}
}
