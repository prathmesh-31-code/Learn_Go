package miscellaneous

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("cannot divide by zero")
		fmt.Println("Error:", err) // printing here
		return 0, err
	}

	result := a / b
	fmt.Println("Division result:", result)
	return result, nil
}
