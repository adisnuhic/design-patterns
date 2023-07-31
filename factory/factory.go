package main

import "fmt"

// Payment is the interface for different payment methods.
type Payment interface {
	Pay(amount float64) string
}

// CreditCard is a concrete payment method (Product).
type CreditCard struct{}

// Pay implements the Pay method for CreditCard.
func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f with Credit Card", amount)
}

// PayPal is a concrete payment method (Product).
type PayPal struct{}

// Pay implements the Pay method for PayPal.
func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f with PayPal", amount)
}

// PaymentFactory is the Factory interface.
type PaymentFactory interface {
	CreatePayment() Payment
}

// CreditCardFactory is a concrete Factory that produces CreditCard instances.
type CreditCardFactory struct{}

// CreatePayment creates a new CreditCard instance (Product).
func (ccf *CreditCardFactory) CreatePayment() Payment {
	return &CreditCard{}
}

// PayPalFactory is a concrete Factory that produces PayPal instances.
type PayPalFactory struct{}

// CreatePayment creates a new PayPal instance (Product).
func (ppf *PayPalFactory) CreatePayment() Payment {
	return &PayPal{}
}

func main() {
	// Let's use the factory to create different payment methods.

	// Create a CreditCardFactory
	creditCardFactory := &CreditCardFactory{}

	// Use the factory to create a CreditCard payment
	creditCardPayment := creditCardFactory.CreatePayment()
	fmt.Println(creditCardPayment.Pay(100.50))

	// Create a PayPalFactory
	payPalFactory := &PayPalFactory{}

	// Use the factory to create a PayPal payment
	payPalPayment := payPalFactory.CreatePayment()
	fmt.Println(payPalPayment.Pay(50.75))
}
