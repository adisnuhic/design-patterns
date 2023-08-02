/*
	Consider a payment processing application that supports multiple payment methods,
	such as credit cards, PayPal, and cryptocurrencies. Each payment method has its own algorithm for processing payments.
*/

package main

import "fmt"

// PaymentStrategy is the interface for different payment methods.
type PaymentStrategy interface {
	ProcessPayment(amount float64) error
}

// CreditCardPayment represents the payment processing using a credit card.
type CreditCardPayment struct{}

func (c *CreditCardPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment for amount $%.2f\n", amount)
	// Simulate processing payment using a credit card.
	// In a real implementation, it would call the actual credit card payment API.
	return nil
}

// PayPalPayment represents the payment processing using PayPal.
type PayPalPayment struct{}

func (p *PayPalPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment for amount $%.2f\n", amount)
	// Simulate processing payment using PayPal.
	// In a real implementation, it would call the actual PayPal API.
	return nil
}

// CryptocurrencyPayment represents the payment processing using cryptocurrencies.
type CryptocurrencyPayment struct{}

func (c *CryptocurrencyPayment) ProcessPayment(amount float64) error {
	fmt.Printf("Processing cryptocurrency payment for amount $%.2f\n", amount)
	// Simulate processing payment using cryptocurrencies.
	// In a real implementation, it would call the actual cryptocurrency payment API.
	return nil
}

// PaymentProcessor is the context that holds the selected payment strategy.
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) SetPaymentStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) ProcessPayment(amount float64) error {
	if p.strategy == nil {
		return fmt.Errorf("payment strategy not set")
	}
	return p.strategy.ProcessPayment(amount)
}

func main() {
	// Client code interacts with the PaymentProcessor and selects a payment strategy dynamically.
	processor := &PaymentProcessor{}

	// Process payment using a credit card.
	processor.SetPaymentStrategy(&CreditCardPayment{})
	processor.ProcessPayment(100.50)

	// Process payment using PayPal.
	processor.SetPaymentStrategy(&PayPalPayment{})
	processor.ProcessPayment(50.75)

	// Process payment using cryptocurrencies.
	processor.SetPaymentStrategy(&CryptocurrencyPayment{})
	processor.ProcessPayment(75.20)
}
