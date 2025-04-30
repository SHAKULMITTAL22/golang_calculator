package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Absolute_d231f0ab10
ROOST_METHOD_SIG_HASH=Absolute_ec3c06e5a3

FUNCTION_DEF=func Absolute(num float64) float64 // Absolute value
*/
func TestAbsolute(t *testing.T) {

	testCases := []struct {
		name            string
		input           float64
		expected        float64
		checkNaN        bool
		checkInf        bool
		expectedInfSign int
		checkSignBit    bool
	}{

		{
			name:     "Positive Input",
			input:    5.7,
			expected: 5.7,
		},

		{
			name:     "Negative Input",
			input:    -10.3,
			expected: 10.3,
		},

		{
			name:         "Zero Input",
			input:        0.0,
			expected:     0.0,
			checkSignBit: true,
		},

		{
			name:         "Negative Zero Input",
			input:        math.Copysign(0.0, -1.0),
			expected:     0.0,
			checkSignBit: true,
		},

		{
			name:            "Positive Infinity Input",
			input:           math.Inf(1),
			expected:        math.Inf(1),
			checkInf:        true,
			expectedInfSign: 1,
		},

		{
			name:            "Negative Infinity Input",
			input:           math.Inf(-1),
			expected:        math.Inf(1),
			checkInf:        true,
			expectedInfSign: 1,
		},

		{
			name:     "NaN Input",
			input:    math.NaN(),
			checkNaN: true,
		},

		{
			name:     "Maximum Float64 Input",
			input:    math.MaxFloat64,
			expected: math.MaxFloat64,
		},

		{
			name:     "Minimum Float64 Input",
			input:    -math.MaxFloat64,
			expected: math.MaxFloat64,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked, failing.")
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: %v", tc.input)

			actual := Absolute(tc.input)

			t.Logf("Actual Output: %v", actual)

			if tc.checkNaN {
				if !math.IsNaN(actual) {
					t.Errorf("Absolute(%v) = %v; want NaN", tc.input, actual)
					t.Logf("Failure Reason: Expected NaN, but got a number.")
				} else {
					t.Logf("Success: Correctly returned NaN for input NaN.")
				}
			} else if tc.checkInf {
				if !math.IsInf(actual, tc.expectedInfSign) {
					expectedSignStr := "+"
					if tc.expectedInfSign == -1 {
						expectedSignStr = "-"
					}
					t.Errorf("Absolute(%v) = %v; want %sInf", tc.input, actual, expectedSignStr)
					t.Logf("Failure Reason: Expected %sInfinity, but got %v.", expectedSignStr, actual)
				} else {
					t.Logf("Success: Correctly returned %sInfinity.", map[int]string{1: "+", -1: "-"}[tc.expectedInfSign])
				}
			} else {

				if actual != tc.expected {
					t.Errorf("Absolute(%v) = %v; want %v", tc.input, actual, tc.expected)
					t.Logf("Failure Reason: Actual result differs from expected result.")
				} else if tc.checkSignBit && math.Signbit(actual) {

					t.Errorf("Absolute(%v) = %v (sign bit set); want +0.0 (sign bit clear)", tc.input, actual)
					t.Logf("Failure Reason: Expected positive zero (+0.0), but got negative zero (-0.0).")
				} else {
					t.Logf("Success: Actual result matches expected result.")
				}
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
		name     string
		num1     int
		num2     int
		expected int
	}

	testCases := []testCase{
		{
			name:     "Scenario 1: Add Two Positive Integers",
			num1:     5,
			num2:     10,
			expected: 15,
		},
		{
			name:     "Scenario 2: Add Two Negative Integers",
			num1:     -5,
			num2:     -10,
			expected: -15,
		},
		{
			name:     "Scenario 3: Add a Positive and a Negative Integer",
			num1:     15,
			num2:     -7,
			expected: 8,
		},
		{
			name:     "Scenario 4: Add a Negative and a Positive Integer",
			num1:     -15,
			num2:     7,
			expected: -8,
		},
		{
			name:     "Scenario 5: Add Zero to a Positive Integer",
			num1:     100,
			num2:     0,
			expected: 100,
		},
		{
			name:     "Scenario 6: Add Zero to a Negative Integer",
			num1:     -100,
			num2:     0,
			expected: -100,
		},
		{
			name:     "Scenario 7: Add Zero to Zero",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Scenario 8: Add Two Numbers Resulting in Zero",
			num1:     42,
			num2:     -42,
			expected: 0,
		},
		{
			name:     "Scenario 9: Add Near Maximum Integer Value (No Overflow)",
			num1:     math.MaxInt - 5,
			num2:     3,
			expected: math.MaxInt - 2,
		},
		{
			name:     "Scenario 10: Add Near Minimum Integer Value (No Underflow)",
			num1:     math.MinInt + 5,
			num2:     -3,
			expected: math.MinInt + 2,
		},
		{
			name:     "Scenario 11: Test Potential Integer Overflow",
			num1:     math.MaxInt,
			num2:     1,
			expected: math.MinInt,
		},
		{
			name:     "Scenario 12: Test Potential Integer Underflow",
			num1:     math.MinInt,
			num2:     -1,
			expected: math.MaxInt,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked, failing scenario: %s", tc.name)
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: num1 = %d, num2 = %d", tc.num1, tc.num2)
			t.Logf("Expected result: %d", tc.expected)

			actual := Add(tc.num1, tc.num2)
			t.Logf("Actual result: %d", actual)

			if actual != tc.expected {

				t.Errorf("Test Case '%s' Failed: Add(%d, %d) = %d; want %d",
					tc.name, tc.num1, tc.num2, actual, tc.expected)
			} else {

				t.Logf("Test Case '%s' Passed.", tc.name)
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

	const tolerance = 1e-9

	testCases := []struct {
		name             string
		num1             float64
		num2             float64
		expectedResult   float64
		expectPanic      bool
		expectedPanicMsg string
		checkInf         bool
		infSign          int
		checkNaN         bool
	}{

		{
			name:           "Scenario 1: Basic Positive Division (Non-Integer Result)",
			num1:           5.0,
			num2:           2.0,
			expectedResult: 2.5,
			expectPanic:    false,
		},

		{
			name:           "Scenario 2: Division Resulting in an Integer",
			num1:           10.0,
			num2:           2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},

		{
			name:             "Scenario 3: Division by Zero (Panic Scenario)",
			num1:             10.0,
			num2:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "division by zero is not allowed",
		},

		{
			name:           "Scenario 4: Division of Zero by a Non-Zero Number",
			num1:           0.0,
			num2:           5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 5: Division with a Negative Numerator",
			num1:           -10.0,
			num2:           2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 6: Division with a Negative Denominator",
			num1:           10.0,
			num2:           -2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 7: Division with Both Numerator and Denominator Negative",
			num1:           -10.0,
			num2:           -2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 8: Division Resulting in a Small Fractional Number",
			num1:           1.0,
			num2:           1000000.0,
			expectedResult: 0.000001,
			expectPanic:    false,
		},

		{
			name:           "Scenario 9: Division Resulting in a Large Number",
			num1:           1000000.0,
			num2:           0.000001,
			expectedResult: 1.0e12,
			expectPanic:    false,
		},

		{
			name:           "Scenario 10: Division by a Very Small Non-Zero Number",
			num1:           1.0,
			num2:           math.SmallestNonzeroFloat64,
			expectPanic:    false,
			checkInf:       true,
			infSign:        1,
			expectedResult: math.Inf(1),
		},

		{
			name:           "Scenario 11: Division Involving Positive Infinity as Numerator",
			num1:           math.Inf(1),
			num2:           2.0,
			expectPanic:    false,
			checkInf:       true,
			infSign:        1,
			expectedResult: math.Inf(1),
		},

		{
			name:           "Scenario 12: Division Involving Positive Infinity as Denominator",
			num1:           10.0,
			num2:           math.Inf(1),
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 13: Division Involving NaN as Numerator",
			num1:           math.NaN(),
			num2:           2.0,
			expectPanic:    false,
			checkNaN:       true,
			expectedResult: math.NaN(),
		},

		{
			name:           "Scenario 14: Division Involving NaN as Denominator",
			num1:           10.0,
			num2:           math.NaN(),
			expectPanic:    false,
			checkNaN:       true,
			expectedResult: math.NaN(),
		},

		{
			name:             "Scenario 15: Division of Infinity by Zero (Panic Scenario)",
			num1:             math.Inf(1),
			num2:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "division by zero is not allowed",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Errorf("Test panicked unexpectedly: %v\n%s", r, string(debug.Stack()))
				}
			}()

			t.Logf("Running test: %s", tc.name)
			t.Logf("Arrange: num1 = %v, num2 = %v", tc.num1, tc.num2)
			if tc.expectPanic {
				t.Logf("Expectation: Panic with message '%s'", tc.expectedPanicMsg)
			} else if tc.checkInf {
				t.Logf("Expectation: Result is %s Infinity", map[int]string{1: "Positive", -1: "Negative"}[tc.infSign])
			} else if tc.checkNaN {
				t.Logf("Expectation: Result is NaN")
			} else {
				t.Logf("Expectation: Result approximately %v", tc.expectedResult)
			}

			if tc.expectPanic {

				var recoveredPanic interface{}
				panicOccurred := false
				func() {
					defer func() {
						if r := recover(); r != nil {
							recoveredPanic = r
							panicOccurred = true
							t.Logf("Act: Call Divide(%v, %v) - Caught expected panic: %v", tc.num1, tc.num2, r)
						}
					}()

					_ = Divide(tc.num1, tc.num2)
				}()

				if !panicOccurred {
					t.Errorf("Assert: FAIL - Expected a panic for Divide(%v, %v), but it did not occur.", tc.num1, tc.num2)
					return
				}

				if msg, ok := recoveredPanic.(string); ok {
					if msg != tc.expectedPanicMsg {
						t.Errorf("Assert: FAIL - Panic message mismatch. Expected: '%s', Got: '%s'", tc.expectedPanicMsg, msg)
					} else {
						t.Logf("Assert: PASS - Panic occurred with the correct message: '%s'", msg)
					}
				} else {

					t.Logf("Assert: PASS - Panic occurred as expected (value: %v). Message type check skipped.", recoveredPanic)
				}

			} else {

				actualResult := Divide(tc.num1, tc.num2)
				t.Logf("Act: Call Divide(%v, %v) -> Result: %v", tc.num1, tc.num2, actualResult)

				if tc.checkNaN {

					if !math.IsNaN(actualResult) {
						t.Errorf("Assert: FAIL - Expected NaN for Divide(%v, %v), but got %v.", tc.num1, tc.num2, actualResult)
					} else {
						t.Logf("Assert: PASS - Result is NaN as expected.")
					}

				} else if tc.checkInf {

					if !math.IsInf(actualResult, tc.infSign) {
						expectedSignStr := map[int]string{1: "Positive", -1: "Negative"}[tc.infSign]
						actualSign := 0
						if math.IsInf(actualResult, 1) {
							actualSign = 1
						} else if math.IsInf(actualResult, -1) {
							actualSign = -1
						}
						actualSignStr := "Not Infinity"
						if actualSign != 0 {
							actualSignStr = map[int]string{1: "Positive", -1: "Negative"}[actualSign]
						}
						t.Errorf("Assert: FAIL - Expected %s Infinity for Divide(%v, %v), but got %v (%s Infinity).", expectedSignStr, tc.num1, tc.num2, actualResult, actualSignStr)
					} else {
						t.Logf("Assert: PASS - Result is %s Infinity as expected.", map[int]string{1: "Positive", -1: "Negative"}[tc.infSign])
					}

				} else {

					diff := math.Abs(tc.expectedResult - actualResult)
					if diff >= tolerance {
						t.Errorf("Assert: FAIL - Result mismatch for Divide(%v, %v). Expected: %v, Got: %v (Difference: %v, Tolerance: %v)", tc.num1, tc.num2, tc.expectedResult, actualResult, diff, tolerance)
					} else {
						t.Logf("Assert: PASS - Result %v is approximately equal to expected %v (within tolerance %v).", actualResult, tc.expectedResult, tolerance)
					}

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

	const epsilon = 1e-9

	testCases := []struct {
		name             string
		num              float64
		base             float64
		expectedResult   float64
		expectPanic      bool
		expectedPanicMsg string
	}{

		{
			name:           "Scenario 1: Standard Integer Inputs (log2(8))",
			num:            8.0,
			base:           2.0,
			expectedResult: 3.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 2: Base 10 and Integer Input (log10(100))",
			num:            100.0,
			base:           10.0,
			expectedResult: 2.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 3: Fractional Result (log3(10))",
			num:            10.0,
			base:           3.0,
			expectedResult: math.Log(10.0) / math.Log(3.0),
			expectPanic:    false,
		},

		{
			name:           "Scenario 4: num = 1 (log5(1))",
			num:            1.0,
			base:           5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 5: num = base (log7.5(7.5))",
			num:            7.5,
			base:           7.5,
			expectedResult: 1.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 11: Very Large Inputs (log2(MaxFloat64))",
			num:            math.MaxFloat64,
			base:           2.0,
			expectedResult: math.Log(math.MaxFloat64) / math.Log(2.0),
			expectPanic:    false,
		},

		{
			name:           "Scenario 12: Very Small Positive Inputs num (log10(SmallestNonzeroFloat64))",
			num:            math.SmallestNonzeroFloat64,
			base:           10.0,
			expectedResult: math.Log(math.SmallestNonzeroFloat64) / math.Log(10.0),
			expectPanic:    false,
		},

		{
			name:           "Scenario 13: Very Small Positive Inputs base (log0.1(0.5))",
			num:            0.5,
			base:           0.1,
			expectedResult: math.Log(0.5) / math.Log(0.1),
			expectPanic:    false,
		},

		{
			name:             "Scenario 6: Panic with num = 0",
			num:              0.0,
			base:             10.0,
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 7: Panic with Negative num",
			num:              -5.0,
			base:             2.0,
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 8: Panic with base = 1",
			num:              10.0,
			base:             1.0,
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 9: Panic with base = 0",
			num:              10.0,
			base:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 10: Panic with Negative base",
			num:              10.0,
			base:             -2.0,
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()

				if tc.expectPanic {

					if r == nil {
						t.Errorf("FAIL: Expected panic for num=%.4g, base=%.4g, but function did not panic.", tc.num, tc.base)
					} else {

						if msg, ok := r.(string); ok && msg == tc.expectedPanicMsg {
							t.Logf("PASS: Correctly panicked with expected message: '%s'", msg)
						} else if ok {
							t.Logf("INFO: Correctly panicked, but message differed. Expected: '%s', Got: '%s'", tc.expectedPanicMsg, msg)

						} else {
							t.Logf("PASS: Correctly panicked, but panic value was not a string: %v", r)
						}
					}
				} else {

					if r != nil {

						t.Errorf("FAIL: Unexpected panic for num=%.4g, base=%.4g: %v\nStack trace:\n%s", tc.num, tc.base, r, string(debug.Stack()))
					}

				}
			}()

			t.Logf("Testing: num=%.4g, base=%.4g", tc.num, tc.base)

			actualResult := Logarithm(tc.num, tc.base)

			if tc.expectPanic {
				t.Errorf("FAIL: Expected panic for num=%.4g, base=%.4g, but function returned %.15f instead.", tc.num, tc.base, actualResult)
				return
			}

			diff := math.Abs(actualResult - tc.expectedResult)
			if diff > epsilon {

				t.Errorf("FAIL: Logarithm(%.4g, %.4g) = %.15f; Want %.15f. Difference (%.15f) exceeds epsilon (%.15f)",
					tc.num, tc.base, actualResult, tc.expectedResult, diff, epsilon)
			} else {

				t.Logf("PASS: Logarithm(%.4g, %.4g) = %.15f. Result is within tolerance %.15f.",
					tc.num, tc.base, actualResult, epsilon)
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
		description    string
		num1           int
		num2           int
		expectedResult int
		expectPanic    bool
	}{

		{
			description:    "Scenario 1: Basic Positive Modulo (10 % 3)",
			num1:           10,
			num2:           3,
			expectedResult: 1,
			expectPanic:    false,
		},

		{
			description:    "Scenario 2: Dividend Smaller Than Divisor (3 % 10)",
			num1:           3,
			num2:           10,
			expectedResult: 3,
			expectPanic:    false,
		},

		{
			description:    "Scenario 3: Zero Dividend (0 % 5)",
			num1:           0,
			num2:           5,
			expectedResult: 0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 4: Zero Remainder (12 % 4)",
			num1:           12,
			num2:           4,
			expectedResult: 0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 5: Negative Dividend, Positive Divisor (-10 % 3)",
			num1:           -10,
			num2:           3,
			expectedResult: -1,
			expectPanic:    false,
		},

		{
			description:    "Scenario 6: Positive Dividend, Negative Divisor (10 % -3)",
			num1:           10,
			num2:           -3,
			expectedResult: 1,
			expectPanic:    false,
		},

		{
			description:    "Scenario 7: Both Negative (-10 % -3)",
			num1:           -10,
			num2:           -3,
			expectedResult: -1,
			expectPanic:    false,
		},

		{
			description:    "Scenario 8: Divisor is 1 (123 % 1)",
			num1:           123,
			num2:           1,
			expectedResult: 0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 9: Divisor is -1 (-45 % -1)",
			num1:           -45,
			num2:           -1,
			expectedResult: 0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 10: Large Integer Values (math.MaxInt % 1000)",
			num1:           math.MaxInt,
			num2:           1000,
			expectedResult: math.MaxInt % 1000,
			expectPanic:    false,
		},

		{
			description:    "Scenario 11: Division by Zero (10 % 0)",
			num1:           10,
			num2:           0,
			expectedResult: 0,
			expectPanic:    true,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil && !tc.expectPanic {

					t.Errorf("Test panicked unexpectedly: %v\nStack trace:\n%s", r, string(debug.Stack()))
				}
			}()

			t.Logf("Running test: %s", tc.description)
			t.Logf("Arrange: num1 = %d, num2 = %d", tc.num1, tc.num2)

			if tc.expectPanic {

				panicOccurred := false
				func() {

					defer func() {
						if r := recover(); r != nil {

							panicOccurred = true
							t.Logf("Act & Assert: Caught expected panic for Modulo(%d, %d). Panic details: %v", tc.num1, tc.num2, r)

						}
					}()

					_ = Modulo(tc.num1, tc.num2)
				}()

				if !panicOccurred {
					t.Errorf("Assert: Failed - Expected a panic for Modulo(%d, %d), but it did not occur.", tc.num1, tc.num2)
				}

			} else {

				actualResult := Modulo(tc.num1, tc.num2)
				t.Logf("Act: Called Modulo(%d, %d)", tc.num1, tc.num2)

				if actualResult != tc.expectedResult {

					t.Errorf("Assert: Failed - Modulo(%d, %d) = %d; want %d", tc.num1, tc.num2, actualResult, tc.expectedResult)
				} else {

					t.Logf("Assert: Success - Modulo(%d, %d) = %d", tc.num1, tc.num2, actualResult)
				}

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

	const epsilon = 1e-9

	testCases := []struct {
		name            string
		num1            float64
		num2            float64
		expected        float64
		useTolerance    bool
		expectedNaN     bool
		expectedInfSign int
	}{

		{
			name:            "Scenario 1: Multiply two positive integers",
			num1:            5.0,
			num2:            4.0,
			expected:        20.0,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 2: Multiply two positive floating-point numbers",
			num1:            2.5,
			num2:            3.5,
			expected:        8.75,
			useTolerance:    true,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 3: Multiply a positive number by a negative number",
			num1:            10.0,
			num2:            -3.5,
			expected:        -35.0,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 4: Multiply two negative numbers",
			num1:            -4.0,
			num2:            -5.5,
			expected:        22.0,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 5: Multiply by zero (first operand)",
			num1:            0.0,
			num2:            123.45,
			expected:        0.0,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 6: Multiply by zero (second operand)",
			num1:            -987.6,
			num2:            0.0,
			expected:        0.0,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 7: Multiply zero by zero",
			num1:            0.0,
			num2:            0.0,
			expected:        0.0,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 8: Multiply by one (identity element)",
			num1:            789.123,
			num2:            1.0,
			expected:        789.123,
			useTolerance:    true,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 9: Multiply by negative one",
			num1:            55.5,
			num2:            -1.0,
			expected:        -55.5,
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 10: Multiply potentially imprecise floats (0.1 * 0.2)",
			num1:            0.1,
			num2:            0.2,
			expected:        0.02,
			useTolerance:    true,
			expectedNaN:     false,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 11: Multiply +Inf by positive number",
			num1:            math.Inf(1),
			num2:            10.0,
			expected:        math.Inf(1),
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 12: Multiply negative number by +Inf",
			num1:            -5.0,
			num2:            math.Inf(1),
			expected:        math.Inf(-1),
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: -1,
		},

		{
			name:            "Scenario 13: Multiply negative number by -Inf",
			num1:            -2.0,
			num2:            math.Inf(-1),
			expected:        math.Inf(1),
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 14: Multiply +Inf by -Inf",
			num1:            math.Inf(1),
			num2:            math.Inf(-1),
			expected:        math.Inf(-1),
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: -1,
		},

		{
			name:            "Scenario 15: Multiply zero by +Inf",
			num1:            0.0,
			num2:            math.Inf(1),
			expected:        math.NaN(),
			useTolerance:    false,
			expectedNaN:     true,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 16: Multiply NaN by number",
			num1:            math.NaN(),
			num2:            100.0,
			expected:        math.NaN(),
			useTolerance:    false,
			expectedNaN:     true,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 17: Multiply number by NaN",
			num1:            -50.0,
			num2:            math.NaN(),
			expected:        math.NaN(),
			useTolerance:    false,
			expectedNaN:     true,
			expectedInfSign: 0,
		},

		{
			name:            "Scenario 18: Multiply large numbers causing positive overflow",
			num1:            math.MaxFloat64,
			num2:            2.0,
			expected:        math.Inf(1),
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 19: Multiply large negative numbers causing negative overflow",
			num1:            -math.MaxFloat64,
			num2:            2.0,
			expected:        math.Inf(-1),
			useTolerance:    false,
			expectedNaN:     false,
			expectedInfSign: -1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Testing scenario: %s", tc.name)
			t.Logf("Inputs: num1 = %v, num2 = %v", tc.num1, tc.num2)

			actual := Multiply(tc.num1, tc.num2)

			if tc.expectedNaN {
				if !math.IsNaN(actual) {
					t.Errorf("FAIL: Expected NaN, but got %v", actual)
				} else {
					t.Logf("PASS: Correctly resulted in NaN.")
				}
				return
			}

			if tc.expectedInfSign != 0 {
				if !math.IsInf(actual, tc.expectedInfSign) {
					expectedSignStr := "+"
					if tc.expectedInfSign < 0 {
						expectedSignStr = "-"
					}
					t.Errorf("FAIL: Expected %sInfinity, but got %v", expectedSignStr, actual)
				} else {
					expectedSignStr := "+"
					if tc.expectedInfSign < 0 {
						expectedSignStr = "-"
					}
					t.Logf("PASS: Correctly resulted in %sInfinity.", expectedSignStr)
				}
				return
			}

			if tc.useTolerance {

				if diff := math.Abs(actual - tc.expected); diff >= epsilon {
					t.Errorf("FAIL: Result %.15f is not within tolerance (%.15f) of expected %.15f. Difference: %.15f", actual, epsilon, tc.expected, diff)
				} else {
					t.Logf("PASS: Result %.15f is within tolerance (%.15f) of expected %.15f.", actual, epsilon, tc.expected)
				}
			} else {

				if actual != tc.expected {
					t.Errorf("FAIL: Expected exactly %v, but got %v", tc.expected, actual)
				} else {
					t.Logf("PASS: Result %v matches expected %v exactly.", actual, tc.expected)
				}
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

	const tolerance = 1e-9

	testCases := []struct {
		name         string
		base         float64
		exponent     float64
		expected     float64
		checkType    string
		description  string
		validation   string
		businessNeed string
	}{

		{
			name:         "Scenario 1: Positive Base, Positive Integer Exponent",
			base:         2.0,
			exponent:     3.0,
			expected:     8.0,
			checkType:    "approx",
			description:  "Test the function with a standard case of a positive base raised to a positive integer exponent.",
			validation:   "Standard mathematical exponentiation dictates 2^3 = 8. Floating-point comparisons require checking for approximate equality due to potential representation inaccuracies.",
			businessNeed: "This is a fundamental test ensuring the core functionality works as expected for common inputs.",
		},

		{
			name:         "Scenario 2: Positive Base, Zero Exponent",
			base:         5.5,
			exponent:     0.0,
			expected:     1.0,
			checkType:    "exact",
			description:  "Test the identity rule where any non-zero number raised to the power of zero equals 1.",
			validation:   "Mathematically, x^0 = 1 for any non-zero x. This specific case often returns an exact 1.0 in floating-point implementations.",
			businessNeed: "Validates correct handling of the zero exponent edge case, which is a common mathematical rule.",
		},

		{
			name:         "Scenario 3: Base One, Any Exponent",
			base:         1.0,
			exponent:     123.45,
			expected:     1.0,
			checkType:    "exact",
			description:  "Test the identity rule where 1 raised to any power equals 1.",
			validation:   "Mathematically, 1^y = 1 for any y. This is a specific behavior defined in \"math.Pow\".",
			businessNeed: "Ensures the function correctly handles the specific edge case of base 1.",
		},

		{
			name:         "Scenario 4: Zero Base, Positive Exponent",
			base:         0.0,
			exponent:     4.0,
			expected:     0.0,
			checkType:    "exact",
			description:  "Test the case where the base is zero and the exponent is positive. The result should be zero.",
			validation:   "Mathematically, 0^y = 0 for any y > 0.",
			businessNeed: "Covers a specific scenario involving zero as the base.",
		},

		{
			name:         "Scenario 5: Positive Base, Negative Integer Exponent",
			base:         4.0,
			exponent:     -2.0,
			expected:     0.0625,
			checkType:    "approx",
			description:  "Test the function with a negative integer exponent, which should result in the reciprocal of the base raised to the positive exponent.",
			validation:   "x^-y = 1 / (x^y). Floating-point comparison requires tolerance.",
			businessNeed: "Ensures correct calculation for negative exponents, representing division or reciprocals.",
		},

		{
			name:         "Scenario 6: Negative Base, Positive Odd Integer Exponent",
			base:         -2.0,
			exponent:     3.0,
			expected:     -8.0,
			checkType:    "approx",
			description:  "Test the function with a negative base raised to a positive odd integer exponent. The result should be negative.",
			validation:   "A negative base raised to an odd integer power results in a negative number. (-2)^3 = -8. Tolerance is used for float comparison.",
			businessNeed: "Validates correct sign handling for negative bases with odd exponents.",
		},

		{
			name:         "Scenario 7: Negative Base, Positive Even Integer Exponent",
			base:         -2.0,
			exponent:     4.0,
			expected:     16.0,
			checkType:    "approx",
			description:  "Test the function with a negative base raised to a positive even integer exponent. The result should be positive.",
			validation:   "A negative base raised to an even integer power results in a positive number. (-2)^4 = 16. Tolerance is used for float comparison.",
			businessNeed: "Validates correct sign handling for negative bases with even exponents.",
		},

		{
			name:         "Scenario 8: Positive Base, Fractional Exponent (Root)",
			base:         9.0,
			exponent:     0.5,
			expected:     3.0,
			checkType:    "approx",
			description:  "Test the function with a fractional exponent, representing a root calculation (e.g., square root).",
			validation:   "x^(1/n) is the nth root of x. 9^0.5 is the square root of 9, which is 3. Tolerance is used for float comparison.",
			businessNeed: "Ensures the function handles non-integer exponents correctly, enabling root calculations.",
		},

		{
			name:         "Scenario 9: Zero Base, Zero Exponent (Special Case)",
			base:         0.0,
			exponent:     0.0,
			expected:     1.0,
			checkType:    "exact",
			description:  "Test the mathematically ambiguous case of 0^0. Go's \"math.Pow(0, 0)\" defines this as 1.",
			validation:   "While mathematically debated, \"math.Pow(0, 0)\" returns 1 by definition in Go's standard library.",
			businessNeed: "Verifies adherence to the specific behavior defined by the underlying \"math.Pow\" function for this edge case.",
		},

		{
			name:         "Scenario 10: Zero Base, Negative Exponent (Infinity)",
			base:         0.0,
			exponent:     -2.0,
			expected:     math.Inf(1),
			checkType:    "posinf",
			description:  "Test raising zero to a negative power, which should result in positive infinity due to division by zero.",
			validation:   "0^-y = 1 / (0^y) = 1/0, which results in positive infinity in floating-point arithmetic. \"math.IsInf\" with sign argument 1 checks for +Inf.",
			businessNeed: "Ensures correct handling of division-by-zero scenarios leading to infinite results, as defined by IEEE 754.",
		},

		{
			name:         "Scenario 11: Negative Base, Non-Integer Exponent (NaN)",
			base:         -4.0,
			exponent:     0.5,
			expected:     math.NaN(),
			checkType:    "nan",
			description:  "Test raising a negative base to a non-integer exponent (e.g., square root of -1). This is undefined in real numbers and should result in NaN (Not a Number).",
			validation:   "The result of raising a negative number to a non-integer power is not a real number. IEEE 754 defines NaN for such results. \"math.IsNaN\" is the standard way to check for NaN.",
			businessNeed: "Verifies correct handling of mathematically undefined operations within the domain of real numbers, returning NaN as expected.",
		},

		{
			name:         "Scenario 12: Base is NaN",
			base:         math.NaN(),
			exponent:     2.0,
			expected:     math.NaN(),
			checkType:    "nan",
			description:  "Test the function's behavior when the base input is NaN. The result should propagate NaN (except for Pow(NaN, 0)).",
			validation:   "Operations involving NaN typically result in NaN. \"math.IsNaN\" checks for this special value.",
			businessNeed: "Ensures correct propagation of invalid numeric states (NaN) through calculations.",
		},

		{
			name:         "Scenario 13: Exponent is NaN",
			base:         2.0,
			exponent:     math.NaN(),
			expected:     math.NaN(),
			checkType:    "nan",
			description:  "Test the function's behavior when the exponent input is NaN. The result should propagate NaN (except for Pow(1, NaN)).",
			validation:   "Operations involving NaN typically result in NaN. \"math.IsNaN\" checks for this special value.",
			businessNeed: "Ensures correct propagation of invalid numeric states (NaN) through calculations.",
		},

		{
			name:         "Scenario 14: Base is Positive Infinity, Positive Exponent",
			base:         math.Inf(1),
			exponent:     2.0,
			expected:     math.Inf(1),
			checkType:    "posinf",
			description:  "Test raising positive infinity to a positive power. The result should be positive infinity.",
			validation:   "Infinity raised to a positive power remains infinity. \"math.IsInf\" checks for infinite values.",
			businessNeed: "Validates handling of infinite inputs according to floating-point standards.",
		},

		{
			name:         "Scenario 15: Base is Positive Infinity, Negative Exponent",
			base:         math.Inf(1),
			exponent:     -2.0,
			expected:     0.0,
			checkType:    "poszero",
			description:  "Test raising positive infinity to a negative power. The result should be positive zero.",
			validation:   "(+Inf)^-y = 1 / ((+Inf)^y) = 1 / +Inf = +0. Checking the sign bit distinguishes +0 from -0.",
			businessNeed: "Validates handling of infinite inputs leading to zero, respecting the sign of zero per IEEE 754.",
		},

		{
			name:         "Scenario 16: Base is Negative Infinity, Positive Odd Integer Exponent",
			base:         math.Inf(-1),
			exponent:     3.0,
			expected:     math.Inf(-1),
			checkType:    "neginf",
			description:  "Test raising negative infinity to a positive odd integer power. The result should be negative infinity.",
			validation:   "(-Inf)^odd = -Inf. \"math.IsInf\" with sign argument -1 checks for -Inf.",
			businessNeed: "Validates handling of negative infinite inputs according to floating-point standards and sign rules.",
		},

		{
			name:         "Scenario 17: Base is Negative Infinity, Positive Even Integer Exponent",
			base:         math.Inf(-1),
			exponent:     2.0,
			expected:     math.Inf(1),
			checkType:    "posinf",
			description:  "Test raising negative infinity to a positive even integer power. The result should be positive infinity.",
			validation:   "(-Inf)^even = +Inf. \"math.IsInf\" with sign argument 1 checks for +Inf.",
			businessNeed: "Validates handling of negative infinite inputs according to floating-point standards and sign rules.",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked unexpectedly")
				}
			}()

			t.Logf("Running Test Case: %s", tc.name)
			t.Logf("Description: %s", tc.description)
			t.Logf("Arrange: base = %v, exponent = %v", tc.base, tc.exponent)
			t.Logf("Expected Check Type: %s, Expected Value (if applicable): %v", tc.checkType, tc.expected)
			t.Logf("Validation Logic: %s", tc.validation)
			t.Logf("Business Need: %s", tc.businessNeed)

			result := Power(tc.base, tc.exponent)
			t.Logf("Act: Called Power(%v, %v). Result: %v", tc.base, tc.exponent, result)

			var success bool
			var failureMsg string

			switch tc.checkType {
			case "approx":
				diff := math.Abs(result - tc.expected)
				success = diff < tolerance
				if !success {
					failureMsg = "Result differs from expected value by more than the tolerance."
				}
			case "exact":
				success = result == tc.expected
				if !success {
					failureMsg = "Result does not exactly match the expected value."
				}
			case "nan":
				success = math.IsNaN(result)
				if !success {
					failureMsg = "Result was expected to be NaN, but it was not."
				}
			case "posinf":
				success = math.IsInf(result, 1)
				if !success {
					failureMsg = "Result was expected to be Positive Infinity, but it was not."
				}
			case "neginf":
				success = math.IsInf(result, -1)
				if !success {
					failureMsg = "Result was expected to be Negative Infinity, but it was not."
				}
			case "poszero":

				success = result == 0.0 && !math.Signbit(result)
				if !success {
					failureMsg = "Result was expected to be Positive Zero (+0.0), but it was not."
				}
			default:
				t.Fatalf("Unknown checkType '%s' in test case '%s'", tc.checkType, tc.name)
			}

			if success {
				t.Logf("Assert: PASSED. Result %v matched expectation for check type '%s'.", result, tc.checkType)
			} else {
				t.Errorf("Assert: FAILED. %s Expected (approx/exact): %v, Got: %v", failureMsg, tc.expected, result)
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

	testCases := []struct {
		name            string
		angle           float64
		expectedSin     float64
		expectedCos     float64
		expectedTan     float64
		tolerance       float64
		checkTanInf     bool
		tanThreshold    float64
		expectNaN       bool
		expectedTanSign int
	}{

		{
			name:        "Scenario 1: Zero Angle",
			angle:       0.0,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
			tolerance:   1e-9,
		},

		{
			name:        "Scenario 2: Pi/4 Angle",
			angle:       math.Pi / 4.0,
			expectedSin: math.Sqrt(2) / 2.0,
			expectedCos: math.Sqrt(2) / 2.0,
			expectedTan: 1.0,
			tolerance:   1e-9,
		},

		{
			name:            "Scenario 3: Pi/2 Angle",
			angle:           math.Pi / 2.0,
			expectedSin:     1.0,
			expectedCos:     0.0,
			tolerance:       1e-9,
			checkTanInf:     true,
			tanThreshold:    1e15,
			expectedTanSign: 1,
		},

		{
			name:        "Scenario 4: Pi Angle",
			angle:       math.Pi,
			expectedSin: 0.0,
			expectedCos: -1.0,
			expectedTan: 0.0,
			tolerance:   1e-9,
		},

		{
			name:            "Scenario 5: 3*Pi/2 Angle",
			angle:           3.0 * math.Pi / 2.0,
			expectedSin:     -1.0,
			expectedCos:     0.0,
			tolerance:       1e-9,
			checkTanInf:     true,
			tanThreshold:    -1e15,
			expectedTanSign: -1,
		},

		{
			name:        "Scenario 6: 2*Pi Angle",
			angle:       2.0 * math.Pi,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
			tolerance:   1e-9,
		},

		{
			name:        "Scenario 7: Negative Angle -Pi/4",
			angle:       -math.Pi / 4.0,
			expectedSin: -math.Sqrt(2) / 2.0,
			expectedCos: math.Sqrt(2) / 2.0,
			expectedTan: -1.0,
			tolerance:   1e-9,
		},

		{
			name:        "Scenario 8: Large Angle 10*Pi",
			angle:       10.0 * math.Pi,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
			tolerance:   1e-9,
		},

		{
			name:      "Scenario 9: NaN Input",
			angle:     math.NaN(),
			expectNaN: true,
		},

		{
			name:      "Scenario 10: Positive Infinity Input",
			angle:     math.Inf(1),
			expectNaN: true,
		},

		{
			name:      "Scenario 11: Negative Infinity Input",
			angle:     math.Inf(-1),
			expectNaN: true,
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
			t.Logf("Input angle: %v", tc.angle)

			sin, cos, tan := SinCosTan(tc.angle)

			if tc.expectNaN {

				if !math.IsNaN(sin) {
					t.Errorf("Expected sin to be NaN, but got %v", sin)
				} else {
					t.Logf("Success: sin is NaN as expected.")
				}
				if !math.IsNaN(cos) {
					t.Errorf("Expected cos to be NaN, but got %v", cos)
				} else {
					t.Logf("Success: cos is NaN as expected.")
				}
				if !math.IsNaN(tan) {
					t.Errorf("Expected tan to be NaN, but got %v", tan)
				} else {
					t.Logf("Success: tan is NaN as expected.")
				}
			} else {

				if diff := math.Abs(sin - tc.expectedSin); diff >= tc.tolerance {
					t.Errorf("FAIL: Sine mismatch. Expected: %v, Got: %v, Diff: %v", tc.expectedSin, sin, diff)
				} else {
					t.Logf("Success: Sine within tolerance. Expected: %v, Got: %v", tc.expectedSin, sin)
				}

				if diff := math.Abs(cos - tc.expectedCos); diff >= tc.tolerance {
					t.Errorf("FAIL: Cosine mismatch. Expected: %v, Got: %v, Diff: %v", tc.expectedCos, cos, diff)
				} else {
					t.Logf("Success: Cosine within tolerance. Expected: %v, Got: %v", tc.expectedCos, cos)
				}

				if tc.checkTanInf {

					if tc.expectedTanSign > 0 {
						if !(tan > tc.tanThreshold) {
							t.Errorf("FAIL: Tangent expected to be large positive ( > %v), but got %v", tc.tanThreshold, tan)
						} else {
							t.Logf("Success: Tangent is large positive as expected. Got: %v", tan)
						}
					} else {
						if !(tan < tc.tanThreshold) {
							t.Errorf("FAIL: Tangent expected to be large negative ( < %v), but got %v", tc.tanThreshold, tan)
						} else {
							t.Logf("Success: Tangent is large negative as expected. Got: %v", tan)
						}
					}

				} else {

					if tc.expectedTan == 0.0 {
						if diff := math.Abs(tan); diff >= tc.tolerance {
							t.Errorf("FAIL: Tangent mismatch. Expected: %v, Got: %v (Abs Diff: %v)", tc.expectedTan, tan, diff)
						} else {
							t.Logf("Success: Tangent within tolerance near zero. Expected: %v, Got: %v", tc.expectedTan, tan)
						}
					} else {
						if diff := math.Abs(tan - tc.expectedTan); diff >= tc.tolerance {
							t.Errorf("FAIL: Tangent mismatch. Expected: %v, Got: %v, Diff: %v", tc.expectedTan, tan, diff)
						} else {
							t.Logf("Success: Tangent within tolerance. Expected: %v, Got: %v", tc.expectedTan, tan)
						}
					}
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

	const epsilon = 1e-9

	testCases := []struct {
		description      string
		num              float64
		expectedResult   float64
		expectPanic      bool
		expectedPanicMsg string
	}{

		{
			description:    "Scenario 1: Positive Perfect Square (4.0)",
			num:            4.0,
			expectedResult: 2.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 2: Positive Non-Perfect Square (2.0)",
			num:            2.0,
			expectedResult: math.Sqrt(2.0),
			expectPanic:    false,
		},

		{
			description:    "Scenario 3: Zero (0.0)",
			num:            0.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 4: One (1.0)",
			num:            1.0,
			expectedResult: 1.0,
			expectPanic:    false,
		},

		{
			description:      "Scenario 5: Negative Number (-4.0)",
			num:              -4.0,
			expectPanic:      true,
			expectedPanicMsg: "square root of a negative number is not defined",
		},

		{
			description:    "Scenario 6: Large Positive Number (math.MaxFloat64)",
			num:            math.MaxFloat64,
			expectedResult: math.Sqrt(math.MaxFloat64),
			expectPanic:    false,
		},

		{
			description:    "Scenario 7: Small Positive Number (1e-10)",
			num:            1e-10,
			expectedResult: math.Sqrt(1e-10),
			expectPanic:    false,
		},

		{
			description:    "Scenario 8: Negative Zero (-0.0)",
			num:            -0.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			defer func() {
				r := recover()
				if tc.expectPanic {
					if r == nil {
						t.Errorf("Expected a panic for input %f, but did not get one.", tc.num)
					} else {

						panicMsg, ok := r.(string)
						if !ok {
							t.Errorf("Expected panic message to be a string, but got type %T: %v", r, r)
						} else if panicMsg != tc.expectedPanicMsg {
							t.Errorf("Expected panic message '%s', but got '%s'", tc.expectedPanicMsg, panicMsg)
						} else {
							t.Logf("Successfully caught expected panic: %v", r)
						}
					}
				} else if r != nil {

					t.Errorf("Unexpected panic for input %f: %v\n%s", tc.num, r, string(debug.Stack()))
				}
			}()

			result := SquareRoot(tc.num)

			if !tc.expectPanic {

				if diff := math.Abs(result - tc.expectedResult); diff > epsilon {
					t.Errorf("For input %f, expected result close to %f, but got %f (difference: %e)",
						tc.num, tc.expectedResult, result, diff)
				} else {
					t.Logf("For input %f, got result %f, which is within tolerance %e of expected %f. Test Passed.",
						tc.num, result, epsilon, tc.expectedResult)
				}

				if tc.num == 0.0 || tc.num == -0.0 {
					if math.Signbit(result) {
						t.Errorf("For input %f, expected result 0.0, but got %f (negative zero)", tc.num, result)
					}
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
		name     string
		num1     int
		num2     int
		expected int
		desc     string
	}{

		{
			name:     "Scenario 1: Positive Subtraction (Larger - Smaller)",
			num1:     10,
			num2:     4,
			expected: 6,
			desc:     "Assert: Result == 6. Logic: 10 - 4 = 6. Importance: Validates basic positive integer subtraction.",
		},

		{
			name:     "Scenario 2: Positive Subtraction (Smaller - Larger)",
			num1:     5,
			num2:     8,
			expected: -3,
			desc:     "Assert: Result == -3. Logic: 5 - 8 = -3. Importance: Ensures correct handling of negative results.",
		},

		{
			name:     "Scenario 3: Subtract Zero from Positive",
			num1:     99,
			num2:     0,
			expected: 99,
			desc:     "Assert: Result == 99. Logic: 99 - 0 = 99. Importance: Confirms identity property with zero.",
		},

		{
			name:     "Scenario 4: Subtract Positive from Zero",
			num1:     0,
			num2:     7,
			expected: -7,
			desc:     "Assert: Result == -7. Logic: 0 - 7 = -7. Importance: Validates subtraction from a zero baseline.",
		},

		{
			name:     "Scenario 5: Subtract Number from Itself",
			num1:     123,
			num2:     123,
			expected: 0,
			desc:     "Assert: Result == 0. Logic: 123 - 123 = 0. Importance: Ensures correct handling of equality cases.",
		},

		{
			name:     "Scenario 6: Subtract Two Negative Integers",
			num1:     -5,
			num2:     -3,
			expected: -2,
			desc:     "Assert: Result == -2. Logic: (-5) - (-3) = -5 + 3 = -2. Importance: Validates subtraction with negative inputs.",
		},

		{
			name:     "Scenario 7: Subtract Positive from Negative",
			num1:     -10,
			num2:     5,
			expected: -15,
			desc:     "Assert: Result == -15. Logic: -10 - 5 = -15. Importance: Ensures correct calculation when decreasing a negative value.",
		},

		{
			name:     "Scenario 8: Subtract Negative from Positive",
			num1:     7,
			num2:     -4,
			expected: 11,
			desc:     "Assert: Result == 11. Logic: 7 - (-4) = 7 + 4 = 11. Importance: Verifies correct handling of double negatives (subtraction becoming addition).",
		},

		{
			name:     "Scenario 9: Subtract Zero from Negative",
			num1:     -50,
			num2:     0,
			expected: -50,
			desc:     "Assert: Result == -50. Logic: -50 - 0 = -50. Importance: Reinforces identity property with zero for negative minuends.",
		},

		{
			name:     "Scenario 10: Subtract Negative from Zero",
			num1:     0,
			num2:     -9,
			expected: 9,
			desc:     "Assert: Result == 9. Logic: 0 - (-9) = 0 + 9 = 9. Importance: Ensures correct calculation relative to zero with negative subtrahends.",
		},

		{
			name:     "Scenario 11: Subtract Negative from MaxInt (Overflow)",
			num1:     math.MaxInt,
			num2:     -1,
			expected: math.MinInt,
			desc:     "Assert: Result == math.MinInt. Logic: math.MaxInt - (-1) = math.MaxInt + 1, which overflows and wraps to math.MinInt. Importance: Tests behavior at the upper integer limit.",
		},

		{
			name:     "Scenario 12: Subtract Positive from MinInt (Underflow)",
			num1:     math.MinInt,
			num2:     1,
			expected: math.MaxInt,
			desc:     "Assert: Result == math.MaxInt. Logic: math.MinInt - 1, which underflows and wraps to math.MaxInt. Importance: Tests behavior at the lower integer limit.",
		},

		{
			name:     "Scenario 13: Subtract MaxInt from Small Positive",
			num1:     1,
			num2:     math.MaxInt,
			expected: 1 - math.MaxInt,
			desc:     "Assert: Result == 1 - math.MaxInt. Logic: Standard subtraction resulting in a large negative number. Importance: Verifies calculations with extremely large subtrahends.",
		},

		{
			name:     "Scenario 14: Subtract MinInt from Zero (Overflow)",
			num1:     0,
			num2:     math.MinInt,
			expected: math.MinInt,
			desc:     "Assert: Result == math.MinInt. Logic: 0 - math.MinInt overflows and wraps to math.MinInt. Importance: Tests overflow when subtracting the most negative number.",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked unexpectedly for inputs num1=%d, num2=%d", tc.num1, tc.num2)
				}
			}()

			t.Logf("Testing: %s - Inputs: num1=%d, num2=%d", tc.name, tc.num1, tc.num2)
			t.Logf("Validation Logic & Importance: %s", tc.desc)

			actual := Subtract(tc.num1, tc.num2)

			if actual != tc.expected {

				t.Errorf("FAIL: Expected result %d, but got %d", tc.expected, actual)
			} else {

				t.Logf("PASS: Expected result %d, got %d", tc.expected, actual)
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

	testCases := []struct {
		name             string
		n                int
		expectedResult   int
		expectedPanicMsg string
	}{

		{
			name:           "Scenario 1: Factorial of Zero",
			n:              0,
			expectedResult: 1,
		},

		{
			name:           "Scenario 2: Factorial of One",
			n:              1,
			expectedResult: 1,
		},

		{
			name:           "Scenario 3: Factorial of 5",
			n:              5,
			expectedResult: 120,
		},

		{
			name:           "Scenario 4: Factorial of 10",
			n:              10,
			expectedResult: 3628800,
		},

		{
			name:             "Scenario 5: Factorial of -1 (Panic)",
			n:                -1,
			expectedPanicMsg: "factorial is not defined for negative numbers",
		},

		{
			name:             "Scenario 6: Factorial of -10 (Panic)",
			n:                -10,
			expectedPanicMsg: "factorial is not defined for negative numbers",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: n=%d", tc.n)

			defer func() {
				r := recover()
				if tc.expectedPanicMsg != "" {

					if r == nil {

						t.Errorf("FAIL: Expected a panic with message '%s', but function did not panic.", tc.expectedPanicMsg)
					} else {

						panicMsg, ok := r.(string)
						if !ok {
							t.Errorf("FAIL: Expected panic message of type string, but got type %T with value '%v'.", r, r)
							t.Logf("Panic stack trace:\n%s", string(debug.Stack()))
						} else if panicMsg != tc.expectedPanicMsg {
							t.Errorf("FAIL: Expected panic message '%s', but got '%s'.", tc.expectedPanicMsg, panicMsg)
							t.Logf("Panic stack trace:\n%s", string(debug.Stack()))
						} else {

							t.Logf("PASS: Successfully caught expected panic: %v", r)
						}
					}
				} else {

					if r != nil {

						t.Errorf("FAIL: Unexpected panic occurred: %v\n%s", r, string(debug.Stack()))

					}

				}
			}()

			actualResult := Factorial(tc.n)

			if tc.expectedPanicMsg == "" {
				if actualResult != tc.expectedResult {

					t.Errorf("FAIL: Factorial(%d): Expected result %d, but got %d.", tc.n, tc.expectedResult, actualResult)
				} else {

					t.Logf("PASS: Factorial(%d): Correctly returned %d.", tc.n, actualResult)
				}
			}

		})
	}
}

/*
ROOST_METHOD_HASH=GCD_1da681d86b
ROOST_METHOD_SIG_HASH=GCD_39a1228f3a

FUNCTION_DEF=func GCD(a, b int) int // Greatest Common Divisor (GCD) using Euclidean algorithm
*/
func TestGcd(t *testing.T) {

	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Scenario 1: Basic Positive Integers",
			a:        54,
			b:        24,
			expected: 6,
		},
		{
			name:     "Scenario 2: Order of Arguments",
			a:        24,
			b:        54,
			expected: 6,
		},
		{
			name:     "Scenario 3: Coprime Positive Integers",
			a:        17,
			b:        23,
			expected: 1,
		},
		{
			name:     "Scenario 4: One Argument is a Multiple of the Other (a > b)",
			a:        30,
			b:        10,
			expected: 10,
		},
		{
			name:     "Scenario 5: One Argument is a Multiple of the Other (b > a)",
			a:        7,
			b:        21,
			expected: 7,
		},
		{
			name:     "Scenario 6: Second Argument is Zero",
			a:        15,
			b:        0,
			expected: 15,
		},
		{
			name:     "Scenario 7: First Argument is Zero",
			a:        0,
			b:        9,
			expected: 9,
		},
		{
			name:     "Scenario 8: Both Arguments are Zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "Scenario 9: Identical Positive Arguments",
			a:        12,
			b:        12,
			expected: 12,
		},
		{
			name:     "Scenario 10: One Positive, One Negative Argument (a > 0, b < 0)",
			a:        54,
			b:        -24,
			expected: 6,
		},
		{
			name:     "Scenario 11: One Negative, One Positive Argument (a < 0, b > 0)",
			a:        -54,
			b:        24,
			expected: -6,
		},
		{
			name:     "Scenario 12: Both Arguments Negative",
			a:        -54,
			b:        -24,
			expected: -6,
		},
		{
			name:     "Scenario 13: Negative Argument with Zero (a < 0, b = 0)",
			a:        -15,
			b:        0,
			expected: -15,
		},
		{
			name:     "Scenario 14: Zero with Negative Argument (a = 0, b < 0)",
			a:        0,
			b:        -9,
			expected: -9,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test failed due to panic: %v", r)
				}
			}()

			t.Logf("Testing scenario: %s (a=%d, b=%d)", tc.name, tc.a, tc.b)

			result := GCD(tc.a, tc.b)

			if result != tc.expected {

				t.Errorf("GCD(%d, %d) failed: expected %d, got %d", tc.a, tc.b, tc.expected, result)
			} else {

				t.Logf("GCD(%d, %d) passed: expected %d, got %d", tc.a, tc.b, tc.expected, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=LCM_85c2702b86
ROOST_METHOD_SIG_HASH=LCM_fb713f0b10

FUNCTION_DEF=func LCM(a, b int) int // Least Common Multiple (LCM) using GCD
*/
func TestLcm(t *testing.T) {

	testCases := []struct {
		name        string
		a           int
		b           int
		expectedLcm int
		expectPanic bool
	}{

		{
			name:        "Scenario 1: Basic Positive Integers (6, 8)",
			a:           6,
			b:           8,
			expectedLcm: 24,
			expectPanic: false,
		},

		{
			name:        "Scenario 2: Coprime Positive Integers (7, 5)",
			a:           7,
			b:           5,
			expectedLcm: 35,
			expectPanic: false,
		},

		{
			name:        "Scenario 3: One Integer is a Multiple (4, 12)",
			a:           4,
			b:           12,
			expectedLcm: 12,
			expectPanic: false,
		},

		{
			name:        "Scenario 4: Equal Positive Integers (10, 10)",
			a:           10,
			b:           10,
			expectedLcm: 10,
			expectPanic: false,
		},

		{
			name:        "Scenario 5: One Input is Zero (9, 0)",
			a:           9,
			b:           0,
			expectedLcm: 0,
			expectPanic: false,
		},
		{
			name:        "Scenario 5: One Input is Zero (0, 5)",
			a:           0,
			b:           5,
			expectedLcm: 0,
			expectPanic: false,
		},

		{
			name:        "Scenario 6: Both Inputs are Zero (0, 0)",
			a:           0,
			b:           0,
			expectedLcm: 0,
			expectPanic: true,
		},

		{
			name:        "Scenario 7: One Negative Input (6, -8)",
			a:           6,
			b:           -8,
			expectedLcm: 24,
			expectPanic: false,
		},
		{
			name: "Scenario 7: One Negative Input (-7, 5)",
			a:    -7,
			b:    5,

			expectedLcm: -35,
			expectPanic: false,
		},

		{
			name:        "Scenario 8: Both Negative Inputs (-6, -8)",
			a:           -6,
			b:           -8,
			expectedLcm: -24,
			expectPanic: false,
		},
		{
			name: "Scenario 8: Both Negative Inputs (-7, -5)",
			a:    -7,
			b:    -5,

			expectedLcm: -35,
			expectPanic: false,
		},

		{
			name:        "Scenario 9: Commutativity Check (9, 6)",
			a:           9,
			b:           6,
			expectedLcm: 18,
			expectPanic: false,
		},
		{
			name: "Scenario 9: Commutativity Check (-9, 6)",
			a:    -9,
			b:    6,

			expectedLcm: 18,
			expectPanic: false,
		},

		{
			name:        "Scenario 10: Large Integers (Overflow Check)",
			a:           math.MaxInt,
			b:           2,
			expectedLcm: -2,

			expectPanic: false,
		},

		{
			name: "Scenario 9b: Commutativity Check (-8, -6)",
			a:    -8,
			b:    -6,

			expectedLcm: -24,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					if !tc.expectPanic {

						t.Errorf("Test %s panicked unexpectedly: %v", tc.name, r)
					} else {

						t.Logf("Test %s correctly panicked as expected.", tc.name)
					}
				} else if tc.expectPanic {

					t.Errorf("Test %s was expected to panic but did not.", tc.name)
				}
			}()

			t.Logf("Testing with a = %d, b = %d", tc.a, tc.b)

			actualLcm := LCM(tc.a, tc.b)

			if tc.expectPanic {

				return
			}

			if actualLcm != tc.expectedLcm {

				t.Errorf("Test %s failed: LCM(%d, %d) = %d; want %d",
					tc.name, tc.a, tc.b, actualLcm, tc.expectedLcm)

				switch tc.name {
				case "Scenario 1: Basic Positive Integers (6, 8)":
					t.Logf("Explanation: Expected LCM using formula (a*b)/GCD(a,b) = (6*8)/GCD(6,8) = 48/2 = 24.")
				case "Scenario 7: One Negative Input (-7, 5)":
					t.Logf("Explanation: Based on formula (a*b)/GCD(a,b) = (-7*5)/GCD(-7,5) = -35/1 = -35. Note: Standard LCM definition is often positive.")
				case "Scenario 8: Both Negative Inputs (-6, -8)":
					t.Logf("Explanation: Based on formula (a*b)/GCD(a,b) = (-6*-8)/GCD(-6,-8) = 48/-2 = -24. Note: Standard LCM definition is often positive.")
				case "Scenario 10: Large Integers (Overflow Check)":
					t.Logf("Explanation: Expected result reflects intermediate overflow in (a*b). (math.MaxInt * 2) overflows to -2 on typical 64-bit systems.")

				default:
					t.Logf("Explanation: Expected result %d based on mathematical definition or specific edge case handling.", tc.expectedLcm)
				}
			} else {

				t.Logf("Test %s passed: LCM(%d, %d) = %d", tc.name, tc.a, tc.b, actualLcm)
			}

			if tc.name == "Scenario 9: Commutativity Check (9, 6)" {
				lcmBA := LCM(tc.b, tc.a)
				if actualLcm != lcmBA {
					t.Errorf("Test %s failed commutativity check: LCM(%d, %d) [%d] != LCM(%d, %d) [%d]",
						tc.name, tc.a, tc.b, actualLcm, tc.b, tc.a, lcmBA)
				} else {
					t.Logf("Test %s passed commutativity check: LCM(%d, %d) == LCM(%d, %d) == %d",
						tc.name, tc.a, tc.b, tc.b, tc.a, actualLcm)
				}
			}

			if tc.name == "Scenario 9b: Commutativity Check (-8, -6)" {
				lcmBA := LCM(tc.b, tc.a)
				if actualLcm != lcmBA {
					t.Errorf("Test %s failed commutativity check: LCM(%d, %d) [%d] != LCM(%d, %d) [%d]",
						tc.name, tc.a, tc.b, actualLcm, tc.b, tc.a, lcmBA)
				} else {
					t.Logf("Test %s passed commutativity check: LCM(%d, %d) == LCM(%d, %d) == %d",
						tc.name, tc.a, tc.b, tc.b, tc.a, actualLcm)
				}
			}
		})
	}
}
