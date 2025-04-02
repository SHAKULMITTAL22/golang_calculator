package calc

import (
	math "math"
	testing "testing"
	debug "runtime/debug"
	fmt "fmt"
	os "os"
	bytes "bytes"
	sync "sync"
	strings "strings"
	time "time"
)








/*
ROOST_METHOD_HASH=Absolute_d231f0ab10
ROOST_METHOD_SIG_HASH=Absolute_ec3c06e5a3

FUNCTION_DEF=func Absolute(num float64) float64 // Absolute value


*/
func TestAbsolute(t *testing.T) {

	t.Parallel()

	testCases := []struct {
		name        string
		input       float64
		expected    float64
		expectPanic bool
	}{

		{
			name:        "Positive Float Input",
			input:       5.67,
			expected:    5.67,
			expectPanic: false,
		},

		{
			name:        "Negative Float Input",
			input:       -8.42,
			expected:    8.42,
			expectPanic: false,
		},

		{
			name:        "Zero Input",
			input:       0.0,
			expected:    0.0,
			expectPanic: false,
		},

		{
			name:        "Large Positive Float Input",
			input:       1e+308,
			expected:    1e+308,
			expectPanic: false,
		},

		{
			name:        "Large Negative Float Input",
			input:       -1e+308,
			expected:    1e+308,
			expectPanic: false,
		},

		{
			name:        "Small Positive Float Input",
			input:       1e-10,
			expected:    1e-10,
			expectPanic: false,
		},

		{
			name:        "Small Negative Float Input",
			input:       -1e-10,
			expected:    1e-10,
			expectPanic: false,
		},

		{
			name:        "NaN Input",
			input:       math.NaN(),
			expected:    math.NaN(),
			expectPanic: false,
		},

		{
			name:        "Positive Infinity",
			input:       math.Inf(1),
			expected:    math.Inf(1),
			expectPanic: false,
		},

		{
			name:        "Negative Infinity",
			input:       math.Inf(-1),
			expected:    math.Inf(1),
			expectPanic: false,
		},

		{
			name:        "Negative Zero Input",
			input:       -0.0,
			expected:    0.0,
			expectPanic: false,
		},
	}

	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Absolute(tc.input)

			t.Logf("Testing %s\nInput: %.6f\nExpected: %.6f\nReceived: %.6f\n", tc.name, tc.input, tc.expected, result)

			if math.IsNaN(tc.expected) {
				if !math.IsNaN(result) {
					t.Errorf("FAILED %s: Expected NaN but got %.6f", tc.name, result)
				} else {
					t.Logf("PASSED %s: Expected NaN and got NaN", tc.name)
				}
			} else if result != tc.expected {
				t.Errorf("FAILED %s: Expected %.6f but got %.6f", tc.name, tc.expected, result)
			} else {
				t.Logf("PASSED %s: Result matched expected value %.6f", tc.name, tc.expected)
			}
		})
	}

}


/*
ROOST_METHOD_HASH=Add_61b83cbd8d
ROOST_METHOD_SIG_HASH=Add_9b37ae6611

FUNCTION_DEF=func Add(num1, num2 int) int // Add two integers


*/
func TestAdd(t *testing.T) {

	testCases := []struct {
		name        string
		num1        int
		num2        int
		expected    int
		shouldPanic bool
	}{
		{
			name:     "Adding two positive numbers",
			num1:     5,
			num2:     10,
			expected: 15,
		},
		{
			name:     "Adding positive and negative numbers",
			num1:     10,
			num2:     -5,
			expected: 5,
		},
		{
			name:     "Adding two negative numbers",
			num1:     -3,
			num2:     -7,
			expected: -10,
		},
		{
			name:     "Adding zero to a number",
			num1:     42,
			num2:     0,
			expected: 42,
		},
		{
			name:     "Adding two zeros",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:        "Adding large positive integers",
			num1:        math.MaxInt - 1,
			num2:        2,
			expected:    math.MaxInt + 1,
			shouldPanic: true,
		},
		{
			name:     "Adding large negative integers",
			num1:     math.MinInt + 1,
			num2:     -1,
			expected: math.MinInt,
		},
		{
			name:        "Integer overflow",
			num1:        math.MaxInt,
			num2:        1,
			expected:    math.MinInt,
			shouldPanic: true,
		},
		{
			name:        "Integer underflow",
			num1:        math.MinInt,
			num2:        -1,
			expected:    math.MaxInt,
			shouldPanic: true,
		},
		{
			name:     "Adding opposite-signed numbers resulting in zero",
			num1:     20,
			num2:     -20,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.shouldPanic {
						t.Logf("Expected panic encountered: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic encountered: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			result := Add(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("For test case '%s', expected %d but got %d", tc.name, tc.expected, result)
			}

			t.Logf("Test '%s' successfully passed. Input (%d + %d) = Expected: %d, Result: %d",
				tc.name, tc.num1, tc.num2, tc.expected, result)
		})
	}

}


/*
ROOST_METHOD_HASH=Divide_6fe509f399
ROOST_METHOD_SIG_HASH=Divide_d926fccfc9

FUNCTION_DEF=func Divide(num1, num2 float64) float64 // Divide two floating-point numbers (with error handling)


*/
func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		num1        float64
		num2        float64
		expected    float64
		expectPanic bool
	}

	testCases := []testCase{
		{
			name:        "Divide two positive numbers",
			num1:        10.0,
			num2:        2.0,
			expected:    5.0,
			expectPanic: false,
		},
		{
			name:        "Divide a positive number by a negative number",
			num1:        10.0,
			num2:        -2.0,
			expected:    -5.0,
			expectPanic: false,
		},
		{
			name:        "Divide two negative numbers",
			num1:        -10.0,
			num2:        -2.0,
			expected:    5.0,
			expectPanic: false,
		},
		{
			name:        "Divide by zero",
			num1:        10.0,
			num2:        0.0,
			expectPanic: true,
		},
		{
			name:        "Divide zero by a number",
			num1:        0.0,
			num2:        5.0,
			expected:    0.0,
			expectPanic: false,
		},
		{
			name:        "Divide two fractional numbers",
			num1:        0.5,
			num2:        0.2,
			expected:    2.5,
			expectPanic: false,
		},
		{
			name:        "Divide a number by 1",
			num1:        123.45,
			num2:        1.0,
			expected:    123.45,
			expectPanic: false,
		},
		{
			name:        "Divide a very large number by a small number",
			num1:        1e10,
			num2:        1e-2,
			expected:    1e12,
			expectPanic: false,
		},
		{
			name:        "Divide two very small numbers",
			num1:        1e-10,
			num2:        1e-5,
			expected:    1e-5,
			expectPanic: false,
		},
		{
			name:        "Divide using identical numbers",
			num1:        123.45,
			num2:        123.45,
			expected:    1.0,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {

		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Expected panic occurred: %v", r)
					} else {
						t.Errorf("Unexpected panic occurred: %v\n%s", r, string(debug.Stack()))
					}
				}
			}()

			var result float64
			if !tc.expectPanic {
				result = Divide(tc.num1, tc.num2)
			} else {

				_ = Divide(tc.num1, tc.num2)
			}

			if !tc.expectPanic {
				if math.Abs(result-tc.expected) > 1e-9 {
					t.Errorf("Test failed for %q: Expected %.10f, but got %.10f", tc.name, tc.expected, result)
				} else {
					t.Logf("Success for %q: Expected %.10f, got %.10f", tc.name, tc.expected, result)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Logarithm_4092f1cba7
ROOST_METHOD_SIG_HASH=Logarithm_0780d00fe8

FUNCTION_DEF=func Logarithm(num, base float64) float64 // Logarithm function (log_base of num)


*/
func TestLogarithm(t *testing.T) {

	type testCase struct {
		description string
		num         float64
		base        float64
		expectPanic bool
		expected    float64
		tolerance   float64
		panicMsg    string
	}

	testCases := []testCase{
		{
			description: "Calculate logarithm with valid positive inputs",
			num:         8.0,
			base:        2.0,
			expectPanic: false,
			expected:    3.0,
			tolerance:   1e-9,
		},
		{
			description: "Panic on zero as the 'num' parameter",
			num:         0.0,
			base:        2.0,
			expectPanic: true,
			panicMsg:    "logarithm is not defined for these values",
		},
		{
			description: "Panic on negative number as the 'num' parameter",
			num:         -5.0,
			base:        2.0,
			expectPanic: true,
			panicMsg:    "logarithm is not defined for these values",
		},
		{
			description: "Panic on base value of 1",
			num:         8.0,
			base:        1.0,
			expectPanic: true,
			panicMsg:    "logarithm is not defined for these values",
		},
		{
			description: "Panic on zero or negative base value",
			num:         8.0,
			base:        -3.0,
			expectPanic: true,
			panicMsg:    "logarithm is not defined for these values",
		},
		{
			description: "Compute logarithm where num equals base",
			num:         5.0,
			base:        5.0,
			expectPanic: false,
			expected:    1.0,
			tolerance:   1e-9,
		},
		{
			description: "Compute logarithm where base is greater than the number",
			num:         2.0,
			base:        4.0,
			expectPanic: false,
			expected:    0.5,
			tolerance:   1e-9,
		},
		{
			description: "Compute logarithm where base is smaller than the number",
			num:         8.0,
			base:        2.0,
			expectPanic: false,
			expected:    3.0,
			tolerance:   1e-9,
		},
		{
			description: "Compute fractional base and number values",
			num:         7.5,
			base:        1.5,
			expectPanic: false,
			expected:    math.Log(7.5) / math.Log(1.5),
			tolerance:   1e-9,
		},
		{
			description: "Edge test - Large numbers with large bases",
			num:         1e10,
			base:        1e5,
			expectPanic: false,
			expected:    2.0,
			tolerance:   1e-9,
		},
		{
			description: "Edge test - Very small values (close to 0 but > 0)",
			num:         1e-10,
			base:        2.0,
			expectPanic: false,
			expected:    math.Log(1e-10) / math.Log(2.0),
			tolerance:   1e-9,
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = oldStdout
		r.Close()
		w.Close()
	}()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					if !tc.expectPanic || r != tc.panicMsg {
						t.Fail()
					}
				}
			}()

			if tc.expectPanic {

				Logarithm(tc.num, tc.base)
			} else {

				actual := Logarithm(tc.num, tc.base)

				if math.Abs(actual-tc.expected) > tc.tolerance {
					t.Errorf("Failure: Expected %v, got %v for num=%v, base=%v", tc.expected, actual, tc.num, tc.base)
				} else {
					t.Logf("Success: Expected %v, got %v within tolerance %v", tc.expected, actual, tc.tolerance)
				}
			}
		})
	}

}


/*
ROOST_METHOD_HASH=Modulo_7e9e651e69
ROOST_METHOD_SIG_HASH=Modulo_502e1458a3

FUNCTION_DEF=func Modulo(num1, num2 int) int // Modulo operation


*/
func TestModulo(t *testing.T) {

	type testCase struct {
		description string
		num1        int
		num2        int
		expected    int
		shouldPanic bool
	}

	tests := []testCase{
		{"Positive Dividend and Positive Divisor", 10, 3, 1, false},
		{"Positive Dividend and Negative Divisor", 10, -3, 1, false},
		{"Negative Dividend and Positive Divisor", -10, 3, -1, false},
		{"Both Dividend and Divisor Are Negative", -10, -3, -1, false},
		{"Dividend Is Zero", 0, 5, 0, false},
		{"Divisor Is One", 123, 1, 0, false},
		{"Divisor Is Negative One", 123, -1, 0, false},
		{"Dividend Equals the Divisor", 7, 7, 0, false},
		{"Dividend Is Smaller than the Divisor", 3, 10, 3, false},
		{"Dividend and Divisor Are Large Positive Integers", 1000000000, 123456, 43264, false},
		{"Divisor Is Zero (Error Case)", 10, 0, 0, true},
		{"Highly Negative Dividend with Positive Divisor", -1000000000, 7, -6, false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution: %v\n%s", r, string(debug.Stack()))
					if !tc.shouldPanic {
						t.Fail()
					}
				}
			}()

			var output string
			oldOut := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			defer func() {
				_ = w.Close()
				os.Stdout = oldOut
				b, _ := fmt.Fscanf(r, "%s")
				output = string(b)
				r.Close()
			}()

			actual := Modulo(tc.num1, tc.num2)

			if tc.shouldPanic {
				t.Logf("Expected panic for %q, but test passed without panic.", tc.description)
				t.Fail()
			} else if actual != tc.expected {
				t.Logf("Test failed for %q. Expected: %d, Got: %d", tc.description, tc.expected, actual)
				t.Fail()
			} else {
				t.Logf("Test passed for %q. Expected and got: %d", tc.description, tc.expected)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Multiply_7a2824e2c7
ROOST_METHOD_SIG_HASH=Multiply_0911ef76c1

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 // Multiply two floating-point numbers


*/
func TestMultiply(t *testing.T) {
	t.Log("Starting table-driven tests for the Multiply function.")

	type testCase struct {
		description string
		num1        float64
		num2        float64
		expected    float64
		shouldPanic bool
	}

	tests := []testCase{
		{"Multiply two positive floating-point numbers", 1.5, 2.0, 3.0, false},
		{"Multiply a positive number by zero", 4.7, 0.0, 0.0, false},
		{"Multiply zero by zero", 0.0, 0.0, 0.0, false},
		{"Multiply two negative floating-point numbers", -3.5, -2.0, 7.0, false},
		{"Multiply a positive number by a negative number", 5.0, -2.0, -10.0, false},
		{"Multiply two extremely large floating-point numbers", 1e10, 5e10, 5e20, false},
		{"Multiply two very small floating-point numbers", 1e-10, 5e-10, 5e-20, false},
		{"Multiply using floating-point precision edge cases", math.MaxFloat64, 2.0, math.MaxFloat64 * 2.0, false},
		{"Multiply using negative precision edge cases", -math.MaxFloat64, 3.0, -math.MaxFloat64 * 3.0, false},
		{"Multiply using NaN (Not-a-Number)", math.NaN(), 2.0, math.NaN(), false},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running test: %s", tc.description)
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered. Recovering test: %v\n%s", r, string(debug.Stack()))
					if !tc.shouldPanic {
						t.Fail()
					}
				}
			}()

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			result := Multiply(tc.num1, tc.num2)

			w.Close()
			os.Stdout = oldStdout
			var capturedOutput string
			fmt.Fscanf(r, "%s", &capturedOutput)

			if math.IsNaN(tc.expected) && math.IsNaN(result) {
				t.Logf("Success: Expected and result are both NaN.")
			} else if result != tc.expected {
				t.Errorf("Failure: For inputs (%v, %v), expected %v but got %v.",
					tc.num1, tc.num2, tc.expected, result)
			} else {
				t.Logf("Success: For inputs (%v, %v), got expected result %v.",
					tc.num1, tc.num2, result)
			}
		})
	}

	t.Log("All test cases executed.")
}


/*
ROOST_METHOD_HASH=Power_1c67a5d8b5
ROOST_METHOD_SIG_HASH=Power_c74b8edd76

FUNCTION_DEF=func Power(base, exponent float64) float64 // Power function


*/
func TestPower(t *testing.T) {

	tests := []struct {
		name     string
		base     float64
		exponent float64
		expected float64
	}{
		{
			name:     "Positive base to positive power",
			base:     2,
			exponent: 3,
			expected: 8,
		},
		{
			name:     "Positive base to negative power",
			base:     2,
			exponent: -2,
			expected: 0.25,
		},
		{
			name:     "Positive base to power of zero",
			base:     5,
			exponent: 0,
			expected: 1,
		},
		{
			name:     "Zero base to positive power",
			base:     0,
			exponent: 4,
			expected: 0,
		},
		{
			name:     "Zero base to power of zero",
			base:     0,
			exponent: 0,
			expected: 1,
		},
		{
			name:     "Positive base to fractional power",
			base:     9,
			exponent: 0.5,
			expected: 3,
		},
		{
			name:     "Negative base to even positive power",
			base:     -2,
			exponent: 4,
			expected: 16,
		},
		{
			name:     "Negative base to odd positive power",
			base:     -3,
			exponent: 3,
			expected: -27,
		},
		{
			name:     "Positive base to very large positive power",
			base:     2,
			exponent: 50,
			expected: 1125899906842624,
		},
		{
			name:     "Positive base to very large negative power",
			base:     2,
			exponent: -50,
			expected: math.Pow(2, -50),
		},
		{
			name:     "Fractional base and positive integer exponent",
			base:     0.5,
			exponent: 2,
			expected: 0.25,
		},
		{
			name:     "Extremely small fractional base and integer exponent",
			base:     0.0001,
			exponent: 3,
			expected: math.Pow(0.0001, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var buf bytes.Buffer
			stdOutBackup := os.Stdout
			os.Stdout = &buf
			defer func() { os.Stdout = stdOutBackup }()

			result := Power(tt.base, tt.exponent)

			_, _ = fmt.Fprintf(os.Stdout, "Running test case: %s\n", tt.name)

			if !almostEqual(result, tt.expected) {
				t.Errorf("Test '%s' failed: Power(%f, %f) = %f, expected %f", tt.name, tt.base, tt.exponent, result, tt.expected)
			} else {
				t.Logf("Test '%s' passed: Power(%f, %f) = %f", tt.name, tt.base, tt.exponent, result)
			}
		})
	}
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}


/*
ROOST_METHOD_HASH=SinCosTan_c242c1aa6d
ROOST_METHOD_SIG_HASH=SinCosTan_0f509380d6

FUNCTION_DEF=func SinCosTan(angle float64) (sin, cos, tan float64) // Trigonometric functions (Sin, Cos, Tan)


*/
func TestSinCosTan(t *testing.T) {

	tests := []struct {
		name        string
		inputAngle  float64
		expectedSin float64
		expectedCos float64
		expectedTan float64
	}{
		{
			name:        "Scenario 1: Common Angle π/4",
			inputAngle:  math.Pi / 4,
			expectedSin: math.Sqrt2 / 2,
			expectedCos: math.Sqrt2 / 2,
			expectedTan: 1.0,
		},
		{
			name:        "Scenario 2: Zero Input Angle",
			inputAngle:  0.0,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
		},
		{
			name:        "Scenario 3: Angle π/2",
			inputAngle:  math.Pi / 2,
			expectedSin: 1.0,
			expectedCos: 0.0,
			expectedTan: math.Inf(1),
		},
		{
			name:        "Scenario 4: Negative Angle -π/4",
			inputAngle:  -math.Pi / 4,
			expectedSin: -math.Sqrt2 / 2,
			expectedCos: math.Sqrt2 / 2,
			expectedTan: -1.0,
		},
		{
			name:        "Scenario 5: Large Angle 4π",
			inputAngle:  4 * math.Pi,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
		},
		{
			name:        "Scenario 6: Small Angle Near Zero",
			inputAngle:  0.0001,
			expectedSin: 0.0001,
			expectedCos: 1.0,
			expectedTan: 0.0001,
		},
		{
			name:        "Scenario 7: Angle Close to π",
			inputAngle:  math.Pi - 0.0001,
			expectedSin: math.Sin(math.Pi - 0.0001),
			expectedCos: math.Cos(math.Pi - 0.0001),
			expectedTan: math.Tan(math.Pi - 0.0001),
		},
		{
			name:        "Scenario 8: Very Large Negative Angle -10π",
			inputAngle:  -10 * math.Pi,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
		},
		{
			name:        "Scenario 9: Higher Precision Angle π/√2",
			inputAngle:  math.Pi / math.Sqrt(2),
			expectedSin: math.Sin(math.Pi / math.Sqrt(2)),
			expectedCos: math.Cos(math.Pi / math.Sqrt(2)),
			expectedTan: math.Tan(math.Pi / math.Sqrt(2)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			actualSin, actualCos, actualTan := SinCosTan(tt.inputAngle)

			const tolerance = 1e-6

			if math.Abs(actualSin-tt.expectedSin) > tolerance {
				t.Errorf("Test %s Failed: Expected sin %.6f, got %.6f", tt.name, tt.expectedSin, actualSin)
			}
			if math.Abs(actualCos-tt.expectedCos) > tolerance {
				t.Errorf("Test %s Failed: Expected cos %.6f, got %.6f", tt.name, tt.expectedCos, actualCos)
			}
			if tt.expectedTan == math.Inf(1) {
				if !math.IsInf(actualTan, 1) {
					t.Errorf("Test %s Failed: Expected tan to be infinity, got %.6f", tt.name, actualTan)
				}
			} else if math.Abs(actualTan-tt.expectedTan) > tolerance {
				t.Errorf("Test %s Failed: Expected tan %.6f, got %.6f", tt.name, tt.expectedTan, actualTan)
			}

			t.Logf("Test %s Passed: Expected and actual values matched within tolerance", tt.name)
		})
	}
}


/*
ROOST_METHOD_HASH=SquareRoot_17095d9165
ROOST_METHOD_SIG_HASH=SquareRoot_232943a56a

FUNCTION_DEF=func SquareRoot(num float64) float64 // Square root (with error handling)


*/
func TestSquareRoot(t *testing.T) {
	t.Run("Scenario 1: Square root of a positive number", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := 4.0
		expected := 2.0

		result := SquareRoot(input)
		if result != expected {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 2: Square root of zero", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := 0.0
		expected := 0.0

		result := SquareRoot(input)
		if result != expected {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 3: Square root of a negative number", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered as expected for negative input. Panic message: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				expected := "square root of a negative number is not defined"
				if fmt.Sprintf("%v", r) != expected {
					t.Errorf("Unexpected panic message for negative input: Expected %v, Got %v", expected, r)
				}
			}
		}()
		input := -4.0
		_ = SquareRoot(input)
		t.Fail()
	})

	t.Run("Scenario 4: Square root of a large positive number", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := 1e10
		expected := 1e5

		result := SquareRoot(input)
		if math.Abs(result-expected) > 1e-6 {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 5: Square root of a small positive number", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := 1e-10
		expected := 1e-5

		result := SquareRoot(input)
		if math.Abs(result-expected) > 1e-12 {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 6: Square root of a non-integer positive number", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := 2.25
		expected := 1.5

		result := SquareRoot(input)
		if result != expected {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 7: Square root of a number close to the machine limit", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := math.MaxFloat64
		expected := math.Sqrt(math.MaxFloat64)

		result := SquareRoot(input)
		if math.Abs(result-expected) > 1e-6 {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 8: Square root of a denormalized floating-point number", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		input := 5e-324
		expected := math.Sqrt(5e-324)

		result := SquareRoot(input)
		if math.Abs(result-expected) > 1e-12 {
			t.Errorf("Test failed for input %v: Expected %v, Got %v", input, expected, result)
		} else {
			t.Logf("Test successful for input %v: Expected %v, Got %v", input, expected, result)
		}
	})

	t.Run("Scenario 9: Concurrent execution of multiple inputs", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered during concurrent execution: %v\nStack Trace:\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()
		inputs := []float64{4.0, 9.0, 16.0}
		expectedResults := []float64{2.0, 3.0, 4.0}

		var wg sync.WaitGroup
		results := make([]float64, len(inputs))
		for i, input := range inputs {
			wg.Add(1)
			go func(i int, input float64) {
				defer wg.Done()
				results[i] = SquareRoot(input)
			}(i, input)
		}
		wg.Wait()

		for i, expected := range expectedResults {
			if results[i] != expected {
				t.Errorf("Concurrent test failed for input %v: Expected %v, Got %v", inputs[i], expected, results[i])
			} else {
				t.Logf("Concurrent test successful for input %v: Expected %v, Got %v", inputs[i], expected, results[i])
			}
		}
	})
}


/*
ROOST_METHOD_HASH=Subtract_58eac52f91
ROOST_METHOD_SIG_HASH=Subtract_b1211baa34

FUNCTION_DEF=func Subtract(num1, num2 int) int // Subtract two integers


*/
func TestSubtract(t *testing.T) {
	type testCase struct {
		name     string
		num1     int
		num2     int
		expected int
	}

	tests := []testCase{
		{"SubtractingTwoPositiveIntegers", 10, 4, 6},
		{"SubtractingNegativeIntegers", -5, -3, -2},
		{"SubtractPositiveFromNegative", -10, 3, -13},
		{"SubtractZeroFromPositive", 42, 0, 42},
		{"SubtractPositiveFromZero", 0, 7, -7},
		{"SubtractEqualIntegers", 15, 15, 0},
		{"HandlingLargePositiveIntegers", 1_000_000, 500_000, 500_000},
		{"HandlingLargeNegativeIntegers", -1_000_000, -500_000, -500_000},
		{"SubtractMinAndMaxIntegers", math.MinInt, math.MaxInt, math.MinInt - math.MaxInt},
		{"SubtractZeroEdgeWithMaxInt", math.MaxInt, 0, math.MaxInt},
		{"SubtractionProducingZero", 123456, 123456, 0},
		{"SubtractionOfPositiveAndNegativeIntegers", 500, -50, 550},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test:\nError: %v\nStack Trace:\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			outputFile, err := os.CreateTemp("", "test_subtract_output_*.txt")
			if err != nil {
				t.Fatalf("Failed to create temp file for output testing: %v", err)
			}
			defer func() {
				outputFile.Close()
				os.Remove(outputFile.Name())
			}()

			expected := tc.expected

			oldStdout := os.Stdout
			os.Stdout = outputFile
			actual := Subtract(tc.num1, tc.num2)
			os.Stdout = oldStdout

			if actual != expected {
				t.Errorf("Test %q failed: Subtract(%d, %d) => got %d, expected %d", tc.name, tc.num1, tc.num2, actual, expected)
			} else {
				t.Logf("Test %q succeeded: Subtract(%d, %d) returned correct result %d", tc.name, tc.num1, tc.num2, actual)
			}

			outputContent, _ := readTempFileContents(outputFile.Name())
			if strings.TrimSpace(outputContent) != fmt.Sprintf("%d", actual) {
				t.Errorf("Output mismatch for test %q: got `%s`, expected `%d`", tc.name, strings.TrimSpace(outputContent), actual)
			}
		})
	}
}

func readTempFileContents(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var contentBuilder strings.Builder
	_, err = fmt.Fscanf(file, "%s", &contentBuilder)
	if err != nil {
		return "", err
	}
	return contentBuilder.String(), nil
}


/*
ROOST_METHOD_HASH=Factorial_68fe6fb960
ROOST_METHOD_SIG_HASH=Factorial_3d037eec72

FUNCTION_DEF=func Factorial(n int) int // Factorial (Recursive)


*/
func TestFactorial(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        int
		expected     int
		expectPanic  bool
		panicMessage string
	}{
		{
			name:         "Scenario 1: Normal Positive Integer",
			input:        5,
			expected:     120,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 2: Factorial of Zero",
			input:        0,
			expected:     1,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 3: Factorial of One",
			input:        1,
			expected:     1,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 4: Negative Input",
			input:        -5,
			expected:     -1,
			expectPanic:  true,
			panicMessage: "factorial is not defined for negative numbers",
		},
		{
			name:         "Scenario 5: Large Positive Integer",
			input:        10,
			expected:     3628800,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 6: Smallest Positive Integer > 1",
			input:        2,
			expected:     2,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 7: Performance of Large Input",
			input:        20,
			expected:     2432902008176640000,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 8: Multiple Calls Test",
			input:        -1,
			expected:     -1,
			expectPanic:  false,
			panicMessage: "",
		},
		{
			name:         "Scenario 9: Factorial of Maximum Practical Integer",
			input:        15,
			expected:     1307674368000,
			expectPanic:  false,
			panicMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectPanic {
						if tt.panicMessage != "" && r != tt.panicMessage {
							t.Errorf("[FAIL] Expected panic message: %v, Got: %v", tt.panicMessage, r)
							t.Fail()
						}
						t.Logf("[PASS] Panic correctly encountered: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("[FAIL] Unexpected Panic Encountered: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			if tt.name == "Scenario 8: Multiple Calls Test" {
				concurrentInputs := []struct {
					input    int
					expected int
				}{
					{0, 1},
					{5, 120},
					{10, 3628800},
				}

				for _, ci := range concurrentInputs {
					result := Factorial(ci.input)
					if result != ci.expected {
						t.Errorf("[FAIL] For input %d, Expected: %d, Got: %d", ci.input, ci.expected, result)
					} else {
						t.Logf("[PASS] For input %d, factorial result validated as %d", ci.input, result)
					}
				}
				return
			}

			start := time.Now()
			actual := Factorial(tt.input)
			duration := time.Since(start)

			if tt.expectPanic {

				return
			}

			if actual != tt.expected {
				t.Errorf("[FAIL] Expected: %d, Got: %d", tt.expected, actual)
			} else {
				t.Logf("[PASS] Factorial of %d validated: %d", tt.input, actual)
			}
			t.Logf("Execution Time: %v", duration)
		})
	}
}


/*
ROOST_METHOD_HASH=GCD_1da681d86b
ROOST_METHOD_SIG_HASH=GCD_39a1228f3a

FUNCTION_DEF=func GCD(a, b int) int // Greatest Common Divisor (GCD) using Euclidean algorithm


*/
func TestGcd(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Positive Integers (56, 98)", 56, 98, 14},
		{"One Zero Input (0, 25)", 0, 25, 25},
		{"One Zero Input Reverse (25, 0)", 25, 0, 25},
		{"Equal Numbers (18, 18)", 18, 18, 18},
		{"Prime and Composite Numbers (7, 20)", 7, 20, 1},
		{"One Negative Input (-24, 36)", -24, 36, 12},
		{"Two Negative Inputs (-48, -18)", -48, -18, 6},
		{"Co-Prime Numbers (13, 27)", 13, 27, 1},
		{"Large Positive Integers (1234567, 7654321)", 1234567, 7654321, 1},
		{"Both Zeros (0, 0)", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing the test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var capturedOutput strings.Builder
			outputWriter := &capturedOutput
			os.Stdout = outputWriter

			actual := GCD(test.a, test.b)

			if actual != test.expected {
				t.Errorf("Test Failed for inputs (%d, %d). Expected: %d, Got: %d", test.a, test.b, test.expected, actual)
			} else {
				t.Logf("Test Passed for inputs (%d, %d): Expected: %d, Got: %d", test.a, test.b, test.expected, actual)
			}

			fmt.Fprintf(outputWriter, "Results validated successfully for [%s]\n", test.name)

			os.Stdout = os.Stdout
		})
	}
}


/*
ROOST_METHOD_HASH=LCM_85c2702b86
ROOST_METHOD_SIG_HASH=LCM_fb713f0b10

FUNCTION_DEF=func LCM(a, b int) int // Least Common Multiple (LCM) using GCD


*/
func TestLcm(t *testing.T) {

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "LCM of two positive integers",
			a:        12,
			b:        18,
			expected: 36,
		},
		{
			name:     "LCM of two prime numbers",
			a:        7,
			b:        11,
			expected: 77,
		},
		{
			name:     "LCM of positive and negative integers",
			a:        8,
			b:        -12,
			expected: 24,
		},
		{
			name:     "LCM of zero and a non-zero integer",
			a:        0,
			b:        10,
			expected: 0,
		},
		{
			name:     "LCM of two negative integers",
			a:        -6,
			b:        -9,
			expected: 18,
		},
		{
			name:     "LCM of identical numbers",
			a:        15,
			b:        15,
			expected: 15,
		},
		{
			name:     "LCM of coprime numbers",
			a:        8,
			b:        15,
			expected: 120,
		},
		{
			name:     "LCM of large integers",
			a:        123456,
			b:        789012,
			expected: 19330524,
		},
		{
			name:     "LCM of a number and one",
			a:        1,
			b:        9,
			expected: 9,
		},
		{
			name:     "LCM of two equal negative numbers",
			a:        -20,
			b:        -20,
			expected: 20,
		},
		{
			name:     "LCM of extremely small numbers (edge case)",
			a:        -1,
			b:        -3,
			expected: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Testing scenario: %s", tt.name)

			backupStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			result := LCM(tt.a, tt.b)

			w.Close()
			os.Stdout = backupStdout

			if result != tt.expected {
				t.Errorf("Test '%s' failed. LCM(%d, %d) = %d; expected %d",
					tt.name, tt.a, tt.b, result, tt.expected)
			} else {
				t.Logf("Test '%s' passed. LCM(%d, %d) = %d", tt.name, tt.a, tt.b, result)
			}

			output := make([]byte, 1024)
			n, _ := fmt.Fscanf(r, "%s")
			fmt.Fprintf(backupStdout, "Validating output (captured=%s, length=%d)", string(output[:n]), n)

		})
	}
}

