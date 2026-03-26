package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	/*
				Datatypes()
				channel()
				fmt.Println("Main Routine Start")
				go goRoutine1()
				go goRoutine2()
				time.Sleep(time.Second * 5)
				fmt.Println("Main Routine End")


			//Interface
			var t text
			p := Person1{name: "Raj"}
			t = p // Person1 satisfies text
			fmt.Println(t.display())

			var x result
			y := Area{length: 3, width: 5}
			x = y // Area satisfies the interface result
			fmt.Println("Area of Rectangle", x.displayResult())


		res, err := divide(10, 5)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Result:", res)
	*/

	//Defer()
	//PanicEx()
	TA()
	TS()
}

// Basic Datatypes
func Datatypes() {

	//Variables
	var num int = 45
	var num2 float32 = 93.354
	var text string = "Good Morning"
	var val bool = true

	name := "Raj" // typed Variable with type inference

	fmt.Println("Integer: ", num)
	fmt.Println("Decimal: ", num2)
	fmt.Println("String: ", text)
	fmt.Println("Boolean Value: ", val)

	fmt.Println("Typed Inference Variable: ", name)
	fmt.Println()

	//Array

	numbers := [10]int{10, 20, 30, 40, 50, 60}
	fmt.Println("Numbers in Array: ", numbers)

	//Multi Dimensional Array
	numbers2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Multi-Dimensional Array: ", numbers2)
}

// Pointers are used to point the memory address of a variable

func Pointers() {
	x := 10
	p := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value via pointer:", *p)

	var pa *int = &x
	ptf(pa)
	fmt.Println("The Value of x after function call: ", x)

}

// Structure is a datatype used to group related data together

type Person struct {
	name string
	age  int
}

func ptf(a *int) {
	*a = 834
}

// Map funciton are like key value pairs

func studentMap() {
	student := map[string]int{
		"Maths":   80,
		"Science": 79,
	}
	student["english"] = 88

	for subject, marks := range student {
		fmt.Println(subject, marks)
	}
	time.Sleep(time.Second * 15)
}

// Channels are used to send and receive data between go routines

func channel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- "hello"
		ch2 <- "world"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)

	case msg2 := <-ch2: //<- channel operator used to receive signals and store in msg
		fmt.Println(msg2)
	}
}

// Go routines are functions that can executed simultaneously

func goRoutine1() {

	t1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second * 5)
		fmt.Printf("Elements from GoRoutine1: %d\n", t1[i])
	}
}

func goRoutine2() {

	t2 := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second * 3)
		fmt.Printf("Elements from GoRoutine1: %d\n ", t2[i])
	}
}

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

// Error handling

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil

}

// Defer, panic, Recover

// Defer - A statement which is used to postpone the execution of a function until the surrounding function returns.
// Follows Last in, First out (LIFO) Policy

func Defer() {
	defer fmt.Println("world")
	fmt.Println("Hello")

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

}

// Panic - A built-in Mechanism used to handle serious, unrecoverable errors by imediately stopping normal program execution and beginning a
// 		   controlled shutdown process.
//       - Aborts the current function execution and starts unwinding the stack

// Recover - A mechanism that regains control over panicking goroutine
// Must be used inside a deffered function, otherwise it does nothing.

func PanicEx() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in PanicEx", r)
		}
	}()
	PanicEx2()
}

func PanicEx2() {
	defer fmt.Println("Defer in Panic Exe2")
	panic("Crash")
}

/* Memory Models basics in GO

Stack:
Go uses stack and heap as primary data structures for memory management
A stack is a data structure which follows a LIFO - Last In First Out Policy.
A stack stores:
- Function calls
- local variables
- Execution context

Heap:
A Heap is a specialised tree-based data structure that satisfies the heap property,
allowing for efficient priority queue management , where the root node always contains
the highest or lowest priority element

A Heap stores:
- Dynamically allocated memory
- Data that outlives function scope
- Shared Data between Go Routines

Heaps are usually used to implement priority queues, where the smallest(Min-Heap) or largest(Max-Heap) element is always at the root of the tree.
*/

// Type Assertions & Type Switch

// Type Assertion - It is used to extract the concrete value from an interface

func TA() {
	var i interface{} = 12
	str := i.(int) //Panics if type doesn't match
	fmt.Println(str)
}

// Type Switch
// A type switch check the type of an interface against multiple types
func TS() {

	var x interface{} = 9.3

	fmt.Printf("%T\n", x)
	switch v := x.(type) {
	case float64:
		fmt.Println("Float type: ", v)

	default:
		fmt.Println("Unknown Type: ", v)
	}

}
