package basics

import "fmt"

// Basic Datatypes
func Datatypes() {

	//Variables
	var num int = 45
	var num2 float32 = 93.354
	var text string = "Good Morning"
	var val bool = true

	name := "Raj" // typed Variable with type inference

	fmt.Println("Integer: ", num)
	fmt.Println("Decimal: ", num2)
	fmt.Println("String: ", text)
	fmt.Println("Boolean Value: ", val)

	fmt.Println("Typed Inference Variable: ", name)
	fmt.Println()

	//Array

	numbers := [10]int{10, 20, 30, 40, 50, 60}
	fmt.Println("Numbers in Array: ", numbers)

	//Multi Dimensional Array
	numbers2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Multi-Dimensional Array: ", numbers2)
}
