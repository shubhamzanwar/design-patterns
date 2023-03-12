# 工厂模式 🏭

_**本文最初发表于此 [博客](shubhamzanwar.com/blog)**_

工厂模式是一种常用的创建型设计模式。 用户通常使用它在多个选项中进行选择。让我们举个例子来理解。

### 宠物店

让我们以一个在宠物店如何工作的场景为例。为了完全理解这一点，我们将从店主(创建工厂的“开发人员”)和客户(使用接口的“用户”)两个角度来观察实现方式。

#### 店主视角

假设你是一家狗店的老板(你只把小狗送人领养)。因为你是在软件世界中，每条狗都是你拥有的`Dog `类的一个实例。现在，当客户到来时，您只需创建一个新的`Dog`实例，并让他们领养。🐶

然而，最近顾客开始要求多样化。他们也在寻找收养猫的选择。😼

作为一个聪明的店主，你已经意识到这种需求只会变得越来越多样化。人们将继续期待更多的变化。😨😤

**_你需要一个健壮的、可扩展的系统来为客户生成新的宠物_**
进入，工厂模式

你列出宠物的所有共同特征。它们可以让你知道它们的名字，它们发出的声音和它们的年龄。该列表允许您创建具有以下功能的接口:

```go
type Pet interface {
	GetName() string
	GetAge() int
	GetSound() string
}
```

现在，你可以创建任意数量具有相同功能的宠物(“实现相同的接口”)。你可以养猫，狗，鱼，鹦鹉，任何东西-只要实现`Pet`接口!😯 现在，让我们创建狗和猫:

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

你还需要一个工厂，它会根据用户的请求返回不同的宠物(狗/猫)。简单地说，如果用户想要一只狗，就给他们一只可爱的狗。🙄🦮

```go
func GetPet(petType string) Pet {
	if petType == "dog" {
		return &Dog{
			pet{
				name:  "Chester",
				age:   2,
				sound: "bark",
			},
		}
	}

	if petType == "cat" {
		return &Cat{
			pet{
				name:  "Mr. Buttons",
				age:   3,
				sound: "meow",
			},
		}
	}

	return nil
}
```

注意`GetPet`'函数如何只告诉它返回`Pet`-而不是显式地返回` Dog `或`Cat`。因此，这个函数是开放扩展的(通过编写更多结构图实现 `Pet `接口)。增加更多的`Pet`类型不会影响现有用户只想要`Dog`。恭喜你!你已经使用工厂模式创建了一个宠物商店。🎉❤️

#### 客户视角

让我们从用户的角度来看这个问题。他们所需要做的就是用他们想要的任何配置(在本例中为`type`)调用`GetPet`函数。通过返回值，他们只知道他们得到了一个`pet`。🤔这在现实世界的意义上可能听起来很奇怪，但在代码方面，最好保持抽象。😌

用户可以随心所欲地“使用”`Pet`。不管他们得到的是什么类型的宠物，这种“用法”都是一样的(因为所有的宠物都实现了公共接口!!)

让我们来测试一下

```go
func describePet(pet Pet) string {
    return fmt.Sprintf("%s is %d years old. Its sound is %s", pet.GetName(), pet.GetAge(), pet.GetSound())
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

输出应该如下所示:

```text
Chester is 2 years old. Its sound is bark
-------------
Mr. Buttons is 3 years old. Its sound is meow
```
