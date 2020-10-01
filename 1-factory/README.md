# Factory Pattern 🏭

_**This post was originally published at this [blog](shubhamzanwar.com/blog)**_

Factory pattern is a commonly used creational design pattern. It is normally used when the user is expected to choose between multiple options.

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

Now, you can create any number of pets that have the same features (`implement the same interface`). You could have cats, dogs, fish, parrots, anything - as long as the implement the `Pet` interface!😯 As of now, let's create the Dog and the Cat:

```go
// pet is a struct that implements Pet interface and
// would be used in any animal struct that we create.
// See `Dog` and `Cat` below
type pet struct {
    name  string
    age   int
    sound string
}

func (p *pet) GetName() string {
    return p.name
}

func (p *pet) GetSound() string {
    return p.sound
}

func (p *pet) GetAge() int {
    return p.age
}

type Dog struct {
    pet
}

type Cat struct {
    pet
}
```

An additional thing you would require is a factory that would return a different pet (dog/cat) based on the user's request. Simply put, if the user asks for a dog, give them a cute doggo, duh.🙄🦮

```go
func GetPet(type string) Pet {
    if type == "dog" {
        return &Dog{
            pet{
                name:  "Chester",
                age:   2,
                sound: "bark",
            },
        }
    }
    if type === "cat" {
        return &Cat{
            pet{
                name:  "Mr. Buttons",
                age:   3,
                sound: "meow",
            },
        }
    }
}
```

Notice how the `GetPet` function only tells that it returns a `Pet` - Not a `Dog` or a `Cat` explicitly. Hence, this function is open to extension (by writing more structs that implement the `Pet` interface). Adding more `Pet` types will not affect the existing users who just want `Dog`s anyway.

Congratulations! you've created a Pet shop using the factory pattern🎉❤️

#### Customer's perspective

Let's look at it from the perspective of the user. All they would need to do is call the `GetPet` function with whatever config (in this case, `type`) they want. In return, all they know is that they are getting a `Pet`.🤔 This may sound bizarre in the real world sense of things, but when it comes to code, it's better to maintain abstraction.😌

The users can freely go about "using" the `Pet` as they like. This "usage" would remain the same irrespective of what type of pet they got back (because all pets implement the common interface!!)

Let's test it out

```go
func describePet(pet Pet) string {
    return fmt.Sprintf("%s is %d years old. It's sound is %s", pet.GetName(), pet.GetAge(), pet.GetSound())
}

func main() {
    petType := "dog"

    dog := GetPet(petType)
    petDescription := describePet(dog)

    fmt.Println(petDescription)
    fmt.Println("-------------")

    petType = "cat"
    cat := GetPet(petType)
    petDescription = describePet(cat)

    fmt.Println(petDescription)
}
```

The output should look like:

```text
Chester is 2 years old. It's sound is bark
-------------
Mr. Buttons is 3 years old. It's sound is meow
```
