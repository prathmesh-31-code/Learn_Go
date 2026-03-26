package advanced

import "fmt"

func TS() {
	var x interface{} = 9.3

	switch v := x.(type) {
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown:", v)
	}
}
