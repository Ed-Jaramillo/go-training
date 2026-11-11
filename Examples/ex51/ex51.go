package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	err := errors.New("no dividing by 0")
	if y == 0 {
		return 0, err
	}
	return x / y, nil
}

func main() {
	for _, y := range []float64{2, 0} {
		if result, err := divide(10, y); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}