package calc

import (
	"fmt"
	"math"
	"os"
	"testing"
	"strings"
)








/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int 

*/
func TestAdd(t *testing.T) {

	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"Adding Two Positive Numbers", 5, 7, 12},
		{"Adding a Positive and a Negative Number", 10, -3, 7},
		{"Adding Two Negative Numbers", -4, -6, -10},
		{"Adding Zero to a Number", 9, 0, 9},
		{"Adding Zero to Zero", 0, 0, 0},
		{"Adding Large Positive Numbers (Overflow Check)", math.MaxInt32, 1, math.MaxInt32 + 1},
		{"Adding Large Negative Numbers (Underflow Check)", math.MinInt32, -1, math.MinInt32 - 1},
		{"Adding Maximum and Minimum Integer Values", math.MaxInt32, math.MinInt32, -1},
		{"Adding Consecutive Numbers", 100, 101, 201},
		{"Adding Two Identical Numbers", 25, 25, 50},
		{"Adding a Small Number to a Large Number", math.MaxInt32, 2, math.MaxInt32 + 2},
		{"Adding Two Random Numbers", 123, 456, 579},
	}

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.num1, tt.num2)

			w.Close()
			os.Stdout = originalStdout
			var output string
			fmt.Fscanf(r, "%s", &output)

			t.Logf("Running test: %s", tt.name)
			t.Logf("Inputs: num1=%d, num2=%d | Expected: %d | Got: %d", tt.num1, tt.num2, tt.expected, result)

			if result != tt.expected {
				t.Errorf("FAILED: %s | Expected %d, but got %d", tt.name, tt.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Divide_052b9c25ea
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64 

*/
func TestDivide(t *testing.T) {

	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{
			name:     "Divide two positive numbers",
			num1:     10.0,
			num2:     2.0,
			expected: 5.0,
		},
		{
			name:     "Divide a positive number by a negative number",
			num1:     10.0,
			num2:     -2.0,
			expected: -5.0,
		},
		{
			name:     "Divide a negative number by a positive number",
			num1:     -10.0,
			num2:     2.0,
			expected: -5.0,
		},
		{
			name:     "Divide two negative numbers",
			num1:     -10.0,
			num2:     -2.0,
			expected: 5.0,
		},
		{
			name:     "Divide by zero",
			num1:     10.0,
			num2:     0.0,
			expected: math.Inf(1),
		},
		{
			name:     "Divide zero by a positive number",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		{
			name:     "Divide zero by a negative number",
			num1:     0.0,
			num2:     -5.0,
			expected: 0.0,
		},
		{
			name:     "Divide a large number by a small number",
			num1:     1e6,
			num2:     2.0,
			expected: 5e5,
		},
		{
			name:     "Divide a small number by a large number",
			num1:     1.0,
			num2:     1e6,
			expected: 1e-6,
		},
		{
			name:     "Divide two identical numbers",
			num1:     7.0,
			num2:     7.0,
			expected: 1.0,
		},
		{
			name:     "Divide by a fraction (decimal number)",
			num1:     10.0,
			num2:     0.5,
			expected: 20.0,
		},
		{
			name:     "Divide a very small number by another small number",
			num1:     1e-6,
			num2:     2e-6,
			expected: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Divide(tt.num1, tt.num2)
			if math.IsInf(tt.expected, 1) || math.IsInf(tt.expected, -1) {

				if !math.IsInf(result, 1) && !math.IsInf(result, -1) {
					t.Errorf("Test failed for %s: expected %v, got %v", tt.name, tt.expected, result)
				}
			} else if math.Abs(result-tt.expected) > 1e-9 {

				t.Errorf("Test failed for %s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Multiply_1585632006
ROOST_METHOD_SIG_HASH=Multiply_d6ab1fb07f

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 

*/
func TestMultiply(t *testing.T) {

	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{
			name:     "Multiply Two Positive Numbers",
			num1:     5,
			num2:     3,
			expected: 15,
		},
		{
			name:     "Multiply a Positive and a Negative Number",
			num1:     6,
			num2:     -2,
			expected: -12,
		},
		{
			name:     "Multiply Two Negative Numbers",
			num1:     -4,
			num2:     -5,
			expected: 20,
		},
		{
			name:     "Multiply by Zero",
			num1:     9,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Multiply by One (Identity Property)",
			num1:     7,
			num2:     1,
			expected: 7,
		},
		{
			name:     "Multiply Two Large Numbers",
			num1:     1e308,
			num2:     2,
			expected: math.Inf(1),
		},
		{
			name:     "Multiply Two Small (Fractional) Numbers",
			num1:     0.1,
			num2:     0.2,
			expected: 0.02,
		},
		{
			name:     "Multiply Two Very Small Numbers (Underflow Check)",
			num1:     1e-308,
			num2:     1e-308,
			expected: 0,
		},
		{
			name:     "Multiply a Number by Itself (Square Calculation)",
			num1:     4,
			num2:     4,
			expected: 16,
		},
		{
			name:     "Multiply Using Maximum Float64 Values (Overflow Check)",
			num1:     math.MaxFloat64,
			num2:     math.MaxFloat64,
			expected: math.Inf(1),
		},
		{
			name:     "Multiply Using Minimum Float64 Values (Negative Overflow Check)",
			num1:     -math.MaxFloat64,
			num2:     -math.MaxFloat64,
			expected: math.Inf(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Testing scenario: %s", tt.name)
			result := Multiply(tt.num1, tt.num2)

			const delta = 1e-10
			if math.IsInf(tt.expected, 1) && math.IsInf(result, 1) {

				t.Logf("Expected +Inf and got +Inf, test passed.")
			} else if math.IsInf(tt.expected, -1) && math.IsInf(result, -1) {

				t.Logf("Expected -Inf and got -Inf, test passed.")
			} else if math.Abs(result-tt.expected) > delta {
				t.Errorf("FAILED: Expected %.10f but got %.10f", tt.expected, result)
			} else {
				t.Logf("PASSED: Expected %.10f and got %.10f", tt.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Subtract_559013d27f
ROOST_METHOD_SIG_HASH=Subtract_29b74c09c9

FUNCTION_DEF=func Subtract(num1, num2 int) int 

*/
func TestSubtract(t *testing.T) {

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"Subtracting Two Positive Numbers", 10, 5, 5},
		{"Subtracting a Larger Number from a Smaller Number", 5, 10, -5},
		{"Subtracting Zero from a Number", 7, 0, 7},
		{"Subtracting a Number from Itself", 6, 6, 0},
		{"Subtracting Negative Numbers", -5, -3, -2},
		{"Subtracting a Negative Number from a Positive Number", 8, -2, 10},
		{"Subtracting a Positive Number from a Negative Number", -4, 3, -7},
		{"Subtracting Zero from Zero", 0, 0, 0},
		{"Subtracting the Largest Positive Integer from Itself", math.MaxInt, math.MaxInt, 0},
		{"Subtracting the Smallest Negative Integer from Itself", math.MinInt, math.MinInt, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := Subtract(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("FAIL: %s - Subtract(%d, %d) = %d; expected %d", tc.name, tc.num1, tc.num2, result, tc.expected)
			} else {
				t.Logf("PASS: %s - Subtract(%d, %d) = %d", tc.name, tc.num1, tc.num2, result)
			}
		})
	}

	outputCapture := captureStdout(func() {
		fmt.Println(Subtract(10, 5))
	})

	expectedOutput := "5\n"
	if outputCapture != expectedOutput {
		t.Errorf("FAIL: Expected output %q but got %q", expectedOutput, outputCapture)
	} else {
		t.Logf("PASS: Correct output captured: %q", outputCapture)
	}
}

