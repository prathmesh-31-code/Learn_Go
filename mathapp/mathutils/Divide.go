package mathutils

import (
	"errors"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot Divide by Zero")
	}
	return a / b, nil
}
