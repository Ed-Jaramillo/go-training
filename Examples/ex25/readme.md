# Passing Variables by Value

Variables in Go are passed by value (exept for a few data types we haven't 
covered yet). "Pass by value" means that when a variable is passed into a function,
that function received a copy of the variable. The function is unable to mutate the caller's original data.


func main() {
	x := 5
	increment(x)
	fmt.Println(x)
	// still prints 5
	// because the increment function
	// received a copy of x
}

func increment(x int) {
	x++
}

# Assignment

* monthlyBillIncrease: Should return the increase in the bill from the
previous to the current month. If the bill decreased, return a negative number.

* getBillForMonth: Should return the total cost for the number of messages sent.


Fix the bugs in the monthlyBillIncrease and getBillForMonth functions. Looks
like whoever wrote the function's bill parameter would be passed by value. It's not actually updating
the lastMonthBill and thisMonthBill variables as intended so monthlyBillIncrease isn't returning the right result.



