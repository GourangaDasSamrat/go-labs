package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) { fmt.Println("Making payment using razorpay, ", amount) }

type stripe struct{}

func (r stripe) pay(amount float32) { fmt.Println("Making payment using stripe, ", amount) }

func main() {
	// stripeGateway := stripe{}
	razorpayGateway := razorpay{}
	newPayment := payment{
		gateway: razorpayGateway,
	}
	newPayment.makePayment(2000)
}
