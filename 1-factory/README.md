# Factory Pattern 🏭

Factory pattern in a commonly used creational design pattern. It is normally used when the user is expected to choose between multiple options.

Let's take an example to understand.

### Pet store

Let's take a scenario of how this would work in a pet store. To understand this completely, we'll look at the implementation from both, the view of the shop-owner(`developer` creating the factory) and the customer (`user` using the interface)

#### Owner's perspective

Assume that you're the owner of a dog store (you only put little doggos up for adoption). Since you're in the software world, every dog is an instance of a `Dog` class that you have. Now, when a customer arrives, you simply create a new instance of `Dog` and let them adopt that.🐶

Lately, however, the customers have started requesting for variety. They're looking for options to adopt cats too.😼

Being a clever shop-owner, you've identified that this demand can only keep getting more varied. People will continue expecting more variety.😨😤

**_You need a robust, scalable system to generate new pets for customers_**
Enter, the factory pattern

You make a list of all the common traits (`features`) of your pets. They allow you to get the name, get the sound they make and get their age. This list allows you to create an interface with the following functions:

```go
type Pet interface {
    GetName() string
    GetAge() string
    GetSound() string
}
```

Now, you can create any number of pets that have the same features (`implement the same interface`). You could have cats, dogs, fish, parrots, anything - as long as the implement the `Pet` interface!😯

An additional thing you would require is a factory that would return a different pet (dog/cat) based on the user's request. Simply put, if the user asks for a dog, give them a cute doggo, duh.🙄🦮

```go
func GetPet(type string) Pet {
    if type == "dog" {
        // return a new Dog
    }
    if type === "cat" {
        // return a new Cat
    }
}
```

Notice how the `GetPet` function only tells that it returns a `Pet` - Not a `Dog` or a `Cat` explicitly. Hence, this function is open to extension (by writing more structs that implement the `Pet` interface). Adding more `Pet` types will not affect the existing users who just want `Dog`s anyway.

Congratulations! you've created a Pet shop using the factory pattern🎉❤️

#### Customer's perspective

Let's look at it from the perspective of the user. All they would need to do is call the `GetPet` function with whatever config (in this case, `type`) they want. In return, all they know is that they are getting a `Pet`.🤔 This may sound bizarre in the real world sense of things, but when it comes to code, it's better to maintain abstraction.😌

The users can freely go about "using" the `Pet` as they like. This "usage" would remain the same irrespective of what type of pet they got back (because all pets implement the common interface!!)
