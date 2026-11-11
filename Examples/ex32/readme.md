# Processing Orders

Management thinks our branding is so creative that our SaaS customers wil pay for Textio merch.


# Assignment

Complete the placeOrder function.

It returns a bool indicating whether the order was successful (true is a success) and float64 representing
the user's balance after the order. THe placeOrder function should always return the account balance regardless of whether it was adjusted.

The amountInStock and calcPrice functions can be used to look up the current stock and price for an item.

*  If the quantity is greater than the amount in stock, the order should be rejected.
* If the user doesn't have enough money in their account, the order should be rejected
* Otherwise, the order should be accepted and you should return the new balance.