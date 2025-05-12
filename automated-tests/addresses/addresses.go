package addresses

import (
	"slices"
	"strings"
)

// AddressType checks if the address type is valid and returns it.
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	lowerCaseAddress := strings.ToLower(address)
	firstAddressWord := strings.Split(lowerCaseAddress, " ")[0]
	validAddress := slices.Contains(validTypes, firstAddressWord)

	if !validAddress {
		return "Invalid Type"
	}

	return strings.Title(firstAddressWord)
}
