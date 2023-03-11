# 抽象工厂模式 🏭

抽象工厂模式是一种创建型的设计模式。当我们需要创建多个类似产品家族时，就会使用它。
让我们以披萨连锁店为例来理解这种模式。

### 披萨店

假设你是一家在镇上开披萨店的老板。其中一项职责是保证某家店的所有食品(在本例中是披萨和蒜蓉面包)属于同一个品牌——要么是达美乐比萨，要么是必胜客。

有很多种方法可以做到。想到的最简单的方法之一是创建一个工厂来生产必胜客或达美乐类型的披萨，以及另一个类似的工厂来生产大蒜面包。

> 注:如果你不确定正常的工厂是如何运作的，你可以去看看[这个例子](https://github.com/coolbook/design-patterns/tree/master/1-1-%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F)

拥有独立工厂的缺点是，现在我们相信，在给定的商店里，用户会选择他们想要的正确类型的披萨和蒜蓉面包。如果他们犯了一个错误，把达美乐披萨和必胜客蒜蓉面包放在一起供应，你的顾客会很生气，你也违反了你与这些连锁店签订的合同。

别担心。 **_有一个更简单的方法_**

你可以为每个单独的品牌创建一个工厂，而不是为每个单独的产品(披萨和大蒜面包)创建一个工厂。强制这些工厂提供“制作披萨”和“制作蒜蓉面包”的条款。

然后，在建立商店时，你可以给商店经理一个必胜客工厂或达美乐工厂，然后自信地相信他们不会不小心混合搭配产品。

让我们把它转换成代码:

在我们开始编写工厂之前，让我们先创建产品:

##### 通用的披萨

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

##### 品牌的披萨

```go
type pizzaHutPizza struct {
    pizza
}

type dominosPizza struct {
    pizza
}
```

##### 通用的蒜蓉面包

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

##### 品牌的蒜蓉面包

```go
type pizzaHutGarlicBread struct {
    garlicBread
}

type dominosGarlicBread struct {
    garlicBread
}
```

我们创造了这两种产品。它们都实现了一个公共接口，使得最终用户更容易使用它们。双关语 😉

现在让我们看看如何为这些产品创建工厂:

##### 通用工厂

```go
type iPizzaFactory interface {
    createPizza() iPizza
    createGarlicBread() iGarlicBread
}
```

现在，`PizzaHutFactory `和`DominosFactory`都可以实现这个接口，以确保它们公开统一的功能

##### 品牌工厂

```go
type PizzaHutFactory struct {}

func (p *PizzaHutFactory) createPizza() iPizza {
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

现在我们可以在这两家工厂中任选一家，然后继续生产披萨或蒜蓉面包，而且绝对可以确保一家工厂生产的任何产品都属于同一个家族/品牌。

**_我们就快完成了_**。让我们通过创建一个将返回我们所选择的工厂的工厂来结束它。困惑吗?🤯花一分钟时间再读一遍这个句子 😋

基本上，把我们的工厂想象成另一个对象。现在，根据我们想要建立的类型或比萨饼店(必胜客或达美乐)，我们可以请求特定的工厂(只是另一个对象)。为了自动获取这些“对象”，我们可以创建另一个工厂，该工厂将返回这些对象中的一个。

一些代码可以帮助你:

##### 工厂的工厂

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

也许这样更有意义 💡😁

要记住的要点是:抽象工厂模式实现了工厂的工厂。这些内部工厂可以用来制造属于特定种类的产品。
