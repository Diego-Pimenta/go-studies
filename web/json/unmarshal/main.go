package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"name"` // `json:"-"` oculta a tag
	Race string `json:"race"`
	Age  int    `json:"age"`
}

func main() {
	dogJSON := `{"name":"Rex","race":"Labrador","age":5}`

	var dog Dog

	if err := json.Unmarshal([]byte(dogJSON), &dog); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog)

	dog2JSON := `{"age":"3","name":"Sheik","race":"Poodle"}`

	dog2 := make(map[string]string) // be cautious with types

	if err := json.Unmarshal([]byte(dog2JSON), &dog2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog2)
}
