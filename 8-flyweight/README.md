# Design Patterns: Flyweight Pattern ⚖️

The Flyweight design pattern is a structural design pattern commonly used when we want to group similar data for multiple objects. The primary intent of this pattern is to store common data for multiple objects in a single place and hence optimise for memory usage.

Let's understand this pattern by taking the case of a sneaker store 👟

### Sneaker Store

Imagine running an online sneaker store. You've noticed that your customers always have a lot of questions and so, you've maintained every little detail about each pair of sneakers you have on display! 🗃 

You have the following data points on **every** single pair:

- Name
- Size
- Color
- Brand
- Price
- Materials Used
- Description
- Sample Image
- laced

You soon realized that for a given _type_ of pair, you're storing the same information again and again. Take the case of the Nike Air Force 1. For any two pairs, the only thing different is the `size` and the `color`. Everything else is duplicate information 😱

This definitely isn't the best way to store the details because it's super memory intensive ad costly. To solve this problem, you decide to create an object to store all the _common_ details of the shoe and store only the `size`, `color` and a reference to the "common object" with every shoe instance.

If all of this is confusing, don't worry. Taking a look at the code should make things easier 😉

```go
type SneakerDetails struct {
	name        string
	brand       string
	price       float64
	materials   []string
	description string
	image       string
	laced       bool
}

var NikeAirForceDetails = SneakerDetails{
	name:        "Nike Air Force 1",
	brand:       "Nike",
	price:       144.99,
	materials:   []string{"leather", "rubber"},
	description: "Originally released in 1982, the Nike Air Force 1 was the first Nike model to feature Air technology.",
	image:       "https://via.placeholder.com/50",
	laced:       true,
}

var AdidasSuperstarDetails = SneakerDetails{
	name:        "Adidas Superstar",
	brand:       "Adidas",
	price:       85.99,
	materials:   []string{"leather", "rubber"},
	description: "Originally made for basketball courts in the '70s.",
	image:       "https://via.placeholder.com/50",
	laced:       true,
}
```

```go
const (
	NikeAirForce = "Nike Air Force 1"
	AdidasSuperstar = "Adidas Superstar"
)

var DetailsMap = map[string]SneakerDetails{
	NikeAirForce:    NikeAirForceDetails,
	AdidasSuperstar: AdidasSuperstarDetails,
}

func getShoeDetails(shoeType string) SneakerDetails {
	return DetailsMap[shoeType]
}
```

```go
type Sneaker struct {
    shoeType string
    size int
    color string
}

func (s *Sneaker) describe() string {
    details := getShoeDetails(s.shoeType)
    return fmt.Sprintf("The %s is of size %d, color %s and costs %.2f. The Manufacturer is %s", details.name, s.size, s.color, details.price, details.brand)
}
```

Looking through the code written so far, we've created the details object for each type of shoe in our shop using a common struct. This way, we can add descriptions for more sneakers in the future 🤔

We've also created a map of the shoe type against the respective type. This way, every instance of the `Sneaker` struct can easily access the corresponding details.

Also notice that our `Sneaker` struct now only contains the `size`, `color` along with the `shoeType` that allows us to fetch it's other details.

Let's see this in action

```go
func main() {
	sneaker1 := Sneaker{
		shoeType: NikeAirForce,
		size:     9,
		color:    "white",
	}

	sneaker2 := Sneaker{
		shoeType: AdidasSuperstar,
		size:     8,
		color:    "black",
	}

	fmt.Println(sneaker1.describe())
	fmt.Println(sneaker2.describe())
}
```

This should give the following output:

```text
The Nike Air Force 1 is of size 9, color white and costs 144.99. The Manufacturer is Nike
The Adidas Superstar is of size 8, color black and costs 85.99. The Manufacturer is Adidas
```

That's about it for the Flyweight pattern! 😁

Note that one added benefit of this pattern (apart from memory efficiency) is also the fact that updating the common details, like price, for a type of sneaker would require you to make the change in only one place - the common details object as opposed to making the change in every individual instance of the sneaker
