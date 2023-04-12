package Simple_Unit_Testing

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(1, 2) != float64(3) {
		t.Errorf("Addition function was incorrect, got: %v, want: %v.", Addition(1, 2), float64(3))
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(2, 1) != float64(1) {
		t.Errorf("Subtraction function was incorrect, got: %v, want: %v.", Subtraction(2, 1), float64(1))
	}
}

func TestDivision(t *testing.T) {
	result, err := Division(6, 2)
	if err != nil {
		t.Errorf("Division function was incorrect, got error: %v", err)
	}

	if result != float64(3) {
		t.Errorf("Division function was incorrect, got: %v, want: %v.", result, float64(3))
	}

	_, err = Division(6, 0)
	if err == nil {
		t.Errorf("Division by zero")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(2, 3) != float64(6) {
		t.Errorf("Multiplication function was incorrect, got: %v, want: %v.", Multiplication(2, 3), float64(6))
	}
}
