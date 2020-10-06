package main

import "fmt"

// StorageSystem defines the generic interface for any storage system
type StorageSystem interface {
	writeInfo(info string)
	readInfo() string
}

// GenericStorageSystem is a simple storage that keeps data in memory
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
