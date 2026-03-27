package calculator

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func App() {
	for {

		fmt.Println("--------------CALCULATOR APP--------------")
		var a, b float64
		fmt.Println("Enter first number: ")
		fmt.Scan(&a)
		fmt.Println("Enter second number: ")
		fmt.Scan(&b)

		// Division error

		if b == 0 {
			err := errors.New("Division by zero is not possible")
			fmt.Println("Error: ", err)
			panic("Stopped Execution")
		}

		fmt.Println("-----------SELECT OPERATION--------------")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")

		var op int

		//Operation Error Handliing
		if op > 4 || op < 0 {
			err2 := errors.New("Invalid Operation number!!")
			fmt.Println("Error: ", err2)
			panic("Stopped Execution")
		}

		fmt.Println("Enter Operation number: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Println("Addition: ", a+b)

		case 2:
			fmt.Println("Subtraction: ", a-b)

		case 3:
			fmt.Println("Multiplication: ", a*b)

		case 4:
			fmt.Println("Division: ", a/b)
		}

		var low string
		fmt.Println("Continue? Type Yes/No: ")
		fmt.Scan(&low)

		c := strings.ToLower(low)

		if c == "yes" {
			continue
		} else if c == "no" {
			fmt.Println("Program successfully terminated")
			os.Exit(1)
		} else {
			panic("Unexpected Exit")
		}
	}
}
