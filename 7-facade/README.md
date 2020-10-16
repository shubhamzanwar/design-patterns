# Design Patterns: Facade Pattern 🎭

The facade pattern is a structural design pattern, commonly used when there is some interaction with a complex external library or service.

In this pattern, we create a class to encapsulate the interactions with the 3rd party library. This class is called the Facade. This allows us to expose a simpler interface to the rest of our application. Developers working on other parts of our system need not learn the 3rd party library; rather, they need to only interface with the facade class we create.

This serves a few crucial purposes:

1. In case the 3rd party system is super complex and requires multiple steps to perform a single action, we can encapsulate all of that in a single function in our facade.
2. We can easily switch out one 3rd party library for another and still maintain the same interface for our facade. This flexibility is one of the main selling points for the facade.
3. We can also filter out any extra functionality from the external service that we don't really require in our application. This is very common while using the facade pattern because external libraries can be very comprehensive and we may not require all the features offered.

Let's take an example of using the facade pattern

### Amazon order system

Let's picture that you're building the module that deals with the placement of orders at Amazon. While building this, you have realized that there are several steps to placing an order; each step involving interactions with several other systems/modules.

> Note: This probably is not how Amazon has implemented it's order system. This is just a simplistic representation of what it _could be_

Let's list out the steps:

1. Authenticate the user - Users service
2. Verify product's availability - Products service
3. Assign product to user - Products service
4. Make payment - Payments service
5. Send notification to dealer - Notification service

All of that on the click of a button 😱

Normally, if you were to directly integrate all the systems and services on the button's click, it would hamper the code's readability and would also make it a nightmare to switch one service for another (Imagine a case where you would revamp the products service)

To solve for all this, we can create a facade class that deals with all this and exposes a function called `placeOrder` that takes a `productId` and `userId` to complete all the above steps 👆🏽

Let's implement this solution in go

```go
type orderFacade struct {
    userService UserService
    productService ProductService
    paymentService PaymentService
    notificationService NotificationService
}

func (o *orderFacade) placeOrder(userId string, productId string) {
    fmt.Println("[Facade] Starting order placement")

    userValid := o.userService.isUserValid(userId)
    productAvailable := o.productService.productAvailable(productId)

    if userValid && productAvailable {
        o.productService.assignProductToUser(productId, userId)
        o.paymentService.makePayment(userId, productId)
        o.notificationService.notifyDealer(productId)
    }
}
```

Let's also take a look at the other services' code

```go
type UserService struct {}

func (u *UserService) isUserValid(userId string) bool {
    fmt.Println("[UserService] validating the user: ", userId);
    // Complex logic for checking validity
    return true
}

type ProductService struct {}

func (p *ProductService) productAvailable(productId string) bool {
    fmt.Println("[ProductService] checking availability of product: ", productId)
    // Complex logic for checking availability
    return true;
}

func (p *ProductService) assignProductToUser(productId string, userId string) {
    fmt.Printf("[ProductService] assigning product %s to user %s\n", productId, userId)
    // complex logic for product assignment
}

type PaymentService struct {}

func (p *PaymentService) makePayment(userId string, productId string) {
    fmt.Printf("[PaymentService] charging user %s for product %s\n", userId, productId)
    // complex logic for making payment
}

type NotificationService struct {}

func (n *NotificationService) notifyDealer(productId string) {
    fmt.Printf("[NotificationService] notifying dealer about sale of product %s\n", productId)
    // complex notification logic
}
```

Oooh, so now we have our dummy services and completely legit facade in place 💁🏻‍♂️ - It's time to test it out

```go
func main() {
    orderModule := &orderFacade{
        userService:  UserService{},
        productService:  ProductService{},
        paymentService:  PaymentService{},
        notificationService:  NotificationService{},
    }

    userId := "test-user-id"
    productId := "test-product-id"


    orderModule.placeOrder(userId, productId)
}
```

When you run your program, you should see something like this in the output:

```text
[Facade] Starting order placement
[UserService] validating the user:  test-user-id
[ProductService] checking availability of product:  test-product-id
[ProductService] assigning product test-product-id to user test-user-id
[PaymentService] charging user test-user-id for product test-product-id
[NotificationService] notifying dealer about sale of product test-product-id
```

You can notice from the code and the output that the facade only is the surface layer of the functionality. The actual action may happen in the external services and it is still abstracted away from the user of the facade by exposing a super simple interface.

That's the facade pattern in a gist 😁
