package miscellaneous

import "fmt"

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
