package calc

import (
	fmt "fmt"
	math "math"
	os "os"
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Absolute_d231f0ab10
ROOST_METHOD_SIG_HASH=Absolute_ec3c06e5a3

FUNCTION_DEF=func Absolute(num float64) float64 // Absolute value
*/
func TestAbsolute(t *testing.T) {

	tests := []struct {
		name   string
		input  float64
		output float64
	}{
		{name: "Positive Number Input", input: 10.5, output: 10.5},
		{name: "Negative Number Input", input: -25.75, output: 25.75},
		{name: "Zero Input", input: 0.0, output: 0.0},
		{name: "Small Positive Decimal Value", input: 0.000123, output: 0.000123},
		{name: "Small Negative Decimal Value", input: -0.000123, output: 0.000123},
		{name: "Large Positive Number", input: 1e15, output: 1e15},
		{name: "Large Negative Number", input: -1e15, output: 1e15},
		{name: "Positive Infinity Input", input: math.Inf(1), output: math.Inf(1)},
		{name: "Negative Infinity Input", input: math.Inf(-1), output: math.Inf(1)},
		{name: "NaN Input", input: math.NaN(), output: math.NaN()},
		{name: "Negative Zero Input", input: -0.0, output: 0.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test for '%s'. Reason: %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Absolute(tc.input)

			if math.IsNaN(tc.output) {

				if !math.IsNaN(result) {
					t.Errorf("Test '%s' failed: expected NaN, got %v", tc.name, result)
				} else {
					t.Logf("Test '%s' passed: correctly returned NaN", tc.name)
				}
			} else if result != tc.output {
				t.Errorf("Test '%s' failed: expected %v, got %v", tc.name, tc.output, result)
			} else {
				t.Logf("Test '%s' passed: expected %v, got %v", tc.name, tc.output, result)
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
	type testCase struct {
		name        string
		num1        int
		num2        int
		expected    int
		expectPanic bool
	}

	testCases := []testCase{
		{
			name:        "Adding two positive integers",
			num1:        3,
			num2:        5,
			expected:    8,
			expectPanic: false,
		},
		{
			name:        "Adding two negative integers",
			num1:        -4,
			num2:        -7,
			expected:    -11,
			expectPanic: false,
		},
		{
			name:        "Adding zero to a positive integer",
			num1:        0,
			num2:        9,
			expected:    9,
			expectPanic: false,
		},
		{
			name:        "Adding zero to a negative integer",
			num1:        0,
			num2:        -9,
			expected:    -9,
			expectPanic: false,
		},
		{
			name:        "Adding large integers near int max value",
			num1:        math.MaxInt32,
			num2:        100,
			expected:    math.MaxInt32 + 100,
			expectPanic: false,
		},
		{
			name:        "Adding opposite integers resulting in zero",
			num1:        15,
			num2:        -15,
			expected:    0,
			expectPanic: false,
		},
		{
			name:        "Adding zero twice",
			num1:        0,
			num2:        0,
			expected:    0,
			expectPanic: false,
		},
		{
			name:        "Adding a positive and smaller negative integer resulting in positive",
			num1:        12,
			num2:        -5,
			expected:    7,
			expectPanic: false,
		},
		{
			name:        "Adding a positive and larger negative integer resulting in negative",
			num1:        3,
			num2:        -10,
			expected:    -7,
			expectPanic: false,
		},
		{
			name:        "Adding extreme values near int min and max",
			num1:        math.MaxInt32,
			num2:        math.MinInt32,
			expected:    math.MaxInt32 + math.MinInt32,
			expectPanic: false,
		},
		{
			name:        "Adding a negative integer and zero",
			num1:        -8,
			num2:        0,
			expected:    -8,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Expected panic occurred: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic occurred: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			var output int

			output = Add(tc.num1, tc.num2)

			if !tc.expectPanic && output != tc.expected {
				t.Errorf("FAILED Test '%s': Add(%d, %d) = %d; expected %d",
					tc.name, tc.num1, tc.num2, output, tc.expected)
			} else {
				t.Logf("PASSED Test '%s': Add(%d, %d) = %d as expected", tc.name, tc.num1, tc.num2, output)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Divide_6fe509f399
ROOST_METHOD_SIG_HASH=Divide_d926fccfc9

FUNCTION_DEF=func Divide(num1, num2 float64) float64 // Divide two floating-point numbers (with error handling)
*/
func TestDivide(t *testing.T) {

	type testCase struct {
		desc         string
		num1         float64
		num2         float64
		expected     float64
		expectPanic  bool
		panicMessage string
	}

	testCases := []testCase{
		{
			desc:        "Scenario 1: Valid positive numbers (10.0 / 5.0)",
			num1:        10.0,
			num2:        5.0,
			expected:    2.0,
			expectPanic: false,
		},
		{
			desc:        "Scenario 2: Negative numerator (-10.0 / 5.0)",
			num1:        -10.0,
			num2:        5.0,
			expected:    -2.0,
			expectPanic: false,
		},
		{
			desc:        "Scenario 3: Negative denominator (10.0 / -5.0)",
			num1:        10.0,
			num2:        -5.0,
			expected:    -2.0,
			expectPanic: false,
		},
		{
			desc:        "Scenario 4: Both negative (-10.0 / -5.0)",
			num1:        -10.0,
			num2:        -5.0,
			expected:    2.0,
			expectPanic: false,
		},
		{
			desc:        "Scenario 5: Zero numerator (0.0 / 5.0)",
			num1:        0.0,
			num2:        5.0,
			expected:    0.0,
			expectPanic: false,
		},
		{
			desc:        "Scenario 6: Large values (1e8 / 2e4)",
			num1:        1e8,
			num2:        2e4,
			expected:    5000.0,
			expectPanic: false,
		},
		{
			desc:        "Scenario 7: Small values (1e-8 / 2e-4)",
			num1:        1e-8,
			num2:        2e-4,
			expected:    0.00005,
			expectPanic: false,
		},
		{
			desc:         "Scenario 8: Division by zero (10.0 / 0.0)",
			num1:         10.0,
			num2:         0.0,
			expectPanic:  true,
			panicMessage: "division by zero is not allowed",
		},
		{
			desc:        "Scenario 9: Division by infinity (10.0 / math.Inf(1))",
			num1:        10.0,
			num2:        math.Inf(1),
			expected:    0.0,
			expectPanic: false,
		},
	}

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Panic encountered: %v \n%s", r, string(debug.Stack()))
						panicMessage := fmt.Sprintf("%v", r)
						if panicMessage != tc.panicMessage {
							t.Errorf("Unexpected panic message. Got: %v, Expected: %v", panicMessage, tc.panicMessage)
						}
					} else {
						t.Logf("Unexpected panic: %v \n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			var actual float64
			actual = Divide(tc.num1, tc.num2)

			if !tc.expectPanic {
				if actual != tc.expected {
					t.Errorf("Test failed for %s. Got: %v, Expected: %v", tc.desc, actual, tc.expected)
				} else {
					t.Logf("Test passed for %s. Got: %v, Expected: %v", tc.desc, actual, tc.expected)
				}
			}
		})
	}

	w.Close()
	stdoutOutput := make([]byte, 1024)
	fmt.Fscanf(r, "%s", &stdoutOutput)
}

/*
ROOST_METHOD_HASH=Logarithm_4092f1cba7
ROOST_METHOD_SIG_HASH=Logarithm_0780d00fe8

FUNCTION_DEF=func Logarithm(num, base float64) float64 // Logarithm function (log_base of num)
*/
func TestLogarithm(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		num             float64
		base            float64
		expected        float64
		expectPanic     bool
		panicMessage    string
		validationNotes string
	}{

		{
			name:            "ValidPositiveInputs",
			num:             100,
			base:            10,
			expected:        2,
			expectPanic:     false,
			validationNotes: "This validates the function under normal positive inputs; expected output log_10(100) = 2.",
		},

		{
			name:            "NegativeNum",
			num:             -5,
			base:            2,
			expectPanic:     true,
			panicMessage:    "logarithm is not defined for these values",
			validationNotes: "Negative num should panic since logarithm is undefined for negative values.",
		},

		{
			name:            "ZeroBase",
			num:             100,
			base:            0,
			expectPanic:     true,
			panicMessage:    "logarithm is not defined for these values",
			validationNotes: "Base cannot be zero as logarithm is undefined.",
		},

		{
			name:            "BaseEqualsOne",
			num:             50,
			base:            1,
			expectPanic:     true,
			panicMessage:    "logarithm is not defined for these values",
			validationNotes: "Base of 1 is invalid for logarithms; should panic.",
		},

		{
			name:            "BaseLessThanNum",
			num:             8,
			base:            2,
			expected:        3,
			expectPanic:     false,
			validationNotes: "Run log_2(8), result must be 3 as per mathematical rules.",
		},

		{
			name:            "BaseGreaterThanNum",
			num:             2,
			base:            10,
			expected:        math.Log(2) / math.Log(10),
			expectPanic:     false,
			validationNotes: "For base > num, log_10(2), result checks range (0, 1).",
		},

		{
			name:            "NumEqualsBase",
			num:             5,
			base:            5,
			expected:        1,
			expectPanic:     false,
			validationNotes: "Logarithmic identity log_b(b) = 1, verifies correctness.",
		},

		{
			name:            "LargeValues",
			num:             1e10,
			base:            1e5,
			expected:        math.Log(1e10) / math.Log(1e5),
			expectPanic:     false,
			validationNotes: "Tests ability to handle large floating-point values without precision issues.",
		},

		{
			name:            "SmallValues",
			num:             0.0001,
			base:            0.01,
			expected:        math.Log(0.0001) / math.Log(0.01),
			expectPanic:     false,
			validationNotes: "Checks function precision for small magnitude values.",
		},

		{
			name:            "IrrationalNumbers",
			num:             math.Pi,
			base:            math.E,
			expected:        math.Log(math.Pi) / math.Log(math.E),
			expectPanic:     false,
			validationNotes: "Tests mathematical versatility with irrational values like π and e.",
		},

		{
			name:            "ZeroNum",
			num:             0,
			base:            2,
			expectPanic:     true,
			panicMessage:    "logarithm is not defined for these values",
			validationNotes: "Num=0 case asserts validity of error handling for undefined logarithm.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Test Case: %s - %s", tt.name, tt.validationNotes)

			defer func() {
				if r := recover(); r != nil {
					if tt.expectPanic {
						t.Logf("Expected panic encountered: %v\nStack Trace:\n%s", r, string(debug.Stack()))
						if r != tt.panicMessage {
							t.Errorf("Panic message mismatch: got '%v', expected '%v'", r, tt.panicMessage)
						}
					} else {
						t.Logf("Unexpected panic: %v\nStack Trace:\n%s", r, string(debug.Stack()))
						t.FailNow()
					}
				} else if tt.expectPanic {
					t.Errorf("Failed to capture expected panic for case: %s", tt.name)
				}
			}()

			stdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			defer func() { os.Stdout = stdout }()

			if !tt.expectPanic {
				result := Logarithm(tt.num, tt.base)
				if math.Abs(result-tt.expected) > 1e-9 {
					t.Errorf("Result mismatch: got %v, expected %v", result, tt.expected)
				} else {
					t.Logf("Logarithm function produced correct result: %v\n", result)
				}
			} else {
				_ = Logarithm(tt.num, tt.base)
				t.Log("Executed panic creating test case")
			}

			w.Close()
			var validatedOutput string
			fmt.Fscanf(r, "%s", &validatedOutput)
			r.Close()
		})
	}
}

/*
ROOST_METHOD_HASH=Multiply_7a2824e2c7
ROOST_METHOD_SIG_HASH=Multiply_0911ef76c1

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 // Multiply two floating-point numbers
*/
func TestMultiply(t *testing.T) {

	tests := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
		validate func(result, expected float64) bool
	}{
		{
			name:     "Scenario 1: Multiply two positive floating-point numbers",
			num1:     5.2,
			num2:     4.3,
			expected: 22.36,
			validate: func(result, expected float64) bool { return math.Abs(result-expected) < 1e-10 },
		},
		{
			name:     "Scenario 2: Multiply a positive number by zero",
			num1:     7.5,
			num2:     0,
			expected: 0,
			validate: func(result, expected float64) bool { return result == expected },
		},
		{
			name:     "Scenario 3: Multiply zero by zero",
			num1:     0,
			num2:     0,
			expected: 0,
			validate: func(result, expected float64) bool { return result == expected },
		},
		{
			name:     "Scenario 4: Multiply two negative floating-point numbers",
			num1:     -3.4,
			num2:     -2.2,
			expected: 7.48,
			validate: func(result, expected float64) bool { return math.Abs(result-expected) < 1e-10 },
		},
		{
			name:     "Scenario 5: Multiply a positive floating-point number by a negative floating-point number",
			num1:     6.5,
			num2:     -2.3,
			expected: -14.95,
			validate: func(result, expected float64) bool { return math.Abs(result-expected) < 1e-10 },
		},
		{
			name:     "Scenario 6: Multiply two extremely large floating-point numbers",
			num1:     1e308,
			num2:     1e308,
			expected: math.Inf(1),
			validate: func(result, expected float64) bool { return math.IsInf(result, 1) },
		},
		{
			name:     "Scenario 7: Multiply two small floating-point numbers",
			num1:     1e-308,
			num2:     2e-308,
			expected: 0,
			validate: func(result, expected float64) bool { return result == expected || math.Abs(result-expected) < 1e-308 },
		},
		{
			name:     "Scenario 8: Multiply a floating-point number by one",
			num1:     5.72,
			num2:     1,
			expected: 5.72,
			validate: func(result, expected float64) bool { return math.Abs(result-expected) < 1e-10 },
		},
		{
			name:     "Scenario 9: Multiply floating-point numbers of opposite signs yielding fractional results",
			num1:     -2.3,
			num2:     0.5,
			expected: -1.15,
			validate: func(result, expected float64) bool { return math.Abs(result-expected) < 1e-10 },
		},
		{
			name:     "Scenario 10: Multiply numbers where one operand is NaN (Not a Number)",
			num1:     math.NaN(),
			num2:     3.5,
			expected: math.NaN(),
			validate: func(result, expected float64) bool { return math.IsNaN(result) },
		},
		{
			name:     "Scenario 11: Multiply numbers where one operand is Infinity",
			num1:     math.Inf(1),
			num2:     2.1,
			expected: math.Inf(1),
			validate: func(result, expected float64) bool { return math.IsInf(result, 1) },
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

			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			defer func() {
				os.Stdout = originalStdout
				w.Close()
			}()

			result := Multiply(tt.num1, tt.num2)

			var output string
			fmt.Fscanf(r, "%s", &output)

			if !tt.validate(result, tt.expected) {
				t.Errorf("Test '%s' failed: expected %v but got %v. Output: %s", tt.name, tt.expected, result, output)
			} else {
				t.Logf("Test '%s' passed. Result: %v, Expected: %v", tt.name, result, tt.expected)
			}

		})
	}
}

/*
ROOST_METHOD_HASH=Power_1c67a5d8b5
ROOST_METHOD_SIG_HASH=Power_c74b8edd76

FUNCTION_DEF=func Power(base, exponent float64) float64 // Power function
*/
func TestPower(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		base           float64
		exponent       float64
		expectedResult float64
		shouldPanic    bool
	}{
		{
			name:           "Positive Base with Positive Exponent",
			base:           2,
			exponent:       3,
			expectedResult: 8,
			shouldPanic:    false,
		},
		{
			name:           "Zero Base with Positive Exponent",
			base:           0,
			exponent:       5,
			expectedResult: 0,
			shouldPanic:    false,
		},
		{
			name:           "Positive Base with Zero Exponent",
			base:           2,
			exponent:       0,
			expectedResult: 1,
			shouldPanic:    false,
		},
		{
			name:           "Zero Base with Zero Exponent",
			base:           0,
			exponent:       0,
			expectedResult: 1,
			shouldPanic:    false,
		},
		{
			name:           "Negative Base with Positive Exponent",
			base:           -2,
			exponent:       3,
			expectedResult: -8,
			shouldPanic:    false,
		},
		{
			name:           "Positive Base with Negative Exponent",
			base:           2,
			exponent:       -2,
			expectedResult: 0.25,
			shouldPanic:    false,
		},
		{
			name:           "Negative Base with Negative Exponent",
			base:           -2,
			exponent:       -2,
			expectedResult: 0.25,
			shouldPanic:    false,
		},
		{
			name:           "Non-Integer Exponent (Fractional Power)",
			base:           16,
			exponent:       0.5,
			expectedResult: 4,
			shouldPanic:    false,
		},
		{
			name:           "Large Base with Small Exponent",
			base:           1e6,
			exponent:       2,
			expectedResult: 1e12,
			shouldPanic:    false,
		},
		{
			name:           "Small Base with Large Exponent",
			base:           0.1,
			exponent:       10,
			expectedResult: 1e-10,
			shouldPanic:    false,
		},
		{
			name:           "Small Base with Negative Exponent",
			base:           0.1,
			exponent:       -2,
			expectedResult: 100,
			shouldPanic:    false,
		},
		{
			name:           "Negative Exponent Edge Case (Exponent -1)",
			base:           5,
			exponent:       -1,
			expectedResult: 0.2,
			shouldPanic:    false,
		},
		{
			name:           "Handling Very Large Exponents",
			base:           2,
			exponent:       1e6,
			expectedResult: math.Inf(1),
			shouldPanic:    true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if test.shouldPanic {
						t.Logf("Panic handled successfully for test: %s. Error: %v\n%s", test.name, r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic for test: %s. Error: %v\n%s", test.name, r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			result := Power(test.base, test.exponent)

			if !test.shouldPanic {
				if math.Abs(result-test.expectedResult) > 1e-9 {
					t.Errorf("Test '%s' failed. Expected: %v, Got: %v", test.name, test.expectedResult, result)
				} else {
					t.Logf("Test '%s' passed successfully. Result: %v", test.name, result)
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=SinCosTan_c242c1aa6d
ROOST_METHOD_SIG_HASH=SinCosTan_0f509380d6

FUNCTION_DEF=func SinCosTan(angle float64) (sin, cos, tan float64) // Trigonometric functions (Sin, Cos, Tan)
*/
func TestSinCosTan(t *testing.T) {

	type testCase struct {
		name         string
		input        float64
		expectedSin  float64
		expectedCos  float64
		expectedTan  float64
		isTanSpecial bool
	}

	tests := []testCase{
		{
			name:         "Scenario 1: Trigonometric values for zero angle",
			input:        0,
			expectedSin:  0,
			expectedCos:  1,
			expectedTan:  0,
			isTanSpecial: false,
		},
		{
			name:         "Scenario 2: Trigonometric values for π/2 radians",
			input:        math.Pi / 2,
			expectedSin:  1,
			expectedCos:  0,
			expectedTan:  math.MaxFloat64,
			isTanSpecial: true,
		},
		{
			name:         "Scenario 3: Trigonometric values for π radians",
			input:        math.Pi,
			expectedSin:  0,
			expectedCos:  -1,
			expectedTan:  0,
			isTanSpecial: false,
		},
		{
			name:         "Scenario 4: Trigonometric values for 3π/2 radians",
			input:        3 * math.Pi / 2,
			expectedSin:  -1,
			expectedCos:  0,
			expectedTan:  -math.MaxFloat64,
			isTanSpecial: true,
		},
		{
			name:         "Scenario 5: Trigonometric values for periodic input (>2π)",
			input:        2*math.Pi + math.Pi/4,
			expectedSin:  math.Sin(math.Pi / 4),
			expectedCos:  math.Cos(math.Pi / 4),
			expectedTan:  math.Tan(math.Pi / 4),
			isTanSpecial: false,
		},
		{
			name:         "Scenario 6: Trigonometric values for negative angles",
			input:        -math.Pi / 4,
			expectedSin:  -math.Sin(math.Pi / 4),
			expectedCos:  math.Cos(math.Pi / 4),
			expectedTan:  -math.Tan(math.Pi / 4),
			isTanSpecial: false,
		},
		{
			name:         "Scenario 7: Handling small angle approximations",
			input:        0.0001,
			expectedSin:  0.0001,
			expectedCos:  1,
			expectedTan:  0.0001,
			isTanSpecial: false,
		},
		{
			name:         "Scenario 8: Handling large negative angles",
			input:        -3 * math.Pi,
			expectedSin:  0,
			expectedCos:  -1,
			expectedTan:  0,
			isTanSpecial: false,
		},
		{
			name:         "Scenario 9: Edge case for undefined tangent (π/2 boundary)",
			input:        math.Pi/2 - 0.0000001,
			expectedSin:  math.Sin(math.Pi/2 - 0.0000001),
			expectedCos:  math.Cos(math.Pi/2 - 0.0000001),
			expectedTan:  math.Tan(math.Pi/2 - 0.0000001),
			isTanSpecial: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, failing test: %v\n%s", r, string(debug.Stack()))
					t.FailNow()
				}
			}()

			sin, cos, tan := SinCosTan(tc.input)

			if math.Abs(sin-tc.expectedSin) > 1e-8 {
				t.Errorf("[FAIL] Sin mismatch for input=%f: expected=%f, got=%f", tc.input, tc.expectedSin, sin)
			} else {
				t.Logf("[PASS] Sin validated for input=%f", tc.input)
			}

			if math.Abs(cos-tc.expectedCos) > 1e-8 {
				t.Errorf("[FAIL] Cos mismatch for input=%f: expected=%f, got=%f", tc.input, tc.expectedCos, cos)
			} else {
				t.Logf("[PASS] Cos validated for input=%f", tc.input)
			}

			if tc.isTanSpecial {
				if tan < math.MaxFloat64 && tan > -math.MaxFloat64 {
					t.Logf("[PASS] Tan approximated correctly for input=%f", tc.input)
				} else {
					t.Errorf("[FAIL] Tan incorrect bounds at input=%f: got=%f", tc.input, tan)
				}
			} else {
				if math.Abs(tan-tc.expectedTan) > 1e-8 {
					t.Errorf("[FAIL] Tan mismatch for input=%f: expected=%f, got=%f", tc.input, tc.expectedTan, tan)
				} else {
					t.Logf("[PASS] Tan validated for input=%f", tc.input)
				}
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
	type testCase struct {
		name        string
		input       float64
		want        float64
		expectPanic bool
	}

	var testCases = []testCase{
		{name: "Positive Number", input: 16.0, want: 4.0, expectPanic: false},
		{name: "Zero", input: 0.0, want: 0.0, expectPanic: false},
		{name: "Small Positive Number", input: 0.25, want: 0.5, expectPanic: false},
		{name: "Large Positive Number", input: 1000000.0, want: 1000.0, expectPanic: false},
		{name: "Negative Number", input: -4.0, want: 0, expectPanic: true},
		{name: "Extremely Small Positive Number", input: 1e-10, want: 1e-5, expectPanic: false},
		{name: "Non-Integer Positive Floating Point", input: 20.25, want: 4.5, expectPanic: false},
		{name: "Minimum Positive Float", input: math.SmallestNonzeroFloat64, want: math.Sqrt(math.SmallestNonzeroFloat64), expectPanic: false},
		{name: "Maximum Float", input: math.MaxFloat64, want: math.Sqrt(math.MaxFloat64), expectPanic: false},
		{name: "NaN Input", input: math.NaN(), want: math.NaN(), expectPanic: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Logf("Panic encountered for %s unexpectedly: %v\n%s", tc.name, r, string(debug.Stack()))
						t.Fail()
					} else {
						t.Logf("Expected panic observed for %s: %v", tc.name, r)
					}
				}
			}()

			var got float64
			if !tc.expectPanic {
				got = SquareRoot(tc.input)
			}

			if tc.expectPanic {
				t.Logf("Test handled panic scenario for %s", tc.name)
				return
			}

			if math.IsNaN(tc.want) {
				if !math.IsNaN(got) {
					t.Errorf("Test %s failed: expected NaN, got %v", tc.name, got)
				} else {
					t.Logf("Test %s passed with NaN handling", tc.name)
				}
			} else if math.Abs(got-tc.want) > 1e-9 {

				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.want, got)
			} else {
				t.Logf("Test %s passed: expected %v, got %v", tc.name, tc.want, got)
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
	type testCase struct {
		name        string
		input       int
		expected    int
		expectPanic bool
	}

	testCases := []testCase{
		{
			name:        "Factorial of a Positive Number",
			input:       5,
			expected:    120,
			expectPanic: false,
		},
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
			name:        "Factorial of a Large Positive Number",
			input:       10,
			expected:    3628800,
			expectPanic: false,
		},
		{
			name:        "Factorial of a Negative Number",
			input:       -5,
			expected:    0,
			expectPanic: true,
		},

		{
			name:        "Factorial for Maximum Integer (Performance Limits)",
			input:       math.MaxInt8,
			expected:    39916800,
			expectPanic: false,
		},
		{
			name:        "Factorial of a Very Small Number Close to Zero",
			input:       2,
			expected:    2,
			expectPanic: false,
		},
		{
			name:        "Sequential Recursive Calls for Small Input",
			input:       3,
			expected:    6,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					if !tc.expectPanic {
						t.Fatalf("Test failed unexpectedly - panic not expected")
					} else {
						t.Logf("Panic as expected for input %d", tc.input)
					}
				}
			}()

			result := 0
			if !tc.expectPanic {
				result = Factorial(tc.input)
				if result != tc.expected {
					t.Errorf("Failed: Input(%d), Expected(%d), Got(%d)", tc.input, tc.expected, result)
				} else {
					t.Logf("Success: Input(%d), Expected(%d), Got(%d)", tc.input, tc.expected, result)
				}
			} else {
				t.Log("Expecting panic, no assertion required for value")
				Factorial(tc.input)
			}
		})
	}

	t.Run("Verify Recursive Results Sequentially", func(t *testing.T) {
		seqInputs := []int{2, 3, 4}
		expectedResults := []int{2, 6, 24}

		for i, input := range seqInputs {
			tcName := fmt.Sprintf("Factorial of Sequential Input %d", input)
			t.Run(tcName, func(t *testing.T) {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Unexpected panic for input %d: %v", input, r)
					}
				}()

				result := Factorial(input)
				expected := expectedResults[i]
				if result != expected {
					t.Errorf("Failed: Input(%d), Expected(%d), Got(%d)", input, expected, result)
				} else {
					t.Logf("Success: Input(%d), Expected(%d), Got(%d)", input, expected, result)
				}
			})
		}
	})
}

/*
ROOST_METHOD_HASH=GCD_1da681d86b
ROOST_METHOD_SIG_HASH=GCD_39a1228f3a

FUNCTION_DEF=func GCD(a, b int) int // Greatest Common Divisor (GCD) using Euclidean algorithm
*/
func TestGcd(t *testing.T) {
	type testCase struct {
		name     string
		a        int
		b        int
		expected int
	}

	testCases := []testCase{
		{"GCD of Two Positive Integers", 48, 18, 6},
		{"GCD of Zero and a Positive Integer", 0, 56, 56},
		{"GCD of Two Negative Integers", -48, -18, 6},
		{"GCD of Zero and Zero", 0, 0, 0},
		{"GCD of One and Any Integer", 1, 99, 1},
		{"GCD of Two Co-Prime Integers", 13, 7, 1},
		{"GCD of Two Large Integers", 12345678, 87654321, 9},
		{"GCD of Two Identical Numbers", 42, 42, 42},
		{"GCD of One Negative and One Positive Integer", -48, 18, 6},
		{"GCD of a Prime Number and a Composite Number", 17, 20, 1},
		{"GCD for One Large and One Small Integer", 100000000, 25, 25},
		{"GCD of a Multiplicative Pair", 15, 45, 15},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			defer func() {

				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			originalStdout := os.Stdout
			defer func() { os.Stdout = originalStdout }()
			stdoutR, stdoutW, _ := os.Pipe()
			os.Stdout = stdoutW

			result := GCD(tc.a, tc.b)

			stdoutW.Close()
			var actualOutput string
			fmt.Fscanf(stdoutR, "%s", &actualOutput)

			if result != tc.expected {
				t.Errorf("Test failed for %s: Got %d, Expected %d", tc.name, result, tc.expected)
			} else {
				t.Logf("Test passed for %s: Got %d as expected", tc.name, result)
			}
			stdoutR.Close()

		})
	}
}

/*
ROOST_METHOD_HASH=LCM_85c2702b86
ROOST_METHOD_SIG_HASH=LCM_fb713f0b10

FUNCTION_DEF=func LCM(a, b int) int // Least Common Multiple (LCM) using GCD
*/
func TestLcm(t *testing.T) {
	type testCase struct {
		name        string
		a, b        int
		expected    int
		shouldPanic bool
	}

	tests := []testCase{
		{
			name:     "Scenario 1: LCM for two positive integers",
			a:        6,
			b:        8,
			expected: 24,
		},
		{
			name:     "Scenario 2: LCM when one input is zero",
			a:        0,
			b:        15,
			expected: 0,
		},
		{
			name:     "Scenario 3: LCM when both inputs are the same",
			a:        7,
			b:        7,
			expected: 7,
		},
		{
			name:     "Scenario 4: LCM for two prime numbers",
			a:        13,
			b:        17,
			expected: 221,
		},
		{
			name:     "Scenario 5: LCM for one negative and one positive integer",
			a:        -4,
			b:        10,
			expected: 20,
		},
		{
			name:     "Scenario 6: LCM for very large numbers",
			a:        1000000,
			b:        1234567,
			expected: 1234567000000 / GCD(1000000, 1234567),
		},
		{
			name:     "Scenario 7: LCM for two negative integers",
			a:        -9,
			b:        -12,
			expected: 36,
		},
		{
			name:     "Scenario 8: LCM for one number being a multiple of the other",
			a:        4,
			b:        12,
			expected: 12,
		},
		{
			name:     "Scenario 9: LCM for two small integers",
			a:        3,
			b:        5,
			expected: 15,
		},
		{
			name:     "Scenario 10: LCM for input where GCD requires multiple steps to compute",
			a:        27,
			b:        36,
			expected: 108,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			defer func() {
				os.Stdout = old
			}()
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected a panic but did not receive one")
					}
				}()
			}

			result := LCM(tc.a, tc.b)

			w.Close()
			var output string
			fmt.Fscan(r, &output)

			if result != tc.expected {
				t.Errorf("Test failed for %s: got %d, expected %d", tc.name, result, tc.expected)
			} else {
				t.Logf("Test passed for %s: got %d, expected %d", tc.name, result, tc.expected)
			}

		})
	}
}
