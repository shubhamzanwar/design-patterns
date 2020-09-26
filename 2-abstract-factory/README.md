# Abstract Factory Pattern 🏭

Abstract Factory pattern is a creational design pattern. It is used when we need to create multiple families of similar products.

Let us understand this pattern using the example of a pizza chain

### Pizza Shops

Assume that you are the head of a business that opens up pizza stores around town. One of your many responsibilities is to maintain all food items (in this case, pizzas and garlic bread) in a given store to belong to the same brand - either dominos or pizza hut.

There are multiple ways you could go about it. One of the simplest that comes to mind is to create a factory for creating pizzas of either pizza hut or dominos type and another similar factory for garlic bread.

> Note: If you're not sure about how normal factories work, you can take a look at [this example](https://github.com/shubhamzanwar/design-patterns/tree/master/1-factory)

The drawback of having separate factories is that now we trust the user, at a given store, to choose the correct type of pizza, garlic bread they want. If they ever make the mistake of serving a dominos pizza with a pizza hut garlic bread, your customers are going to be mad and you'd also be violating your make-believe contract with each of these chains.

Don't worry. **_There's a simpler way_**

Instead of creating a factory over each individual product (pizzas and garlic bread), you could create a factory for each individual brand. Enforce both these factories to have provision for `creating a pizza` and `creating a garlic bread`.

Then, while setting up the store, you could give the store manager either a pizza hut factory or a dominos factory and then confidently rest that they won't be accidentally mixing-and-matching products.

Let's translate this to code:

Before we begin writing the factories, let's first create the products:

##### Generic Pizza

```go
type iPizza interface {
    GetPrice() float64
    GetName() string
    GetToppings() []string
}

type pizza struct {
    name     string
    price    float64
    toppings []string
}

func (p *pizza) GetName() string {
    return p.name
}

func (p *pizza) GetPrice() float64 {
    return p.price
}

func (p *pizza) GetToppings() []string {
    return p.toppings
}
```

##### Branded Pizzas

```go
type pizzaHutPizza struct {
    pizza
}

type dominosPizza struct {
    pizza
}
```

##### Generic Garlic Bread

```go
type iGarlicBread interface {
    GetPrice() float64
    GetName() string
}

type garlicBread struct {
    name  string
    price float64
}

func (g *garlicBread) GetName() string {
    return g.name
}

func (g *garlicBread) GetPrice() float64 {
    return g.price
}
```

##### Branded Garlic Bread

```go
type pizzaHutGarlicBread struct {
    garlicBread
}

type dominosGarlicBread struct {
    garlicBread
}
```

We have created both our products. They both implement a common interface, making it easier for the end user to consume them. Pun intended 😉

Let's now look at how we can create the factories for each of these:

##### Generic Factory

```go
type iPizzaFactory interface {
    createPizza() iPizza
    createGarlicBread() iGarlicBread
}
```

Now, both, `PizzaHutFactory` and `DominosFactory` can implement this interface to ensure that they expose uniform functionality

##### Branded Factories

```go
type PizzaHutFactory struct {}

func (p *PizzaHutFactory) createPizza(): iPizza {
    return &pizzaHutPizza{
        pizza{
            name:     "pepperoni",
            price:    230.3,
            toppings: []string{"olives", "mozzarella", "pork"},
        },
    }
}

func (p *pizzaHutFactory) createGarlicBread() iGarlicBread {
    return &pizzaHutGarlicBread{
        garlicBread{
            name:  "garlic bread",
            price: 180.99,
        },
    }
}
```

```go
type dominosFactory struct{}

func (d *dominosFactory) createPizza() iPizza {
    return &dominosPizza{
        pizza{
            name:     "margherita",
            price:    200.5,
            toppings: []string{"tomatoes", "basil", "olive oil"},
        },
    }
}

func (d *dominosFactory) createGarlicBread() iGarlicBread {
    return &dominosGarlicBread{
        garlicBread{
            name:  "cheesy bread sticks",
            price: 150.00,
        },
    }
}
```

We can now choose between either of the factories and then continue creating either pizzas or garlic bread and be absolutely sure that any product from a factory will always belong to the same family/brand of products.

**_We're almost there_**. Let's wrap this up by creating a factory that will return the factory of our choice. Confusing? 🤯 Give it a minute and read that sentence again 😋

Basically, think of our factories as just another object. Now, based on what type or pizza store we want to set up, (pizza hut or dominos), we can make request for that particular factory (just another object). To automatically get these "objects", we can create another factory that would return one of these.

Some code to help you out:

##### Factory of factories

```go
func getPizzaFactory(chain string) (iPizzaFactory, error) {
    if chain == "P" {
        return &pizzaHutFactory{}, nil
    }
    if chain == "D" {
        return &dominosFactory{}, nil
    }
    return nil, fmt.Errorf("Enter a valid chain type next time")
}
```

Maybe this makes more sense 💡😁

The main point to remember is this: The abstract factory pattern implements a factory of factories. These inner factories can be used to create products that belong to a particular kind.
