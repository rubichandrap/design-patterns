
package strategy

import "fmt"

type PaymentStrategy interface {
    Pay(amount float64)
}

type CreditCard struct{}
func (c CreditCard) Pay(amount float64) {
    fmt.Printf("Paid $%.2f using Credit Card
", amount)
}

type PayPal struct{}
func (p PayPal) Pay(amount float64) {
    fmt.Printf("Paid $%.2f using PayPal
", amount)
}

type ShoppingCart struct {
    strategy PaymentStrategy
}

func (s *ShoppingCart) SetStrategy(ps PaymentStrategy) {
    s.strategy = ps
}

func (s ShoppingCart) Checkout(amount float64) {
    s.strategy.Pay(amount)
}
