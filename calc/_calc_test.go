package calc

import (
	fmt "fmt"
	math "math"
	os "os"
	testing "testing"
	debug "runtime/debug"
	bytes "bytes"
)








/*
ROOST_METHOD_HASH=Absolute_f8af7505a1
ROOST_METHOD_SIG_HASH=Absolute_4bad226818

FUNCTION_DEF=func Absolute(num float64) float64 

*/
func TestAbsolute(t *testing.T) {

	testCases := []struct {
		name      string
		input     float64
		expected  float64
		expectNaN bool
	}{

		{"PositiveNumber", 42.5, 42.5, false},

		{"NegativeNumber", -42.5, 42.5, false},

		{"ZeroValue", 0.0, 0.0, false},

		{"LargeNegativeMagnitude", -1e10, 1e10, false},

		{"SmallFractionalValue", -0.000001, 0.000001, false},

		{"PositiveInfinity", math.Inf(1), math.Inf(1), false},
		{"NegativeInfinity", math.Inf(-1), math.Inf(1), false},

		{"NaNValue", math.NaN(), math.NaN(), true},

		{"NegativeZero", -0.0, 0.0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			oldStdout := os.Stdout
			defer func() { os.Stdout = oldStdout }()
			r, w, _ := os.Pipe()
			os.Stdout = w

			actual := Absolute(tc.input)

			w.Close()
			var output string
			fmt.Fscan(r, &output)

			if tc.expectNaN {
				if !math.IsNaN(actual) {
					t.Errorf(
						"Test %s: Expected NaN for input %.6f, got %.6f. Captured Output: %s",
						tc.name, tc.input, actual, output,
					)
				} else {
					t.Logf(
						"Test %s Passed: Expected NaN, Matched successfully for input %.6f. Captured Output: %s",
						tc.name, tc.input, output,
					)
				}
			} else {
				if actual != tc.expected {
					t.Errorf(
						"Test %s Failed: Expected %.6f, Got %.6f for input %.6f. Captured Output: %s",
						tc.name, tc.expected, actual, tc.input, output,
					)
				} else {
					t.Logf(
						"Test %s Passed: Expected %.6f, Matched successfully for input %.6f. Captured Output: %s",
						tc.name, tc.expected, tc.input, output,
					)
				}
			}
		})
	}

}


/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int 

*/
func TestAdd(t *testing.T) {

	tests := []struct {
		name        string
		num1        int
		num2        int
		expected    int
		description string
	}{
		{name: "Scenario1", num1: 3, num2: 5, expected: 8, description: "Adding two positive numbers"},
		{name: "Scenario2", num1: -4, num2: -7, expected: -11, description: "Adding two negative numbers"},
		{name: "Scenario3", num1: 9, num2: -2, expected: 7, description: "Adding one positive and one negative number"},
		{name: "Scenario4", num1: 10, num2: 0, expected: 10, description: "Adding zero to an integer"},
		{name: "Scenario5", num1: 0, num2: 0, expected: 0, description: "Adding two zeros"},
		{name: "Scenario6", num1: math.MaxInt32, num2: 1, expected: math.MaxInt32 + 1, description: "Adding a large number with a small number"},
		{name: "Scenario7", num1: math.MaxInt32, num2: math.MaxInt32, expected: math.MaxInt32 + math.MaxInt32, description: "Adding two large numbers, possible overflow"},
		{name: "Scenario8", num1: math.MinInt32, num2: -1, expected: math.MinInt32 + -1, description: "Adding near negative integer boundary"},
		{name: "Scenario9", num1: 6, num2: 6, expected: 12, description: "Adding two identical numbers"},
		{name: "Scenario10", num1: 8, num2: -8, expected: 0, description: "Adding opposite numbers to result in zero"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			output := captureStdout(test.num1, test.num2)
			actualResult := Add(test.num1, test.num2)

			if actualResult != test.expected {
				t.Errorf("Test '%s' failed: %s. Expected %d, got %d", test.name, test.description, test.expected, actualResult)
			} else {
				t.Logf("Test '%s' succeeded: %s. Expected %d, got %d", test.name, test.description, test.expected, actualResult)
			}

			t.Logf("Captured Output: %s", output)
		})
	}
}

func captureStdout(num1, num2 int) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	fmt.Fprintf(w, "%d + %d = %d\n", num1, num2, Add(num1, num2))

	w.Close()
	os.Stdout = stdout

	var result string
	fmt.Fscanf(r, "%s", &result)
	return result
}


/*
ROOST_METHOD_HASH=Divide_f2ddee767d
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64 

*/
func TestDivide(t *testing.T) {

	tests := []struct {
		name         string
		num1         float64
		num2         float64
		expected     float64
		expectPanic  bool
		panicMessage string
	}{
		{
			name:        "Scenario 1: Test Division with Regular Positive Numbers",
			num1:        10,
			num2:        2,
			expected:    5,
			expectPanic: false,
		},
		{
			name:        "Scenario 2: Test Division with Regular Negative Numbers",
			num1:        -10,
			num2:        -2,
			expected:    5,
			expectPanic: false,
		},
		{
			name:        "Scenario 3: Division by Positive Fraction",
			num1:        10,
			num2:        0.5,
			expected:    20,
			expectPanic: false,
		},
		{
			name:        "Scenario 4: Division by Negative Number",
			num1:        10,
			num2:        -2,
			expected:    -5,
			expectPanic: false,
		},
		{
			name:        "Scenario 5: Division Resulting in Infinity",
			num1:        math.MaxFloat64,
			num2:        0.00001,
			expected:    math.MaxFloat64 / 0.00001,
			expectPanic: false,
		},
		{
			name:         "Scenario 6: Division by Zero (Error Handling)",
			num1:         10,
			num2:         0,
			expected:     0,
			expectPanic:  true,
			panicMessage: "division by zero is not allowed",
		},
		{
			name:        "Scenario 7: Division of Zero by Non-Zero Number",
			num1:        0,
			num2:        5,
			expected:    0,
			expectPanic: false,
		},
		{
			name:         "Scenario 8: Division of Negative Number by Zero (Error Handling)",
			num1:         -10,
			num2:         0,
			expected:     0,
			expectPanic:  true,
			panicMessage: "division by zero is not allowed",
		},
		{
			name:        "Scenario 9: Division of Two Large Numbers",
			num1:        math.MaxFloat64,
			num2:        math.MaxFloat64,
			expected:    1,
			expectPanic: false,
		},
		{
			name:        "Scenario 10: Division by Very Small Fraction",
			num1:        5,
			num2:        0.000001,
			expected:    5000000,
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {

				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Logf("Unexpected panic occurred: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
					if r != tt.panicMessage {
						t.Errorf("Expected panic message '%s', but got '%v'", tt.panicMessage, r)
					}
					t.Logf("Panic handled as expected: %v", r)
				} else if tt.expectPanic {
					t.Errorf("Expected panic but it did not occur")
				}
			}()

			var result float64
			if !tt.expectPanic {
				result = Divide(tt.num1, tt.num2)
			}

			if !tt.expectPanic && result != tt.expected {
				t.Errorf("Test failed for %s: expected %f, got %f", tt.name, tt.expected, result)
			} else if !tt.expectPanic {
				t.Logf("Test passed for %s: expected %f, got %f", tt.name, tt.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Factorial_89543dc467
ROOST_METHOD_SIG_HASH=Factorial_9b038c83eb

FUNCTION_DEF=func Factorial(n int) int 

*/
func TestFactorial(t *testing.T) {
	type testCase struct {
		name        string
		input       int
		expected    int
		expectPanic bool
	}

	tests := []testCase{
		{
			name:        "Factorial of Zero",
			input:       0,
			expected:    1,
			expectPanic: false,
		},
		{
			name:        "Factorial of One",
			input:       1,
			expected:    1,
			expectPanic: false,
		},
		{
			name:        "Factorial of Positive Number",
			input:       5,
			expected:    120,
			expectPanic: false,
		},
		{
			name:        "Factorial of Large Positive Number",
			input:       20,
			expected:    2432902008176640000,
			expectPanic: false,
		},
		{
			name:        "Factorial of Negative Number",
			input:       -5,
			expected:    0,
			expectPanic: true,
		},
		{
			name:        "Factorial of Large Number Exceeding Integer Limit",
			input:       21,
			expected:    math.MaxInt64,
			expectPanic: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Panic was correctly triggered for input %d: %v", tc.input, r)
						t.Fail()
					} else {
						t.Errorf("Unexpected panic for input %d: %v\n%s", tc.input, r, string(debug.Stack()))
					}
				}
			}()

			result := Factorial(tc.input)
			if tc.expectPanic {
				t.Errorf("Expected panic for input %d but function did not panic", tc.input)
			} else if result != tc.expected {
				t.Errorf("Factorial(%d) = %d; expected %d", tc.input, result, tc.expected)
			} else {
				t.Logf("Test %s succeeded for input %d, output %d", tc.name, tc.input, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=GCD_6cf0472095
ROOST_METHOD_SIG_HASH=GCD_ab1c91475d

FUNCTION_DEF=func GCD(a, b int) int 

*/
func TestGcd(t *testing.T) {

	type testCase struct {
		name     string
		a        int
		b        int
		expected int
	}

	tests := []testCase{
		{
			name:     "Two positive integers",
			a:        60,
			b:        48,
			expected: 12,
		},
		{
			name:     "One number is zero",
			a:        0,
			b:        42,
			expected: 42,
		},
		{
			name:     "Two prime numbers",
			a:        13,
			b:        17,
			expected: 1,
		},
		{
			name:     "Two identical numbers",
			a:        24,
			b:        24,
			expected: 24,
		},
		{
			name:     "One number is negative",
			a:        -12,
			b:        18,
			expected: 6,
		},
		{
			name:     "Small numbers including zero",
			a:        0,
			b:        1,
			expected: 1,
		},
		{
			name:     "Two large numbers",
			a:        12345678,
			b:        8765432,
			expected: 2,
		},
		{
			name:     "Numbers where one is a divisor",
			a:        84,
			b:        42,
			expected: 42,
		},
		{
			name:     "Smallest negative and positive numbers",
			a:        -1,
			b:        1,
			expected: 1,
		},
		{
			name:     "Both numbers are zero",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tc := range tests {

		func() {
			var buf bytes.Buffer
			stdout := os.Stdout
			defer func() { os.Stdout = stdout }()
			os.Stdout = &buf

			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						t.Logf("Panic encountered in test '%s'.\nDetails: %v\nStack trace: %s", tc.name, r, string(debug.Stack()))
						t.Fail()
					}
				}()

				result := GCD(tc.a, tc.b)

				if result != tc.expected {
					t.Errorf("Test '%s' failed. Expected %d but got %d.", tc.name, tc.expected, result)
				} else {
					t.Logf("Test '%s' succeeded. Expected %d and got %d.", tc.name, tc.expected, result)
				}

				fmt.Fprintf(&buf, "Result: %d\n", result)

			})
		}()
	}
}

