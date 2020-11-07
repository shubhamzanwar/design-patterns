# Design Patterns: Chain of Responsibility ⛓

The Chain of responsibility pattern is a behavioral design pattern that deals with breaking down a complex problem into sequentially occurring steps. At each step, there is a possibility of moving to the next step or exiting from the chain.

To understand this better, let's understand how a basic call center would work. Think of all the stages you'd go through when you ring up the customer care of an electronics store to get some help regarding a new product you bought.

### Customer Care

A generic customer care you follow the following steps:

1. The user would start off by interacting with an automated voice assistant. The user can choose to continue or drop off.
2. Next, the user would be redirected to the call center associate. Based on the query, the associate can choose to resolve it or progress the user to the next level.
3. The final stage; the user is redirected to the manager - who can take care of almost anything.

**There's also a special case, if the customer is a high priority customer, the system skips the conversation with the associate and directly connects the customer to the manager.**

Let's imagine how we would code this. 🤔 A naïve approach would be to have a single function that the user enters and then based on the user's status and interactions, different sub-functions (voice assistant, associate or manager) are called.

However, as our system grows and we want to add more features, this approach would make the function difficult to manage and the sub-functions difficult to reuse. More importantly, this approach **violates the SOLID programming principle of separating concerns**. 🤕

As an alternative, we could design our sub-functions in such a way that they would immediately call the next step in the chain as soon as their execution is over. In that way, the functionality for every step is maintained in a separate function but also it's easier to swap pieces out from the chain/insert more steps in the chain 🔁

Let's check out the implementation for this in go:

```go
type step interface {
    run(*customer)
    setNextStep(step)
}
```

```go
type customer struct {
    name string
    isHighPriority bool
}
```

```go
type voiceAssistant struct {
    next step
}

func (v *voiceAssistant) run(cust *customer) {
    fmt.Println("[Voice Assistant] Serving the customer: ", cust.name)
    v.next.run(cust)
}

func (v *voiceAssistant) setNextStep(next step) {
    v.next = next
}
```

```go
type associate struct {
    next step
}

func (a *associate) run(cust *customer) {
    if cust.isHighPriority {
        fmt.Println("Redirecting customer directly to manager")
        a.next.run(cust)
        return
    }
    fmt.Println("[Associate] Serving the customer: ", cust.name)
    a.next.run(cust)
}

func (a *associate) setNextStep(next step) {
    a.next = next
}
```

```go
type manager struct {
    next step
}

func (a *manager) run(cust *customer) {
    fmt.Println("[Manager] Serving the customer: ", cust.name)
}

func (a *manager) setNextStep(next step) {
    a.next = next
}
```

At this point, each link of our chain is now created. We need to arrange them properly so that they follow the desired workflow: `Voice Assistant -> Associate(optional) -> Manager`. Let's do this in our main function for two different customers and see the result

```go
func main() {
    m := &manager{}

    assoc := &associate{};
    assoc.setNextStep(m)

    va := &voiceAssistant{};
    va.setNextStep(assoc);

    // Chain formation complete

    // Start chain execution for normal customer
    normalCust := &customer{
        name: "Bob"
    }

    va.run(normalCust);

    fmt.Println("===================")

    // Start chain execution for high priority customer
    highPriorityCust := &customer{
        name: "John",
        isHighPriority: true,
    }

    va.run(highPriorityCust)
}
```

This should give you the following output:

```text
[Voice Assistant] Serving the customer:  Bob
[Associate] Serving the customer:  Bob
[Manager] Serving the customer:  Bob
===================
[Voice Assistant] Serving the customer:  John
Redirecting customer directly to manager
[Manager] Serving the customer:  John
```

That's the Chain of Responsibility pattern for you! 😁
