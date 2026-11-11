package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(input int) int {
		sum += input
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}