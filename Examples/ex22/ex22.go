package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64
	
	// don't edit above this line
	finalCost = bulkMessageCost
	switch isPremiumUser {
	case true:
		finalCost = finalCost - (finalCost * discountRate)
	}
	
	switch {
	case accountBalance >= finalCost:
		accountBalance = accountBalance - finalCost
		fmt.Println(purchaseSuccessMessage)
	default:
		fmt.Println(insufficientFundMessage)
	}
}

