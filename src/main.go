package main

//EXAMPLE CODE

import (
	"fmt"

	"github.com/Heaplyn/GoDatabase/src/go_database"
)

// test
func main() {
	fmt.Println("Hello, Go!")

	err := go_database.RegisterDatabase("app")
	if err != nil {
		fmt.Println("Error registering database:", err)
		return
	}

	err = go_database.SetData("app", "Alice", 3242)
	if err != nil {
		fmt.Println("Error setting data:", err)
		return
	}

	go_database.GetData("app", "Alice")

	fmt.Println("Successfully inserted/updated data for Alice!")
}
