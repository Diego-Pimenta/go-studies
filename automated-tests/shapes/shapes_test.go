package shapes_test

import (
	. "automated-tests/shapes"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{Radius: 5}
		expectedArea := float64(math.Pi * 25)
		actualArea := circle.Area()

		if actualArea != expectedArea {
			// when using Fatalf(), the code stops when it finds the error
			t.Fatalf("Expected area %f, but got %f", expectedArea, actualArea)
		}
	})

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 4, Height: 5}
		expectedArea := float64(20)
		actualArea := rectangle.Area()

		if actualArea != expectedArea {
			t.Fatalf("Expected area %f, but got %f", expectedArea, actualArea)
		}
	})
}
