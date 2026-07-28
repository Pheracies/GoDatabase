package main

//EXAMPLE CODE

import (
	"fmt"

	"github.com/Heaplyn/GoDatabase/src/go_database"
)

type thing struct {
	Defer bool
}

// test
func main() {
	fmt.Println("Hello, Go!")

	b, err := go_database.Encrypt(thing{Defer: true})
	var t thing
	t, err = go_database.Decrypt[thing](b)
	fmt.Println(b)
	fmt.Println(t.Defer)
	err = go_database.RegisterDatabase("app", &go_database.DatabaseOptions[int]{
		DefaultValue: 20,
	})

	errExtra := go_database.RegisterDatabase("app", &go_database.DatabaseOptions[int]{
		DefaultValue: 300,
	})
	if err != nil {
		fmt.Println("Error registering database:", err)
		return
	}
	if errExtra != nil {
		fmt.Println("Error registering database:", errExtra)
		return
	}

	err = go_database.SetData("app", "Alice3", 3242)
	if err != nil {
		fmt.Println("Error setting data:", err)
		return
	}

	_, err = go_database.GetData[int]("app", "Alice4")
	fmt.Println(err)
	if err != nil {
		fmt.Println("Error getting data:", err)
		return
	}

	Thing3, err3 := go_database.GetData[int]("app", "Alice3")
	fmt.Println(Thing3)
	if err3 != nil {
		fmt.Println("Error getting data:", err3)
		return
	}

	fmt.Println("Successfully inserted/updated data for Alice!")
}
