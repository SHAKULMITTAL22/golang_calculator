package calc

import (
	fmt "fmt"
	math "math"
	os "os"
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
		{
			name:     "Adding Two Positive Numbers",
			num1:     5,
			num2:     10,
			expected: 15,
		},
		{
			name:     "Adding a Positive and a Negative Number",
			num1:     10,
			num2:     -3,
			expected: 7,
		},
		{
			name:     "Adding Two Negative Numbers",
			num1:     -4,
			num2:     -6,
			expected: -10,
		},
		{
			name:     "Adding Zero to a Number",
			num1:     8,
			num2:     0,
			expected: 8,
		},
		{
			name:     "Adding Two Zeroes",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Adding Large Positive Numbers",
			num1:     1000000,
			num2:     2000000,
			expected: 3000000,
		},
		{
			name:     "Adding Large Negative Numbers",
			num1:     -1000000,
			num2:     -2000000,
			expected: -3000000,
		},
		{
			name:     "Adding the Maximum and Minimum Integer Values",
			num1:     math.MaxInt,
			num2:     math.MinInt,
			expected: -1,
		},
		{
			name:     "Adding Two Identical Numbers",
			num1:     7,
			num2:     7,
			expected: 14,
		},
		{
			name:     "Adding Two Opposite Numbers",
			num1:     9,
			num2:     -9,
			expected: 0,
		},
		{
			name:     "Adding Smallest and Largest Possible Integers",
			num1:     math.MinInt,
			num2:     math.MaxInt,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			old := os.Stdout
			_, wPipe, _ := os.Pipe()
			os.Stdout = wPipe

			result := Add(tt.num1, tt.num2)

			wPipe.Close()
			os.Stdout = old

			if result != tt.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s passed: expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestDivide(t *testing.T) {

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
	}()

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
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Divide(tt.num1, tt.num2)

			if math.IsInf(tt.expected, 0) {
				if !math.IsInf(result, 0) {
					t.Errorf("Expected %v but got %v", tt.expected, result)
				}
			} else if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}

			t.Logf("Test %s passed. Expected: %v, Got: %v", tt.name, tt.expected, result)
		})
	}

	w.Close()
	var output string
	fmt.Fscanf(r, "%s", &output)
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{
			name:     "Multiply two positive numbers",
			num1:     3.5,
			num2:     2.0,
			expected: 7.0,
		},
		{
			name:     "Multiply a positive and a negative number",
			num1:     4.0,
			num2:     -2.5,
			expected: -10.0,
		},
		{
			name:     "Multiply two negative numbers",
			num1:     -3.0,
			num2:     -2.0,
			expected: 6.0,
		},
		{
			name:     "Multiply by zero",
			num1:     5.0,
			num2:     0.0,
			expected: 0.0,
		},
		{
			name:     "Multiply by one",
			num1:     7.0,
			num2:     1.0,
			expected: 7.0,
		},
		{
			name:     "Multiply two large numbers",
			num1:     1e10,
			num2:     2e10,
			expected: 2e20,
		},
		{
			name:     "Multiply two small numbers",
			num1:     1e-10,
			num2:     2e-10,
			expected: 2e-20,
		},
		{
			name:     "Multiply a number by itself (square)",
			num1:     5.0,
			num2:     5.0,
			expected: 25.0,
		},
		{
			name:     "Multiply by a fraction",
			num1:     8.0,
			num2:     0.5,
			expected: 4.0,
		},
		{
			name:     "Multiply a very large and very small number",
			num1:     1e10,
			num2:     1e-10,
			expected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Multiply(tt.num1, tt.num2)

			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s passed: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestSubtract(t *testing.T) {

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
	}()

	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Subtracting Two Positive Numbers",
			num1:     10,
			num2:     5,
			expected: 5,
		},
		{
			name:     "Subtracting a Larger Number from a Smaller Number",
			num1:     5,
			num2:     10,
			expected: -5,
		},
		{
			name:     "Subtracting Zero from a Number",
			num1:     7,
			num2:     0,
			expected: 7,
		},
		{
			name:     "Subtracting a Number from Itself",
			num1:     8,
			num2:     8,
			expected: 0,
		},
		{
			name:     "Subtracting Negative Numbers",
			num1:     -3,
			num2:     -7,
			expected: 4,
		},
		{
			name:     "Subtracting a Positive Number from a Negative Number",
			num1:     -5,
			num2:     3,
			expected: -8,
		},
		{
			name:     "Subtracting a Negative Number from a Positive Number",
			num1:     6,
			num2:     -4,
			expected: 10,
		},
		{
			name:     "Subtracting Zero from Zero",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Subtracting the Minimum Integer Value",
			num1:     math.MinInt,
			num2:     -1,
			expected: math.MinInt + 1,
		},
		{
			name:     "Subtracting the Maximum Integer Value",
			num1:     math.MaxInt,
			num2:     1,
			expected: math.MaxInt - 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Subtract(tt.num1, tt.num2)

			if result != tt.expected {
				t.Errorf("Test failed for %s: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test passed for %s: expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}

	w.Close()
	var output string
	fmt.Fscanf(r, "%s", &output)
}
