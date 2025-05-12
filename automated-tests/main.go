package main

import (
	"automated-tests/addresses"
	"fmt"
)

func main() {
	addressType := addresses.AddressType("Rua das Flores")
	fmt.Println(addressType)
}
