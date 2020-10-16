package main

import "fmt"

// UserService is a dummy service to mck user journeys
type UserService struct{}

func (u *UserService) isUserValid(userID string) bool {
	fmt.Println("[UserService] validating the user: ", userID)
	// Complex logic for checking validity
	return true
}

// ProductService is a dummy service to mck user journeys
type ProductService struct{}

func (p *ProductService) productAvailable(productID string) bool {
	fmt.Println("[ProductService] checking availability of product: ", productID)
	// Complex logic for checking availability
	return true
}

func (p *ProductService) assignProductToUser(productID string, userID string) {
	fmt.Printf("[ProductService] assigning product %s to user %s\n", productID, userID)
	// complex logic for product assignment
}

// PaymentService is a dummy service to mck user journeys
type PaymentService struct{}

func (p *PaymentService) makePayment(userID string, productID string) {
	fmt.Printf("[PaymentService] charging user %s for product %s\n", userID, productID)
	// complex logic for making payment
}

// NotificationService is a dummy service to mck user journeys
type NotificationService struct{}

func (n *NotificationService) notifyDealer(productID string) {
	fmt.Printf("[NotificationService] notifying dealer about sale of product %s\n", productID)
	// complex notification logic
}
