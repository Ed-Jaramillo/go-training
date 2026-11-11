package main
// import "fmt"

func bulkSend(numMessages int) float64 {
	//var totalCost float64
	//var cost float64
	var decimalCost float64
	for i := 0; i < numMessages; i++ {
		decimalCost += 1.00 + (float64(i)/100.00)
	}
	return decimalCost
}