# Builder Pattern 👷🏽‍♂️

The builder pattern is a creational design pattern. It is used when the creation of a product involves multiple, independent steps.

What's the one thing that immediately comes to your mind when you think of something that requires multiple steps to be "built"? That's right, subway!! 🥙🤤
You choose your bread, you choose if you want cheese, you choose your toppings and finally your sauces!

I'm sure you understand the process of creating a sub. Let's see what it would look like if we coded it using the Builder pattern:

### Subway

Let us assume that we have 2 different types(recipes) for a sub - Veggie Delight and Chicken Teriyaki. Each of these would have a "Builder" dedicated to them.

The interface that these builders implement will be the same - That way, we can easily add more types(recipes) to our menu.

```go
type sub struct {
    bread string
    hasCheese bool
    toppings []string
    sauces []string
}
```

```go
type iSubBuilder interface {
    setBread()
    setCheese()
    setToppings()
    setSauces()
    getSub() sub
}
```

```go
type veggieDelightBuilder struct {
    sub
}

func (v *veggieDelightBuilder) setBread() {
    v.sub.bread = "parmesan oregano"
}

func (v *veggieDelightBuilder) setCheese() {
    v.sub.hasCheese = false
}

func (v *veggieDelightBuilder) setToppings() {
    v.sub.toppings = []string{"olives", "tomatoes", "onions", "jalapeños"}
}

func (v *veggieDelightBuilder) setSauces() {
    v.sub.sauces = []string{"south west"}
}

func (v *veggieDelightBuilder) getSub() sub {
    return v.sub
}
```

```go
type chickenTeriyakiBuilder struct {
    sub
}

func (c *chickenTeriyakiBuilder) setBread() {
    c.sub.bread = "italian"
}

func (c *chickenTeriyakiBuilder) setCheese() {
    c.sub.hasCheese = true
}

func (c *chickenTeriyakiBuilder) setToppings() {
    c.sub.toppings = []string{"roasted chicken", "olives", "onions", "jalapeños"}
}

func (c *chickenTeriyakiBuilder) setSauces() {
    c.sub.sauces = []string{"chilli", "bbq"}
}

func (c *chickenTeriyakiBuilder) getSub() sub {
    return c.sub
}
```

We've successfully created the sub and the builders. As you may have observed, we can easily create another recipe and have it implement all the functions.

Now, let's create an optional director struct that would accept accept a builder and then build subs for us.

>The director is not always part of the builder pattern. You could look at it as an added layer of abstraction for cleaner code

```go
type director struct {
    builder iSubBuilder
}

func (d *director) setBuilder(builder iSubBuilder) {
    d.builder = builder
}

func (d *director) buildSub() sub {
    d.builder.setBread()
    d.builder.setCheese()
    d.builder.setToppings()
    d.builder.setSauces()

    return d.builder.getSub()
}
```

**_Aaand we're done_** 🎉

Let's see it in action:

```go
func describeSub(sub sub) {
    fmt.Printf("bread: %s, cheese: %t, toppings: %s, sauces: %s\n", sub.bread, sub.hasCheese, strings.Join(sub.toppings, ", "), strings.Join(sub.sauces, ", "))
}

func main() {
    veggieDelight := &veggieDelightBuilder{}
    director := &director {
        builder: veggieDelight,
    }
    veggieDelightSub := director.buildSub()
    describeSub(veggieDelightSub)

    fmt.Println("------------")

    director.setBuilder(&chickenTeriyakiBuilder{})
    chickenTeriyakiSub := director.buildSub()

    describeSub(chickenTeriyakiSub)
}
```

Running this program should get you this in your terminal

```
bread: parmesan oregano, cheese: false, toppings: olives, tomatoes, onions, jalapeños, sauces: south west
------------
bread: italian, cheese: true, toppings: roasted chicken, olives, onions, jalapeños, sauces: chilli, bbq
```