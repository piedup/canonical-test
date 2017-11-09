package main

import (
	"fmt"

	cj "github.com/gibson042/canonicaljson-go"
)

// Person is a person
type Person struct {
	Name string `json:"name"`
}

func main() {
	person := Person{
		Name: "me",
	}
	_, err := cj.Marshal(person)
	if err != nil {
		fmt.Println("Error")
	}
}
