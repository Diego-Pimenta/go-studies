package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  int    `json:"age"`
}

func main() {
	dog := Dog{"Rex", "Labrador", 5}

	dogJSON, err := json.Marshal(dog) // converts from struct to JSON
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dogJSON))

	dog2 := map[string]string{
		"name": "Sheik",
		"race": "Poodle",
		"age":  "3",
	}

	dog2JSON, err := json.Marshal(dog2) // converts from map to JSON
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dog2JSON))
}
