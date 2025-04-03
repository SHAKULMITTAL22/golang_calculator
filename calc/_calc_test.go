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


/*
ROOST_METHOD_HASH=LCM_6035446662
ROOST_METHOD_SIG_HASH=LCM_121c872fbf

FUNCTION_DEF=func LCM(a, b int) int 

*/
func TestLcm(t *testing.T) {
	testCases := []struct {
		description string
		a           int
		b           int
		expected    int
	}{
		{
			description: "Scenario 1: Calculate LCM of two positive integers",
			a:           12,
			b:           15,
			expected:    60,
		},
		{
			description: "Scenario 2: Calculate LCM when one input is zero",
			a:           0,
			b:           15,
			expected:    0,
		},
		{
			description: "Scenario 3: Calculate LCM when both inputs are zero",
			a:           0,
			b:           0,
			expected:    0,
		},
		{
			description: "Scenario 4: Calculate LCM for two equal numbers",
			a:           15,
			b:           15,
			expected:    15,
		},
		{
			description: "Scenario 5: Calculate LCM when one number is a multiple of the other",
			a:           6,
			b:           18,
			expected:    18,
		},
		{
			description: "Scenario 6: Calculate LCM for large integers",
			a:           123456,
			b:           789012,
			expected:    8117355456,
		},
		{
			description: "Scenario 7: Calculate LCM for negative integers",
			a:           -12,
			b:           -15,
			expected:    60,
		},
		{
			description: "Scenario 8: Calculate LCM for one positive and one negative integer",
			a:           -12,
			b:           15,
			expected:    60,
		},
		{
			description: "Scenario 9: Evaluate LCM when inputs are prime numbers",
			a:           7,
			b:           13,
			expected:    91,
		},
		{
			description: "Scenario 10: Evaluate LCM when one input is 1",
			a:           1,
			b:           50,
			expected:    50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			actual := LCM(tc.a, tc.b)

			fmt.Fprintf(w, "LCM(%d, %d) = %d\n", tc.a, tc.b, actual)

			w.Close()
			os.Stdout = oldStdout

			if actual != tc.expected {
				t.Errorf("Failed %s: expected %d but got %d", tc.description, tc.expected, actual)
			} else {
				t.Logf("Success: %s. %d matches expected %d.", tc.description, actual, tc.expected)
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
		expected    float64
		expectPanic bool
	}{

		{"Positive num and base", 100, 10, math.Log(100) / math.Log(10), false},

		{"Negative num", -5, 2, 0, true},

		{"Negative base", 10, -2, 0, true},

		{"Zero num", 0, 5, 0, true},

		{"Zero base", 10, 0, 0, true},

		{"Base equals 1", 10, 1, 0, true},

		{"Non-integer base", 100, 2.5, math.Log(100) / math.Log(2.5), false},

		{"Num equals base", 10, 10, 1, false},

		{"Large num", 1e10, 10, math.Log(1e10) / math.Log(10), false},

		{"Small num", 1e-10, 10, math.Log(1e-10) / math.Log(10), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectPanic {
						t.Logf("Panic expected: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic for %s with num=%v, base=%v\n%s", tt.name, tt.num, tt.base, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			var buf bytes.Buffer
			output := os.Stdout
			os.Stdout = &buf
			defer func() { os.Stdout = output }()

			var actual float64
			if !tt.expectPanic {
				actual = Logarithm(tt.num, tt.base)
				if math.Abs(actual-tt.expected) > 1e-9 {
					t.Logf("Test %s failed: expected=%v, actual=%v", tt.name, tt.expected, actual)
					t.Fail()
				} else {
					t.Logf("Test %s succeeded: expected=%v, actual=%v", tt.name, tt.expected, actual)
				}
			} else {

				Logarithm(tt.num, tt.base)
			}
		})
	}

	defer func() {
		if buf.Len() > 0 {
			fmt.Fprintln(output, buf.String())
		}
	}()
}


/*
ROOST_METHOD_HASH=Modulo_eb9c4baeed
ROOST_METHOD_SIG_HASH=Modulo_09898f6fed

FUNCTION_DEF=func Modulo(num1, num2 int) int 

*/
func TestModulo(t *testing.T) {
	tests := []struct {
		name           string
		num1           int
		num2           int
		expectedResult int
		expectPanic    bool
	}{
		{
			name:           "Scenario 1: Modulo operation with positive integers",
			num1:           10,
			num2:           3,
			expectedResult: 1,
			expectPanic:    false,
		},
		{
			name:           "Scenario 2: Modulo operation with num1 equal to num2",
			num1:           5,
			num2:           5,
			expectedResult: 0,
			expectPanic:    false,
		},
		{
			name:           "Scenario 3: Modulo operation with num1 less than num2",
			num1:           3,
			num2:           10,
			expectedResult: 3,
			expectPanic:    false,
		},
		{
			name:        "Scenario 4: Modulo operation with zero divisor",
			num1:        10,
			num2:        0,
			expectPanic: true,
		},
		{
			name:           "Scenario 5: Modulo operation with zero dividend",
			num1:           0,
			num2:           5,
			expectedResult: 0,
			expectPanic:    false,
		},
		{
			name:           "Scenario 6: Modulo operation with negative dividend",
			num1:           -10,
			num2:           3,
			expectedResult: -1,
			expectPanic:    false,
		},
		{
			name:           "Scenario 7: Modulo operation with negative divisor",
			num1:           10,
			num2:           -3,
			expectedResult: 1,
			expectPanic:    false,
		},
		{
			name:           "Scenario 8: Modulo operation with both negative numbers",
			num1:           -10,
			num2:           -3,
			expectedResult: -1,
			expectPanic:    false,
		},
		{
			name:           "Scenario 9: Modulo with large numbers",
			num1:           math.MaxInt64,
			num2:           math.MaxInt64 - 1,
			expectedResult: 1,
			expectPanic:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					if tt.expectPanic {
						t.Log("Expected panic occurred. Test passed.")
					} else {
						t.Fail()
					}
				}
			}()

			var got int

			func() {
				if tt.expectPanic {
					got = Modulo(tt.num1, tt.num2)
				} else {
					got = Modulo(tt.num1, tt.num2)
				}
			}()

			if !tt.expectPanic {
				if got != tt.expectedResult {
					t.Errorf("Failed %s. Got %d; Expected %d", tt.name, got, tt.expectedResult)
				} else {
					t.Logf("Success %s. Got %d as expected.", tt.name, tt.expectedResult)
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

	tests := []struct {
		name      string
		num1      float64
		num2      float64
		expected  float64
		expectNaN bool
		expectInf string
	}{
		{"Scenario 1: Positive x Positive", 3.5, 2.0, 7.0, false, ""},
		{"Scenario 2: Positive x Negative", 3.5, -2.0, -7.0, false, ""},
		{"Scenario 3: Negative x Negative", -3.5, -2.0, 7.0, false, ""},
		{"Scenario 4a: Zero x Positive", 0.0, 3.5, 0.0, false, ""},
		{"Scenario 4b: Zero x Negative", 0.0, -3.5, 0.0, false, ""},
		{"Scenario 4c: Zero x Zero", 0.0, 0.0, 0.0, false, ""},
		{"Scenario 5: Large x Large", math.MaxFloat64 / 2, math.MaxFloat64 / 2, (math.MaxFloat64 / 2) * (math.MaxFloat64 / 2), false, ""},
		{"Scenario 6: Small x Small", 1e-10, 2e-10, 2e-20, false, ""},
		{"Scenario 7: Overflow to Infinity", math.MaxFloat64, 2.0, 0.0, false, "positive"},
		{"Scenario 8: Identity Multiplication by One", 42.123, 1.0, 42.123, false, ""},
		{"Scenario 9a: NaN x Valid Float", math.NaN(), 3.5, 0.0, true, ""},
		{"Scenario 9b: NaN x NaN", math.NaN(), math.NaN(), 0.0, true, ""},
		{"Scenario 10a: Negative Infinity x Positive Float", math.Inf(-1), 3.5, math.Inf(-1), false, "negative"},
		{"Scenario 10b: Negative Infinity x Negative Float", math.Inf(-1), -3.5, math.Inf(1), false, "positive"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			actual := Multiply(test.num1, test.num2)

			if test.expectNaN {
				if !math.IsNaN(actual) {
					t.Errorf("Expected NaN but got %v", actual)
				} else {
					t.Logf("Success: Expected NaN and got NaN")
				}
				return
			}

			if test.expectInf == "positive" {
				if actual != math.Inf(1) {
					t.Errorf("Expected positive infinity but got %v", actual)
				} else {
					t.Logf("Success: Expected positive infinity")
				}
				return
			} else if test.expectInf == "negative" {
				if actual != math.Inf(-1) {
					t.Errorf("Expected negative infinity but got %v", actual)
				} else {
					t.Logf("Success: Expected negative infinity")
				}
				return
			}

			if actual != test.expected {
				t.Errorf("Test '%s' failed: expected %v but got %v", test.name, test.expected, actual)
			} else {
				t.Logf("Test '%s' succeeded: expected %v and got %v", test.name, test.expected, actual)
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

	tests := []struct {
		name        string
		base        float64
		exponent    float64
		expected    float64
		expectPanic bool
	}{
		{
			name:        "Positive Base and Positive Exponent",
			base:        2.0,
			exponent:    3.0,
			expected:    8.0,
			expectPanic: false,
		},
		{
			name:        "Negative Base and Positive Exponent",
			base:        -2.0,
			exponent:    3.0,
			expected:    -8.0,
			expectPanic: false,
		},
		{
			name:        "Positive Base and Negative Exponent",
			base:        2.0,
			exponent:    -2.0,
			expected:    0.25,
			expectPanic: false,
		},
		{
			name:        "Negative Base and Negative Exponent",
			base:        -2.0,
			exponent:    -3.0,
			expected:    -0.125,
			expectPanic: false,
		},
		{
			name:        "Base Zero and Positive Exponent",
			base:        0.0,
			exponent:    5.0,
			expected:    0.0,
			expectPanic: false,
		},
		{
			name:        "Base Zero and Negative Exponent",
			base:        0.0,
			exponent:    -1.0,
			expected:    math.Inf(1),
			expectPanic: false,
		},
		{
			name:        "Base One and Any Exponent",
			base:        1.0,
			exponent:    5.0,
			expected:    1.0,
			expectPanic: false,
		},
		{
			name:        "Exponent Zero and Any Base",
			base:        10.0,
			exponent:    0.0,
			expected:    1.0,
			expectPanic: false,
		},
		{
			name:        "Very Large Base and Exponent",
			base:        1e10,
			exponent:    1e5,
			expected:    math.Inf(1),
			expectPanic: false,
		},
		{
			name:        "Small Base and Large Negative Exponent",
			base:        0.0001,
			exponent:    -1e5,
			expected:    math.Inf(1),
			expectPanic: false,
		},
		{
			name:        "Fractional Base and Exponent",
			base:        0.5,
			exponent:    2.5,
			expected:    math.Pow(0.5, 2.5),
			expectPanic: false,
		},
		{
			name:        "Very Small Base and Positive Exponent",
			base:        1e-10,
			exponent:    5.0,
			expected:    math.Pow(1e-10, 5.0),
			expectPanic: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Test %s encountered a panic: %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			actual := Power(tc.base, tc.exponent)

			if tc.expectPanic {
				t.Fatalf("Test %s failed: Expected panic but did not get one", tc.name)
			} else if math.IsNaN(tc.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("Test %s failed: Expected NaN but got %v", tc.name, actual)
				}
			} else if math.IsInf(tc.expected, 0) {
				if !math.IsInf(actual, 0) {
					t.Errorf("Test %s failed: Expected infinity but got %v", tc.name, actual)
				}
			} else if math.Abs(actual-tc.expected) > 1e-9 {
				t.Errorf("Test %s failed: Expected %v but got %v", tc.name, tc.expected, actual)
			} else {
				t.Logf("Test %s passed. Expected %v got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

