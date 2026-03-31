package main

import (
	"fmt"
	"learn_go/concurrency"
)

func main() {

	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	for _, num := range nums {
		go concurrency.Squares(num, ch)

	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
}
