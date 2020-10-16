package main

func main() {
	orderModule := &orderFacade{
		userService:         UserService{},
		productService:      ProductService{},
		paymentService:      PaymentService{},
		notificationService: NotificationService{},
	}

	userID := "test-user-id"
	productID := "test-product-id"

	orderModule.placeOrder(userID, productID)
}
