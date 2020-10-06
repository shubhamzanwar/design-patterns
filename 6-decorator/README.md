# Decorator Pattern 📦

The Decorator pattern is a structural design pattern that is used when we need to _enhance_ the functionality of a given object _without changing it's interface._

In this pattern, we usually have a base/bare-bones class that implements a certain interface. Along with this, we have multiple "wrapper" classes that allow users to enhance the methods of the base class without changing the interface.

_This means that even the "Wrapper" classes implement the same interface_

Let's consider an example to understand this better. We will discuss about a usecase that requires storing user information to a storage system and what happens when we need to encrypt some data before storing. 🗃

### Storage system

Assume that you have a storage system for your websites audience that stores the users' data: `email`, in our case. Being a good developer, you had designed this module to be super generic and it is now being used by multiple teams in your organization. 😎

Everything sounds great? Here's the problem. You have a new requirement - Only your team needs to encrypt the users' data before storing it. 🤯

You obviously cannot modify the existing module. Too many teams are using it and they'll kill you if you do 🙃. You also cannot re-write everything from scratch. You don't have enough time for that and it also isn't the best way of solving this.

You remember a colleague telling you about the decorator pattern and you decide to put it into action 💡. Here's how it goes:

1. You have an interface, `StorageSystem` that contains two methods, `writeInfo` and `readInfo`.
2. The struct you had previously created by implementing `StorageSystem` is `GenericStorageSystem` - Assume that this just stores the information in memory.
3. To fulfill the requirement, you still need some struct to implement `StorageSystem` while also encrypting the data. Yes, I repeat, you can't just create a new struct from scratch. You need to use `GenericStorageSystem`.
4. You decide to create `EncryptedStorageSystem`. It wraps over the `GenericStorageSystem` and implements the `writeInfo` and `readInfo` which call `writeInfo` and `readInfo` of `GenericStorageSystem` respectively _after_ encrypting/decrypting the data.

Voila! You've saved a bunch time and also have a scalable solution. Here, have a cookie for your efforts 🍪

Let's see how this looks in Go code:

```go
type StorageSystem interface {
    writeInfo(info string)
    readInfo() string
}

type GenericStorageSystem struct {
    info string
}

func (g *GenericStorageSystem) writeInfo(info string) {
    // write information to file system as is.
    fmt.Println("Writing info to the file system: ", info)
    g.info = info
}

func (g *GenericStorageSystem) readInfo() string {
    // Return info from file system as is.
    fmt.Println("Reading info from the file system: ", g.info)
    return g.info
}
```

```go
type EncryptedStorageSystem struct {
    storageSystem StorageSystem
}

func (e *EncryptedStorageSystem) writeInfo(info string) {
    fmt.Println("Encrypting the data")
    encryptedInfo := encrypt(info)
    e.storageSystem.writeInfo(encryptedInfo)
}

func (e *EncryptedStorageSystem) readInfo() string {
    info := e.storageSystem.readInfo()
    decryptedInfo := decrypt(info)
    fmt.Println("Decrypting info from the file system: ", decryptedInfo)
    return decryptedInfo
}
```

Let's also create dummy `encrypt` and `decrypt` functions. These just rotate the string in different directions

```go
func encrypt (info string) string {
    return info[1:len(info)] + string(info[0])
}


func decrypt (info string) string {
    return string(info[len(info) - 1]) + info[0:len(info)-1]
}
```

Aaaand we're done 💃🏻. Let's check it out in action:

```go
func main() {
    info := "kingslayer@gmail.com"
    genericStorage := GenericStorageSystem{}

    genericStorage.writeInfo(info)
    genericStorage.readInfo()

    fmt.Println("------------")

    encryptedStorage = EncryptedStorage{
        storageSystem: genericStorage,
    }

    encryptedStorage.writeInfo(info)
    encryptedStorage.readInfo()
}
```

You should see this in the output:

```text
Writing info to the file system:  kingslayer@gmail.com
Reading info from the file system:  kingslayer@gmail.com
------------
Encrypting the data
Writing info to the file system:  ingslayer@gmail.comk
Reading info from the file system:  ingslayer@gmail.comk
Decrypting info from the file system:  kingslayer@gmail.com
```

Notice how the end result in both the cases are similar. You wrote a string, and got back the same string when you requested it. However, in the second case, the data was encrypted while storing and decrypted while reading - The most important part is that all of this is abstracted from the user.

There you have it - The decorator pattern 🎉

Now that you have a good understanding of the decorator pattern, let's discuss why it's a good thing:

- Notice how you did not have to create a new storage from scratch. This saved a bunch of time and also didn't change the API for your storage system. Hence, no unhappy users.
- Also notice that the `EncryptedStorageSystem` implements the `StorageSystem` interface and wraps around any `StorageSystem`; not just `GenericStorageSystem`. This is pretty powerful because you will be able to create multiple recursive wrappers. You can now, even _compress_ your data after encrypting, if you're into that 💪🏽😁

> Note: The Decorator pattern is different from the [Adapter pattern](../4-adapter/README.md) because the adapter pattern converts one struct into another. Hence, there is a change in API. The decorator pattern, on the other hand, simply wraps around the existing struct without changing the API.
