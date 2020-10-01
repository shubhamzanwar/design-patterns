# Bridge Pattern 🌉

The bridge is a structural design pattern that helps in splitting up a complex piece of code into multiple parts so that each part can be developed, scaled and maintained independently.

You will commonly see this employed when you have classes in your code that have multiple tangential functionalities. In such a case, it is better practice to split the functionalities up into smaller classes and build a mechanism (_a bridge, maybe?_) to connect them

Let us take a look at an example to understand better

### Employee Management system

Assume that your workplace needs an employee management system. As part of this requirement, you need to categorize your workforce by gender and their function/department.

Currently, you have the following genders and department:

**Genders**

- Male
- Female

**Department**

- Developer
- Product manager (PM)
- Human Resources (HR)

A naive way of approaching this problem would be to implement multiple structs, like, `MaleDeveloper`, `FemaleDeveloper`, `MalePM`, `FemalePM`, and so on. This is fine if we have already defined a set of genders and departments. However, that's never realistically the case.

We can always grow departments to include more functions and we can always onboard employees who identify as non-binary (let's not be ignorant 🏳️‍🌈). The problem that would arise then is that, for every new gender we onboard, we would have to create 3 new structs - one for each combination of the new gender and existing departments.

This number would only increase as our organization grows and it's going to be a nightmare to maintain these structs 😨

> Note: Go does not have classes. Structs are a way of implementing classes in golang. If you're working with another programming language, think `classes` every time I say `structs`

Enter, the Bridge pattern. According to the Bridge pattern, we should separate the different dimensions of a struct into separate sub structs. In this case, we would have separate structs for each gender and separate structs for each department.

Let's take a look at this:

```go
type Gender interface {
    describePerson() string
}

type Department interface {
    getDepartmentName() string
}
```

Now that we have the two interfaces defined, we can expect any future departments/genders to implement them. Currently, though, let's implement the existing struct we know using these interfaces:

```go
// male.go

type Male struct {
    name string
    age int
    department Department
}

func (m *Male) describePerson() string {
    description := fmt.Sprintf("%s is a %d years old man\n", m.name, m.age)
    description  = description + fmt.Sprintf("He works in the %s department", m.department.getDepartmentName())
    return description
}
```

```go
// female.go

type Female struct {
    name string
    age int
    department Department
}

func (f *Female) describePerson() string {
    description := fmt.Sprintf("%s is a %d years old woman\n", f.name, f.age)
    description  = description + fmt.Sprintf("She works in the %s department", f.department.getDepartmentName())
    return description
}
```

```go
// departments.go

type Developer struct { }
type PM struct { }
type HR struct { }

func(d Developer) getDepartmentName() string {
    return "software development"
}

func (pm PM) getDepartmentName() string {
    return "product management"
}

func (hr HR) getDepartmentName() string {
    return "human resources"
}
```

As you can see, we have successfully separated the functionalities by gender and department. Notice however, that we have a `department` key in both the gender structs. This key is, in fact, the **_bridge_** between the two dimensions, gender and department.

Now, every time we need to create a new gender or department, we only have to create one additional class and the bridge takes care of the rest.

Time to test our code out 🧪

```go
//main.go

func main() {
    maleDeveloper := Male{
        name: "John",
        age: "22",
        department: Developer{},
    }

    fmt.Println(maleDeveloper.describePerson())
    fmt.Println("-------------")

    femalePM := Female{
        name: "Natalie",
        age: "24",
        department: PM{},
    }

    fmt.Println(femalePM.describePerson())
}
```

The output of this program should look like this

```text
John is a 22 years old man
He works in the software development department
-------------
Natalie is a 24 years old woman
She works in the product management department
```

That's the bridge pattern for you 😉
