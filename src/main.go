package main

//EXAMPLE CODE

import (
	"fmt"
	"godb/src/go_database"
)

func main() {
	fmt.Println("Hello, Go!")

	err := go_database.RegisterDatabase("app")
	if err != nil {
		fmt.Println("Error registering database:", err)
		return
	}

	err = go_database.SetData("app", "Alice", 5590)
	if err != nil {
		fmt.Println("Error setting data:", err)
		return
	}

	go_database.GetData("app", "Alice")

	fmt.Println("Successfully inserted/updated data for Alice!")
}
