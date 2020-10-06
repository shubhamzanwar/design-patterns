package main

import "fmt"

func main() {
	info := "kingslayer@gmail.com"

	genericStorage := GenericStorageSystem{}

	genericStorage.writeInfo(info)
	genericStorage.readInfo()

	fmt.Println("------------")

	encryptedStorage := EncryptedStorageSystem{
		storageSystem: &genericStorage,
	}

	encryptedStorage.writeInfo(info)
	encryptedStorage.readInfo()
}
