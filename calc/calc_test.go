package calc

import (
	math "math"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Multiply_7a2824e2c7
ROOST_METHOD_SIG_HASH=Multiply_0911ef76c1

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 // Multiply two floating-point numbers
*/
func TestMultiply(t *testing.T) {

	testCases := []struct {
		name           string
		num1           float64
		num2           float64
		expectedResult float64
	}{
		{"Test for Positive Numbers", 2, 3, 6},
		{"Test for Negative Numbers", -2, -3, 6},
		{"Test for Positive and Negative Numbers", 2, -3, -6},
		{"Test for Decimal Cases", 0.2, 0.3, 0.06},
		{"Test for Zero Cases", 6, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()
			result := Multiply(tc.num1, tc.num2)

			if math.Abs(result-tc.expectedResult) > 1e-9 {
				t.Fatalf("Test Failed. Expected %v, got %v", tc.expectedResult, result)
			} else {
				t.Logf("Test Passed. Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=SquareRoot_17095d9165
ROOST_METHOD_SIG_HASH=SquareRoot_232943a56a

FUNCTION_DEF=func SquareRoot(num float64) float64 // Square root (with error handling)
*/
func TestSquareRoot(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    float64
		Expected float64
		IsPanic  bool
	}{
		{
			Name:     "Scenario 1: Testing normal operation - Positive Numbers",
			Input:    4.0,
			Expected: 2.0,
			IsPanic:  false,
		},
		{
			Name:     "Scenario 2: Testing edge case - Zero",
			Input:    0.0,
			Expected: 0.0,
			IsPanic:  false,
		},
		{
			Name:    "Scenario 3: Testing error handling - Negative Numbers",
			Input:   -1.0,
			IsPanic: true,
		},
		{
			Name:     "Scenario 4: Testing edge case - Very Large Numbers",
			Input:    1.0e+308,
			Expected: math.Sqrt(1.0e+308),
			IsPanic:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.IsPanic {
						t.Logf("A panic was expected and occurred successfully during the test. Reason: %v", r)
					} else {
						t.Errorf("The test failed as an unexpected panic occurred. Reason: %v", r)
					}
				}
			}()

			res := SquareRoot(tc.Input)
			if res != tc.Expected && !tc.IsPanic {
				t.Errorf("The expected result (%v) does not match the obtained one (%v)", tc.Expected, res)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Subtract_58eac52f91
ROOST_METHOD_SIG_HASH=Subtract_b1211baa34

FUNCTION_DEF=func Subtract(num1, num2 int) int // Subtract two integers
*/
func TestSubtract(t *testing.T) {

	scenarios := []struct {
		Description string
		Num1        int
		Num2        int
		Expected    int
	}{
		{
			Description: "Subtraction of two positive numbers",
			Num1:        5,
			Num2:        3,
			Expected:    2,
		},
		{
			Description: "Subtraction of two negative numbers",
			Num1:        -5,
			Num2:        -3,
			Expected:    -2,
		},
		{
			Description: "Subtracting zero from a number",
			Num1:        5,
			Num2:        0,
			Expected:    5,
		},
		{
			Description: "Subtracting a number from zero",
			Num1:        0,
			Num2:        5,
			Expected:    -5,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test: %v\n", r)
					t.Fail()
				}
			}()

			result := Subtract(tt.Num1, tt.Num2)
			if result != tt.Expected {
				t.Errorf("Failed: %s: Subtract(%d, %d): expected %d, received %d", tt.Description, tt.Num1, tt.Num2, tt.Expected, result)
			} else {
				t.Logf("Success: %s: Subtract(%d, %d): expected %d, received %d", tt.Description, tt.Num1, tt.Num2, tt.Expected, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Factorial_68fe6fb960
ROOST_METHOD_SIG_HASH=Factorial_3d037eec72

FUNCTION_DEF=func Factorial(n int) int // Factorial (Recursive)
*/
func TestFactorial(t *testing.T) {

	tests := []struct {
		name         string
		input        int
		expected     int
		expectPanic  bool
		panicMessage string
	}{
		{
			name:     "Factorial of positive number",
			input:    5,
			expected: 120,
		},
		{
			name:     "Factorial of zero",
			input:    0,
			expected: 1,
		},
		{
			name:         "Factorial of negative number",
			input:        -5,
			expectPanic:  true,
			panicMessage: "factorial is not defined for negative numbers",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if test.expectPanic {
						if r != test.panicMessage {
							t.Fatalf("expected panic message %q but got %q", test.panicMessage, r)
						} else {
							t.Logf("expected panic occurred: %v", r)
						}
					} else {
						t.Fatalf("unexpected panic occurred: %v", r)
					}
				} else if test.expectPanic {
					t.Fatal("expected a panic, but no panic occurred")
				}
			}()

			result := Factorial(test.input)

			if result != test.expected {
				t.Errorf("expected %d but got %d", test.expected, result)
			} else {
				t.Logf("Success: expected %d, received %d", test.expected, result)
			}
		})
	}
}
