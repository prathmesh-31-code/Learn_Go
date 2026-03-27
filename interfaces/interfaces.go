package interfaces

// Interfaces :- A interface holds type and value
// A Interface defines a set of methods without providing their code
// just like data abstraction

type Text interface {
	display() string
}

type Person struct {
	Name string
}

func (p Person) Display() string {
	return p.Name
}

type Result interface {
	displayResult() int
}

type Area struct { // Area type that implements the result interface
	length int
	width  int
}

func (a Area) DisplayResult() int {
	return a.length * a.width
}
