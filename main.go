package main

import (
	"fmt"
	"learn_go/basics"
)

func main() {

	// Nested Structure example
	e := basics.Employee{
		Name: "Vivek",
		Age:  21,
		Address: basics.Address{
			City: "Pune",
		},
	}
	fmt.Println("Details of Employee: ", e.Name, e.Age, e.Address.City)

	// Value Receiver Example
	c := basics.Car{Name: "Mercedes"}
	c.ChangeCar("BMW")
	fmt.Println("Printing Car Name: ", c.Name) // Value of the car does not change

	// Pointer Receiver Example
	f := basics.Fruit{Color: "Red"}
	f.ChangeColour("Green") // Color Changes to Green
	fmt.Println("Color of Fruit: ", f.Color)

	// Composition example

	d := basics.Dog{}
	d.Name = "Clove"
	fmt.Println(d.Name)

}
