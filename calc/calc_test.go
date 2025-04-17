package calc

import (
	fmt "fmt"
	math "math"
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Absolute_f8af7505a1
ROOST_METHOD_SIG_HASH=Absolute_4bad226818

FUNCTION_DEF=func Absolute(num float64) float64
*/
func TestAbsolute(t *testing.T) {

	testCases := []struct {
		name          string
		input         float64
		expected      float64
		expectNaN     bool
		expectInf     bool
		infSign       int
		checkSignBit  bool
		expectedPanic bool
	}{

		{
			name:     "Positive Input",
			input:    123.45,
			expected: 123.45,
		},

		{
			name:     "Negative Input",
			input:    -987.65,
			expected: 987.65,
		},

		{
			name:     "Zero Input",
			input:    0.0,
			expected: 0.0,
		},

		{
			name:         "Negative Zero Input",
			input:        math.Copysign(0.0, -1),
			expected:     0.0,
			checkSignBit: true,
		},

		{
			name:      "Positive Infinity Input",
			input:     math.Inf(1),
			expectInf: true,
			infSign:   1,
		},

		{
			name:      "Negative Infinity Input",
			input:     math.Inf(-1),
			expectInf: true,
			infSign:   1,
		},

		{
			name:      "NaN Input",
			input:     math.NaN(),
			expectNaN: true,
		},

		{
			name:     "Maximum Positive Float64 Input",
			input:    math.MaxFloat64,
			expected: math.MaxFloat64,
		},

		{
			name:     "Most Negative Float64 Input",
			input:    -math.MaxFloat64,
			expected: math.MaxFloat64,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test case '%s' panicked unexpectedly.", tc.name)
				}

			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: %v", tc.input)

			actual := Absolute(tc.input)

			if tc.expectNaN {
				if !math.IsNaN(actual) {
					t.Errorf("Test Case '%s' Failed: Expected NaN, but got %v", tc.name, actual)
				} else {
					t.Logf("Test Case '%s' Passed: Correctly resulted in NaN", tc.name)
				}
			} else if tc.expectInf {
				if !math.IsInf(actual, tc.infSign) {
					expectedSign := "positive"
					if tc.infSign == -1 {
						expectedSign = "negative"
					}
					t.Errorf("Test Case '%s' Failed: Expected %s infinity, but got %v", tc.name, expectedSign, actual)
				} else {
					expectedSign := "positive"
					if tc.infSign == -1 {
						expectedSign = "negative"
					}
					t.Logf("Test Case '%s' Passed: Correctly resulted in %s infinity", tc.name, expectedSign)
				}
			} else {

				if actual != tc.expected {
					t.Errorf("Test Case '%s' Failed: Input=%v, Expected=%v, Got=%v", tc.name, tc.input, tc.expected, actual)
				} else {

					if tc.checkSignBit && math.Signbit(actual) {
						t.Errorf("Test Case '%s' Failed: Expected positive zero (+0.0), but got negative zero (-0.0). Input=%v, Got=%v", tc.name, tc.input, actual)
					} else {
						t.Logf("Test Case '%s' Passed: Input=%v, Expected=%v, Got=%v", tc.name, tc.input, tc.expected, actual)
					}
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

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{

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
			num2:     -10,
			expected: 5,
		},

		{
			name:     "Scenario 4: Add a Negative and a Positive Integer",
			num1:     -15,
			num2:     10,
			expected: -5,
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
			name:     "Scenario 7: Add Two Zeros",
			num1:     0,
			num2:     0,
			expected: 0,
		},

		{
			name:     "Scenario 8: Add Numbers Resulting in Zero",
			num1:     50,
			num2:     -50,
			expected: 0,
		},

		{
			name:     "Scenario 9: Add Numbers Approaching Integer Maximum (No Overflow)",
			num1:     math.MaxInt / 2,
			num2:     math.MaxInt / 3,
			expected: (math.MaxInt / 2) + (math.MaxInt / 3),
		},

		{
			name:     "Scenario 10: Add Numbers Approaching Integer Minimum (No Underflow)",
			num1:     math.MinInt / 2,
			num2:     math.MinInt / 3,
			expected: (math.MinInt / 2) + (math.MinInt / 3),
		},

		{
			name:     "Scenario 11: Add math.MaxInt and Zero",
			num1:     math.MaxInt,
			num2:     0,
			expected: math.MaxInt,
		},

		{
			name:     "Scenario 12: Add math.MinInt and Zero",
			num1:     math.MinInt,
			num2:     0,
			expected: math.MinInt,
		},

		{
			name:     "Scenario 13: Test Potential Integer Overflow (Wrap Around)",
			num1:     math.MaxInt,
			num2:     1,
			expected: math.MinInt,
		},

		{
			name:     "Scenario 14: Test Potential Integer Underflow (Wrap Around)",
			num1:     math.MinInt,
			num2:     -1,
			expected: math.MaxInt,
		},

		{
			name:     "Scenario 15: Add math.MaxInt and math.MinInt",
			num1:     math.MaxInt,
			num2:     math.MinInt,
			expected: -1,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked unexpectedly: %v", r)
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: num1 = %d, num2 = %d", tc.num1, tc.num2)
			t.Logf("Expected result: %d", tc.expected)

			actual := Add(tc.num1, tc.num2)
			t.Logf("Actual result: %d", actual)

			if actual != tc.expected {

				t.Errorf("FAIL: For inputs num1=%d, num2=%d: expected %d, but got %d",
					tc.num1, tc.num2, tc.expected, actual)
			} else {

				t.Logf("PASS: Test case '%s' passed.", tc.name)
			}

			fmt.Println()
		})
	}
}

/*
ROOST_METHOD_HASH=Divide_f2ddee767d
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64
*/
func TestDivide(t *testing.T) {

	const tolerance = 1e-9

	testCases := []struct {
		description      string
		num1             float64
		num2             float64
		expectedResult   float64
		expectPanic      bool
		expectedPanicMsg string
		checkInf         bool
		infSign          int
	}{

		{
			description:    "Scenario 1: Basic Positive Division (10.0 / 2.0)",
			num1:           10.0,
			num2:           2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 2: Basic Negative Numerator Division (-10.0 / 2.0)",
			num1:           -10.0,
			num2:           2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 3: Basic Negative Denominator Division (10.0 / -2.0)",
			num1:           10.0,
			num2:           -2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 4: Basic Both Negative Division (-10.0 / -2.0)",
			num1:           -10.0,
			num2:           -2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 5: Division Resulting in a Fraction (5.0 / 2.0)",
			num1:           5.0,
			num2:           2.0,
			expectedResult: 2.5,
			expectPanic:    false,
		},

		{
			description:    "Scenario 6: Division with Zero Numerator (0.0 / 5.0)",
			num1:           0.0,
			num2:           5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			description:      "Scenario 7: Division by Zero (10.0 / 0.0)",
			num1:             10.0,
			num2:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "division by zero is not allowed",
		},

		{
			description:      "Scenario 8: Division by Zero with Negative Numerator (-10.0 / 0.0)",
			num1:             -10.0,
			num2:             0.0,
			expectPanic:      true,
			expectedPanicMsg: "division by zero is not allowed",
		},

		{
			description:    "Scenario 9: Division with Large Numbers (1.0e15 / 1.0e5)",
			num1:           1.0e15,
			num2:           1.0e5,
			expectedResult: 1.0e10,
			expectPanic:    false,
		},

		{
			description:    "Scenario 10: Division Resulting in Very Small Number (1.0e-10 / 1.0e10)",
			num1:           1.0e-10,
			num2:           1.0e10,
			expectedResult: 1.0e-20,
			expectPanic:    false,
		},

		{
			description:    "Scenario 11: Division involving math.MaxFloat64 (MaxFloat64 / 2.0)",
			num1:           math.MaxFloat64,
			num2:           2.0,
			expectedResult: math.MaxFloat64 / 2.0,
			expectPanic:    false,
		},

		{
			description:    "Scenario 12: Division by math.MaxFloat64 (1.0 / MaxFloat64)",
			num1:           1.0,
			num2:           math.MaxFloat64,
			expectedResult: 1.0 / math.MaxFloat64,
			expectPanic:    false,
		},

		{
			description:    "Scenario 13: Division involving Positive Infinity (Inf / 100.0)",
			num1:           math.Inf(1),
			num2:           100.0,
			expectedResult: math.Inf(1),
			expectPanic:    false,
			checkInf:       true,
			infSign:        1,
		},

		{
			description:    "Scenario 14: Division by Positive Infinity (100.0 / Inf)",
			num1:           100.0,
			num2:           math.Inf(1),
			expectedResult: 0.0,
			expectPanic:    false,
		},
	}

	for _, tc := range testCases {

		tc := tc

		t.Run(tc.description, func(t *testing.T) {

			defer func() {
				r := recover()

				if tc.expectPanic {
					if r == nil {

						t.Errorf("FAIL: Expected a panic for scenario '%s', but did not get one.", tc.description)
					} else {

						if panicMsg, ok := r.(string); ok {
							if panicMsg == tc.expectedPanicMsg {
								t.Logf("PASS: Successfully caught expected panic: %v", r)
							} else {
								t.Errorf("FAIL: Expected panic message '%s' but got '%s'", tc.expectedPanicMsg, panicMsg)
							}
						} else {

							t.Errorf("FAIL: Expected panic message '%s' but got panic with different type: %T, value: %v", tc.expectedPanicMsg, r, r)
						}
					}
				} else if r != nil {

					t.Errorf("FAIL: Unexpected panic encountered for scenario '%s': %v\nStack trace:\n%s", tc.description, r, string(debug.Stack()))
				}

			}()

			t.Logf("Running test case: %s", tc.description)
			t.Logf("Arrange: num1 = %v, num2 = %v", tc.num1, tc.num2)

			actualResult := Divide(tc.num1, tc.num2)
			t.Logf("Act: Called Divide(%v, %v), got result: %v", tc.num1, tc.num2, actualResult)

			if !tc.expectPanic {

				if tc.checkInf {
					if math.IsInf(actualResult, tc.infSign) {
						infType := "Positive"
						if tc.infSign == -1 {
							infType = "Negative"
						}
						t.Logf("PASS: Assert: Result is %s Infinity as expected.", infType)
					} else {
						infType := "Positive"
						if tc.infSign == -1 {
							infType = "Negative"
						}
						t.Errorf("FAIL: Assert: Expected %s Infinity, but got %v.", infType, actualResult)
					}
				} else {

					if math.Abs(actualResult-tc.expectedResult) < tolerance || (actualResult == 0 && tc.expectedResult == 0) {
						t.Logf("PASS: Assert: Result %v is within tolerance %v of expected %v.", actualResult, tolerance, tc.expectedResult)
					} else {
						t.Errorf("FAIL: Assert: Expected %v, got %v (difference %v, tolerance %v)", tc.expectedResult, actualResult, math.Abs(actualResult-tc.expectedResult), tolerance)
					}
				}
			} else {

				t.Logf("Assert: Skipped result check as panic was expected.")
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

	testCases := []struct {
		name     string
		n        int
		want     int
		panicMsg string
	}{
		{
			name:     "Scenario 1: Factorial of Zero",
			n:        0,
			want:     1,
			panicMsg: "",
		},
		{
			name:     "Scenario 2: Factorial of One",
			n:        1,
			want:     1,
			panicMsg: "",
		},
		{
			name:     "Scenario 3: Factorial of a Small Positive Number (5)",
			n:        5,
			want:     120,
			panicMsg: "",
		},
		{
			name:     "Scenario 4: Factorial of a Larger Positive Number (10)",
			n:        10,
			want:     3628800,
			panicMsg: "",
		},
		{
			name:     "Scenario 5: Factorial of a Negative Number (-1)",
			n:        -1,
			want:     0,
			panicMsg: "factorial is not defined for negative numbers",
		},
		{
			name:     "Scenario 6: Factorial of Another Negative Number (-10)",
			n:        -10,
			want:     0,
			panicMsg: "factorial is not defined for negative numbers",
		},
		{
			name: "Scenario 7: Factorial Near Integer Overflow Boundary (20)",

			n:        20,
			want:     2432902008176640000,
			panicMsg: "",
		},
	}

	for _, tc := range testCases {

		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.panicMsg != "" {

					if r == nil {
						t.Errorf("FAIL: Expected panic with message '%s', but Factorial(%d) did not panic.", tc.panicMsg, tc.n)
					} else {

						recoveredMsg := fmt.Sprintf("%v", r)
						if recoveredMsg != tc.panicMsg {
							t.Errorf("FAIL: Expected panic message '%s', but got '%s' for Factorial(%d).", tc.panicMsg, recoveredMsg, tc.n)
						} else {
							t.Logf("PASS: Factorial(%d) correctly panicked with message: '%s'.", tc.n, tc.panicMsg)
						}
					}
				} else {

					if r != nil {
						t.Errorf("FAIL: Factorial(%d) panicked unexpectedly. Recovered: %v\nStack trace:\n%s", tc.n, r, string(debug.Stack()))

					}

				}
			}()

			t.Logf("Running test case: %s - Input: Factorial(%d)", tc.name, tc.n)

			got := Factorial(tc.n)

			if tc.panicMsg == "" {
				if got != tc.want {
					t.Errorf("FAIL: Factorial(%d) = %d; want %d", tc.n, got, tc.want)
				} else {
					t.Logf("PASS: Factorial(%d) = %d, as expected.", tc.n, got)
				}
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
		scenario string
	}

	testTable := []testCase{
		{
			name:     "Scenario 1: Basic Positive (a > b)",
			a:        54,
			b:        24,
			expected: 6,
			scenario: "Test GCD with a > b (54, 24). Expected GCD is 6.",
		},
		{
			name:     "Scenario 2: Basic Positive (b > a)",
			a:        24,
			b:        54,
			expected: 6,
			scenario: "Test GCD with b > a (24, 54). Expected GCD is 6 (commutative property).",
		},
		{
			name:     "Scenario 3: One Multiple of Other",
			a:        48,
			b:        12,
			expected: 12,
			scenario: "Test GCD where a is a multiple of b (48, 12). Expected GCD is the smaller number (12).",
		},
		{
			name:     "Scenario 4: Coprime Numbers",
			a:        17,
			b:        23,
			expected: 1,
			scenario: "Test GCD with coprime numbers (17, 23). Expected GCD is 1.",
		},
		{
			name:     "Scenario 5: Equal Positive Numbers",
			a:        30,
			b:        30,
			expected: 30,
			scenario: "Test GCD with equal positive numbers (30, 30). Expected GCD is the number itself (30).",
		},
		{
			name:     "Scenario 6: One Argument is Zero (b=0)",
			a:        42,
			b:        0,
			expected: 42,
			scenario: "Test GCD base case where b is 0 (42, 0). Expected GCD is a (42).",
		},
		{
			name:     "Scenario 7: One Argument is Zero (a=0)",
			a:        0,
			b:        15,
			expected: 15,
			scenario: "Test GCD where a is 0 (0, 15). Expected GCD is b (15) after one recursion.",
		},
		{
			name:     "Scenario 8: Both Arguments Zero",
			a:        0,
			b:        0,
			expected: 0,
			scenario: "Test GCD where both a and b are 0 (0, 0). Expected GCD is 0 based on implementation's base case.",
		},
		{
			name:     "Scenario 9: One Negative Argument (b < 0)",
			a:        54,
			b:        -24,
			expected: 6,
			scenario: "Test GCD with one negative argument (54, -24). Expected GCD is 6, considering Go's modulo behavior.",
		},
		{
			name:     "Scenario 10: Both Negative Arguments",
			a:        -54,
			b:        -24,
			expected: -6,
			scenario: "Test GCD with both negative arguments (-54, -24). Expected GCD is -6, highlighting Go's modulo behavior.",
		},
		{
			name:     "Scenario 11: Large Integer Inputs",
			a:        196350,
			b:        135135,
			expected: 135,
			scenario: "Test GCD with large integer inputs (196350, 135135). Expected GCD is 135.",
		},
	}

	for _, tc := range testTable {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked, failing.")
				}
			}()

			t.Logf("Running test: %s", tc.scenario)
			t.Logf("Arrange: a = %d, b = %d", tc.a, tc.b)

			actual := GCD(tc.a, tc.b)
			t.Logf("Act: Called GCD(%d, %d). Result: %d", tc.a, tc.b, actual)

			if actual != tc.expected {

				t.Errorf("Assert: FAILED. Expected %d, but got %d.", tc.expected, actual)
			} else {

				t.Logf("Assert: PASSED. Expected %d, got %d.", tc.expected, actual)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=LCM_6035446662
ROOST_METHOD_SIG_HASH=LCM_121c872fbf

FUNCTION_DEF=func LCM(a, b int) int
*/
func TestLcm(t *testing.T) {

	type lcmTestCase struct {
		name        string
		a           int
		b           int
		expected    int
		expectPanic bool
	}

	testCases := []lcmTestCase{

		{name: "Scenario 1: Basic Positive Integers (4, 6)", a: 4, b: 6, expected: 12, expectPanic: false},

		{name: "Scenario 2: One Multiple of Other (3, 9)", a: 3, b: 9, expected: 9, expectPanic: false},

		{name: "Scenario 3: Co-prime Positive (7, 5)", a: 7, b: 5, expected: 35, expectPanic: false},

		{name: "Scenario 4: Identical Positive (10, 10)", a: 10, b: 10, expected: 10, expectPanic: false},

		{name: "Scenario 5: One Input Zero (15, 0)", a: 15, b: 0, expected: 0, expectPanic: false},
		{name: "Scenario 5: One Input Zero (0, 15)", a: 0, b: 15, expected: 0, expectPanic: false},

		{name: "Scenario 6: Both Inputs Zero (0, 0)", a: 0, b: 0, expected: 0, expectPanic: true},

		{name: "Scenario 7: Mixed Signs (6, -9)", a: 6, b: -9, expected: -18, expectPanic: false},
		{name: "Scenario 7: Mixed Signs (-6, 9)", a: -6, b: 9, expected: -18, expectPanic: false},

		{name: "Scenario 8: Both Negative (-4, -6)", a: -4, b: -6, expected: 12, expectPanic: false},

		{name: "Scenario 9: Includes 1 (1, 13)", a: 1, b: 13, expected: 13, expectPanic: false},
		{name: "Scenario 9: Includes 1 (13, 1)", a: 13, b: 1, expected: 13, expectPanic: false},

		{name: "Scenario 10: Includes -1 (-1, 13)", a: -1, b: 13, expected: -13, expectPanic: false},
		{name: "Scenario 10: Includes -1 (13, -1)", a: 13, b: -1, expected: -13, expectPanic: false},
		{name: "Scenario 10: Includes -1 (-1, -13)", a: -1, b: -13, expected: 13, expectPanic: false},

		{
			name:        "Scenario 11: Large Inputs (Potential Intermediate Overflow)",
			a:           3 * (1 << 30),
			b:           5 * (1 << 30),
			expected:    15 * (1 << 30),
			expectPanic: false,
		},

		{name: "Edge Case: Zero and Negative (0, -5)", a: 0, b: -5, expected: 0, expectPanic: false},
		{name: "Edge Case: MaxInt and 1", a: math.MaxInt, b: 1, expected: math.MaxInt, expectPanic: false},

		{name: "Edge Case: MinInt and 1", a: math.MinInt, b: 1, expected: math.MinInt, expectPanic: false},

		{name: "Edge Case: MaxInt and 2 (Overflow Wrap-around)", a: math.MaxInt, b: 2, expected: -2, expectPanic: false},

		{name: "Edge Case: MinInt and 2", a: math.MinInt, b: 2, expected: math.MinInt, expectPanic: false},

		{name: "Edge Case: MinInt and -1 (Intermediate Overflow Panic)", a: math.MinInt, b: -1, expected: 0, expectPanic: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if r != nil {

					if tc.expectPanic {
						t.Logf("SUCCESS: Caught expected panic for inputs a=%d, b=%d. Panic: %v", tc.a, tc.b, r)

					} else {

						t.Errorf("FAIL: Unexpected panic for inputs a=%d, b=%d. Panic: %v\n%s", tc.a, tc.b, r, string(debug.Stack()))
					}
				} else {

					if tc.expectPanic {

						t.Errorf("FAIL: Expected panic did not occur for inputs a=%d, b=%d", tc.a, tc.b)
					}

				}
			}()

			t.Logf("Running test case: %s (a=%d, b=%d)", tc.name, tc.a, tc.b)

			actual := LCM(tc.a, tc.b)

			if !tc.expectPanic {

				if actual != tc.expected {
					t.Errorf("FAIL: LCM(%d, %d) returned %d, but expected %d", tc.a, tc.b, actual, tc.expected)
				} else {
					t.Logf("SUCCESS: LCM(%d, %d) correctly returned %d", tc.a, tc.b, actual)
				}
			}

		})
	}

	t.Log("----------------------------------------------------------------------")
	t.Log("Test Suite Notes:")
	t.Logf("NOTE: The LCM function implementation `(a * b) / GCD(a, b)` is used.")
	t.Logf("  - Limitation: This formula is susceptible to intermediate integer overflow if `a * b` exceeds the limits of the 'int' type (%d to %d), even if the final LCM result is representable. See 'Scenario 11' and 'Edge Case: MaxInt and 2'.", math.MinInt, math.MaxInt)
	t.Logf("  - Limitation: The intermediate calculation `a * b` might also overflow and cause a panic before division for extreme values like `MinInt * -1`. See 'Edge Case: MinInt and -1'.")
	t.Logf("  - Suggestion: An alternative formula like `(a / GCD(a, b)) * b` (handling potential division by zero if GCD is zero, or if b is zero) can mitigate intermediate overflow but requires careful implementation regarding zero inputs.")
	t.Log("NOTE: These tests rely on an assumed `GCD` function within the 'calc' package.")
	t.Log("  - Assumption: `GCD` correctly handles positive, negative, and zero inputs according to standard mathematical definitions (GCD(a,b)=GCD(|a|,|b|), GCD(a,0)=|a|, GCD(0,0)=0).")
	t.Log("  - Assumption: `GCD` handles large integer inputs without unexpected overflow issues itself.")
	t.Log("  - Verification: The correctness of `LCM` tests depends heavily on the correctness of the underlying `GCD` implementation.")
	t.Log("----------------------------------------------------------------------")
}

/*
ROOST_METHOD_HASH=Logarithm_546f6d96c4
ROOST_METHOD_SIG_HASH=Logarithm_ddbb699678

FUNCTION_DEF=func Logarithm(num, base float64) float64
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
			name:           "Scenario 1: Basic Integer Result (log2(8))",
			num:            8.0,
			base:           2.0,
			expectedResult: 3.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 2: Fractional Result (log2(10))",
			num:            10.0,
			base:           2.0,
			expectedResult: math.Log(10.0) / math.Log(2.0),
			expectPanic:    false,
		},

		{
			name:           "Scenario 3: Base > Number (log10(5))",
			num:            5.0,
			base:           10.0,
			expectedResult: math.Log(5.0) / math.Log(10.0),
			expectPanic:    false,
		},

		{
			name:           "Scenario 4: Number == Base (log10(10))",
			num:            10.0,
			base:           10.0,
			expectedResult: 1.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 5: Number == 1 (log5(1))",
			num:            1.0,
			base:           5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 6: Base between 0 and 1 (log0.5(8))",
			num:            8.0,
			base:           0.5,
			expectedResult: -3.0,
			expectPanic:    false,
		},

		{
			name:             "Scenario 7: Panic on Zero Number (log10(0))",
			num:              0.0,
			base:             10.0,
			expectedResult:   math.NaN(),
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 8: Panic on Negative Number (log10(-5))",
			num:              -5.0,
			base:             10.0,
			expectedResult:   math.NaN(),
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 9: Panic on Zero Base (log0(10))",
			num:              10.0,
			base:             0.0,
			expectedResult:   math.NaN(),
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 10: Panic on Negative Base (log-2(10))",
			num:              10.0,
			base:             -2.0,
			expectedResult:   math.NaN(),
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:             "Scenario 11: Panic on Base == 1 (log1(10))",
			num:              10.0,
			base:             1.0,
			expectedResult:   math.NaN(),
			expectPanic:      true,
			expectedPanicMsg: "logarithm is not defined for these values",
		},

		{
			name:           "Scenario 12: Very Large Number (log10(1e50))",
			num:            1e50,
			base:           10.0,
			expectedResult: 50.0,
			expectPanic:    false,
		},

		{
			name:           "Scenario 13: Very Small Positive Number (log10(1e-50))",
			num:            1e-50,
			base:           10.0,
			expectedResult: -50.0,
			expectPanic:    false,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.expectPanic {

					if r == nil {
						t.Errorf("FAIL: Expected panic for num=%.2f, base=%.2f, but did not panic.", tc.num, tc.base)
					} else {

						if msg, ok := r.(string); !ok || msg != tc.expectedPanicMsg {
							t.Errorf("FAIL: Expected panic message '%s', but got '%v' (type %T).", tc.expectedPanicMsg, r, r)
						} else {

							t.Logf("PASS: Correctly panicked with message '%s' for num=%.2f, base=%.2f.", tc.expectedPanicMsg, tc.num, tc.base)
						}
					}
				} else {

					if r != nil {

						t.Errorf("FAIL: Unexpected panic for num=%.2f, base=%.2f: %v\n%s", tc.num, tc.base, r, string(debug.Stack()))
					}

				}
			}()

			actualResult := Logarithm(tc.num, tc.base)

			if !tc.expectPanic {

				if !approxEqual(actualResult, tc.expectedResult, epsilon) {
					t.Errorf("FAIL: Logarithm(%.2f, %.2f) = %.15f; want %.15f (Difference: %.15f)",
						tc.num, tc.base, actualResult, tc.expectedResult, math.Abs(actualResult-tc.expectedResult))
				} else {

					t.Logf("PASS: Logarithm(%.2f, %.2f) = %.15f.", tc.num, tc.base, actualResult)
				}
			}

		})
	}
}

func approxEqual(a, b, epsilon float64) bool {

	if math.IsNaN(a) || math.IsNaN(b) {
		return math.IsNaN(a) && math.IsNaN(b)
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return math.IsInf(a, 0) && math.IsInf(b, 0) && (math.Signbit(a) == math.Signbit(b))
	}
	return math.Abs(a-b) < epsilon
}

/*
ROOST_METHOD_HASH=Modulo_eb9c4baeed
ROOST_METHOD_SIG_HASH=Modulo_09898f6fed

FUNCTION_DEF=func Modulo(num1, num2 int) int
*/
func TestModulo(t *testing.T) {

	testCases := []struct {
		name        string
		num1        int
		num2        int
		expected    int
		expectPanic bool
	}{

		{
			name:        "Scenario 1: Basic Positive Modulo",
			num1:        10,
			num2:        3,
			expected:    1,
			expectPanic: false,
		},

		{
			name:        "Scenario 2: Positive Dividend Smaller",
			num1:        3,
			num2:        10,
			expected:    3,
			expectPanic: false,
		},

		{
			name:        "Scenario 3: Zero Result",
			num1:        10,
			num2:        5,
			expected:    0,
			expectPanic: false,
		},

		{
			name:        "Scenario 4: Negative Dividend",
			num1:        -10,
			num2:        3,
			expected:    -1,
			expectPanic: false,
		},

		{
			name:        "Scenario 5: Negative Divisor",
			num1:        10,
			num2:        -3,
			expected:    1,
			expectPanic: false,
		},

		{
			name:        "Scenario 6: Both Negative",
			num1:        -10,
			num2:        -3,
			expected:    -1,
			expectPanic: false,
		},

		{
			name:        "Scenario 7: Zero Dividend",
			num1:        0,
			num2:        5,
			expected:    0,
			expectPanic: false,
		},

		{
			name:        "Scenario 8: Zero Divisor Panic",
			num1:        10,
			num2:        0,
			expectPanic: true,
		},

		{
			name:        "Scenario 9: Both Zero Panic",
			num1:        0,
			num2:        0,
			expectPanic: true,
		},

		{
			name:        "Scenario 10: Divisor One",
			num1:        123,
			num2:        1,
			expected:    0,
			expectPanic: false,
		},

		{
			name:        "Scenario 11: Divisor Minus One",
			num1:        123,
			num2:        -1,
			expected:    0,
			expectPanic: false,
		},

		{
			name:        "Scenario 12: MaxInt Dividend",
			num1:        math.MaxInt,
			num2:        10,
			expected:    math.MaxInt % 10,
			expectPanic: false,
		},

		{
			name:        "Scenario 13: MinInt Dividend",
			num1:        math.MinInt,
			num2:        10,
			expected:    math.MinInt % 10,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if tc.expectPanic {
					if r == nil {

						t.Errorf("FAIL: Expected a panic for Modulo(%d, %d) but did not get one.", tc.num1, tc.num2)
					} else {

						t.Logf("PASS: Successfully caught expected panic for Modulo(%d, %d): %v", tc.num1, tc.num2, r)
					}
				} else {
					if r != nil {

						t.Errorf("FAIL: Unexpected panic for Modulo(%d, %d): %v\n%s", tc.num1, tc.num2, r, string(debug.Stack()))
					}

				}
			}()

			t.Logf("Running test case: %s, Inputs: num1=%d, num2=%d", tc.name, tc.num1, tc.num2)

			actual := Modulo(tc.num1, tc.num2)

			if !tc.expectPanic {
				if actual != tc.expected {
					t.Errorf("FAIL: Modulo(%d, %d) = %d; want %d", tc.num1, tc.num2, actual, tc.expected)
				} else {
					t.Logf("PASS: Modulo(%d, %d) = %d. Result matches expected value.", tc.num1, tc.num2, actual)
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

	const tolerance = 1e-9

	testCases := []struct {
		name            string
		num1            float64
		num2            float64
		expected        float64
		useTolerance    bool
		expectNaN       bool
		expectInf       bool
		expectedInfSign int
	}{

		{
			name:     "Scenario 1: Positive * Positive",
			num1:     5.0,
			num2:     4.0,
			expected: 20.0,
		},

		{
			name:     "Scenario 2: Positive * Negative",
			num1:     6.0,
			num2:     -3.0,
			expected: -18.0,
		},

		{
			name:     "Scenario 3: Negative * Negative",
			num1:     -7.0,
			num2:     -2.0,
			expected: 14.0,
		},

		{
			name:     "Scenario 4a: NonZero * Zero",
			num1:     123.45,
			num2:     0.0,
			expected: 0.0,
		},
		{
			name:     "Scenario 4b: Zero * NonZero",
			num1:     0.0,
			num2:     -99.9,
			expected: 0.0,
		},

		{
			name:     "Scenario 5a: NonOne * One",
			num1:     987.65,
			num2:     1.0,
			expected: 987.65,
		},
		{
			name:     "Scenario 5b: One * NonOne",
			num1:     1.0,
			num2:     -123.45,
			expected: -123.45,
		},

		{
			name:         "Scenario 6: Non-Integer Floats",
			num1:         2.5,
			num2:         3.5,
			expected:     8.75,
			useTolerance: true,
		},

		{
			name:            "Scenario 7a: Positive Overflow to +Inf",
			num1:            math.MaxFloat64 / 2,
			num2:            3.0,
			expectInf:       true,
			expectedInfSign: 1,
		},
		{
			name:            "Scenario 7b: Negative Overflow to -Inf",
			num1:            math.MaxFloat64,
			num2:            -2.0,
			expectInf:       true,
			expectedInfSign: -1,
		},

		{
			name:     "Scenario 8: Underflow to Zero",
			num1:     1e-200,
			num2:     1e-200,
			expected: 0.0,
		},

		{
			name:            "Scenario 9a: Finite Positive * +Inf",
			num1:            5.0,
			num2:            math.Inf(1),
			expectInf:       true,
			expectedInfSign: 1,
		},
		{
			name:            "Scenario 9b: Finite Negative * +Inf",
			num1:            -5.0,
			num2:            math.Inf(1),
			expectInf:       true,
			expectedInfSign: -1,
		},
		{
			name:            "Scenario 9c: +Inf * +Inf",
			num1:            math.Inf(1),
			num2:            math.Inf(1),
			expectInf:       true,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 10a: Finite Positive * -Inf",
			num1:            5.0,
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: -1,
		},
		{
			name:            "Scenario 10b: Finite Negative * -Inf",
			num1:            -5.0,
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: 1,
		},
		{
			name:            "Scenario 10c: -Inf * -Inf",
			num1:            math.Inf(-1),
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: 1,
		},
		{
			name:            "Scenario 10d: +Inf * -Inf",
			num1:            math.Inf(1),
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: -1,
		},

		{
			name:      "Scenario 11a: Finite * NaN",
			num1:      10.0,
			num2:      math.NaN(),
			expectNaN: true,
		},
		{
			name:      "Scenario 11b: NaN * Finite",
			num1:      math.NaN(),
			num2:      -5.0,
			expectNaN: true,
		},
		{
			name:      "Scenario 11c: NaN * NaN",
			num1:      math.NaN(),
			num2:      math.NaN(),
			expectNaN: true,
		},
		{
			name:      "Scenario 11d: Inf * NaN",
			num1:      math.Inf(1),
			num2:      math.NaN(),
			expectNaN: true,
		},

		{
			name:      "Scenario 12a: Zero * +Inf",
			num1:      0.0,
			num2:      math.Inf(1),
			expectNaN: true,
		},
		{
			name:      "Scenario 12b: -Inf * Zero",
			num1:      math.Inf(-1),
			num2:      0.0,
			expectNaN: true,
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

			t.Logf("Testing scenario: %s", tc.name)
			t.Logf("Input: num1 = %v, num2 = %v", tc.num1, tc.num2)

			result := Multiply(tc.num1, tc.num2)

			if tc.expectNaN {

				if !math.IsNaN(result) {
					t.Errorf("FAIL: Expected NaN, but got %v. Logic: Operations involving NaN (or 0 * Inf) should result in NaN (IEEE 754).", result)
				} else {
					t.Logf("PASS: Got NaN as expected.")
				}
			} else if tc.expectInf {

				expectedSignStr := "+"
				if tc.expectedInfSign == -1 {
					expectedSignStr = "-"
				}
				if !math.IsInf(result, tc.expectedInfSign) {
					t.Errorf("FAIL: Expected %sInfinity, but got %v. Logic: Multiplication resulted in overflow or involved Infinity according to IEEE 754 rules.", expectedSignStr, result)
				} else {
					t.Logf("PASS: Got %sInfinity as expected.", expectedSignStr)
				}
			} else if tc.useTolerance {

				if diff := math.Abs(result - tc.expected); diff > tolerance {
					t.Errorf("FAIL: Expected result within tolerance %v of %v, but got %v (difference %v). Logic: Standard multiplication expected, allowing for minor floating-point inaccuracies.", tolerance, tc.expected, result, diff)
				} else {
					t.Logf("PASS: Result %v is within tolerance %v of expected %v.", result, tolerance, tc.expected)
				}
			} else {

				if result != tc.expected {
					t.Errorf("FAIL: Expected exactly %v, but got %v. Logic: Based on standard mathematical rules for multiplication (integers, zero, one).", tc.expected, result)
				} else {
					t.Logf("PASS: Got exact expected result %v.", result)
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=Power_2542f03efe
ROOST_METHOD_SIG_HASH=Power_feb9859574

FUNCTION_DEF=func Power(base, exponent float64) float64
*/
func TestPower(t *testing.T) {

	testCases := []struct {
		name        string
		base        float64
		exponent    float64
		expected    float64
		expectNaN   bool
		expectInf   int
		description string
	}{

		{
			name:        "Scenario 1: Positive Base, Positive Integer Exponent",
			base:        2.0,
			exponent:    3.0,
			expected:    8.0,
			description: "Test basic exponentiation: 2^3 should be 8.",
		},

		{
			name:        "Scenario 2: Positive Base, Positive Fractional Exponent",
			base:        9.0,
			exponent:    0.5,
			expected:    3.0,
			description: "Test fractional exponent (square root): 9^0.5 should be 3.",
		},

		{
			name:        "Scenario 3: Positive Base, Zero Exponent",
			base:        5.0,
			exponent:    0.0,
			expected:    1.0,
			description: "Test zero exponent: 5^0 should be 1.",
		},

		{
			name:        "Scenario 4: Positive Base, Negative Integer Exponent",
			base:        2.0,
			exponent:    -2.0,
			expected:    0.25,
			description: "Test negative exponent: 2^-2 should be 1/(2^2) = 0.25.",
		},

		{
			name:        "Scenario 5: Zero Base, Positive Exponent",
			base:        0.0,
			exponent:    5.0,
			expected:    0.0,
			description: "Test zero base with positive exponent: 0^5 should be 0.",
		},

		{
			name:        "Scenario 6: Zero Base, Zero Exponent",
			base:        0.0,
			exponent:    0.0,
			expected:    1.0,
			description: "Test zero base with zero exponent: 0^0 should be 1 (by math.Pow convention).",
		},

		{
			name:        "Scenario 7: Zero Base, Negative Exponent",
			base:        0.0,
			exponent:    -2.0,
			expectInf:   1,
			description: "Test zero base with negative exponent: 0^-2 should be +Inf.",
		},

		{
			name:        "Scenario 8: Negative Base, Positive Odd Integer Exponent",
			base:        -2.0,
			exponent:    3.0,
			expected:    -8.0,
			description: "Test negative base with odd exponent: (-2)^3 should be -8.",
		},

		{
			name:        "Scenario 9: Negative Base, Positive Even Integer Exponent",
			base:        -2.0,
			exponent:    2.0,
			expected:    4.0,
			description: "Test negative base with even exponent: (-2)^2 should be 4.",
		},

		{
			name:        "Scenario 10: Negative Base, Zero Exponent",
			base:        -5.0,
			exponent:    0.0,
			expected:    1.0,
			description: "Test negative base with zero exponent: (-5)^0 should be 1.",
		},

		{
			name:        "Scenario 11: Negative Base, Negative Odd Integer Exponent",
			base:        -2.0,
			exponent:    -3.0,
			expected:    -0.125,
			description: "Test negative base with negative odd exponent: (-2)^-3 should be 1/(-8) = -0.125.",
		},

		{
			name:        "Scenario 12: Negative Base, Negative Even Integer Exponent",
			base:        -2.0,
			exponent:    -2.0,
			expected:    0.25,
			description: "Test negative base with negative even exponent: (-2)^-2 should be 1/4 = 0.25.",
		},

		{
			name:        "Scenario 13: Negative Base, Fractional Exponent",
			base:        -4.0,
			exponent:    0.5,
			expectNaN:   true,
			description: "Test negative base with fractional exponent: (-4)^0.5 should be NaN (complex result).",
		},

		{
			name:        "Scenario 14: Base One, Any Exponent",
			base:        1.0,
			exponent:    123.45,
			expected:    1.0,
			description: "Test base 1: 1^123.45 should be 1.",
		},

		{
			name:        "Scenario 15: Base NaN",
			base:        math.NaN(),
			exponent:    2.0,
			expectNaN:   true,
			description: "Test NaN base: NaN^2 should be NaN.",
		},

		{
			name:        "Scenario 16: Exponent NaN",
			base:        2.0,
			exponent:    math.NaN(),
			expectNaN:   true,
			description: "Test NaN exponent: 2^NaN should be NaN.",
		},

		{
			name:        "Scenario 17.1: Base Positive Infinity, Positive Exponent",
			base:        math.Inf(1),
			exponent:    2.0,
			expectInf:   1,
			description: "Test +Inf base with positive exponent: (+Inf)^2 should be +Inf.",
		},
		{
			name:        "Scenario 17.2: Base Positive Infinity, Negative Exponent",
			base:        math.Inf(1),
			exponent:    -2.0,
			expected:    0.0,
			description: "Test +Inf base with negative exponent: (+Inf)^-2 should be 0.",
		},

		{
			name:        "Scenario 18.1: Base > 1, Exponent Positive Infinity",
			base:        2.0,
			exponent:    math.Inf(1),
			expectInf:   1,
			description: "Test base > 1 with +Inf exponent: 2^+Inf should be +Inf.",
		},
		{
			name:        "Scenario 18.2: Base < 1, Exponent Positive Infinity",
			base:        0.5,
			exponent:    math.Inf(1),
			expected:    0.0,
			description: "Test 0 < base < 1 with +Inf exponent: 0.5^+Inf should be 0.",
		},
		{
			name:        "Scenario 18.3: Base 1, Exponent Positive Infinity",
			base:        1.0,
			exponent:    math.Inf(1),
			expected:    1.0,
			description: "Test base 1 with +Inf exponent: 1^+Inf should be 1.",
		},
		{
			name:        "Scenario 18.4: Base -1, Exponent Positive Infinity",
			base:        -1.0,
			exponent:    math.Inf(1),
			expected:    1.0,
			description: "Test base -1 with +Inf exponent: (-1)^+Inf should be 1.",
		},

		{
			name:        "Scenario 19.1: Base > 1, Exponent Negative Infinity",
			base:        2.0,
			exponent:    math.Inf(-1),
			expected:    0.0,
			description: "Test base > 1 with -Inf exponent: 2^-Inf should be 0.",
		},
		{
			name:        "Scenario 19.2: Base < 1, Exponent Negative Infinity",
			base:        0.5,
			exponent:    math.Inf(-1),
			expectInf:   1,
			description: "Test 0 < base < 1 with -Inf exponent: 0.5^-Inf should be +Inf.",
		},
		{
			name:        "Scenario 19.3: Base 1, Exponent Negative Infinity",
			base:        1.0,
			exponent:    math.Inf(-1),
			expected:    1.0,
			description: "Test base 1 with -Inf exponent: 1^-Inf should be 1.",
		},
		{
			name:        "Scenario 19.4: Base -1, Exponent Negative Infinity",
			base:        -1.0,
			exponent:    math.Inf(-1),
			expected:    1.0,
			description: "Test base -1 with -Inf exponent: (-1)^-Inf should be 1.",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test panicked unexpectedly for base=%v, exponent=%v", tc.base, tc.exponent)
				}
			}()

			t.Logf("Running %s: base=%v, exponent=%v. Description: %s", tc.name, tc.base, tc.exponent, tc.description)

			result := Power(tc.base, tc.exponent)

			if tc.expectNaN {
				if !math.IsNaN(result) {
					t.Errorf("Power(%v, %v) = %v; expected NaN, Description: %s", tc.base, tc.exponent, result, tc.description)
				} else {
					t.Logf("Power(%v, %v) = NaN (expected). Success.", tc.base, tc.exponent)
				}
			} else if tc.expectInf != 0 {
				if !math.IsInf(result, tc.expectInf) {
					expectedInfStr := "+Inf"
					if tc.expectInf < 0 {
						expectedInfStr = "-Inf"
					}
					t.Errorf("Power(%v, %v) = %v; expected %s, Description: %s", tc.base, tc.exponent, result, expectedInfStr, tc.description)
				} else {
					expectedInfStr := "+Inf"
					if tc.expectInf < 0 {
						expectedInfStr = "-Inf"
					}
					t.Logf("Power(%v, %v) = %s (expected). Success.", tc.base, tc.exponent, expectedInfStr)
				}
			} else {

				if result != tc.expected {

					t.Errorf("Power(%v, %v) = %v; expected %v, Description: %s", tc.base, tc.exponent, result, tc.expected, tc.description)
				} else {
					t.Logf("Power(%v, %v) = %v (expected). Success.", tc.base, tc.exponent, result)
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=SinCosTan_c6521a7850
ROOST_METHOD_SIG_HASH=SinCosTan_6ec04d6e93

FUNCTION_DEF=func SinCosTan(angle float64) (sin, cos, tan float64)
*/
func TestSinCosTan(t *testing.T) {

	const epsilon = 1e-9

	testCases := []struct {
		name        string
		angle       float64
		expectedSin float64
		expectedCos float64
		expectedTan float64
		checkTanInf int
		checkNaN    bool
		description string
		validation  string
		importance  string
	}{

		{
			name:        "Scenario 1: Zero Angle",
			angle:       0.0,
			expectedSin: 0.0,
			expectedCos: 1.0,
			expectedTan: 0.0,
			checkTanInf: 0,
			checkNaN:    false,
			description: "Verify the function's output when the input angle is 0 radians.",
			validation:  "Based on trigonometric identities, sin(0) = 0, cos(0) = 1, tan(0) = 0. Direct comparison (within tolerance for floats) is used.",
			importance:  "Ensures the function handles the baseline case correctly, crucial for any application relying on trigonometric calculations.",
		},

		{
			name:        "Scenario 2: Pi/2 Angle",
			angle:       math.Pi / 2.0,
			expectedSin: 1.0,
			expectedCos: 0.0,

			checkTanInf: 1,
			checkNaN:    false,
			description: "Verify the function's output for an angle of Pi/2 radians (90 degrees).",
			validation:  "sin(pi/2) = 1, cos(pi/2) = 0. tan(x) = sin(x)/cos(x) approaches +infinity as cos(x) approaches 0 from positive side. Assert sin/cos with tolerance. Assert tan is large positive (> 1e16) due to floating point representation of Pi.",
			importance:  "Tests handling of quadrant angles and the specific case where tangent involves division by near-zero.",
		},

		{
			name:        "Scenario 3: Pi Angle",
			angle:       math.Pi,
			expectedSin: 0.0,
			expectedCos: -1.0,
			expectedTan: 0.0,
			checkTanInf: 0,
			checkNaN:    false,
			description: "Verify the function's output for an angle of Pi radians (180 degrees).",
			validation:  "Based on trigonometric identities, sin(pi) = 0, cos(pi) = -1, tan(pi) = 0. Floating-point representation requires tolerance checks, especially for expected zeros.",
			importance:  "Ensures correctness for another standard quadrant angle, important for calculations involving rotations or oscillations.",
		},

		{
			name:        "Scenario 4: 3*Pi/2 Angle",
			angle:       3.0 * math.Pi / 2.0,
			expectedSin: -1.0,
			expectedCos: 0.0,

			checkTanInf: -1,
			checkNaN:    false,
			description: "Verify the function's output for an angle of 3*Pi/2 radians (270 degrees).",
			validation:  "sin(3pi/2) = -1, cos(3pi/2) = 0. tan(x) = sin(x)/cos(x) approaches -infinity as cos(x) approaches 0 from negative side. Assert sin/cos with tolerance. Assert tan is large negative (< -1e16).",
			importance:  "Complements the Pi/2 test by checking the other case where tangent approaches infinity, ensuring robustness around division-by-near-zero.",
		},

		{
			name:        "Scenario 5: Negative Angle (-Pi/4)",
			angle:       -math.Pi / 4.0,
			expectedSin: -math.Sqrt2 / 2.0,
			expectedCos: math.Sqrt2 / 2.0,
			expectedTan: -1.0,
			checkTanInf: 0,
			checkNaN:    false,
			description: "Verify the function's output for a negative angle, specifically -Pi/4 radians (-45 degrees).",
			validation:  "Tests correct handling of negative inputs based on odd/even properties: sin(-x)=-sin(x), cos(-x)=cos(x), tan(-x)=-tan(x). Comparison uses tolerance.",
			importance:  "Ensures the function works correctly for all valid angle inputs, including negative ones.",
		},

		{
			name:        "Scenario 6: Angle Greater Than 2*Pi (5*Pi/2)",
			angle:       5.0 * math.Pi / 2.0,
			expectedSin: 1.0,
			expectedCos: 0.0,

			checkTanInf: 1,
			checkNaN:    false,
			description: "Verify the function's behavior with an angle outside the standard 0 to 2*Pi range (5*Pi/2).",
			validation:  "Trigonometric functions are periodic (2*Pi for sin/cos, Pi for tan). Results for angle 'a' should equal results for 'a mod 2*Pi'. 5*Pi/2 is equivalent to Pi/2.",
			importance:  "Confirms correct handling of angles beyond the principal range, important for cumulative rotation or phase calculations.",
		},

		{
			name:        "Scenario 7: Arbitrary Angle (Pi/6)",
			angle:       math.Pi / 6.0,
			expectedSin: 0.5,
			expectedCos: math.Sqrt(3) / 2.0,
			expectedTan: 1.0 / math.Sqrt(3),
			checkTanInf: 0,
			checkNaN:    false,
			description: "Verify the function's output for a common, non-quadrantal angle like Pi/6 radians (30 degrees).",
			validation:  "Tests a standard, well-known angle value where all results are finite and non-zero. Requires tolerance-based comparison.",
			importance:  "Provides confidence in the function's accuracy for typical, non-edge-case inputs.",
		},

		{
			name:        "Scenario 8: NaN Input",
			angle:       math.NaN(),
			checkNaN:    true,
			description: "Verify the function's behavior when the input angle is Not-a-Number (NaN).",
			validation:  "Standard behavior for math.Sin/Cos/Tan with NaN input is to return NaN. The function should propagate this. Use math.IsNaN() for assertion.",
			importance:  "Important for robustness. Ensures predictable handling of invalid floating-point inputs according to IEEE 754 standards.",
		},

		{
			name:        "Scenario 9: Positive Infinity Input",
			angle:       math.Inf(1),
			checkNaN:    true,
			description: "Verify the function's behavior when the input angle is positive infinity.",
			validation:  "math.Sin/Cos/Tan return NaN for infinite inputs as the functions are undefined at infinity. Use math.IsNaN() for assertion.",
			importance:  "Tests handling of non-finite floating-point values, ensuring predictable behavior with inputs resulting from overflows or divisions by zero.",
		},

		{
			name:        "Scenario 10: Negative Infinity Input",
			angle:       math.Inf(-1),
			checkNaN:    true,
			description: "Verify the function's behavior when the input angle is negative infinity.",
			validation:  "Similar to positive infinity, math.Sin/Cos/Tan return NaN for negative infinity inputs. Use math.IsNaN() for assertion.",
			importance:  "Complements the positive infinity test, ensuring consistent handling of non-finite inputs.",
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

			t.Logf("--- Running Test: %s ---", tc.name)
			t.Logf("Description: %s", tc.description)
			t.Logf("Input Angle: %v", tc.angle)
			t.Logf("Validation Logic: %s", tc.validation)
			t.Logf("Importance: %s", tc.importance)

			sin, cos, tan := SinCosTan(tc.angle)

			if tc.checkNaN {

				if !math.IsNaN(sin) {
					t.Errorf("FAIL: Expected sin to be NaN, but got %v", sin)
				} else {
					t.Logf("Success: sin is NaN as expected.")
				}
				if !math.IsNaN(cos) {
					t.Errorf("FAIL: Expected cos to be NaN, but got %v", cos)
				} else {
					t.Logf("Success: cos is NaN as expected.")
				}
				if !math.IsNaN(tan) {
					t.Errorf("FAIL: Expected tan to be NaN, but got %v", tan)
				} else {
					t.Logf("Success: tan is NaN as expected.")
				}
			} else {

				if !approxEqual(sin, tc.expectedSin, epsilon) {
					t.Errorf("FAIL: Sin mismatch. Expected: ~%v (tolerance %v), Got: %v", tc.expectedSin, epsilon, sin)
				} else {
					t.Logf("Success: Sin (%v) matches expected value ~%v within tolerance %v.", sin, tc.expectedSin, epsilon)
				}

				if !approxEqual(cos, tc.expectedCos, epsilon) {
					t.Errorf("FAIL: Cos mismatch. Expected: ~%v (tolerance %v), Got: %v", tc.expectedCos, epsilon, cos)
				} else {
					t.Logf("Success: Cos (%v) matches expected value ~%v within tolerance %v.", cos, tc.expectedCos, epsilon)
				}

				switch tc.checkTanInf {
				case 1:

					const largePositiveThreshold = 1e16
					if !(tan > largePositiveThreshold) {
						t.Errorf("FAIL: Tan mismatch. Expected: very large positive (> %v), Got: %v", largePositiveThreshold, tan)
					} else {
						t.Logf("Success: Tan (%v) is very large positive (> %v) as expected.", tan, largePositiveThreshold)
					}
				case -1:

					const largeNegativeThreshold = -1e16
					if !(tan < largeNegativeThreshold) {
						t.Errorf("FAIL: Tan mismatch. Expected: very large negative (< %v), Got: %v", largeNegativeThreshold, tan)
					} else {
						t.Logf("Success: Tan (%v) is very large negative (< %v) as expected.", tan, largeNegativeThreshold)
					}
				case 0:
					if !approxEqual(tan, tc.expectedTan, epsilon) {
						t.Errorf("FAIL: Tan mismatch. Expected: ~%v (tolerance %v), Got: %v", tc.expectedTan, epsilon, tan)
					} else {
						t.Logf("Success: Tan (%v) matches expected value ~%v within tolerance %v.", tan, tc.expectedTan, epsilon)
					}
				}
			}
			t.Logf("--- Finished Test: %s ---", tc.name)
		})
	}
}

/*
ROOST_METHOD_HASH=SquareRoot_600b6ad663
ROOST_METHOD_SIG_HASH=SquareRoot_5aa1e1a6d6

FUNCTION_DEF=func SquareRoot(num float64) float64
*/
func TestSquareRoot(t *testing.T) {

	testCases := []struct {
		name             string
		input            float64
		expectedOutput   float64
		tolerance        float64
		expectPanic      bool
		expectedPanicMsg string
		checkNaN         bool
		checkInf         bool
	}{
		{
			name:           "Scenario 1: Test with a positive perfect square",
			input:          9.0,
			expectedOutput: 3.0,
			tolerance:      0.0,
			expectPanic:    false,
		},
		{
			name:           "Scenario 2: Test with a positive non-perfect square",
			input:          2.0,
			expectedOutput: math.Sqrt(2.0),
			tolerance:      1e-9,
			expectPanic:    false,
		},
		{
			name:           "Scenario 3: Test with zero input",
			input:          0.0,
			expectedOutput: 0.0,
			tolerance:      0.0,
			expectPanic:    false,
		},
		{
			name:             "Scenario 4: Test with a negative input (Expect Panic)",
			input:            -4.0,
			expectPanic:      true,
			expectedPanicMsg: "square root of a negative number is not defined",
		},
		{
			name:           "Scenario 5: Test with a very large positive input",
			input:          math.MaxFloat64,
			expectedOutput: math.Sqrt(math.MaxFloat64),
			tolerance:      1e-9,
			expectPanic:    false,
		},
		{
			name:           "Scenario 6: Test with a very small positive input (subnormal)",
			input:          math.SmallestNonzeroFloat64,
			expectedOutput: math.Sqrt(math.SmallestNonzeroFloat64),
			tolerance:      1e-9,
			expectPanic:    false,
		},
		{
			name:        "Scenario 7: Test with Positive Infinity input",
			input:       math.Inf(1),
			expectPanic: false,
			checkInf:    true,
		},
		{
			name:             "Scenario 8: Test with Negative Infinity input (Expect Panic)",
			input:            math.Inf(-1),
			expectPanic:      true,
			expectedPanicMsg: "square root of a negative number is not defined",
		},
		{
			name:        "Scenario 9: Test with NaN input",
			input:       math.NaN(),
			expectPanic: false,
			checkNaN:    true,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil && !tc.expectPanic {

					t.Errorf("Test panicked unexpectedly: %v\n%s", r, string(debug.Stack()))
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input: %v", tc.input)

			if tc.expectPanic {

				defer func() {
					r := recover()
					if r == nil {

						t.Errorf("Expected a panic for input %v, but function did not panic", tc.input)
						return
					}

					panicMsg, ok := r.(string)
					if !ok {
						t.Errorf("Panic occurred, but recovered value was not a string: %v (type: %T)", r, r)
						return
					}
					if panicMsg != tc.expectedPanicMsg {
						t.Errorf("Panic message mismatch: expected %q, got %q", tc.expectedPanicMsg, panicMsg)
						return
					}

					t.Logf("Successfully caught expected panic: %q", panicMsg)
				}()

				SquareRoot(tc.input)

			} else {

				actualOutput := SquareRoot(tc.input)

				if tc.checkNaN {
					if !math.IsNaN(actualOutput) {
						t.Errorf("Expected NaN for input %v, but got %v", tc.input, actualOutput)
					} else {
						t.Logf("Successfully verified NaN for input %v", tc.input)
					}
					return
				}

				if tc.checkInf {

					if !math.IsInf(actualOutput, 1) {
						t.Errorf("Expected Positive Infinity for input %v, but got %v", tc.input, actualOutput)
					} else {
						t.Logf("Successfully verified Positive Infinity for input %v", tc.input)
					}
					return
				}

				if tc.tolerance == 0.0 {

					if actualOutput != tc.expectedOutput {
						t.Errorf("For input %v: expected exactly %v, but got %v", tc.input, tc.expectedOutput, actualOutput)
					} else {
						t.Logf("For input %v: successfully verified exact output %v", tc.input, actualOutput)
					}
				} else {

					diff := math.Abs(actualOutput - tc.expectedOutput)
					if diff >= tc.tolerance {
						t.Errorf("For input %v: expected %v (within tolerance %v), but got %v (difference %v)",
							tc.input, tc.expectedOutput, tc.tolerance, actualOutput, diff)
					} else {
						t.Logf("For input %v: successfully verified output %v is within tolerance %v of expected %v",
							tc.input, actualOutput, tc.tolerance, tc.expectedOutput)
					}
				}
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

		{
			name:     "Scenario 1: Positive larger minus positive smaller",
			num1:     10,
			num2:     3,
			expected: 7,
		},

		{
			name:     "Scenario 2: Positive smaller minus positive larger",
			num1:     5,
			num2:     12,
			expected: -7,
		},

		{
			name:     "Scenario 3: Zero minus positive",
			num1:     0,
			num2:     9,
			expected: -9,
		},

		{
			name:     "Scenario 4: Positive minus zero",
			num1:     15,
			num2:     0,
			expected: 15,
		},

		{
			name:     "Scenario 5: Negative minus negative",
			num1:     -5,
			num2:     -8,
			expected: 3,
		},

		{
			name:     "Scenario 6: Positive minus negative",
			num1:     10,
			num2:     -4,
			expected: 14,
		},

		{
			name:     "Scenario 7: Negative minus positive",
			num1:     -7,
			num2:     3,
			expected: -10,
		},

		{
			name:     "Scenario 8: Number minus itself",
			num1:     42,
			num2:     42,
			expected: 0,
		},

		{
			name:     "Scenario 9: MaxInt minus small positive",
			num1:     math.MaxInt,
			num2:     1,
			expected: math.MaxInt - 1,
		},

		{
			name:     "Scenario 10: MinInt minus positive (underflow wrap)",
			num1:     math.MinInt,
			num2:     1,
			expected: math.MaxInt,
		},

		{
			name:     "Scenario 11: MaxInt minus negative (overflow wrap)",
			num1:     math.MaxInt,
			num2:     -1,
			expected: math.MinInt,
		},

		{
			name:     "Scenario 12: Zero minus MinInt",
			num1:     0,
			num2:     math.MinInt,
			expected: math.MinInt,
		},

		{
			name:     "Scenario 13: MinInt minus MaxInt (extreme underflow wrap)",
			num1:     math.MinInt,
			num2:     math.MaxInt,
			expected: 1,
		},

		{
			name:     "Scenario 14: MaxInt minus MinInt (extreme overflow wrap)",
			num1:     math.MaxInt,
			num2:     math.MinInt,
			expected: -1,
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

			t.Logf("Running test: %s", tc.name)
			t.Logf("Arrange: num1 = %d, num2 = %d", tc.num1, tc.num2)
			t.Logf("Arrange: expected = %d", tc.expected)

			actual := Subtract(tc.num1, tc.num2)
			t.Logf("Act: Subtract(%d, %d) returned %d", tc.num1, tc.num2, actual)

			if actual != tc.expected {

				t.Errorf("Assert: FAILED - For inputs num1=%d, num2=%d: expected %d, but got %d", tc.num1, tc.num2, tc.expected, actual)
			} else {

				t.Logf("Assert: PASSED - Result %d matches expected %d", actual, tc.expected)
			}
		})
	}
}
