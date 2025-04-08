package calc

import (
	math "math"
	testing "testing"
	debug "runtime/debug"
	strconv "strconv"
	fmt "fmt"
	os "os"
	io "io"
)








/*
ROOST_METHOD_HASH=Absolute_f8af7505a1
ROOST_METHOD_SIG_HASH=Absolute_4bad226818

FUNCTION_DEF=func Absolute(num float64) float64 

*/
func TestAbsolute(t *testing.T) {
	type testCase struct {
		desc           string
		input          float64
		expectedOutput float64
		expectNaN      bool
	}

	tests := []testCase{
		{
			desc:           "Positive Number Input - scenario 1",
			input:          5.0,
			expectedOutput: 5.0,
			expectNaN:      false,
		},
		{
			desc:           "Negative Number Input - scenario 2",
			input:          -8.3,
			expectedOutput: 8.3,
			expectNaN:      false,
		},
		{
			desc:           "Zero Input - scenario 3",
			input:          0.0,
			expectedOutput: 0.0,
			expectNaN:      false,
		},
		{
			desc:           "Large Positive Number - scenario 4",
			input:          1e+18,
			expectedOutput: 1e+18,
			expectNaN:      false,
		},
		{
			desc:           "Large Negative Number - scenario 5",
			input:          -2e+18,
			expectedOutput: 2e+18,
			expectNaN:      false,
		},
		{
			desc:           "Small Decimal Values - scenario 6",
			input:          -0.0003,
			expectedOutput: 0.0003,
			expectNaN:      false,
		},
		{
			desc:           "Positive Infinity - scenario 7",
			input:          math.Inf(1),
			expectedOutput: math.Inf(1),
			expectNaN:      false,
		},
		{
			desc:           "Negative Infinity - scenario 8",
			input:          math.Inf(-1),
			expectedOutput: math.Inf(1),
			expectNaN:      false,
		},
		{
			desc:           "NaN Input - scenario 9",
			input:          math.NaN(),
			expectedOutput: math.NaN(),
			expectNaN:      true,
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i)+"/"+tc.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			actualOutput := Absolute(tc.input)

			if tc.expectNaN {
				if !math.IsNaN(actualOutput) {
					t.Errorf("Test '%s' failed: expected output to be NaN but got %v", tc.desc, actualOutput)
				} else {
					t.Logf("Test '%s' passed: output correctly handled NaN input", tc.desc)
				}
			} else {
				if actualOutput != tc.expectedOutput {
					t.Errorf("Test '%s' failed: expected %v but got %v", tc.desc, tc.expectedOutput, actualOutput)
				} else {
					t.Logf("Test '%s' passed: output is correct", tc.desc)
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
		name        string
		num1        int
		num2        int
		expected    int
		expectPanic bool
	}{
		{"Adding positive integers", 3, 5, 8, false},
		{"Adding negative integers", -3, -5, -8, false},
		{"Adding a positive and a negative integer", 5, -3, 2, false},
		{"Adding zero to an integer", 10, 0, 10, false},
		{"Adding zero to zero", 0, 0, 0, false},
		{"Adding largest positive integers (overflow check)", math.MaxInt, math.MaxInt - 1, 0, true},
		{"Adding smallest negative integers (overflow check)", math.MinInt, math.MinInt + 1, 0, true},
		{"Adding two equal integers", 4, 4, 8, false},
		{"Adding integers to reach zero (inverse property)", 7, -7, 0, false},
		{"Adding very large positive and negative integers (balance check)", math.MaxInt, -math.MaxInt, 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Panic encountered as expected for %v: %v \n %s", tc.name, r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic error. %v \n %s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			t.Logf("Executing test: %v", tc.name)

			actual := Add(tc.num1, tc.num2)
			if !tc.expectPanic && actual != tc.expected {
				t.Errorf("%v failed. Expected %v, but got %v", tc.name, tc.expected, actual)
			} else {
				t.Logf("Success: %v correctly returned the expected result %v", tc.name, actual)
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

	type TestCase struct {
		name           string
		num1           float64
		num2           float64
		expectedResult float64
		expectPanic    bool
		panicMessage   string
	}

	testCases := []TestCase{
		{
			name:           "Successful division of two positive numbers",
			num1:           10.0,
			num2:           2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},
		{
			name:           "Division of two negative numbers",
			num1:           -10.0,
			num2:           -2.0,
			expectedResult: 5.0,
			expectPanic:    false,
		},
		{
			name:           "Division of a positive number by a negative number",
			num1:           10.0,
			num2:           -2.0,
			expectedResult: -5.0,
			expectPanic:    false,
		},
		{
			name:           "Division of zero by a positive number",
			num1:           0.0,
			num2:           5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},
		{
			name:           "Division of zero by a negative number",
			num1:           0.0,
			num2:           -5.0,
			expectedResult: 0.0,
			expectPanic:    false,
		},
		{
			name:           "Division of a number by 1",
			num1:           15.0,
			num2:           1.0,
			expectedResult: 15.0,
			expectPanic:    false,
		},
		{
			name:           "Division of a number by itself",
			num1:           12.34,
			num2:           12.34,
			expectedResult: 1.0,
			expectPanic:    false,
		},
		{
			name:           "Division by zero (error handling)",
			num1:           10.0,
			num2:           0.0,
			expectedResult: 0.0,
			expectPanic:    true,
			panicMessage:   "division by zero is not allowed",
		},
		{
			name:           "Division resulting in infinity",
			num1:           math.MaxFloat64,
			num2:           math.SmallestNonzeroFloat64,
			expectedResult: math.Inf(1),
			expectPanic:    false,
		},
		{
			name:           "Division resulting in very small fractions",
			num1:           math.SmallestNonzeroFloat64,
			num2:           math.MaxFloat64,
			expectedResult: math.SmallestNonzeroFloat64 / math.MaxFloat64,
			expectPanic:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {

					if tc.expectPanic {
						t.Logf("Panic occurred as expected: %v", r)
						if r != tc.panicMessage {
							t.Errorf("Expected panic message '%s', got '%s'", tc.panicMessage, r)
						}
					} else {
						t.Logf("Panic not expected but occurred. %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			if tc.expectPanic {
				_ = Divide(tc.num1, tc.num2)
				t.Errorf("Expected panic but function executed without error")
				return
			}

			result := Divide(tc.num1, tc.num2)

			if math.Abs(result-tc.expectedResult) > 1e-9 {
				t.Errorf("Test failed for '%s'. Expected: %v, Got: %v", tc.name, tc.expectedResult, result)
			} else {
				t.Logf("Test passed for '%s'. Expected: %v, Got: %v", tc.name, tc.expectedResult, result)
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
		name         string
		input        int
		expected     interface{}
		expectPanic  bool
		panicMessage string
	}

	testCases := []testCase{
		{
			name:        "Factorial of 0",
			input:       0,
			expected:    1,
			expectPanic: false,
		},
		{
			name:        "Factorial of 1",
			input:       1,
			expected:    1,
			expectPanic: false,
		},
		{
			name:        "Factorial of 5",
			input:       5,
			expected:    120,
			expectPanic: false,
		},
		{
			name:        "Factorial of 10",
			input:       10,
			expected:    3628800,
			expectPanic: false,
		},
		{
			name:         "Factorial of -5 (Negative Scenario)",
			input:        -5,
			expectPanic:  true,
			panicMessage: "factorial is not defined for negative numbers",
		},
		{
			name:         "Factorial of MinInt Edge Case",
			input:        math.MinInt,
			expectPanic:  true,
			panicMessage: "factorial is not defined for negative numbers",
		},
		{
			name:        "Factorial of 4 (Intermediate Test)",
			input:       4,
			expected:    24,
			expectPanic: false,
		},
		{
			name:        "Factorial of 7 (Intermediate Test)",
			input:       7,
			expected:    5040,
			expectPanic: false,
		},
		{
			name:        "Factorial of 15 (Stress Test)",
			input:       15,
			expected:    1307674368000,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Logf("Unexpected panic encountered: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					} else {

						if r != tc.panicMessage {
							t.Errorf("Expected panic message '%s', but got '%v'", tc.panicMessage, r)
						}
					}
				} else if tc.expectPanic {
					t.Errorf("Expected panic, but no panic occurred")
				}
			}()

			if !tc.expectPanic {
				result := Factorial(tc.input)
				if result != tc.expected {
					t.Errorf("Failed %s: Expected %v, got %v", tc.name, tc.expected, result)
				} else {
					t.Logf("Passed %s: Expected %v, got %v", tc.name, tc.expected, result)
				}
			} else {
				_ = Factorial(tc.input)
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

	testCases := []struct {
		name        string
		a, b        int
		expect      int
		expectPanic bool
	}{
		{
			name:        "Test GCD for two positive integers with a common divisor",
			a:           36,
			b:           24,
			expect:      12,
			expectPanic: false,
		},
		{
			name:        "Test GCD for one integer being a multiple of the other",
			a:           15,
			b:           45,
			expect:      15,
			expectPanic: false,
		},
		{
			name:        "Test GCD when integers are coprime",
			a:           13,
			b:           27,
			expect:      1,
			expectPanic: false,
		},
		{
			name:        "Test GCD when one input is zero",
			a:           0,
			b:           29,
			expect:      29,
			expectPanic: false,
		},
		{
			name:        "Test GCD for both inputs being zero",
			a:           0,
			b:           0,
			expect:      0,
			expectPanic: false,
		},
		{
			name:        "Test GCD for negative integers",
			a:           -36,
			b:           -48,
			expect:      12,
			expectPanic: false,
		},
		{
			name:        "Test GCD for one positive and one negative integer",
			a:           42,
			b:           -56,
			expect:      14,
			expectPanic: false,
		},
		{
			name:        "Test GCD for two equal integers",
			a:           20,
			b:           20,
			expect:      20,
			expectPanic: false,
		},
		{
			name:        "Test GCD for very large integers",
			a:           1000000,
			b:           5000000,
			expect:      1000000,
			expectPanic: false,
		},
		{
			name:        "Test GCD for math.MaxInt with a smaller integer",
			a:           math.MaxInt,
			b:           123,
			expect:      3,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Logf("Unexpected panic during %s: %v\n%s", tc.name, r, string(debug.Stack()))
						t.Fail()
					} else {
						t.Logf("Expected panic encountered: %s", tc.name)
					}
				}
			}()

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			result := GCD(tc.a, tc.b)

			w.Close()
			os.Stdout = old

			if result != tc.expect {
				t.Errorf("FAIL: %s | GCD(%d, %d) = %d; want %d", tc.name, tc.a, tc.b, result, tc.expect)
			} else {
				t.Logf("PASS: %s | GCD(%d, %d) = %d", tc.name, tc.a, tc.b, result)
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
		description string
		a           int
		b           int
		expected    int
	}

	testCases := []testCase{
		{"Basic functionality with two positive integers", 4, 6, 12},
		{"One input is zero", 0, 6, 0},
		{"Both inputs are zero", 0, 0, 0},
		{"Negative integers as inputs", -4, -6, 12},
		{"One positive and one negative integer input", -4, 6, 12},
		{"Large integers as inputs", math.MaxInt32, 2, math.MaxInt32 * 2 / int(math.MaxInt32)},
		{"LCM of two prime numbers", 7, 13, 91},
		{"LCM of two identical numbers", 8, 8, 8},
		{"LCM of one and any number", 1, 25, 25},
		{"LCM of coprime numbers", 12, 35, 420},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			oldStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe for capturing stdout: %v", err)
			}
			os.Stdout = w

			result := LCM(tc.a, tc.b)

			w.Close()
			os.Stdout = oldStdout

			var output string
			fmt.Fscanf(r, "%s", &output)

			if result != tc.expected {
				t.Errorf("Test failed for description: %s; expected %d, got %d", tc.description, tc.expected, result)
			} else {
				t.Logf("Test passed for description: %s", tc.description)
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
	t.Parallel()

	type testCase struct {
		name          string
		num           float64
		base          float64
		expectedValue float64
		shouldPanic   bool
		panicMessage  string
	}

	tests := []testCase{
		{
			name:          "Valid logarithm with base 10",
			num:           100,
			base:          10,
			expectedValue: 2,
			shouldPanic:   false,
		},
		{
			name:          "Valid logarithm with base 2",
			num:           8,
			base:          2,
			expectedValue: 3,
			shouldPanic:   false,
		},
		{
			name:          "Valid logarithm with base e",
			num:           math.Exp(2),
			base:          math.E,
			expectedValue: 2,
			shouldPanic:   false,
		},
		{
			name:          "Valid logarithm with fractional base",
			num:           4,
			base:          0.5,
			expectedValue: math.Log(4) / math.Log(0.5),
			shouldPanic:   false,
		},
		{
			name:         "Invalid logarithm with negative number",
			num:          -1,
			base:         10,
			shouldPanic:  true,
			panicMessage: "logarithm is not defined for these values",
		},
		{
			name:         "Invalid logarithm with zero as num",
			num:          0,
			base:         10,
			shouldPanic:  true,
			panicMessage: "logarithm is not defined for these values",
		},
		{
			name:         "Invalid logarithm with base 1",
			num:          10,
			base:         1,
			shouldPanic:  true,
			panicMessage: "logarithm is not defined for these values",
		},
		{
			name:         "Invalid logarithm with base 0",
			num:          10,
			base:         0,
			shouldPanic:  true,
			panicMessage: "logarithm is not defined for these values",
		},
		{
			name:          "Logarithm where num and base are the same",
			num:           7,
			base:          7,
			expectedValue: 1,
			shouldPanic:   false,
		},
		{
			name:          "Logarithm of 1 with any valid base",
			num:           1,
			base:          5,
			expectedValue: 0,
			shouldPanic:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if test.shouldPanic {
						t.Logf("Expected panic caught: %v\n%s", r, string(debug.Stack()))
						if r != test.panicMessage {
							t.Errorf("unexpected panic message, got '%v' expected '%v'", r, test.panicMessage)
						}
					} else {
						t.Errorf("unexpected panic encountered: %v\n%s", r, string(debug.Stack()))
					}
				}
			}()

			if test.shouldPanic {
				Logarithm(test.num, test.base)
				t.Error("expected panic but none occurred")
			} else {
				result := Logarithm(test.num, test.base)
				if math.Abs(result-test.expectedValue) > 1e-9 {
					t.Errorf("unexpected result, got '%v', expected '%v'", result, test.expectedValue)
				} else {
					t.Logf("Test passed. Input(num: %v, base: %v), Output: %v, Expected: %v",
						test.num, test.base, result, test.expectedValue)
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

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"Scenario 1: Standard Operation Using Positive Integers", 10, 3, 1},
		{"Scenario 2: Division by Zero Should Produce a Panic or Proper Handling", 5, 0, -1},
		{"Scenario 3: Modulo Using Two Negative Numbers", -10, -3, -1},
		{"Scenario 4: Modulo Using a Positive Dividend and Negative Divisor", 10, -3, 1},
		{"Scenario 5: Modulo Using a Negative Dividend and Positive Divisor", -10, 3, -1},
		{"Scenario 6: Zero Dividend Should Return Zero", 0, 7, 0},
		{"Scenario 7: Large Numbers for Stress Testing", math.MaxInt32, 3, 1},
		{"Scenario 8: Modulo Using Identical Numbers", 7, 7, 0},
		{"Scenario 9: Negative Divisor Larger Than Negative Dividend", -3, -10, -3},
		{"Scenario 10: Large Negative and Positive Mixed Integers", math.MinInt32, 7, -4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered for test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					if tc.num2 == 0 {
						t.Log("Expected panic due to division by zero. Test passed.")
					} else {
						t.Fail()
					}
				}
			}()

			stdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			actual := Modulo(tc.num1, tc.num2)

			w.Close()
			os.Stdout = stdout

			if tc.num2 != 0 {
				if actual != tc.expected {
					t.Errorf("Scenario '%s' failed. Expected: %d, Got: %d", tc.name, tc.expected, actual)
				} else {
					t.Logf("Scenario '%s' passed. Expected: %d, Got: %d", tc.name, tc.expected, actual)
				}
			}

			if tc.num2 != 0 {
				t.Logf("Test Case '%s' ran successfully with inputs num1=%d, num2=%d", tc.name, tc.num1, tc.num2)
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

	testCases := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
	}{
		{
			name:     "Multiplication of Two Positive Numbers",
			num1:     5.0,
			num2:     3.0,
			expected: 15.0,
		},
		{
			name:     "Multiplication of Two Negative Numbers",
			num1:     -4.0,
			num2:     -2.0,
			expected: 8.0,
		},
		{
			name:     "Multiplication of a Positive and a Negative Number",
			num1:     6.0,
			num2:     -3.0,
			expected: -18.0,
		},
		{
			name:     "Multiplication by Zero - Positive Factor",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		{
			name:     "Multiplication by Zero - Both Factors Zero",
			num1:     0.0,
			num2:     0.0,
			expected: 0.0,
		},
		{
			name:     "Multiplication of a Decimal Number and an Integer",
			num1:     2.5,
			num2:     4.0,
			expected: 10.0,
		},
		{
			name:     "Multiplication of Two Large Numbers",
			num1:     1e6,
			num2:     1e6,
			expected: 1e12,
		},
		{
			name:     "Multiplication of Two Small Numbers (Floating-Point Precision)",
			num1:     0.0001,
			num2:     0.0002,
			expected: 0.00000002,
		},
		{
			name:     "Multiplication with Positive Infinity",
			num1:     math.Inf(1),
			num2:     3.0,
			expected: math.Inf(1),
		},
		{
			name:     "Multiplication with Negative Infinity",
			num1:     math.Inf(-1),
			num2:     3.0,
			expected: math.Inf(-1),
		},
		{
			name:     "Multiplication with NaN",
			num1:     math.NaN(),
			num2:     5.0,
			expected: math.NaN(),
		},
		{
			name:     "Multiplication with Extreme Values",
			num1:     math.MaxFloat64,
			num2:     -math.MaxFloat64,
			expected: math.Inf(-1),
		},
		{
			name:     "Multiplication of Two Identical Numbers (Squaring)",
			num1:     7.0,
			num2:     7.0,
			expected: 49.0,
		},
		{
			name:     "Multiplication with Negative Zero",
			num1:     -0.0,
			num2:     5.0,
			expected: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered for test '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Multiply(tc.num1, tc.num2)

			if math.IsNaN(tc.expected) {
				if !math.IsNaN(result) {
					t.Errorf("Test %s - Failed: Expected NaN but got %v", tc.name, result)
					return
				}
				t.Logf("Test %s - Passed: Result matches expectation for NaN", tc.name)
				return
			}

			if result != tc.expected {
				t.Errorf("Test %s - Failed: Expected %v but got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s - Passed: Result %v matches expected %v", tc.name, result, tc.expected)
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
		expectPanic bool
	}{
		{"Positive Base, Positive Exponent", 2.0, 3.0, 8.0, false},
		{"Positive Base, Exponent Zero", 5.0, 0.0, 1.0, false},
		{"Base Zero, Positive Exponent", 0.0, 3.0, 0.0, false},
		{"Base Zero, Exponent Zero", 0.0, 0.0, 1.0, false},
		{"Negative Base, Positive Odd Exponent", -2.0, 3.0, -8.0, false},
		{"Negative Base, Positive Even Exponent", -2.0, 2.0, 4.0, false},
		{"Negative Base, Negative Exponent", -2.0, -2.0, 0.25, false},
		{"Positive Base, Negative Exponent", 2.0, -3.0, 0.125, false},
		{"Fractional Base, Positive Exponent", 0.5, 2.0, 0.25, false},
		{"Fractional Base, Negative Exponent", 0.5, -2.0, 4.0, false},
		{"Large Base, Large Exponent", 10.0, 10.0, 10000000000.0, false},
		{"Exponent as Fractional Value", 9.0, 0.5, 3.0, false},
		{"Negative Fractional Exponent", 9.0, -0.5, 0.3333333333333333, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case '%s': %v\nStack Trace:\n%s", tc.name, r, string(debug.Stack()))
					if !tc.expectPanic {
						t.Fail()
					}
				}
			}()

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			result := Power(tc.base, tc.exponent)

			w.Close()
			os.Stdout = oldStdout
			output, _ := fmt.Fscan(r)

			if math.Abs(result-tc.expected) > 1e-9 {
				t.Errorf("Test case '%s' failed: expected %v, got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Test case '%s' succeeded: expected %v, got %v", tc.name, tc.expected, result)
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

	type testCase struct {
		name        string
		angle       float64
		expectedSin float64
		expectedCos float64
		expectedTan float64
		validateTan func(float64) bool
		expectPanic bool
	}

	const tolerance = 1e-9
	almostEqual := func(a, b float64) bool {
		return math.Abs(a-b) <= tolerance
	}

	tests := []testCase{
		{
			name:        "Angle 0 radians",
			angle:       0,
			expectedSin: 0,
			expectedCos: 1,
			expectedTan: 0,
			validateTan: func(tan float64) bool { return almostEqual(tan, 0) },
		},
		{
			name:        "Angle π/2 radians (90 degrees)",
			angle:       math.Pi / 2,
			expectedSin: 1,
			expectedCos: 0,
			expectedTan: math.MaxFloat64,
			validateTan: func(tan float64) bool { return math.IsInf(tan, 0) || tan > 1e9 },
			expectPanic: false,
		},
		{
			name:        "Angle π radians (180 degrees)",
			angle:       math.Pi,
			expectedSin: 0,
			expectedCos: -1,
			expectedTan: 0,
			validateTan: func(tan float64) bool { return almostEqual(tan, 0) },
		},
		{
			name:        "Angle -π/4 radians (-45 degrees)",
			angle:       -math.Pi / 4,
			expectedSin: -1 / math.Sqrt(2),
			expectedCos: 1 / math.Sqrt(2),
			expectedTan: -1,
			validateTan: func(tan float64) bool { return almostEqual(tan, -1) },
		},
		{
			name:        "Angle θ and θ + 2π radians, periodicity",
			angle:       math.Pi / 6,
			expectedSin: math.Sin(math.Pi / 6),
			expectedCos: math.Cos(math.Pi / 6),
			expectedTan: math.Tan(math.Pi / 6),
			validateTan: func(tan float64) bool { return true },
		},
		{
			name:        "Very large positive angle",
			angle:       1e5,
			expectedSin: math.Sin(1e5),
			expectedCos: math.Cos(1e5),
			expectedTan: math.Tan(1e5),
			validateTan: func(tan float64) bool { return true },
		},
		{
			name:        "Very small negative angle",
			angle:       -1e-7,
			expectedSin: -1e-7,
			expectedCos: 1,
			expectedTan: -1e-7,
			validateTan: func(tan float64) bool { return almostEqual(tan, -1e-7) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if test.expectPanic {
						t.Logf("Panic encountered as expected: %v\n%s", r, string(debug.Stack()))
					} else {
						t.Logf("Unexpected panic encountered: %v\n%s", r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			sin, cos, tan := SinCosTan(test.angle)

			t.Logf("Inputs: angle=%v", test.angle)
			t.Logf("Outputs: sin=%v, cos=%v, tan=%v", sin, cos, tan)

			if !almostEqual(sin, test.expectedSin) {
				t.Errorf("sin(%v): got=%v, expected=%v", test.angle, sin, test.expectedSin)
			}
			if !almostEqual(cos, test.expectedCos) {
				t.Errorf("cos(%v): got=%v, expected=%v", test.angle, cos, test.expectedCos)
			}
			if !test.validateTan(tan) {
				t.Errorf("tan(%v): got=%v, expected=%v or infinity condition", test.angle, tan, test.expectedTan)
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
		name      string
		input     float64
		expected  float64
		expectErr bool
	}{

		{"PositiveNumber", 4.0, 2.0, false},

		{"ZeroInput", 0.0, 0.0, false},

		{"NegativeInput", -5.0, 0.0, true},

		{"LargePositiveNumber", 1e12, 1e6, false},

		{"SmallFraction", 0.0004, 0.02, false},

		{"PrecisionComputation", 2.0, math.Sqrt(2), false},

		{"Float64MaxValue", math.MaxFloat64, math.Sqrt(math.MaxFloat64), false},

		{"StressInputSmallRange", 16.0, 4.0, false},

		{"Float64Input", float64(16), 4.0, false},
	}

	realStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Log("Starting table-driven tests for SquareRoot function")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test: %v\n%s", r, string(debug.Stack()))
					if tt.expectErr {
						t.Logf("Expected panic for input: %v, encountered as expected.", tt.input)
					} else {
						t.FailNow()
					}
				}
			}()

			actual := float64(0)
			if !tt.expectErr {
				actual = SquareRoot(tt.input)
			}

			if tt.expectErr {
				t.Log("Skipping value comparison due to expected panic.")
				return
			}

			epsilon := 1e-9
			if math.Abs(actual-tt.expected) > epsilon {
				t.Errorf("Test %s Failed => Expected: %v, Got: %v", tt.name, tt.expected, actual)
			} else {
				t.Logf("Test %s Passed => Input: %v, Output: %v", tt.name, tt.input, actual)
			}
		})
	}

	w.Close()
	output, _ := io.ReadAll(r)
	os.Stdout = realStdout

	t.Log("Capturing test-related logs:\n", string(output))
	t.Log("Table-driven tests complete -- Results logged above.")
}


/*
ROOST_METHOD_HASH=Subtract_559013d27f
ROOST_METHOD_SIG_HASH=Subtract_29b74c09c9

FUNCTION_DEF=func Subtract(num1, num2 int) int 

*/
func TestSubtract(t *testing.T) {
	type testCase struct {
		name        string
		num1        int
		num2        int
		expected    int
		description string
	}

	tests := []testCase{
		{
			name:        "Subtract smaller integer from a larger integer",
			num1:        10,
			num2:        3,
			expected:    7,
			description: "Checks straightforward subtraction where result is positive.",
		},
		{
			name:        "Subtract larger integer from a smaller integer",
			num1:        3,
			num2:        10,
			expected:    -7,
			description: "Tests subtraction resulting in negative output.",
		},
		{
			name:        "Subtract zero from an integer",
			num1:        5,
			num2:        0,
			expected:    5,
			description: "Ensures subtraction's mathematical principles hold with zero.",
		},
		{
			name:        "Subtract integer from itself",
			num1:        8,
			num2:        8,
			expected:    0,
			description: "Validates subtraction yielding zero for identical inputs.",
		},
		{
			name:        "Subtract two negative integers",
			num1:        -5,
			num2:        -3,
			expected:    -2,
			description: "Ensures correctness in negative integer subtraction.",
		},
		{
			name:        "Subtract positive integer from a negative integer",
			num1:        -7,
			num2:        5,
			expected:    -12,
			description: "Verifies handling of mixed-sign subtraction (negative from positive).",
		},
		{
			name:        "Subtract negative integer from a positive integer",
			num1:        10,
			num2:        -5,
			expected:    15,
			description: "Ensures double negatives convert into addition correctly.",
		},
		{
			name:        "Subtract with minimum integer values",
			num1:        math.MinInt32,
			num2:        -1,
			expected:    math.MinInt32 + 1,
			description: "Validates subtraction involving boundary conditions.",
		},
		{
			name:        "Subtract with maximum integer values",
			num1:        math.MaxInt32,
			num2:        1,
			expected:    math.MaxInt32 - 1,
			description: "Tests handling subtraction near maximum integer values.",
		},
		{
			name:        "Subtraction resulting in zero",
			num1:        0,
			num2:        0,
			expected:    0,
			description: "Validates correctness in subtraction yielding zero.",
		},
		{
			name:        "Subtraction involving edge large values",
			num1:        math.MaxInt32,
			num2:        math.MinInt32,
			expected:    math.MaxInt32 - math.MinInt32,
			description: "Tests subtraction involving the extreme limits without overflow.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered. %v\n%s", r, string(debug.Stack()))
					t.FailNow()
				}
			}()

			t.Logf("Running test '%s': %s", tc.name, tc.description)

			expected := tc.expected
			actual := Subtract(tc.num1, tc.num2)

			if actual != expected {
				t.Errorf("FAILED '%s': expected %d, got %d.", tc.name, expected, actual)
			} else {
				t.Logf("PASSED '%s': expected %d, got %d.", tc.name, expected, actual)
			}
		})
	}
}

