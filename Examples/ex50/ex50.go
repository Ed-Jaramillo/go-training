package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("cannot divide %v by zero", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	if result, err := divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	if result, err := divide(10, 0); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}