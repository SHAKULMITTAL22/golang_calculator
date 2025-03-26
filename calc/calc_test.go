package calc

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	os "os"
	debug "runtime/debug"
	strings "strings"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int
*/
func TestAdd(t *testing.T) {

	var outputBuffer bytes.Buffer

	oldStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull)
	defer func() {
		os.Stdout = oldStdout
	}()

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Add Two Positive Integers",
			num1:     5,
			num2:     10,
			expected: 15,
		},
		{
			name:     "Add Two Negative Integers",
			num1:     -5,
			num2:     -10,
			expected: -15,
		},
		{
			name:     "Add a Positive and a Negative Integer",
			num1:     10,
			num2:     -5,
			expected: 5,
		},
		{
			name:     "Add Zero to an Integer",
			num1:     0,
			num2:     7,
			expected: 7,
		},
		{
			name:     "Add an Integer to Itself",
			num1:     8,
			num2:     8,
			expected: 16,
		},
		{
			name:     "Add the Maximum Integer Value in Go",
			num1:     math.MaxInt,
			num2:     0,
			expected: math.MaxInt,
		},
		{
			name:     "Adding the Minimum Integer Value in Go",
			num1:     math.MinInt,
			num2:     0,
			expected: math.MinInt,
		},
		{
			name:     "Add Two Large Numbers",
			num1:     math.MaxInt - 1000,
			num2:     999,
			expected: math.MaxInt - 1,
		},
		{
			name:     "Adding Two Small Negative Numbers",
			num1:     -1,
			num2:     -2,
			expected: -3,
		},
		{
			name:     "Adding Same Opposite Numbers",
			num1:     10,
			num2:     -10,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution. Test failed.\nDetails: %v\nStack Trace:\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Add(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Test '%s' Failed: Expected %d, Got %d. num1: %d, num2: %d", tc.name, tc.expected, result, tc.num1, tc.num2)
				return
			}

			t.Logf("Test '%s' Passed: %d + %d = %d", tc.name, tc.num1, tc.num2, result)
		})
	}

	if outputBuffer.Len() > 0 {
		outputContent := outputBuffer.String()
		t.Logf("Captured console output during tests:\n%s", outputContent)
	}
}

/*
ROOST_METHOD_HASH=Divide_052b9c25ea
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64
*/
func TestDivide(t *testing.T) {

	type testCase struct {
		name        string
		num1        float64
		num2        float64
		expected    float64
		expectPanic bool
	}

	testCases := []testCase{
		{name: "Scenario 1: Two Positive Numbers", num1: 10.0, num2: 2.0, expected: 5.0},
		{name: "Scenario 2: Two Negative Numbers", num1: -10.0, num2: -2.0, expected: 5.0},
		{name: "Scenario 3: Positive by Negative Number", num1: 10.0, num2: -2.0, expected: -5.0},
		{name: "Scenario 4: Floating-Point Result", num1: 7.0, num2: 2.0, expected: 3.5},
		{name: "Scenario 5: Division by Zero", num1: 10.0, num2: 0.0, expectPanic: true},
		{name: "Scenario 6: Zero Divided by Non-Zero Number", num1: 0.0, num2: 3.0, expected: 0.0},
		{name: "Scenario 7: Very Large Inputs", num1: 1e308, num2: 1e308, expected: 1.0},
		{name: "Scenario 8: Very Small Inputs", num1: 1e-308, num2: 2e-308, expected: 0.5},
		{name: "Scenario 9: Identical Numerator and Denominator", num1: 5.0, num2: 5.0, expected: 1.0},
		{name: "Scenario 10: Extremely Uneven Division", num1: 1.0, num2: 1e10, expected: 1e-10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Panic expected and encountered successfully for '%s' scenario: %v\n%s", tc.name, r, string(debug.Stack()))
						return
					} else {
						t.Logf("Unexpected panic for '%s' scenario: %v\n%s", tc.name, r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			var sb strings.Builder
			writer := &sb

			fmt.Fprintf(writer, "Executing test case: %s\n", tc.name)

			var result float64
			if !tc.expectPanic {
				result = Divide(tc.num1, tc.num2)

				if result != tc.expected {
					t.Errorf("Scenario '%s' failed. Expected: %v, Got: %v", tc.name, tc.expected, result)
				} else {
					t.Logf("Scenario '%s' succeeded. Expected result matched: %v", tc.name, result)
				}
			} else {
				Divide(tc.num1, tc.num2)
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
		name         string
		num1         float64
		num2         float64
		expected     float64
		expectNaN    bool
		expectInf    bool
		description  string
		verifyResult func(actual, expected float64) bool
	}{
		{
			name:         "Scenario 1: Multiplication of Two Positive Numbers",
			num1:         2.5,
			num2:         5.0,
			expected:     12.5,
			description:  "Test to ensure the function correctly multiplies two positive floating-point numbers.",
			verifyResult: func(actual, expected float64) bool { return actual == expected },
		},
		{
			name:         "Scenario 2: Multiplication of a Positive Number and Zero",
			num1:         5.2,
			num2:         0.0,
			expected:     0.0,
			description:  "Check the function's behavior when one operand is zero.",
			verifyResult: func(actual, expected float64) bool { return actual == expected },
		},
		{
			name:         "Scenario 3: Multiplication of Two Negative Numbers",
			num1:         -4.5,
			num2:         -2.0,
			expected:     9.0,
			description:  "Verify multiplication of two negative numbers produces a positive result.",
			verifyResult: func(actual, expected float64) bool { return actual == expected },
		},
		{
			name:         "Scenario 4: Multiplication of a Positive and Negative Number",
			num1:         7.0,
			num2:         -3.0,
			expected:     -21.0,
			description:  "Ensure the function computes correctly for positive and negative numbers.",
			verifyResult: func(actual, expected float64) bool { return actual == expected },
		},
		{
			name:         "Scenario 5: Multiplication of Very Large Numbers",
			num1:         1e+308,
			num2:         2.0,
			expected:     math.Inf(1),
			expectInf:    true,
			description:  "Check for overflow behavior with large numbers.",
			verifyResult: func(actual, expected float64) bool { return math.IsInf(actual, 1) },
		},
		{
			name:         "Scenario 6: Multiplication of Very Small Numbers",
			num1:         1e-308,
			num2:         2.0,
			expected:     2e-308,
			description:  "Test behavior for small numbers near float64 precision limits.",
			verifyResult: func(actual, expected float64) bool { return math.Abs(actual-expected) < 1e-320 },
		},
		{
			name:         "Scenario 7: Multiplication by One",
			num1:         7.25,
			num2:         1.0,
			expected:     7.25,
			description:  "Test compliance with the identity property of multiplication.",
			verifyResult: func(actual, expected float64) bool { return actual == expected },
		},
		{
			name:         "Scenario 8: Multiplication of Numbers Resulting in Zero",
			num1:         1e-308,
			num2:         1e+308,
			expected:     1.0,
			description:  "Test for underflow resulting from extreme input values.",
			verifyResult: func(actual, expected float64) bool { return math.Abs(actual-expected) < 1e-15 },
		},
		{
			name:         "Scenario 9: Multiplication Resulting in Negative Zero",
			num1:         -0.0,
			num2:         2.0,
			expected:     -0.0,
			description:  "Validate edge cases where the result is negative zero.",
			verifyResult: func(actual, expected float64) bool { return math.Signbit(actual) && actual == expected },
		},
		{
			name:         "Scenario 10: Multiplication with High Precision Floating-Point Numbers",
			num1:         0.12345,
			num2:         0.67890,
			expected:     0.083863905,
			description:  "Test high-precision inputs for floating-point calculations.",
			verifyResult: func(actual, expected float64) bool { return math.Abs(actual-expected) < 1e-9 },
		},
		{
			name:         "Scenario 11: Multiplication with NaN Inputs",
			num1:         math.NaN(),
			num2:         5.0,
			expectNaN:    true,
			description:  "Check behavior when one or both inputs are NaN.",
			verifyResult: func(actual, _ float64) bool { return math.IsNaN(actual) },
		},
		{
			name:         "Scenario 12: Multiplication with Infinite Inputs",
			num1:         math.Inf(1),
			num2:         -3.0,
			expected:     math.Inf(-1),
			expectInf:    true,
			description:  "Validate behavior when one operand is infinity.",
			verifyResult: func(actual, expected float64) bool { return math.IsInf(actual, -1) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during %s. %v\n%s", tt.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Log(tt.description)

			actual := Multiply(tt.num1, tt.num2)

			if tt.expectNaN {
				if !math.IsNaN(actual) {
					t.Errorf("Test %s failed: expected NaN but got %v", tt.name, actual)
				}
			} else if tt.expectInf {
				if !math.IsInf(actual, int(tt.expected)/int(math.Abs(tt.expected))) {
					t.Errorf("Test %s failed: expected %v infinity but got %v", tt.name, tt.expected, actual)
				}
			} else if !tt.verifyResult(actual, tt.expected) {
				t.Errorf("Test %s failed: expected %v but got %v", tt.name, tt.expected, actual)
			} else {
				t.Logf("Test %s passed successfully.", tt.name)
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
	type testCase struct {
		name     string
		num1     int
		num2     int
		expected int
	}

	testCases := []testCase{
		{"Subtract Two Positive Integers", 10, 5, 5},
		{"Subtract Two Negative Integers", -10, -5, -5},
		{"Subtract Positive and Negative Integer", -10, 5, -15},
		{"Subtract Negative and Positive Integer", 10, -5, 15},
		{"Subtract Zero from an Integer", 25, 0, 25},
		{"Subtract Integer from Zero", 0, 7, -7},
		{"Subtract Two Zero Values", 0, 0, 0},
		{"Subtract Maximum Integer Value from Another Integer", math.MaxInt32, 10, math.MaxInt32 - 10},
		{"Subtract Minimum Integer Value from Another Integer", math.MinInt32, 5, math.MinInt32 - 5},
		{"Subtract Very Large Integers", 999999999, 888888888, 111111111},
		{"Validate Output Type Consistency", -5, 5, -5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Subtract(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed: result is %d as expected", tc.name, result)
			}

			if fmt.Sprintf("%T", result) != "int" {
				t.Errorf("Test %s failed: result type mismatch, expected int but got %T", tc.name, result)
			}
		})
	}

}
