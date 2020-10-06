package main

import "fmt"

func encrypt(info string) string {
	return info[1:len(info)] + string(info[0])
}

func decrypt(info string) string {
	return string(info[len(info)-1]) + info[0:len(info)-1]
}

// EncryptedStorageSystem is a decorator that encrypts data before storing it
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
