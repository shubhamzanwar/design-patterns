package main

import "fmt"

type orderFacade struct {
	userService         UserService
	productService      ProductService
	paymentService      PaymentService
	notificationService NotificationService
}

func (o *orderFacade) placeOrder(userID string, productID string) {
	fmt.Println("[Facade] Starting order placement")

	userValid := o.userService.isUserValid(userID)
	productAvailable := o.productService.productAvailable(productID)

	if userValid && productAvailable {
		o.productService.assignProductToUser(productID, userID)
		o.paymentService.makePayment(userID, productID)
		o.notificationService.notifyDealer(productID)
	}
}
