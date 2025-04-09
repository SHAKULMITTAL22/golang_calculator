package calc

import (
	fmt "fmt"
	math "math"
	os "os"
	testing "testing"
	debug "runtime/debug"
	strings "strings"
	bytes "bytes"
)








/*
ROOST_METHOD_HASH=Absolute_f8af7505a1
ROOST_METHOD_SIG_HASH=Absolute_4bad226818

FUNCTION_DEF=func Absolute(num float64) float64 

*/
func TestAbsolute(t *testing.T) {

	type testCase struct {
		name     string
		input    float64
		expected float64
	}

	tests := []testCase{
		{name: "Positive Number Input", input: 5.0, expected: 5.0},
		{name: "Negative Number Input", input: -8.3, expected: 8.3},
		{name: "Zero Input", input: 0.0, expected: 0.0},
		{name: "Large Positive Number", input: 1e+18, expected: 1e+18},
		{name: "Large Negative Number", input: -2e+18, expected: 2e+18},
		{name: "Small Decimal Values", input: -0.0003, expected: 0.0003},
		{name: "Positive Infinity", input: math.Inf(1), expected: math.Inf(1)},
		{name: "Negative Infinity", input: math.Inf(-1), expected: math.Inf(1)},
		{name: "NaN Input", input: math.NaN(), expected: math.NaN()},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			actual := Absolute(tc.input)

			w.Close()
			os.Stdout = oldStdout

			if math.IsNaN(tc.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("FAILED %s: Expected NaN but got %v", tc.name, actual)
				} else {
					t.Logf("PASSED %s: NaN correctly handled.", tc.name)
				}
			} else if actual != tc.expected {
				t.Errorf("FAILED %s: Expected %v but got %v", tc.name, tc.expected, actual)
			} else {
				t.Logf("PASSED %s: Input %v produced expected output %v", tc.name, tc.input, actual)
			}

			t.Logf("Test case '%s' executed with input: %v, expected: %v, actual: %v", tc.name, tc.input, tc.expected, actual)
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
		name        string
		num1        int
		num2        int
		expected    int
		expectPanic bool
		explanation string
	}{
		{
			name:        "Adding positive integers",
			num1:        3,
			num2:        5,
			expected:    8,
			expectPanic: false,
			explanation: "Adding 3 and 5 should return 8 demonstrating normal addition of positive integers",
		},
		{
			name:        "Adding negative integers",
			num1:        -3,
			num2:        -5,
			expected:    -8,
			expectPanic: false,
			explanation: "Adding -3 and -5 should return -8 as addition of negatives must follow standard arithmetic",
		},
		{
			name:        "Adding positive and negative integers",
			num1:        5,
			num2:        -3,
			expected:    2,
			expectPanic: false,
			explanation: "Adding 5 and -3 should return 2 demonstrating subtraction-like results",
		},
		{
			name:        "Adding zero to an integer",
			num1:        10,
			num2:        0,
			expected:    10,
			expectPanic: false,
			explanation: "Adding 0 to 10 should return 10, validating identity property",
		},
		{
			name:        "Adding zero to zero",
			num1:        0,
			num2:        0,
			expected:    0,
			expectPanic: false,
			explanation: "Adding 0 to 0 must return 0 as addition with zeros should preserve identity",
		},
		{
			name:        "Adding largest positive integers (overflow check)",
			num1:        math.MaxInt,
			num2:        math.MaxInt - 1,
			expected:    math.MaxInt + (math.MaxInt - 1),
			expectPanic: false,
			explanation: "It checks for integer overflow behavior when summing massive integers",
		},
		{
			name:        "Adding smallest negative integers (negative overflow check)",
			num1:        math.MinInt,
			num2:        math.MinInt + 1,
			expected:    math.MinInt + (math.MinInt + 1),
			expectPanic: false,
			explanation: "It examines the addition of extreme negative values and potential overflow conditions",
		},
		{
			name:        "Adding two equal integers",
			num1:        4,
			num2:        4,
			expected:    8,
			expectPanic: false,
			explanation: "Adding identical integers tests symmetry and basic correctness logic",
		},
		{
			name:        "Adding integers to reach zero (inverse property)",
			num1:        7,
			num2:        -7,
			expected:    0,
			expectPanic: false,
			explanation: "Adding an integer and its negation should yield 0, verifying the inverse property",
		},
		{
			name:        "Adding largest positive and negative integers (balance check)",
			num1:        math.MaxInt,
			num2:        -math.MaxInt,
			expected:    0,
			expectPanic: false,
			explanation: "Adding maximum positive integer and its negative counterpart should result in zero",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test %q: %v\n%s", tt.name, r, string(debug.Stack()))
					if !tt.expectPanic {
						t.FailNow()
					}
				}
			}()

			actual := Add(tt.num1, tt.num2)

			if actual != tt.expected {
				t.Errorf("Test: %q - FAILED\nExpected result: %d; Actual result: %d\nReason: %s", tt.name, tt.expected, actual, tt.explanation)
			} else {
				t.Logf("Test: %q - PASSED\nExpected result: %d; Actual result: %d\nReason: %s", tt.name, tt.expected, actual, tt.explanation)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Divide_f2ddee767d
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64 

*/
func TestDivide(t *testing.T) {
	type testCase struct {
		description string
		num1        float64
		num2        float64
		expected    float64
		expectPanic bool
	}

	tests := []testCase{

		{
			description: "Divide two positive numbers",
			num1:        10.0,
			num2:        2.0,
			expected:    5.0,
			expectPanic: false,
		},

		{
			description: "Divide two negative numbers",
			num1:        -10.0,
			num2:        -2.0,
			expected:    5.0,
			expectPanic: false,
		},

		{
			description: "Divide a positive number by a negative number",
			num1:        10.0,
			num2:        -2.0,
			expected:    -5.0,
			expectPanic: false,
		},

		{
			description: "Divide zero by a positive number",
			num1:        0.0,
			num2:        5.0,
			expected:    0.0,
			expectPanic: false,
		},

		{
			description: "Divide zero by a negative number",
			num1:        0.0,
			num2:        -5.0,
			expected:    0.0,
			expectPanic: false,
		},

		{
			description: "Divide a number by 1",
			num1:        15.0,
			num2:        1.0,
			expected:    15.0,
			expectPanic: false,
		},

		{
			description: "Divide a number by itself",
			num1:        12.34,
			num2:        12.34,
			expected:    1.0,
			expectPanic: false,
		},

		{
			description: "Divide by zero",
			num1:        10.0,
			num2:        0.0,
			expected:    0.0,
			expectPanic: true,
		},

		{
			description: "Division resulting in infinity",
			num1:        math.MaxFloat64,
			num2:        math.SmallestNonzeroFloat64,
			expected:    math.Inf(1),
			expectPanic: false,
		},

		{
			description: "Division resulting in very small fractions",
			num1:        math.SmallestNonzeroFloat64,
			num2:        math.MaxFloat64,
			expected:    0.0,
			expectPanic: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if test.expectPanic {
						t.Logf("Panic encountered as expected. Panic message: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			if !test.expectPanic {

				stdoutBackup := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				defer func() { os.Stdout = stdoutBackup }()

				actual := Divide(test.num1, test.num2)

				fmt.Fprintf(w, "Result: %f\n", actual)
				w.Close()
				var output string
				fmt.Fscanf(r, "%s", &output)

				if actual != test.expected {
					t.Errorf("Failed %s. Expected: %f, Got: %f", test.description, test.expected, actual)
				} else {
					t.Logf("Passed %s. Expected: %f, Got: %f", test.description, test.expected, actual)
				}
			} else {

				Divide(test.num1, test.num2)
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
		input       int
		expected    int
		expectPanic bool
		panicMsg    string
	}

	testData := []testCase{
		{input: 0, expected: 1, expectPanic: false, panicMsg: ""},
		{input: 1, expected: 1, expectPanic: false, panicMsg: ""},
		{input: 5, expected: 120, expectPanic: false, panicMsg: ""},
		{input: 10, expected: 3628800, expectPanic: false, panicMsg: ""},
		{input: -5, expected: 0, expectPanic: true, panicMsg: "factorial is not defined for negative numbers"},
		{input: 20, expected: 2432902008176640000, expectPanic: false, panicMsg: ""},
		{input: 4, expected: 24, expectPanic: false, panicMsg: ""},
		{input: 7, expected: 5040, expectPanic: false, panicMsg: ""},
		{input: math.MinInt, expected: 0, expectPanic: true, panicMsg: "factorial is not defined for negative numbers"},
	}

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		w.Close()
		os.Stdout = stdOut
	}()

	for _, tt := range testData {
		t.Run(fmt.Sprintf("Input_%d", tt.input), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectPanic {
						if r == tt.panicMsg {
							t.Logf("Successfully detected expected panic for input %d: %v", tt.input, r)
						} else {
							t.Errorf("Unexpected panic message for input %d. Got: %v, Expected: %v", tt.input, r, tt.panicMsg)
						}
					} else {
						t.Logf("Unexpected panic encountered: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			var result int
			if !tt.expectPanic {
				result = Factorial(tt.input)
			}

			if !tt.expectPanic {
				if result == tt.expected {
					t.Logf("Test passed for input %d: Result = %d, Expected = %d", tt.input, result, tt.expected)
				} else {
					t.Errorf("Test failed for input %d: Result = %d, Expected = %d", tt.input, result, tt.expected)
				}
			}
		})
	}

	w.Close()
	var output string
	fmt.Fscanf(r, "%s", &output)
	t.Logf("Captured output: \n%s", output)
}


/*
ROOST_METHOD_HASH=GCD_6cf0472095
ROOST_METHOD_SIG_HASH=GCD_ab1c91475d

FUNCTION_DEF=func GCD(a, b int) int 

*/
func TestGcd(t *testing.T) {
	type testData struct {
		a, b   int
		expect int
		desc   string
	}

	tests := []testData{
		{36, 24, 12, "Scenario 1: Test GCD for common divisor"},
		{15, 45, 15, "Scenario 2: Test GCD for multiple relationship"},
		{13, 27, 1, "Scenario 3: Test GCD for coprime integers"},
		{0, 29, 29, "Scenario 4: Test GCD when one integer is zero"},
		{0, 0, 0, "Scenario 5: Test GCD when both integers are zero"},
		{-36, -48, 12, "Scenario 6: Test GCD for negative integers"},
		{42, -56, 14, "Scenario 7: Test GCD for positive-negative integers"},
		{20, 20, 20, "Scenario 8: Test GCD for equal integers"},
		{1_000_000, 5_000_000, 1_000_000, "Scenario 9: Test GCD for large integers"},
		{math.MaxInt, 123, 3, "Scenario 10: Test GCD with math.MaxInt edge case"},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		w.Close()
		os.Stdout = old
	}()

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered while testing %q. Reason: %v\nTrace:\n%s", test.desc, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			t.Logf("Running test: %s", test.desc)
			result := GCD(test.a, test.b)

			w.Close()
			out, _ := os.ReadFile(r.Name())
			fmt.Fprintf(os.Stderr, "Captured Output: %s\n", string(out))

			if result != test.expect {
				t.Errorf("Test Failed for %s: Input (a=%d, b=%d), Expected GCD=%d, Got GCD=%d", test.desc, test.a, test.b, test.expect, result)
			} else {
				t.Logf("Test Passed for %s: Input (a=%d, b=%d), Expected GCD=%d, Got GCD=%d", test.desc, test.a, test.b, test.expect, result)
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

	type testCase struct {
		name     string
		a        int
		b        int
		expected int
	}

	testCases := []testCase{
		{"Basic functionality with two positive integers", 4, 6, 12},
		{"One input is zero", 0, 6, 0},
		{"Both inputs are zero", 0, 0, 0},
		{"Negative integers as inputs", -4, -6, 12},
		{"One positive and one negative integer input", -4, 6, 12},
		{"Large integers as inputs", math.MaxInt32, 2, math.MaxInt32 * 2 / GCD(math.MaxInt32, 2)},
		{"LCM of two prime numbers", 7, 13, 91},
		{"LCM of two identical numbers", 8, 8, 8},
		{"LCM of one and any number", 1, 25, 25},
		{"LCM of coprime numbers", 12, 35, 420},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test '%s'. Reason: %v\nStack trace:\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := LCM(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("Test '%s' failed. Expected: %d, Got: %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test '%s' passed. Output matched expected value: %d", tc.name, tc.expected)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Logarithm_546f6d96c4
ROOST_METHOD_SIG_HASH=Logarithm_ddbb699678

FUNCTION_DEF=func Logarithm(num, base float64) float64 

*/
func TestLogarithm(t *testing.T) {

	tests := []struct {
		name        string
		num         float64
		base        float64
		expect      interface{}
		expectPanic bool
	}{
		{
			name:   "Valid logarithm with base 10",
			num:    100,
			base:   10,
			expect: math.Log10(100),
		},
		{
			name:   "Valid logarithm with base 2",
			num:    8,
			base:   2,
			expect: 3.0,
		},
		{
			name:   "Logarithm with base e (natural logarithm)",
			num:    math.Exp(2),
			base:   math.E,
			expect: 2.0,
		},
		{
			name:   "Logarithm with fractional base",
			num:    4,
			base:   0.5,
			expect: math.Log(4) / math.Log(0.5),
		},
		{
			name:        "Invalid logarithm with negative number",
			num:         -5,
			base:        10,
			expectPanic: true,
		},
		{
			name:        "Invalid logarithm with zero for num",
			num:         0,
			base:        10,
			expectPanic: true,
		},
		{
			name:        "Invalid logarithm with base equal to 1",
			num:         10,
			base:        1,
			expectPanic: true,
		},
		{
			name:        "Invalid logarithm with base as zero",
			num:         10,
			base:        0,
			expectPanic: true,
		},
		{
			name:   "Logarithm where num and base are the same",
			num:    7,
			base:   7,
			expect: 1.0,
		},
		{
			name:   "Logarithm with num of 1",
			num:    1,
			base:   5,
			expect: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Logf("Panic encountered unexpectedly. %v\n%s", r, string(debug.Stack()))
						t.Fail()
					} else {
						t.Logf("Panic expected and captured successfully. %v", r)
					}
				}
			}()

			r, w, _ := os.Pipe()
			os.Stdout = w

			var result interface{}
			if !tt.expectPanic {
				result = Logarithm(tt.num, tt.base)
			}

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = os.Stdout
			t.Log(string(out))

			if !tt.expectPanic {
				if math.Abs(result.(float64)-tt.expect.(float64)) > 1e-9 {
					t.Errorf("Expected %v, got %v", tt.expect, result)
				} else {
					t.Logf("Success for test %s. Result: %v", tt.name, result)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Modulo_eb9c4baeed
ROOST_METHOD_SIG_HASH=Modulo_09898f6fed

FUNCTION_DEF=func Modulo(num1, num2 int) int 

*/
func TestModulo(t *testing.T) {

	type ModuloTestCase struct {
		Name           string
		Num1           int
		Num2           int
		ExpectedResult int
		ExpectPanic    bool
	}

	testCases := []ModuloTestCase{
		{
			Name:           "Positive Number Modulo Positive Number",
			Num1:           10,
			Num2:           3,
			ExpectedResult: 1,
			ExpectPanic:    false,
		},
		{
			Name:           "Division by Zero",
			Num1:           5,
			Num2:           0,
			ExpectedResult: 0,
			ExpectPanic:    true,
		},
		{
			Name:           "Negative Number Modulo Negative Number",
			Num1:           -10,
			Num2:           -3,
			ExpectedResult: -1,
			ExpectPanic:    false,
		},
		{
			Name:           "Positive Dividend Modulo Negative Divisor",
			Num1:           10,
			Num2:           -3,
			ExpectedResult: 1,
			ExpectPanic:    false,
		},
		{
			Name:           "Negative Dividend Modulo Positive Divisor",
			Num1:           -10,
			Num2:           3,
			ExpectedResult: -1,
			ExpectPanic:    false,
		},
		{
			Name:           "Zero Dividend",
			Num1:           0,
			Num2:           7,
			ExpectedResult: 0,
			ExpectPanic:    false,
		},
		{
			Name:           "Large Number Stress Test",
			Num1:           math.MaxInt32,
			Num2:           3,
			ExpectedResult: 1,
			ExpectPanic:    false,
		},
		{
			Name:           "Identical Numbers",
			Num1:           7,
			Num2:           7,
			ExpectedResult: 0,
			ExpectPanic:    false,
		},
		{
			Name:           "Negative Divisor Larger Than Negative Dividend",
			Num1:           -3,
			Num2:           -10,
			ExpectedResult: -3,
			ExpectPanic:    false,
		},
		{
			Name:           "Large Negative Dividend",
			Num1:           math.MinInt32,
			Num2:           7,
			ExpectedResult: -2,
			ExpectPanic:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if tc.ExpectPanic {
						t.Logf("Expected panic occurred: %v", r)
					} else {
						t.Logf("Unexpected panic occurred: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			var result int
			if !tc.ExpectPanic {
				result = Modulo(tc.Num1, tc.Num2)
			} else {

				result = Modulo(tc.Num1, tc.Num2)
			}

			if result != tc.ExpectedResult {
				t.Errorf("Test '%s' failed. Expected result: %v, Got: %v", tc.Name, tc.ExpectedResult, result)
				return
			}

			t.Logf("Test '%s' passed. Expected result: %v", tc.Name, tc.ExpectedResult)
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
		name     string
		num1     float64
		num2     float64
		expected float64
		isNaN    bool
	}

	testCases := []testCase{
		{
			name:     "Scenario 1: Multiplication of Two Positive Numbers",
			num1:     5.0,
			num2:     3.0,
			expected: 15.0,
		},
		{
			name:     "Scenario 2: Multiplication of Two Negative Numbers",
			num1:     -4.0,
			num2:     -2.0,
			expected: 8.0,
		},
		{
			name:     "Scenario 3: Multiplication of a Positive and a Negative Number",
			num1:     6.0,
			num2:     -3.0,
			expected: -18.0,
		},
		{
			name:     "Scenario 4: Multiplication by Zero",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		{
			name:     "Scenario 5: Multiplication of a Decimal Number and an Integer",
			num1:     2.5,
			num2:     4.0,
			expected: 10.0,
		},
		{
			name:     "Scenario 6: Multiplication of Two Large Numbers",
			num1:     1e6,
			num2:     1e6,
			expected: 1e12,
		},
		{
			name:     "Scenario 7: Multiplication of Two Small Numbers (With Floating-Point Precision)",
			num1:     0.0001,
			num2:     0.0002,
			expected: 0.00000002,
		},
		{
			name:     "Scenario 8: Multiplication of Infinite Values (Positive Infinity)",
			num1:     math.Inf(1),
			num2:     3.0,
			expected: math.Inf(1),
		},
		{
			name:     "Scenario 9: Multiplication of Infinite Values (Negative Infinity)",
			num1:     math.Inf(-1),
			num2:     3.0,
			expected: math.Inf(-1),
		},
		{
			name:  "Scenario 10: Multiplication of NaN Value",
			num1:  math.NaN(),
			num2:  5.0,
			isNaN: true,
		},
		{
			name:     "Scenario 11: Multiplication with Extreme Positive Number and Negative Numbers",
			num1:     math.MaxFloat64,
			num2:     -math.MaxFloat64,
			expected: math.Inf(-1),
		},
		{
			name:     "Scenario 12: Multiplication of Two Identical Numbers (Squaring)",
			num1:     7.0,
			num2:     7.0,
			expected: 49.0,
		},
		{
			name:     "Scenario 13: Multiplication with Negative Zero",
			num1:     -0.0,
			num2:     5.0,
			expected: -0.0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Multiply(tc.num1, tc.num2)

			if tc.isNaN {
				if !math.IsNaN(result) {
					t.Errorf("Test %s failed; expected NaN but got %v", tc.name, result)
				} else {
					t.Logf("Test %s passed; result is NaN as expected", tc.name)
				}
			} else if result != tc.expected && !(math.IsInf(tc.expected, 0) && math.IsInf(result, 0)) {
				t.Errorf("Test %s failed; expected %v but got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed; got expected result %v", tc.name, result)
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
	t.Parallel()

	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		_ = w.Close()
		os.Stdout = origStdout
	}()

	type testCase struct {
		base        float64
		exponent    float64
		expected    float64
		description string
	}

	tests := []testCase{
		{base: 2.0, exponent: 3.0, expected: 8.0, description: "Positive Base and Positive Exponent"},
		{base: 5.0, exponent: 0.0, expected: 1.0, description: "Positive Base and Exponent Zero"},
		{base: 0.0, exponent: 3.0, expected: 0.0, description: "Base Zero and Positive Exponent"},
		{base: 0.0, exponent: 0.0, expected: 1.0, description: "Base Zero and Exponent Zero"},
		{base: -2.0, exponent: 3.0, expected: -8.0, description: "Negative Base and Odd Positive Exponent"},
		{base: -2.0, exponent: 2.0, expected: 4.0, description: "Negative Base and Even Positive Exponent"},
		{base: -2.0, exponent: -2.0, expected: 0.25, description: "Negative Base and Negative Exponent"},
		{base: 2.0, exponent: -3.0, expected: 0.125, description: "Positive Base and Negative Exponent"},
		{base: 0.5, exponent: 2.0, expected: 0.25, description: "Fractional Base and Positive Exponent"},
		{base: 0.5, exponent: -2.0, expected: 4.0, description: "Fractional Base and Negative Exponent"},
		{base: 10.0, exponent: 10.0, expected: 10000000000.0, description: "Large Base and Large Exponent"},
		{base: 9.0, exponent: 0.5, expected: 3.0, description: "Exponent as Fractional Value"},
		{base: 9.0, exponent: -0.5, expected: 1.0 / 3.0, description: "Negative Fractional Exponents"},
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

			actual := Power(tc.base, tc.exponent)

			if math.Abs(actual-tc.expected) > 1e-9 {
				t.Errorf("Failed %s: expected %.4f but got %.4f", tc.description, tc.expected, actual)
			} else {
				t.Logf("Passed %s: expected %.4f, got %.4f", tc.description, tc.expected, actual)
			}
		})
	}

	_ = w.Close()
	var output []byte
	if _, err := fmt.Fscanf(r, "%s\n", &output); err == nil {
		t.Logf("Output captured: %s", string(output))
	} else {
		t.Logf("No output captured: %v", err)
	}
}


/*
ROOST_METHOD_HASH=SinCosTan_c6521a7850
ROOST_METHOD_SIG_HASH=SinCosTan_6ec04d6e93

FUNCTION_DEF=func SinCosTan(angle float64) (sin, cos, tan float64) 

*/
func TestSinCosTan(t *testing.T) {
	type TestCase struct {
		Angle       float64
		ExpectedSin float64
		ExpectedCos float64
		ExpectedTan float64
		Description string
		ExpectPanic bool
	}

	testCases := []TestCase{
		{
			Angle:       0,
			ExpectedSin: 0,
			ExpectedCos: 1,
			ExpectedTan: 0,
			Description: "Validate Outputs for an Angle of 0 Radians",
		},
		{
			Angle:       math.Pi / 2,
			ExpectedSin: 1,
			ExpectedCos: 0,
			ExpectedTan: 0,
			Description: "Validate Outputs for π/2 Radians (90 Degrees)",
			ExpectPanic: true,
		},
		{
			Angle:       math.Pi,
			ExpectedSin: 0,
			ExpectedCos: -1,
			ExpectedTan: 0,
			Description: "Validate Outputs for π Radians (180 Degrees)",
		},
		{
			Angle:       -math.Pi / 4,
			ExpectedSin: -1 / math.Sqrt(2),
			ExpectedCos: 1 / math.Sqrt(2),
			ExpectedTan: -1,
			Description: "Validate Outputs for -π/4 Radians (-45 Degrees)",
		},
		{
			Angle:       math.Pi / 6,
			ExpectedSin: math.Sin(math.Pi / 6),
			ExpectedCos: math.Cos(math.Pi / 6),
			ExpectedTan: math.Tan(math.Pi / 6),
			Description: "Validate Periodicity of Sine and Cosine",
		},
		{
			Angle:       1e5,
			ExpectedSin: math.Sin(1e5),
			ExpectedCos: math.Cos(1e5),
			ExpectedTan: math.Tan(1e5),
			Description: "Verify Behavior at Very Large Positive Angles",
		},
		{
			Angle:       -1e-7,
			ExpectedSin: -1e-7,
			ExpectedCos: 1,
			ExpectedTan: -1e-7,
			Description: "Verify Behavior at Very Small Negative Angles",
		},
	}

	assertApproxEqual := func(t *testing.T, actual, expected float64, tolerance float64, valueName string) {
		if math.Abs(actual-expected) > tolerance {
			t.Errorf("%s mismatch: got %.12f, expected %.12f (tolerance: %.12f)", valueName, actual, expected, tolerance)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered while testing %v.\n%v\n%s", tc.Description, r, string(debug.Stack()))
					if tc.ExpectPanic {
						t.Log("Test expected panic and succeeded.")
						return
					}
					t.Fail()
				}
			}()

			sin, cos, tan := SinCosTan(tc.Angle)

			t.Logf("Testing: %v", tc.Description)
			assertApproxEqual(t, sin, tc.ExpectedSin, 1e-12, "Sin")
			assertApproxEqual(t, cos, tc.ExpectedCos, 1e-12, "Cos")

			if tc.ExpectPanic {
				t.Log("Skipping tan validation as it tends to infinity for π/2.")
			} else {
				assertApproxEqual(t, tan, tc.ExpectedTan, 1e-12, "Tan")
			}
		})
	}
}


/*
ROOST_METHOD_HASH=SquareRoot_600b6ad663
ROOST_METHOD_SIG_HASH=SquareRoot_5aa1e1a6d6

FUNCTION_DEF=func SquareRoot(num float64) float64 

*/
func TestSquareRoot(t *testing.T) {

	tests := []struct {
		name   string
		input  float64
		expect interface{}
	}{
		{
			name:   "TestPositiveNumber",
			input:  4.0,
			expect: 2.0,
		},
		{
			name:   "TestZeroInput",
			input:  0.0,
			expect: 0.0,
		},
		{
			name:   "TestNegativeInputWithPanic",
			input:  -5.0,
			expect: "square root of a negative number is not defined",
		},
		{
			name:   "TestLargePositiveNumber",
			input:  1e12,
			expect: math.Sqrt(1e12),
		},
		{
			name:   "TestSmallFraction",
			input:  0.0004,
			expect: 0.02,
		},
		{
			name:   "TestPrecisionOfComputation",
			input:  2.0,
			expect: math.Sqrt(2.0),
		},
		{
			name:   "TestFloat64Max",
			input:  math.MaxFloat64,
			expect: math.Sqrt(math.MaxFloat64),
		},
	}

	t.Log("Starting tests for SquareRoot function.")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if expectedMessage, ok := tt.expect.(string); ok {
						t.Logf("Panic encountered with message: %v", r)
						t.Logf("\nStack Trace:\n%s", string(debug.Stack()))
						if !strings.Contains(fmt.Sprintf("%v", r), expectedMessage) {
							t.Errorf("Expected panic message: %v, but got: %v", expectedMessage, r)
						}
					} else {
						t.Errorf("Unexpected panic with message: %v", r)
					}
					t.FailNow()
				}
			}()

			t.Logf("Testing scenario: %s", tt.name)
			result := SquareRoot(tt.input)

			if expectedResult, ok := tt.expect.(float64); ok {
				epsilon := 1e-9
				if math.Abs(result-expectedResult) > epsilon {
					t.Errorf("Failed test %s: Result %.12f differs from expected %.12f by more than %.12f.", tt.name, result, expectedResult, epsilon)
				} else {
					t.Logf("Passed test %s: Result %.12f matches expected %.12f (within %.12f tolerance).", tt.name, result, expectedResult, epsilon)
				}
			}
		})
	}

	t.Log("All tests completed.")
}


/*
ROOST_METHOD_HASH=Subtract_559013d27f
ROOST_METHOD_SIG_HASH=Subtract_29b74c09c9

FUNCTION_DEF=func Subtract(num1, num2 int) int 

*/
func TestSubtract(t *testing.T) {

	type testCase struct {
		name           string
		num1           int
		num2           int
		expectedResult int
		expectedLog    string
	}

	testCases := []testCase{
		{
			name:           "Subtracting smaller integer from larger integer",
			num1:           10,
			num2:           3,
			expectedResult: 7,
			expectedLog:    "Result matches expected value for straightforward subtraction",
		},
		{
			name:           "Subtracting larger integer from smaller integer",
			num1:           3,
			num2:           10,
			expectedResult: -7,
			expectedLog:    "Function handles negative results correctly",
		},
		{
			name:           "Subtracting zero from an integer",
			num1:           5,
			num2:           0,
			expectedResult: 5,
			expectedLog:    "Respects mathematical principles for subtraction involving zero",
		},
		{
			name:           "Subtracting an integer from itself",
			num1:           8,
			num2:           8,
			expectedResult: 0,
			expectedLog:    "Identical values subtracted yield zero",
		},
		{
			name:           "Subtracting two negative integers",
			num1:           -5,
			num2:           -3,
			expectedResult: -2,
			expectedLog:    "Correctly calculates subtraction of negative integers",
		},
		{
			name:           "Subtracting a positive integer from a negative integer",
			num1:           -7,
			num2:           5,
			expectedResult: -12,
			expectedLog:    "Handles mixed-sign inputs correctly",
		},
		{
			name:           "Subtracting a negative integer from a positive integer",
			num1:           10,
			num2:           -5,
			expectedResult: 15,
			expectedLog:    "Correctly interprets subtraction involving double negatives",
		},
		{
			name:           "Subtracting with minimum integer values",
			num1:           math.MinInt32,
			num2:           -1,
			expectedResult: math.MinInt32 + 1,
			expectedLog:    "Properly handles edge case with minimum integer values",
		},
		{
			name:           "Subtracting with maximum integer values",
			num1:           math.MaxInt32,
			num2:           1,
			expectedResult: math.MaxInt32 - 1,
			expectedLog:    "Properly handles edge case with maximum integer values",
		},
		{
			name:           "Subtraction resulting in zero",
			num1:           0,
			num2:           0,
			expectedResult: 0,
			expectedLog:    "Respects subtraction property when input is zero",
		},
		{
			name:           "Subtraction involving large range values near overflow",
			num1:           math.MaxInt32,
			num2:           math.MinInt32,
			expectedResult: math.MaxInt32 - math.MinInt32,
			expectedLog:    "Handles edge case with very large values",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic occurred during test execution: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var buffer bytes.Buffer
			temp := os.Stdout
			os.Stdout = &buffer
			defer func() { os.Stdout = temp }()

			result := Subtract(tc.num1, tc.num2)

			if result != tc.expectedResult {
				t.Errorf("Expected result: %d, Got: %d", tc.expectedResult, result)
			}

			t.Logf(tc.expectedLog)
		})
	}
}

