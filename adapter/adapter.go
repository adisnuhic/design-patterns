/*
	Let's consider a real-world example where we have two third-party payment gateways, PayPal and Stripe.
	Each payment gateway has its own interface for processing payments. We want to integrate both payment gateways
	into our application but present a common interface for processing payments.
	We can use the Adapter pattern to achieve this.
*/

package main

import (
	"fmt"
)

// PaymentGateway is the common interface expected by the payment application.
type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

// PayPalGateway is the PayPal payment gateway with its specific API.
type PayPalGateway struct{}

// MakePayment is the method for processing payment using PayPal's API.
func (pg *PayPalGateway) MakePayment(amount float64) error {
	// Simulate processing payment with PayPal.
	// In a real implementation, it would call the actual PayPal API.
	fmt.Printf("Processing payment of $%.2f with PayPal\n", amount)
	return nil
}

// StripeGateway is the Stripe payment gateway with its specific API.
type StripeGateway struct{}

// Charge is the method for processing payment using Stripe's API.
func (sg *StripeGateway) Charge(amount float64) error {
	// Simulate processing payment with Stripe.
	// In a real implementation, it would call the actual Stripe API.
	fmt.Printf("Charging payment of $%.2f with Stripe\n", amount)
	return nil
}

// PayPalAdapter adapts PayPalGateway to the common PaymentGateway interface.
type PayPalAdapter struct {
	PayPal *PayPalGateway
}

// ProcessPayment implements the common PaymentGateway interface using PayPalGateway's API.
func (pa *PayPalAdapter) ProcessPayment(amount float64) error {
	if pa.PayPal == nil {
		pa.PayPal = &PayPalGateway{}
	}
	return pa.PayPal.MakePayment(amount)
}

// StripeAdapter adapts StripeGateway to the common PaymentGateway interface.
type StripeAdapter struct {
	Stripe *StripeGateway
}

// ProcessPayment implements the common PaymentGateway interface using StripeGateway's API.
func (sa *StripeAdapter) ProcessPayment(amount float64) error {
	if sa.Stripe == nil {
		sa.Stripe = &StripeGateway{}
	}
	return sa.Stripe.Charge(amount)
}

func main() {
	// Create the payment application using the common PaymentGateway interface.
	paymentApp := &PaymentApplication{}

	// Use PayPal via the PayPalAdapter.
	paymentApp.SetPaymentGateway(&PayPalAdapter{})
	paymentApp.ProcessPayment(100.50)

	// Use Stripe via the StripeAdapter.
	paymentApp.SetPaymentGateway(&StripeAdapter{})
	paymentApp.ProcessPayment(50.75)
}

// PaymentApplication is the payment application.
type PaymentApplication struct {
	paymentGateway PaymentGateway
}

// SetPaymentGateway sets the payment gateway to be used by the application.
func (app *PaymentApplication) SetPaymentGateway(pg PaymentGateway) {
	app.paymentGateway = pg
}

// ProcessPayment processes the payment for the specified amount.
func (app *PaymentApplication) ProcessPayment(amount float64) {
	if app.paymentGateway == nil {
		fmt.Println("Payment gateway not set")
		return
	}
	err := app.paymentGateway.ProcessPayment(amount)
	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}
