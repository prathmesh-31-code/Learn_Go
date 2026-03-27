package concurrency

import (
	"fmt"
	"time"
)

// Go routines are functions that can executed simultaneously

func GoRoutine1() {

	t1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second * 5)
		fmt.Printf("Elements from GoRoutine1: %d\n", t1[i])
	}
}

func GoRoutine2() {

	t2 := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second * 3)
		fmt.Printf("Elements from GoRoutine1: %d\n ", t2[i])
	}
}
