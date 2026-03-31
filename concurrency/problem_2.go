/*
Problem 2: Parallel Function Calls
Create 3 functions:

task1() → sleeps 1 sec
task2() → sleeps 2 sec
task3() → sleeps 3 sec
Run them concurrently and print:
All tasks completed
*/

package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup is actually a type of counter which blocks the execution of function (or might say A goroutine) until its internal counter become 0.
/*
WaitGroup exports 3 methods.

1	Add(int)	 It increases WaitGroup counter by given integer value.
2	Done()	 It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
3	Wait()	It Blocks the execution until it's internal counter becomes 0.
*/
func Task1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("Task 1 Execution completed")
}

func Task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println("Task 2 Execution completed")
}

func Task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
	fmt.Println("Task 3 Execution completed")
}

/*

	fmt.Println("Main task execution started")
	var wg sync.WaitGroup

	wg.Add(3)
	go concurrency.Task1(&wg)
	go concurrency.Task2(&wg)
	go concurrency.Task3(&wg)
	wg.Wait()

	fmt.Println("Main task execution completed")

*/
