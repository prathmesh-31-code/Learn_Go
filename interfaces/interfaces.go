package interfaces

// Interfaces :- A interface holds type and value
// A Interface defines a set of methods without providing their code
// just like data abstraction

type text interface {
	display() string
}

type Person1 struct {
	name string
}

func (p Person1) display() string {
	return p.name
}

type result interface {
	displayResult() int
}

type Area struct { // Area type that implements the result interface
	length int
	width  int
}

func (a Area) displayResult() int {
	return a.length * a.width
}
