package basics

import "fmt"

// Pointers are used to point the memory address of a variable

func Pointers() {
	x := 10
	p := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value via pointer:", *p)

	var pa *int = &x
	ptf(pa)
	fmt.Println("The Value of x after function call: ", x)

}

type Person struct {
	name string
	age  int
}

func ptf(a *int) {
	*a = 834
}
