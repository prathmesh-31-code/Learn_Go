package basics

import "fmt"

// --------------------
// 1. Person (Normal Implementation)
// --------------------

type Person1 struct {
	Name string
	Age  int
}

// Display prints the person's details
func (p Person1) Display() {
	fmt.Println(p.Name, p.Age)
}

// --------------------
// 2. Car (Value Receiver)
// --------------------

// Value Receiver:
// - A copy of the struct is made
// - Original data is NOT modified
// - Safe and predictable
//
// Usage:
// - When you don’t want to change the struct
// - Struct is small
// - Read-only behavior

type Car struct {
	Name string
}

// ChangeCar attempts to modify the name
// But since it's a value receiver, it modifies only a copy
func (c Car) ChangeCar(Name string) {
	c.Name = Name
}

// --------------------
// 3. Fruit (Pointer Receiver)
// --------------------

// Pointer Receiver:
// - Works on original struct
// - Modifications persist
//
// Usage:
// - When you need to modify fields
// - Struct is large
// - For consistency across methods

type Fruit struct {
	Color string
}

// ChangeColour modifies the original struct
func (f *Fruit) ChangeColour(Color string) {
	f.Color = Color
}

// --------------------------
// ----NESTED STRUCTURES-----
// --------------------------

type Address struct {
	City string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

// --------------------------
// ---STRUCTURE EMBEDDING----
// --------------------------

// Go's Inheritance

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}


// main function Code

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