package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	c, ok := e.(email)
	if ok {
		return c.toAddress, c.cost()
	}
	f, ok := e.(sms)
	if ok {
		return f.toPhoneNumber, f.cost()
	}
	return "", 0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body string
	toAddress string
}

type sms struct {
	isSubscribed bool
	body string
	toPhoneNumber string
}

type invalid struct{}

func (em email) cost() float64 {
	if !em.isSubscribed {
		return float64(len(em.body)) * .1
	}
	return float64(len(em.body)) * 3 
}

func (sm sms) cost() float64 {
	if !sm.isSubscribed {
		return float64(len(sm.body)) * .1
	}
	return float64(len(sm.body)) * 3
}

func (inv invalid) cost() float64 {
	return 0.0	
}

func main() {
	expenses := []expense{
		email{isSubscribed: true, body: "Hello there!", toAddress: "a@b.com"},
		sms{isSubscribed: false, body: "Hi!", toPhoneNumber: "555-1234"},
		invalid{},
	}

	for _, e := range expenses {
		to, cost := getExpenseReport(e)
		fmt.Printf("to: %q, cost: %.2f\n", to, cost)
	}
}
