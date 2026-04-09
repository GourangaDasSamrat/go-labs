package main

import (
	"fmt"
)

/*
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Delivered
	Canceled
)
*/

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Delivered OrderStatus = "delivered"
	Canceled  OrderStatus = "canceled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Change order status to ", status)
}

func main() {
	changeOrderStatus(Canceled)
}
