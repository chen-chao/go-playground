package builtin

import (
	"fmt"
	"testing"
)

func divide(denomenator, numerator int) (result int, err error) {
	// recover called within a defered function will end the
	// panicking state of the function containing the defered
	// statement, and returns the panic value

	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
			return
		}
	}()

	return numerator / denomenator, nil
}

func TestRecover(t *testing.T) {
	_, err := divide(1, 1)
	fmt.Println("error of 1/1: ", err)

	_, err = divide(0, 1)
	fmt.Println("error of 1/0: ", err)
}
