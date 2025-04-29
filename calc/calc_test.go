package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
)

const float64EqualityThreshold = 1e-9

/*
ROOST_METHOD_HASH=Add_61b83cbd8d
ROOST_METHOD_SIG_HASH=Add_9b37ae6611

FUNCTION_DEF=func Add(num1, num2 int) int // Add two integers
*/
func TestAdd(t *testing.T) {

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{

		{
			name:     "Scenario 1: Positive + Positive",
			num1:     5,
			num2:     10,
			expected: 15,
		},

		{
			name:     "Scenario 2: Negative + Negative",
			num1:     -5,
			num2:     -10,
			expected: -15,
		},

		{
			name:     "Scenario 3: Positive + Negative (Positive Result)",
			num1:     10,
			num2:     -5,
			expected: 5,
		},

		{
			name:     "Scenario 4: Positive + Negative (Negative Result)",
			num1:     -10,
			num2:     5,
			expected: -5,
		},

		{
			name:     "Scenario 5: Positive + Zero",
			num1:     7,
			num2:     0,
			expected: 7,
		},

		{
			name:     "Scenario 6: Negative + Zero",
			num1:     -7,
			num2:     0,
			expected: -7,
		},

		{
			name:     "Scenario 7: Zero + Zero",
			num1:     0,
			num2:     0,
			expected: 0,
		},

		{
			name:     "Scenario 8: Near MaxInt (No Overflow)",
			num1:     math.MaxInt - 5,
			num2:     3,
			expected: math.MaxInt - 2,
		},

		{
			name:     "Scenario 9: Near MinInt (No Underflow)",
			num1:     math.MinInt + 5,
			num2:     -3,
			expected: math.MinInt + 2,
		},

		{
			name:     "Scenario 10: Potential Overflow (MaxInt + 1)",
			num1:     math.MaxInt,
			num2:     1,
			expected: math.MinInt,
		},

		{
			name:     "Scenario 11: Potential Underflow (MinInt - 1)",
			num1:     math.MinInt,
			num2:     -1,
			expected: math.MaxInt,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: num1 = %d, num2 = %d", tc.num1, tc.num2)

			actual := Add(tc.num1, tc.num2)

			if actual != tc.expected {

				t.Errorf("FAIL: Expected result %d, but got %d", tc.expected, actual)
			} else {

				t.Logf("PASS: Expected result %d, got %d", tc.expected, actual)
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

	testCases := []struct {
		name             string
		num1             float64
		num2             float64
		expectedResult   float64
		expectPanic      bool
		expectedPanicMsg string
		useTolerance     bool
		tolerance        float64
		checkNaN         bool
		checkInf         int
	}{

		{
			name:           "Scenario 1: Positive Integers",
			num1:           10.0,
			num2:           2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 2: Non-Integer Float Result",
			num1:           5.0,
			num2:           2.0,
			expectedResult: 2.5,
			expectPanic:    false,
		},

		{
			name:           "Scenario 3: Negative Numerator",
			num1:           -10.0,
			num2:           2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 4: Negative Denominator",
			num1:           10.0,
			num2:           -2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 5: Both Negative",
			num1:           -10.0,
			num2:           -2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 6: Zero Numerator",
			num1:           0.0,
			num2:           5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			name:             "Scenario 7: Division By Zero Panic",
			num1:             10.0,
			num2:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "division by zero is not allowed",
		},

		{
			name:           "Scenario 8: Repeating Decimal Precision",
			num1:           1.0,
			num2:           3.0,
			expectedResult: 0.3333333333333333,
			expectPanic:    false,
			useTolerance:   true,
			tolerance:      1e-9,
		},

		{
			name:           "Scenario 9: Large Numbers",
			num1:           math.MaxFloat64,
			num2:           2.0,
			expectedResult: math.MaxFloat64 / 2.0,
			expectPanic:    false,

			useTolerance: true,
			tolerance:    1e-9,
		},

		{
			name:        "Scenario 10: Positive Infinity Result",
			num1:        1.0,
			num2:        math.SmallestNonzeroFloat64,
			expectPanic: false,
			checkInf:    1,
		},

		{
			name:        "Scenario 11: Negative Infinity Result",
			num1:        -1.0,
			num2:        math.SmallestNonzeroFloat64,
			expectPanic: false,
			checkInf:    -1,
		},

		{
			name:        "Scenario 12: NaN Input",
			num1:        math.NaN(),
			num2:        5.0,
			expectPanic: false,
			checkNaN:    true,
		},

		{
			name:             "Scenario 13: Zero Divided By Zero Panic",
			num1:             0.0,
			num2:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "division by zero is not allowed",
		},

		{
			name:        "Edge Case: NaN Denominator",
			num1:        5.0,
			num2:        math.NaN(),
			expectPanic: false,
			checkNaN:    true,
		},

		{
			name:        "Edge Case: Positive Infinity Numerator",
			num1:        math.Inf(1),
			num2:        2.0,
			expectPanic: false,
			checkInf:    1,
		},

		{
			name:           "Edge Case: Positive Infinity Denominator",
			num1:           10.0,
			num2:           math.Inf(1),
			expectedResult: 0.0,
			expectPanic:    false,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.expectPanic {
					if r == nil {

						t.Errorf("FAIL: Expected a panic for input (%v, %v) but did not get one.", tc.num1, tc.num2)
					} else {

						panicMsg, ok := r.(string)
						if !ok || panicMsg != tc.expectedPanicMsg {
							t.Errorf("FAIL: Panic message mismatch for input (%v, %v). Got: %v (%T), Want: %q", tc.num1, tc.num2, r, r, tc.expectedPanicMsg)
						} else {
							t.Logf("PASS: Correctly panicked with message %q for input (%v, %v).", tc.expectedPanicMsg, tc.num1, tc.num2)
						}
					}
				} else if r != nil {

					t.Errorf("FAIL: Did not expect panic for input (%v, %v), but got: %v\n%s", tc.num1, tc.num2, r, string(debug.Stack()))
				}

			}()

			t.Logf("Running test: %s - Input: num1=%v, num2=%v", tc.name, tc.num1, tc.num2)

			actualResult := Divide(tc.num1, tc.num2)

			if !tc.expectPanic {

				if tc.checkNaN {
					if !math.IsNaN(actualResult) {
						t.Errorf("FAIL: Expected NaN for input (%v, %v), but got %v", tc.num1, tc.num2, actualResult)
					} else {
						t.Logf("PASS: Correctly resulted in NaN for input (%v, %v).", tc.num1, tc.num2)
					}
					return
				}

				if tc.checkInf != 0 {
					if !math.IsInf(actualResult, tc.checkInf) {
						expectedInfStr := "+Inf"
						if tc.checkInf == -1 {
							expectedInfStr = "-Inf"
						}
						t.Errorf("FAIL: Expected %s for input (%v, %v), but got %v", expectedInfStr, tc.num1, tc.num2, actualResult)
					} else {
						expectedInfStr := "+Inf"
						if tc.checkInf == -1 {
							expectedInfStr = "-Inf"
						}
						t.Logf("PASS: Correctly resulted in %s for input (%v, %v).", expectedInfStr, tc.num1, tc.num2)
					}
					return
				}

				if tc.useTolerance {
					diff := math.Abs(actualResult - tc.expectedResult)
					if diff >= tc.tolerance {
						t.Errorf("FAIL: Result outside tolerance for input (%v, %v). Got: %v, Want: %v (Tolerance: %v, Diff: %v)", tc.num1, tc.num2, actualResult, tc.expectedResult, tc.tolerance, diff)
					} else {
						t.Logf("PASS: Result within tolerance for input (%v, %v). Got: %v, Want: %v (Tolerance: %v)", tc.num1, tc.num2, actualResult, tc.expectedResult, tc.tolerance)
					}
				} else {

					if actualResult != tc.expectedResult {
						t.Errorf("FAIL: Exact result mismatch for input (%v, %v). Got: %v, Want: %v", tc.num1, tc.num2, actualResult, tc.expectedResult)
					} else {
						t.Logf("PASS: Correct exact result for input (%v, %v). Got: %v", tc.num1, tc.num2, actualResult)
					}
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

	testCases := []struct {
		name           string
		num1           int
		num2           int
		expectedResult int
		expectPanic    bool
		description    string
	}{
		{
			name:           "Scenario 1: Basic Positive Modulo",
			num1:           10,
			num2:           3,
			expectedResult: 1,
			expectPanic:    false,
			description:    "Test modulo with positive dividend > divisor.",
		},
		{
			name:           "Scenario 2: Positive Modulo Dividend Smaller",
			num1:           3,
			num2:           10,
			expectedResult: 3,
			expectPanic:    false,
			description:    "Test modulo with positive dividend < divisor.",
		},
		{
			name:           "Scenario 3: Zero Dividend",
			num1:           0,
			num2:           5,
			expectedResult: 0,
			expectPanic:    false,
			description:    "Test modulo with zero dividend and positive divisor.",
		},
		{
			name:           "Scenario 4: Zero Result",
			num1:           10,
			num2:           5,
			expectedResult: 0,
			expectPanic:    false,
			description:    "Test modulo where dividend is a multiple of divisor.",
		},
		{
			name:           "Scenario 5: Negative Dividend",
			num1:           -10,
			num2:           3,
			expectedResult: -1,
			expectPanic:    false,
			description:    "Test modulo with negative dividend and positive divisor.",
		},
		{
			name:           "Scenario 6: Negative Divisor",
			num1:           10,
			num2:           -3,
			expectedResult: 1,
			expectPanic:    false,
			description:    "Test modulo with positive dividend and negative divisor.",
		},
		{
			name:           "Scenario 7: Negative Dividend and Divisor",
			num1:           -10,
			num2:           -3,
			expectedResult: -1,
			expectPanic:    false,
			description:    "Test modulo with both negative dividend and divisor.",
		},
		{
			name:           "Scenario 8: Zero Divisor (Expect Panic)",
			num1:           10,
			num2:           0,
			expectedResult: 0,
			expectPanic:    true,
			description:    "Test modulo with zero divisor, expecting a runtime panic.",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.expectPanic {
					if r == nil {

						t.Errorf("FAIL: %s. Expected a panic for Modulo(%d, %d), but did not get one.", tc.description, tc.num1, tc.num2)
					} else {

						t.Logf("PASS: %s. Successfully caught expected panic for Modulo(%d, %d). Panic: %v", tc.description, tc.num1, tc.num2, r)
					}
				} else {
					if r != nil {

						t.Errorf("FAIL: %s. Did not expect a panic for Modulo(%d, %d), but got one: %v\nStack trace:\n%s", tc.description, tc.num1, tc.num2, r, string(debug.Stack()))

					}

				}
			}()

			t.Logf("Running test case: %s. Inputs: num1=%d, num2=%d", tc.name, tc.num1, tc.num2)
			actualResult := Modulo(tc.num1, tc.num2)

			if !tc.expectPanic {
				if actualResult != tc.expectedResult {

					t.Errorf("FAIL: %s. Modulo(%d, %d) = %d; want %d", tc.description, tc.num1, tc.num2, actualResult, tc.expectedResult)
				} else {

					t.Logf("PASS: %s. Modulo(%d, %d) = %d.", tc.description, tc.num1, tc.num2, actualResult)
				}
			} else if recover() == nil {

				t.Logf("INFO: %s. Modulo(%d, %d) completed execution, but panic was expected (final check in defer).", tc.description, tc.num1, tc.num2)
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

	type testCase struct {
		name        string
		num1        float64
		num2        float64
		expected    float64
		expectNaN   bool
		expectInf   bool
		infSign     int
		description string
	}

	testCases := []testCase{
		{
			name:        "Scenario 1: Multiply two positive integers",
			num1:        5.0,
			num2:        4.0,
			expected:    20.0,
			description: "Test basic multiplication with two positive whole numbers.",
		},
		{
			name:        "Scenario 2: Multiply two positive floating-point numbers",
			num1:        2.5,
			num2:        3.5,
			expected:    8.75,
			description: "Test multiplication with two positive numbers having fractional parts.",
		},
		{
			name:        "Scenario 3: Multiply a positive and a negative number",
			num1:        6.0,
			num2:        -3.0,
			expected:    -18.0,
			description: "Test multiplication logic when one number is positive and the other is negative.",
		},
		{
			name:        "Scenario 4: Multiply two negative numbers",
			num1:        -7.0,
			num2:        -2.0,
			expected:    14.0,
			description: "Test multiplication logic when both numbers are negative.",
		},
		{
			name:        "Scenario 5: Multiply by zero",
			num1:        123.45,
			num2:        0.0,
			expected:    0.0,
			description: "Test multiplication where one of the operands is zero.",
		},
		{
			name:        "Scenario 6: Multiply zero by zero",
			num1:        0.0,
			num2:        0.0,
			expected:    0.0,
			description: "Test the specific case where both operands are zero.",
		},
		{
			name:        "Scenario 7: Multiply by Positive Infinity",
			num1:        10.0,
			num2:        math.Inf(1),
			expectInf:   true,
			infSign:     1,
			description: "Test multiplication involving positive infinity (positive * +Inf = +Inf).",
		},
		{
			name:        "Scenario 8: Multiply by Negative Infinity",
			num1:        5.0,
			num2:        math.Inf(-1),
			expectInf:   true,
			infSign:     -1,
			description: "Test multiplication involving negative infinity (positive * -Inf = -Inf).",
		},
		{
			name:        "Scenario 9: Multiply Infinity by Zero",
			num1:        math.Inf(1),
			num2:        0.0,
			expectNaN:   true,
			description: "Test multiplication of infinity by zero (Inf * 0 = NaN).",
		},
		{
			name:        "Scenario 10: Multiply Infinity by Infinity (+Inf * -Inf)",
			num1:        math.Inf(1),
			num2:        math.Inf(-1),
			expectInf:   true,
			infSign:     -1,
			description: "Test multiplication of two infinite values (+Inf * -Inf = -Inf).",
		},
		{
			name:        "Scenario 10b: Multiply Infinity by Infinity (+Inf * +Inf)",
			num1:        math.Inf(1),
			num2:        math.Inf(1),
			expectInf:   true,
			infSign:     1,
			description: "Test multiplication of two infinite values (+Inf * +Inf = +Inf).",
		},
		{
			name:        "Scenario 10c: Multiply Infinity by Infinity (-Inf * -Inf)",
			num1:        math.Inf(-1),
			num2:        math.Inf(-1),
			expectInf:   true,
			infSign:     1,
			description: "Test multiplication of two infinite values (-Inf * -Inf = +Inf).",
		},
		{
			name:        "Scenario 11: Multiply involving NaN (Number * NaN)",
			num1:        15.5,
			num2:        math.NaN(),
			expectNaN:   true,
			description: "Test multiplication where one operand is NaN (Num * NaN = NaN).",
		},
		{
			name:        "Scenario 11b: Multiply involving NaN (NaN * Number)",
			num1:        math.NaN(),
			num2:        -20.0,
			expectNaN:   true,
			description: "Test multiplication where one operand is NaN (NaN * Num = NaN).",
		},
		{
			name:        "Scenario 11c: Multiply involving NaN (NaN * NaN)",
			num1:        math.NaN(),
			num2:        math.NaN(),
			expectNaN:   true,
			description: "Test multiplication where both operands are NaN (NaN * NaN = NaN).",
		},
		{
			name:        "Scenario 11d: Multiply involving NaN (NaN * Inf)",
			num1:        math.NaN(),
			num2:        math.Inf(1),
			expectNaN:   true,
			description: "Test multiplication where operands are NaN and Inf (NaN * Inf = NaN).",
		},
		{
			name:        "Scenario 12: Multiply large numbers resulting in overflow (Infinity)",
			num1:        math.MaxFloat64 / 10,
			num2:        100.0,
			expectInf:   true,
			infSign:     1,
			description: "Test multiplication of large numbers causing overflow to +Infinity.",
		},
		{
			name:        "Scenario 12b: Multiply large negative numbers resulting in overflow (Infinity)",
			num1:        -math.MaxFloat64 / 10,
			num2:        -100.0,
			expectInf:   true,
			infSign:     1,
			description: "Test multiplication of large negative numbers causing overflow to +Infinity.",
		},
		{
			name:        "Scenario 12c: Multiply large numbers resulting in negative overflow (-Infinity)",
			num1:        math.MaxFloat64 / 10,
			num2:        -100.0,
			expectInf:   true,
			infSign:     -1,
			description: "Test multiplication of large numbers causing overflow to -Infinity.",
		},
		{

			name:        "Scenario 13: Multiply small numbers resulting in underflow (Zero)",
			num1:        math.SmallestNonzeroFloat64 * 10,
			num2:        0.01,
			expected:    0.0,
			description: "Test multiplication of small numbers potentially causing underflow to zero.",
		},
		{
			name:        "Scenario 13b: Multiply small negative numbers resulting in underflow (Zero)",
			num1:        -math.SmallestNonzeroFloat64 * 10,
			num2:        -0.01,
			expected:    0.0,
			description: "Test multiplication of small negative numbers potentially causing underflow to zero.",
		},
		{
			name:        "Scenario 13c: Multiply small numbers resulting in negative underflow (Zero)",
			num1:        math.SmallestNonzeroFloat64 * 10,
			num2:        -0.01,
			expected:    -0.0,
			description: "Test multiplication of small numbers potentially causing underflow to negative zero.",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked, failing.")
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Description: %s", tc.description)
			t.Logf("Inputs: num1 = %v, num2 = %v", tc.num1, tc.num2)

			result := Multiply(tc.num1, tc.num2)

			if tc.expectNaN {
				if !math.IsNaN(result) {
					t.Errorf("Multiply(%v, %v) = %v; want NaN", tc.num1, tc.num2, result)
					t.Logf("Failure reason: Expected NaN, but got a number.")
				} else {
					t.Logf("Success: Multiply(%v, %v) correctly resulted in NaN.", tc.num1, tc.num2)
				}
			} else if tc.expectInf {
				if !math.IsInf(result, tc.infSign) {
					expectedSignStr := "positive"
					if tc.infSign < 0 {
						expectedSignStr = "negative"
					}
					t.Errorf("Multiply(%v, %v) = %v; want %s infinity", tc.num1, tc.num2, result, expectedSignStr)
					t.Logf("Failure reason: Expected %s infinity, but got %v.", expectedSignStr, result)
				} else {
					expectedSignStr := "positive"
					if tc.infSign < 0 {
						expectedSignStr = "negative"
					}
					t.Logf("Success: Multiply(%v, %v) correctly resulted in %s infinity.", tc.num1, tc.num2, expectedSignStr)
				}
			} else {

				if tc.expected == 0 && result == 0 {

					if math.Signbit(tc.expected) != math.Signbit(result) {
						t.Errorf("Multiply(%v, %v) = %v; want %v (sign difference in zero)", tc.num1, tc.num2, result, tc.expected)
						t.Logf("Failure reason: Expected %v, but got %v. Sign bit of zero differs.", tc.expected, result)
					} else if !floatsAlmostEqual(result, tc.expected) {

						t.Errorf("Multiply(%v, %v) = %v; want %v (within tolerance %e)", tc.num1, tc.num2, result, tc.expected, float64EqualityThreshold)
						t.Logf("Failure reason: Result %v is not approximately equal to expected %v.", result, tc.expected)
					} else {
						t.Logf("Success: Multiply(%v, %v) = %v, which is approximately equal to expected %v.", tc.num1, tc.num2, result, tc.expected)
					}
				} else if !floatsAlmostEqual(result, tc.expected) {
					t.Errorf("Multiply(%v, %v) = %v; want %v (within tolerance %e)", tc.num1, tc.num2, result, tc.expected, float64EqualityThreshold)
					t.Logf("Failure reason: Result %v is not approximately equal to expected %v.", result, tc.expected)
				} else {
					t.Logf("Success: Multiply(%v, %v) = %v, which is approximately equal to expected %v.", tc.num1, tc.num2, result, tc.expected)
				}
			}
		})
	}
}

func floatsAlmostEqual(a, b float64) bool {

	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return a == b
	}

	if math.IsNaN(a) || math.IsNaN(b) {
		return math.IsNaN(a) && math.IsNaN(b)
	}

	return math.Abs(a-b) <= float64EqualityThreshold
}

/*
ROOST_METHOD_HASH=Power_1c67a5d8b5
ROOST_METHOD_SIG_HASH=Power_c74b8edd76

FUNCTION_DEF=func Power(base, exponent float64) float64 // Power function
*/
func TestPower(t *testing.T) {

	const tolerance = 1e-9

	type testCase struct {
		name            string
		base            float64
		exponent        float64
		expectedResult  float64
		expectNaN       bool
		expectedInfSign int
		useTolerance    bool
	}

	testCases := []testCase{

		{
			name:           "Scenario 1: Positive Base, Positive Integer Exponent",
			base:           2.0,
			exponent:       3.0,
			expectedResult: 8.0,
			useTolerance:   false,
		},

		{
			name:           "Scenario 2: Positive Base, Zero Exponent",
			base:           10.5,
			exponent:       0.0,
			expectedResult: 1.0,
			useTolerance:   false,
		},

		{
			name:           "Scenario 3: Positive Base, Positive Fractional Exponent (Root)",
			base:           9.0,
			exponent:       0.5,
			expectedResult: 3.0,
			useTolerance:   true,
		},

		{
			name:           "Scenario 4: Positive Base, Negative Integer Exponent",
			base:           4.0,
			exponent:       -2.0,
			expectedResult: 0.0625,
			useTolerance:   true,
		},

		{
			name:           "Scenario 5: Zero Base, Positive Exponent",
			base:           0.0,
			exponent:       5.0,
			expectedResult: 0.0,
			useTolerance:   false,
		},

		{
			name:           "Scenario 6: Zero Base, Zero Exponent",
			base:           0.0,
			exponent:       0.0,
			expectedResult: 1.0,
			useTolerance:   false,
		},

		{
			name:            "Scenario 7: Zero Base, Negative Exponent",
			base:            0.0,
			exponent:        -3.0,
			expectedInfSign: 1,
			useTolerance:    false,
		},

		{
			name:           "Scenario 8: Negative Base, Positive Even Integer Exponent",
			base:           -2.0,
			exponent:       4.0,
			expectedResult: 16.0,
			useTolerance:   false,
		},

		{
			name:           "Scenario 9: Negative Base, Positive Odd Integer Exponent",
			base:           -2.0,
			exponent:       3.0,
			expectedResult: -8.0,
			useTolerance:   false,
		},

		{
			name:         "Scenario 10: Negative Base, Non-Integer Exponent",
			base:         -4.0,
			exponent:     0.5,
			expectNaN:    true,
			useTolerance: false,
		},

		{
			name:           "Scenario 11: Base = 1, Negative Exponent",
			base:           1.0,
			exponent:       -123.45,
			expectedResult: 1.0,
			useTolerance:   false,
		},
		{
			name:           "Scenario 11: Base = 1, Zero Exponent",
			base:           1.0,
			exponent:       0.0,
			expectedResult: 1.0,
			useTolerance:   false,
		},
		{
			name:           "Scenario 11: Base = 1, Positive Infinity Exponent",
			base:           1.0,
			exponent:       math.Inf(1),
			expectedResult: 1.0,
			useTolerance:   false,
		},
		{
			name:         "Scenario 11: Base = 1, NaN Exponent",
			base:         1.0,
			exponent:     math.NaN(),
			expectNaN:    true,
			useTolerance: false,
		},

		{
			name:           "Scenario 12: Base = -1, Even Integer Exponent",
			base:           -1.0,
			exponent:       6.0,
			expectedResult: 1.0,
			useTolerance:   false,
		},
		{
			name:           "Scenario 12: Base = -1, Odd Integer Exponent",
			base:           -1.0,
			exponent:       7.0,
			expectedResult: -1.0,
			useTolerance:   false,
		},

		{
			name:         "Scenario 13: Base = -1, Non-Integer Exponent",
			base:         -1.0,
			exponent:     2.5,
			expectNaN:    true,
			useTolerance: false,
		},

		{
			name:         "Scenario 14: NaN Base, Non-Zero Exponent",
			base:         math.NaN(),
			exponent:     2.0,
			expectNaN:    true,
			useTolerance: false,
		},
		{
			name:           "Scenario 14: NaN Base, Zero Exponent",
			base:           math.NaN(),
			exponent:       0.0,
			expectedResult: 1.0,
			useTolerance:   false,
		},

		{
			name:         "Scenario 15: Non-One Base, NaN Exponent",
			base:         5.0,
			exponent:     math.NaN(),
			expectNaN:    true,
			useTolerance: false,
		},

		{
			name:            "Scenario 16: Base |x| > 1, +Inf Exponent",
			base:            2.0,
			exponent:        math.Inf(1),
			expectedInfSign: 1,
			useTolerance:    false,
		},
		{
			name:           "Scenario 16: Base |x| < 1, +Inf Exponent",
			base:           0.5,
			exponent:       math.Inf(1),
			expectedResult: 0.0,
			useTolerance:   false,
		},
		{
			name:            "Scenario 16: Negative Base |x| > 1, +Inf Exponent",
			base:            -2.0,
			exponent:        math.Inf(1),
			expectedInfSign: 1,
			useTolerance:    false,
		},
		{
			name:           "Scenario 16: Negative Base |x| < 1, +Inf Exponent",
			base:           -0.5,
			exponent:       math.Inf(1),
			expectedResult: 0.0,
			useTolerance:   false,
		},

		{
			name:           "Scenario 17: Base |x| > 1, -Inf Exponent",
			base:           2.0,
			exponent:       math.Inf(-1),
			expectedResult: 0.0,
			useTolerance:   false,
		},
		{
			name:            "Scenario 17: Base |x| < 1, -Inf Exponent",
			base:            0.5,
			exponent:        math.Inf(-1),
			expectedInfSign: 1,
			useTolerance:    false,
		},
		{
			name:           "Scenario 17: Negative Base |x| > 1, -Inf Exponent",
			base:           -2.0,
			exponent:       math.Inf(-1),
			expectedResult: 0.0,
			useTolerance:   false,
		},
		{
			name:            "Scenario 17: Negative Base |x| < 1, -Inf Exponent",
			base:            -0.5,
			exponent:        math.Inf(-1),
			expectedInfSign: 1,
			useTolerance:    false,
		},

		{
			name:            "Scenario 18: +Inf Base, Positive Exponent",
			base:            math.Inf(1),
			exponent:        2.0,
			expectedInfSign: 1,
			useTolerance:    false,
		},
		{
			name:           "Scenario 18: +Inf Base, Negative Exponent",
			base:           math.Inf(1),
			exponent:       -2.0,
			expectedResult: 0.0,
			useTolerance:   false,
		},
		{
			name:            "Scenario 18: -Inf Base, Positive Odd Integer Exponent",
			base:            math.Inf(-1),
			exponent:        3.0,
			expectedInfSign: -1,
			useTolerance:    false,
		},
		{
			name:            "Scenario 18: -Inf Base, Positive Even Integer Exponent",
			base:            math.Inf(-1),
			exponent:        2.0,
			expectedInfSign: 1,
			useTolerance:    false,
		},
		{
			name:           "Scenario 18: -Inf Base, Negative Odd Integer Exponent",
			base:           math.Inf(-1),
			exponent:       -3.0,
			expectedResult: 0.0,
			useTolerance:   false,
		},
		{
			name:           "Scenario 18: -Inf Base, Negative Even Integer Exponent",
			base:           math.Inf(-1),
			exponent:       -2.0,
			expectedResult: 0.0,
			useTolerance:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Testing scenario: %s", tc.name)
			t.Logf("Arrange: base = %v, exponent = %v", tc.base, tc.exponent)

			actualResult := Power(tc.base, tc.exponent)
			t.Logf("Act: Called Power(%v, %v), Result: %v", tc.base, tc.exponent, actualResult)

			if tc.expectNaN {
				if !math.IsNaN(actualResult) {
					t.Errorf("Assert: FAILED. Expected NaN, but got %v", actualResult)
				} else {
					t.Logf("Assert: PASSED. Correctly resulted in NaN.")
				}
				return
			}
			if math.IsNaN(actualResult) && !tc.expectNaN {
				t.Errorf("Assert: FAILED. Did not expect NaN, but got NaN. Inputs: base=%v, exponent=%v", tc.base, tc.exponent)
				return
			}

			if tc.expectedInfSign != 0 {
				if !math.IsInf(actualResult, tc.expectedInfSign) {
					expectedInfStr := "+Inf"
					if tc.expectedInfSign == -1 {
						expectedInfStr = "-Inf"
					}
					t.Errorf("Assert: FAILED. Expected %s, but got %v", expectedInfStr, actualResult)
				} else {
					expectedInfStr := "+Inf"
					if tc.expectedInfSign == -1 {
						expectedInfStr = "-Inf"
					}
					t.Logf("Assert: PASSED. Correctly resulted in %s.", expectedInfStr)
				}
				return
			}
			if math.IsInf(actualResult, 0) && tc.expectedInfSign == 0 {
				infSign := 1
				if math.IsInf(actualResult, -1) {
					infSign = -1
				}
				t.Errorf("Assert: FAILED. Expected finite result %v, but got Inf (sign %d)", tc.expectedResult, infSign)
				return
			}

			if tc.useTolerance {
				diff := math.Abs(actualResult - tc.expectedResult)
				if diff >= tolerance {
					t.Errorf("Assert: FAILED. Result %.10f is not within tolerance %.10f of expected %.10f (Difference: %.10f)",
						actualResult, tolerance, tc.expectedResult, diff)
				} else {
					t.Logf("Assert: PASSED. Result %.10f is within tolerance %.10f of expected %.10f.",
						actualResult, tolerance, tc.expectedResult)
				}
			} else {

				if actualResult != tc.expectedResult {
					t.Errorf("Assert: FAILED. Expected exactly %v, but got %v", tc.expectedResult, actualResult)
				} else {
					t.Logf("Assert: PASSED. Result %v matches expected %v.", actualResult, tc.expectedResult)
				}
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

	testCases := []struct {
		name           string
		num1           int
		num2           int
		expectedResult int
	}{
		{
			name:           "Scenario 1: Subtracting two positive integers resulting in a positive number",
			num1:           10,
			num2:           3,
			expectedResult: 7,
		},
		{
			name:           "Scenario 2: Subtracting two positive integers resulting in a negative number",
			num1:           5,
			num2:           8,
			expectedResult: -3,
		},
		{
			name:           "Scenario 3: Subtracting a positive integer from a negative integer",
			num1:           -5,
			num2:           3,
			expectedResult: -8,
		},
		{
			name:           "Scenario 4: Subtracting a negative integer from a positive integer",
			num1:           10,
			num2:           -3,
			expectedResult: 13,
		},
		{
			name:           "Scenario 5: Subtracting two negative integers",
			num1:           -5,
			num2:           -8,
			expectedResult: 3,
		},
		{
			name:           "Scenario 6: Subtracting zero from a positive integer",
			num1:           7,
			num2:           0,
			expectedResult: 7,
		},
		{
			name:           "Scenario 7: Subtracting a positive integer from zero",
			num1:           0,
			num2:           9,
			expectedResult: -9,
		},
		{
			name:           "Scenario 8: Subtracting zero from zero",
			num1:           0,
			num2:           0,
			expectedResult: 0,
		},
		{
			name:           "Scenario 9: Subtracting a number from itself",
			num1:           123,
			num2:           123,
			expectedResult: 0,
		},
		{
			name:           "Scenario 10: Subtracting near the maximum integer limit (no overflow)",
			num1:           math.MaxInt,
			num2:           10,
			expectedResult: math.MaxInt - 10,
		},
		{
			name:           "Scenario 11: Subtracting near the minimum integer limit (no underflow)",
			num1:           math.MinInt,
			num2:           -10,
			expectedResult: math.MinInt + 10,
		},
		{
			name: "Scenario 12: Subtracting to potentially cause integer underflow (wrap-around)",
			num1: math.MinInt,
			num2: 1,

			expectedResult: math.MaxInt,
		},
		{
			name: "Scenario 13: Subtracting to potentially cause integer overflow (wrap-around)",
			num1: math.MaxInt,
			num2: -1,

			expectedResult: math.MinInt,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Testing with: num1 = %d, num2 = %d", tc.num1, tc.num2)

			actualResult := Subtract(tc.num1, tc.num2)

			if actualResult != tc.expectedResult {

				t.Errorf("FAIL: Expected result %d, but got %d", tc.expectedResult, actualResult)
			} else {

				t.Logf("PASS: Expected result %d, got %d", tc.expectedResult, actualResult)
			}
		})
	}
}
