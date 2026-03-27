package main

import (
	"fmt"
	"learn_go/advanced"
	"learn_go/basics"
	"learn_go/interfaces"
	"learn_go/miscellaneous"
)

func main() {
	advanced.Channel()
	basics.Datatypes()
	//concurrency.GoRoutine1()

	p := interfaces.Person1{Name: "Prathmesh"}
	fmt.Println(p.Display())

	miscellaneous.Divide(18, 3)
}
