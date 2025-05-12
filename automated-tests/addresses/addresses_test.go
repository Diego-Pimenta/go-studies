// it's mandatory to tests follow this the naming convention: classname_test.go
package addresses

import "testing"

type testScenario struct {
	testAddress  string
	expectedType string
}

func TestAddressType(t *testing.T) { // parameter has to be of type *testing.T
	scenarios := []testScenario{
		{"Avenida Santos Lopes", "Avenida"},
		{"Rodovia Bela Castro", "Rodovia"},
		{"Praça das Rosas", "Invalid Type"},
		{"RUA AMARELINHA", "Rua"},
		{"AVENIDA AZUL", "Avenida"},
		{"", "Invalid Type"},
	}

	for _, s := range scenarios {
		actualAddressType := AddressType(s.testAddress)

		if actualAddressType != s.expectedType {
			t.Errorf("Expected %s, but got %s", s.expectedType, actualAddressType)
		}
	}
}
