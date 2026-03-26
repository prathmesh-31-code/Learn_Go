package basics

import "fmt"

type Person1 struct {
	name string
	age  int
}

func (p Person1) display() {
	fmt.Println(p.name, p.age)
}
