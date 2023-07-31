package main

import (
	"fmt"
	"sync"
)

/*
	In Go (Golang), the Singleton pattern can be implemented using a combination of package-level variables and functions.
	Since Go doesn't have traditional classes, we can use the package scope to create a single instance and provide a function
	to access that instance. Here's an example of implementing a Singleton in Go:
*/

type Database struct {
	Connection string
}

var (
	db   *Database
	once sync.Once
)

func GetInstance() *Database {
	// Create the singleton instance only once
	once.Do(func() {
		fmt.Println("Triggered only once...")

		db = &Database{
			Connection: "test",
		}
	})

	return db
}

func main() {

	instance1 := GetInstance()
	instance2 := GetInstance()
	instance3 := GetInstance()

	// all 3 pointers are pointing to the same underlying Databse object
	// if we use "&" to print, it will print memory address of the pointer itself not the address of the actual Database object
	fmt.Println(instance1)
	fmt.Println(instance2)
	fmt.Println(instance3)

}
