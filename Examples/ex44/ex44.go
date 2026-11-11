package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name string
	hourlyPay int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	employees := []employee{
		contractor{name: "Carl", hourlyPay: 50, hoursPerYear: 1000},
		fullTime{name: "Fiona", salary: 80000},
	}

	for _, e := range employees {
		fmt.Printf("%s earns %d\n", e.getName(), e.getSalary())
	}
}

