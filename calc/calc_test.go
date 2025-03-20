package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
)








/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int 

*/
func TestAdd(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic encountered, failing test: %v\n%s", r, string(debug.Stack()))
			t.Fail()
		}
	}()

	testCases := []struct {
		name        string
		num1        int
		num2        int
		expected    int
		expectPanic bool
	}{
		{
			name:        "Adding two positive integers",
			num1:        10,
			num2:        5,
			expected:    15,
			expectPanic: false,
		},
		{
			name:        "Adding a positive integer and a negative integer",
			num1:        20,
			num2:        -10,
			expected:    10,
			expectPanic: false,
		},
		{
			name:        "Adding two negative integers",
			num1:        -7,
			num2:        -3,
			expected:    -10,
			expectPanic: false,
		},
		{
			name:        "Adding zero to a positive integer",
			num1:        25,
			num2:        0,
			expected:    25,
			expectPanic: false,
		},
		{
			name:        "Adding zero to a negative integer",
			num1:        -12,
			num2:        0,
			expected:    -12,
			expectPanic: false,
		},
		{
			name:        "Adding two zeros",
			num1:        0,
			num2:        0,
			expected:    0,
			expectPanic: false,
		},
		{
			name:        "Adding the maximum integer value",
			num1:        math.MaxInt,
			num2:        0,
			expected:    math.MaxInt,
			expectPanic: false,
		},
		{
			name:        "Adding the minimum integer value",
			num1:        math.MinInt,
			num2:        0,
			expected:    math.MinInt,
			expectPanic: false,
		},
		{
			name:        "Adding two large integers exceeding the range",
			num1:        math.MaxInt,
			num2:        1,
			expected:    0,
			expectPanic: true,
		},
		{
			name:        "Adding two small integers close to the minimum range",
			num1:        math.MinInt,
			num2:        -1,
			expected:    0,
			expectPanic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Panic recovered as expected: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			result := Add(tc.num1, tc.num2)

			if tc.expectPanic {
				t.Logf("Expected panic, but function returned a value: %d", result)
				t.Fail()
			} else if result != tc.expected {
				t.Errorf("Test failed: %s\nExpected: %d, Got: %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test passed: %s\nExpected: %d, Got: %d", tc.name, tc.expected, result)
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
	type testCase struct {
		name        string
		num1        float64
		num2        float64
		expected    float64
		shouldPanic bool
	}

	tests := []testCase{
		{
			name:     "Scenario 1: Division of two positive numbers",
			num1:     10.0,
			num2:     2.0,
			expected: 5.0,
		},
		{
			name:     "Scenario 2: Division with zero as numerator",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		{
			name:        "Scenario 3: Division by zero",
			num1:        5.0,
			num2:        0.0,
			shouldPanic: true,
		},
		{
			name:     "Scenario 4: Division of a negative number by a positive number",
			num1:     -10.0,
			num2:     2.0,
			expected: -5.0,
		},
		{
			name:     "Scenario 5: Division of a positive number by a negative number",
			num1:     10.0,
			num2:     -2.0,
			expected: -5.0,
		},
		{
			name:     "Scenario 6: Division of two negative numbers",
			num1:     -10.0,
			num2:     -2.0,
			expected: 5.0,
		},
		{
			name:     "Scenario 7: Division of a very large number by a very small number",
			num1:     1e10,
			num2:     1e-5,
			expected: 1e15,
		},
		{
			name:     "Scenario 8: Division of a number by itself",
			num1:     7.5,
			num2:     7.5,
			expected: 1.0,
		},
		{
			name:     "Scenario 9: Division involving very small positive numbers",
			num1:     1e-7,
			num2:     1e-9,
			expected: 100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.shouldPanic {
						t.Logf("Test passed: Expected panic encountered. %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Panic encountered during test case: %s - Panic: %v\n%s", tt.name, r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			if tt.shouldPanic {

				_ = Divide(tt.num1, tt.num2)
				t.Fail()
			} else {

				result := Divide(tt.num1, tt.num2)
				if math.Abs(result-tt.expected) > 1e-9 {
					t.Errorf("Test %s failed: Expected result %.9f, got %.9f", tt.name, tt.expected, result)
				} else {
					t.Logf("Test %s passed: Expected %.9f, got %.9f", tt.name, tt.expected, result)
				}
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
	type testCase struct {
		description string
		num1        float64
		num2        float64
		expected    float64
	}

	tests := []testCase{
		{
			description: "Multiply two positive numbers",
			num1:        4.5,
			num2:        2.0,
			expected:    9.0,
		},
		{
			description: "Multiply a positive number by zero",
			num1:        7.8,
			num2:        0.0,
			expected:    0.0,
		},
		{
			description: "Multiply a negative number by zero",
			num1:        -3.2,
			num2:        0.0,
			expected:    0.0,
		},
		{
			description: "Multiply two negative numbers",
			num1:        -4.5,
			num2:        -3.0,
			expected:    13.5,
		},
		{
			description: "Multiply positive and negative numbers",
			num1:        6.0,
			num2:        -2.0,
			expected:    -12.0,
		},
		{
			description: "Multiply floating-point numbers with decimals",
			num1:        2.5,
			num2:        1.6,
			expected:    4.0,
		},
		{
			description: "Multiply by one (identity property)",
			num1:        8.7,
			num2:        1.0,
			expected:    8.7,
		},
		{
			description: "Multiply two zeroes",
			num1:        0.0,
			num2:        0.0,
			expected:    0.0,
		},
		{
			description: "Multiply small fractional numbers",
			num1:        0.0002,
			num2:        0.0004,
			expected:    0.00000008,
		},
		{
			description: "Multiply very large numbers",
			num1:        1e+10,
			num2:        2.5,
			expected:    2.5e+10,
		},
		{
			description: "Multiply a number by negative one (sign inversion)",
			num1:        15.0,
			num2:        -1.0,
			expected:    -15.0,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, failing test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Multiply(tc.num1, tc.num2)
			if math.Abs(result-tc.expected) > 1e-7 {
				t.Errorf("[FAILED] %s: Multiply(%v, %v) = %v; expected = %v", tc.description, tc.num1, tc.num2, result, tc.expected)
			} else {
				t.Logf("[PASSED] %s: Multiply(%v, %v) = %v; expected = %v", tc.description, tc.num1, tc.num2, result, tc.expected)
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
		name           string
		num1, num2     int
		expectedOutput int
		shouldPanic    bool
	}

	testCases := []testCase{
		{
			name:           "Subtracting two positive integers",
			num1:           10,
			num2:           5,
			expectedOutput: 5,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting a positive integer from zero",
			num1:           0,
			num2:           7,
			expectedOutput: -7,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting a negative integer from a positive integer",
			num1:           10,
			num2:           -5,
			expectedOutput: 15,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting zero from a positive integer",
			num1:           15,
			num2:           0,
			expectedOutput: 15,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting a positive integer from itself",
			num1:           8,
			num2:           8,
			expectedOutput: 0,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting two negative integers",
			num1:           -4,
			num2:           -6,
			expectedOutput: 2,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting zero from zero",
			num1:           0,
			num2:           0,
			expectedOutput: 0,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting a large positive integer from a small positive integer",
			num1:           3,
			num2:           10,
			expectedOutput: -7,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting maximum integer values",
			num1:           2147483647,
			num2:           1,
			expectedOutput: 2147483646,
			shouldPanic:    false,
		},
		{
			name:           "Subtracting minimum integer values",
			num1:           -2147483648,
			num2:           1,
			expectedOutput: -2147483647,
			shouldPanic:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actualOutput := Subtract(tc.num1, tc.num2)

			if actualOutput != tc.expectedOutput {
				t.Errorf("Test '%s' failed. Expected %d, got %d", tc.name, tc.expectedOutput, actualOutput)
			} else {
				t.Logf("Test '%s' passed. Correct output: %d", tc.name, actualOutput)
			}
		})
	}
}

