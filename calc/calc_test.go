package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"Adding Two Positive Numbers", 5, 10, 15},
		{"Adding a Positive and a Negative Number", 10, -3, 7},
		{"Adding Two Negative Numbers", -4, -6, -10},
		{"Adding Zero to a Number", 8, 0, 8},
		{"Adding Two Zeroes", 0, 0, 0},
		{"Adding Large Positive Numbers", 1000000, 2000000, 3000000},
		{"Adding Large Negative Numbers", -1000000, -2000000, -3000000},
		{"Adding the Maximum and Minimum Integer Values", math.MaxInt, math.MinInt, -1},
		{"Adding Two Identical Numbers", 7, 7, 14},
		{"Adding Two Opposite Numbers", 9, -9, 0},
		{"Adding Smallest and Largest Possible Integers", math.MinInt, math.MaxInt, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Panic encountered: %v\n%s", r, string(debug.Stack()))
				}
			}()

			result := Add(tt.num1, tt.num2)

			if result != tt.expected {
				t.Fatalf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{"Division of Two Positive Numbers", 10, 2, 5},
		{"Division of a Positive Number by a Negative Number", 10, -2, -5},
		{"Division of a Negative Number by a Positive Number", -10, 2, -5},
		{"Division of Two Negative Numbers", -10, -2, 5},
		{"Division by Zero", 10, 0, math.Inf(1)},
		{"Division of Zero by a Nonzero Number", 0, 10, 0},
		{"Division of a Large Number by a Small Number", 1e10, 1, 1e10},
		{"Division of a Small Number by a Large Number", 1, 1e10, 1e-10},
		{"Division of Two Identical Numbers", 5, 5, 1},
		{"Division of a Floating-Point Number by Itself", 5.5, 5.5, 1},
		{"Division Resulting in a Fraction", 1, 3, 1.0 / 3.0},
		{"Division of Very Small Floating-Point Numbers", 1e-10, 1e-5, 1e-5},
		{"Division of Maximum Float64 Value by 1", math.MaxFloat64, 1, math.MaxFloat64},
		{"Division of Minimum Float64 Value by 1", math.SmallestNonzeroFloat64, 1, math.SmallestNonzeroFloat64},
		{"Division of a Number by a Very Large Number", 1, math.MaxFloat64, 1 / math.MaxFloat64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Panic encountered: %v\n%s", r, string(debug.Stack()))
				}
			}()

			result := Divide(tt.num1, tt.num2)

			if math.IsInf(tt.expected, 0) {
				if !math.IsInf(result, 0) {
					t.Fatalf("Expected %v but got %v", tt.expected, result)
				}
			} else if math.Abs(result-tt.expected) > 1e-9 {
				t.Fatalf("Expected %v but got %v", tt.expected, result)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{"Multiply two positive numbers", 3.5, 2.0, 7.0},
		{"Multiply a positive and a negative number", 4.0, -2.5, -10.0},
		{"Multiply two negative numbers", -3.0, -2.0, 6.0},
		{"Multiply by zero", 5.0, 0.0, 0.0},
		{"Multiply by one", 7.0, 1.0, 7.0},
		{"Multiply two large numbers", 1e10, 2e10, 2e20},
		{"Multiply two small numbers", 1e-10, 2e-10, 2e-20},
		{"Multiply a number by itself (square)", 5.0, 5.0, 25.0},
		{"Multiply by a fraction", 8.0, 0.5, 4.0},
		{"Multiply a very large and very small number", 1e10, 1e-10, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Panic encountered: %v\n%s", r, string(debug.Stack()))
				}
			}()

			result := Multiply(tt.num1, tt.num2)

			if math.Abs(result-tt.expected) > 1e-9 {
				t.Fatalf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"Subtracting Two Positive Numbers", 10, 5, 5},
		{"Subtracting a Larger Number from a Smaller Number", 5, 10, -5},
		{"Subtracting Zero from a Number", 7, 0, 7},
		{"Subtracting a Number from Itself", 8, 8, 0},
		{"Subtracting Negative Numbers", -3, -7, 4},
		{"Subtracting a Positive Number from a Negative Number", -5, 3, -8},
		{"Subtracting a Negative Number from a Positive Number", 6, -4, 10},
		{"Subtracting Zero from Zero", 0, 0, 0},
		{"Subtracting the Minimum Integer Value", math.MinInt, -1, math.MinInt + 1},
		{"Subtracting the Maximum Integer Value", math.MaxInt, 1, math.MaxInt - 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Panic encountered: %v\n%s", r, string(debug.Stack()))
				}
			}()

			result := Subtract(tt.num1, tt.num2)

			if result != tt.expected {
				t.Fatalf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
