package main

import (
	"github.com/rubichandrap/design-patterns/internal/behavioral/strategy"
)

func main() {
	cart := strategy.ShoppingCart{}
	cart.SetStrategy(strategy.CreditCard{})
	cart.Checkout(120.5)

	cart.SetStrategy(strategy.PayPal{})
	cart.Checkout(89.9)
}