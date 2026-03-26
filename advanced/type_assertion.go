package advanced

import "fmt"

func TA() {
	var i interface{} = 12
	val := i.(int)
	fmt.Println(val)
}
