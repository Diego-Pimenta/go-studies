// it's mandatory to tests follow this the naming convention: classname_test.go
package addresses_test

import (
	. "automated-tests/addresses" // create an alias for the package
	"testing"
)

type testScenario struct {
	testAddress  string
	expectedType string
}

/*
"go test" runs the tests in the current package
"go test ./..." runs all tests in the current directory and subdirectories

flag "-v" shows the details of the tests
flag "--cover" shows the coverage of the tests

"go test --coverprofile coverage.txt" creates a coverage file
"go tool cover --html=coverage.txt" shows the coverage in a browser
*/

func TestAddressType(t *testing.T) { // parameter has to be of type *testing.T
	t.Parallel() // tests with this function will run in parallel

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
