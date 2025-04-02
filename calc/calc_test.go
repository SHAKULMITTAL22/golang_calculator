package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
	fmt "fmt"
	os "os"
	bytes "bytes"
	sync "sync"
)








/*
ROOST_METHOD_HASH=Add_61b83cbd8d
ROOST_METHOD_SIG_HASH=Add_9b37ae6611

FUNCTION_DEF=func Add(num1, num2 int) int // Add two integers


*/
func TestAdd(t *testing.T) {

	type addTestCase struct {
		name   string
		input1 int
		input2 int
		expect int
		desc   string
	}

	tests := []addTestCase{
		{
			name:   "AddTwoPositiveNumbers",
			input1: 5,
			input2: 10,
			expect: 15,
			desc:   "Validates addition of two positive integers (5 + 10 = 15).",
		},
		{
			name:   "AddPositiveAndNegativeNumbers",
			input1: 10,
			input2: -5,
			expect: 5,
			desc:   "Checks addition of positive and negative integers (10 + -5 = 5).",
		},
		{
			name:   "AddTwoNegativeNumbers",
			input1: -3,
			input2: -7,
			expect: -10,
			desc:   "Validates addition of two negative integers (-3 + -7 = -10).",
		},
		{
			name:   "AddZeroToInteger",
			input1: 42,
			input2: 0,
			expect: 42,
			desc:   "Validates behavior when adding zero to a non-zero integer (42 + 0 = 42).",
		},
		{
			name:   "AddTwoZeros",
			input1: 0,
			input2: 0,
			expect: 0,
			desc:   "Verifies addition of two zeros (0 + 0 = 0).",
		},
		{
			name:   "AddTwoLargePositiveIntegers",
			input1: math.MaxInt - 1,
			input2: 2,
			expect: math.MaxInt,
			desc:   "Tests addition of large positive integers close to maximum int limit.",
		},
		{
			name:   "AddTwoLargeNegativeIntegers",
			input1: math.MinInt + 1,
			input2: -1,
			expect: math.MinInt,
			desc:   "Tests addition of large negative integers close to minimum int limit.",
		},
		{
			name:   "OverflowWithTwoIntegers",
			input1: math.MaxInt,
			input2: 1,
			expect: math.MinInt,
			desc:   "Tests behavior when addition results in integer overflow.",
		},
		{
			name:   "UnderflowWithTwoIntegers",
			input1: math.MinInt,
			input2: -1,
			expect: math.MaxInt,
			desc:   "Tests behavior when addition results in integer underflow.",
		},
		{
			name:   "AddOppositeNumbersResultingZero",
			input1: 20,
			input2: -20,
			expect: 0,
			desc:   "Validates addition of two integers of equal magnitude but opposite signs.",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test %s: %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Add(tc.input1, tc.input2)

			if result != tc.expect {
				t.Errorf("FAIL [%s]: %s | Input: (%d, %d) | Expected: %d | Got: %d",
					tc.name, tc.desc, tc.input1, tc.input2, tc.expect, result)
			} else {
				t.Logf("PASS [%s]: %s | Input: (%d, %d) | Result: %d",
					tc.name, tc.desc, tc.input1, tc.input2, result)
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

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
			t.Fail()
		}
	}()

	type testCase struct {
		name        string
		num1        float64
		num2        float64
		expected    float64
		shouldPanic bool
	}

	testCases := []testCase{
		{name: "Divide two positive numbers", num1: 10, num2: 2, expected: 5, shouldPanic: false},
		{name: "Divide positive by negative number", num1: 10, num2: -2, expected: -5, shouldPanic: false},
		{name: "Divide two negative numbers", num1: -10, num2: -2, expected: 5, shouldPanic: false},
		{name: "Divide by zero", num1: 10, num2: 0, expected: 0, shouldPanic: true},
		{name: "Divide zero by a number", num1: 0, num2: 2, expected: 0, shouldPanic: false},
		{name: "Divide two fractional numbers", num1: 0.5, num2: 0.2, expected: 2.5, shouldPanic: false},
		{name: "Divide a number by 1", num1: 7.5, num2: 1, expected: 7.5, shouldPanic: false},
		{name: "Divide a very large number by small number", num1: 1e10, num2: 1e-2, expected: 1e12, shouldPanic: false},
		{name: "Divide two very small numbers", num1: 1e-10, num2: 1e-5, expected: 1e-5, shouldPanic: false},
		{name: "Divide identical numbers", num1: 6, num2: 6, expected: 1, shouldPanic: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.shouldPanic {
						t.Logf("Successfully caught expected panic for %s: %v", tc.name, r)
					} else {
						t.Logf("Unexpected panic for %s: %v\n%s", tc.name, r, string(debug.Stack()))
						t.Fail()
					}
				}
			}()

			fmt.Fprintf(os.Stdout, "Testing: %s\n", tc.name)

			var result float64
			if !tc.shouldPanic {
				result = Divide(tc.num1, tc.num2)
			}

			if tc.shouldPanic {

				return
			}

			if math.Abs(result-tc.expected) > 1e-9 {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
				t.Logf("Details: Inputs num1=%v, num2=%v, Expected outcome=%v", tc.num1, tc.num2, tc.expected)
			} else {
				t.Logf("Test %s succeeded: expected %v, got %v", tc.name, tc.expected, result)
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
		name          string
		input         int
		expected      int
		expectPanic   bool
		panicMessage  string
		validateError string
	}

	testCases := []testCase{
		{
			name:        "Scenario 1: Factorial(5)",
			input:       5,
			expected:    120,
			expectPanic: false,
		},
		{
			name:        "Scenario 2: Factorial(0)",
			input:       0,
			expected:    1,
			expectPanic: false,
		},
		{
			name:        "Scenario 3: Factorial(1)",
			input:       1,
			expected:    1,
			expectPanic: false,
		},
		{
			name:         "Scenario 4: Factorial(-5)",
			input:        -5,
			expected:     0,
			expectPanic:  true,
			panicMessage: "factorial is not defined for negative numbers",
		},
		{
			name:        "Scenario 5: Factorial(10)",
			input:       10,
			expected:    3628800,
			expectPanic: false,
		},
		{
			name:        "Scenario 6: Factorial(2)",
			input:       2,
			expected:    2,
			expectPanic: false,
		},
		{
			name:        "Scenario 7: Performance testing at boundary - Factorial(20)",
			input:       20,
			expected:    2432902008176640000,
			expectPanic: false,
		},
		{
			name:        "Scenario 8: Factorial(15)",
			input:       15,
			expected:    1307674368000,
			expectPanic: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if testCase.expectPanic {
						if r != testCase.panicMessage {
							t.Errorf("Expected panic message '%s', got '%v'", testCase.panicMessage, r)
						}
					} else {
						t.Errorf("Unexpected panic occurred: %v", r)
					}
				} else {
					if testCase.expectPanic {
						t.Errorf("Expected panic did not occur for input %d", testCase.input)
					}
				}
			}()

			result := Factorial(testCase.input)
			if result != testCase.expected {
				t.Errorf("For %d expected %d, but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=GCD_1da681d86b
ROOST_METHOD_SIG_HASH=GCD_39a1228f3a

FUNCTION_DEF=func GCD(a, b int) int // Greatest Common Divisor (GCD) using Euclidean algorithm


*/
func TestGCD(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expectedGCD int
	}{
		{name: "GCD of 0 and 0", a: 0, b: 0, expectedGCD: 0},
		{name: "GCD of 0 and positive number", a: 0, b: 9, expectedGCD: 9},
		{name: "GCD of positive number and 0", a: 7, b: 0, expectedGCD: 7},
		{name: "GCD of two positive numbers", a: 56, b: 14, expectedGCD: 14},
		{name: "GCD of two other positive numbers", a: 100, b: 75, expectedGCD: 25},
		{name: "GCD of larger first number", a: 81, b: 27, expectedGCD: 27},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GCD(tt.a, tt.b)
			if got != tt.expectedGCD {
				t.Errorf("GCD(%d, %d) = %d; expected %d", tt.a, tt.b, got, tt.expectedGCD)
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

	type TestData struct {
		a        int
		b        int
		expected int
		name     string
	}

	tests := []TestData{
		{a: 12, b: 18, expected: 36, name: "LCM of Two Positive Integers"},
		{a: 7, b: 11, expected: 77, name: "LCM of Two Prime Numbers"},
		{a: 8, b: -12, expected: 24, name: "LCM of a Positive and Negative Integer"},
		{a: 0, b: 10, expected: 0, name: "LCM of Zero and a Non-Zero Integer"},
		{a: -6, b: -9, expected: 18, name: "LCM of Two Negative Integers"},
		{a: 15, b: 15, expected: 15, name: "LCM of Identical Numbers"},
		{a: 8, b: 15, expected: 120, name: "LCM of Coprime Numbers"},
		{a: 123456, b: 789012, expected: 8117355456, name: "LCM of Large Integers"},
		{a: 1, b: 9, expected: 9, name: "LCM of a Number and One"},
		{a: -20, b: -20, expected: 20, name: "LCM of Two Equal Negative Numbers"},
		{a: -1, b: -3, expected: 3, name: "LCM of Extremely Small Numbers"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {

				if r := recover(); r != nil {
					t.Errorf("Panic encountered in scenario '%s': %v\n%s", tc.name, r, string(debug.Stack()))
				}
			}()

			output := LCM(tc.a, tc.b)

			if output != tc.expected {
				t.Errorf("FAILED scenario '%s': LCM(%d, %d) = %d; expected %d", tc.name, tc.a, tc.b, output, tc.expected)
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
	tolerance := 1e-9

	tests := []struct {
		name        string
		num         float64
		base        float64
		expectPanic bool
		expected    float64
	}{

		{"Valid positive inputs", 8.0, 2.0, false, 3.0},

		{"Panic on zero num", 0.0, 2.0, true, 0},

		{"Panic on negative num", -5.0, 2.0, true, 0},

		{"Panic on base 1", 8.0, 1.0, true, 0},

		{"Panic on base zero", 8.0, 0.0, true, 0},
		{"Panic on negative base", 8.0, -3.0, true, 0},

		{"Num equals base", 5.0, 5.0, false, 1.0},

		{"Base greater than num", 2.0, 4.0, false, 0.5},

		{"Base smaller than num", 8.0, 2.0, false, 3.0},

		{"Fractional num and base", 7.5, 1.5, false, math.Log(7.5) / math.Log(1.5)},

		{"Large num and base", 1e10, 1e5, false, math.Log(1e10) / math.Log(1e5)},

		{"Small num and base", 1e-10, 2.0, false, math.Log(1e-10) / math.Log(2.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectPanic {
						t.Logf("Expected panic encountered: %v\n%s", r, string(debug.Stack()))
						return
					}
					t.Logf("Unexpected panic encountered: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			if tt.expectPanic {

				func() {
					_ = Logarithm(tt.num, tt.base)
				}()
				t.Errorf("Expected panic but function did not panic with inputs num=%v and base=%v", tt.num, tt.base)
			} else {

				result := Logarithm(tt.num, tt.base)
				if math.Abs(result-tt.expected) > tolerance {
					t.Errorf("Failed test: %v | Expected: %v, Actual: %v", tt.name, tt.expected, result)
				} else {
					t.Logf("Test succeeded! Input: num=%v, base=%v | Expected: %v, Actual: %v", tt.num, tt.base, tt.expected, result)
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
		name          string
		dividend      int
		divisor       int
		expectPanic   bool
		expectedValue int
	}

	testCases := []testCase{
		{"Positive Dividend and Positive Divisor", 10, 3, false, 1},
		{"Positive Dividend and Negative Divisor", 10, -3, false, 1},
		{"Negative Dividend and Positive Divisor", -10, 3, false, -1},
		{"Both Dividend and Divisor Are Negative", -10, -3, false, -1},
		{"Dividend Is Zero", 0, 5, false, 0},
		{"Divisor Is One", 123, 1, false, 0},
		{"Divisor Is Negative One", 123, -1, false, 0},
		{"Dividend Equals Divisor", 7, 7, false, 0},
		{"Dividend Is Smaller than Divisor", 3, 10, false, 3},
		{"Dividend and Divisor Are Large Positive Integers", 1_000_000_000, 123_456, false, 64000},
		{"Divisor Is Zero (Error Case)", 10, 0, true, 0},
		{"Highly Negative Dividend with Positive Divisor", -1_000_000_000, 7, false, -6},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered. Expected %v panic: %v\n%s", tc.expectPanic, r, string(debug.Stack()))
					if !tc.expectPanic {
						t.Fail()
					}
				}
			}()

			result := func() int {
				return Modulo(tc.dividend, tc.divisor)
			}()

			if !tc.expectPanic {
				if result != tc.expectedValue {
					t.Errorf("FAIL: Test '%s', Got %d, Expected %d", tc.name, result, tc.expectedValue)
				} else {
					t.Logf("PASS: Test '%s', Got %d as expected", tc.name, result)
				}
			}
		})
	}
	t.Log("Modulo function tests executed successfully. Validate unexpected failures or skipped cases.")
}


/*
ROOST_METHOD_HASH=Multiply_7a2824e2c7
ROOST_METHOD_SIG_HASH=Multiply_0911ef76c1

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 // Multiply two floating-point numbers


*/
func TestMultiply(t *testing.T) {

	tests := []struct {
		name      string
		num1      float64
		num2      float64
		expected  float64
		expectNaN bool
	}{

		{"PositiveMultiplication", 1.5, 2.0, 3.0, false},

		{"MultiplicationPositiveByZero", 4.7, 0.0, 0.0, false},

		{"MultiplicationZeroByZero", 0.0, 0.0, 0.0, false},

		{"NegativeMultiplication", -3.5, -2.0, 7.0, false},

		{"PositiveByNegativeMultiplication", 5.0, -2.0, -10.0, false},

		{"LargeNumberMultiplication", 1e10, 5e10, 5e20, false},

		{"SmallNumberMultiplication", 1e-10, 5e-10, 5e-20, false},

		{"MaxFloatMultiplication", math.MaxFloat64, 2.0, math.Inf(1), false},

		{"NegativeMaxFloatMultiplication", -math.MaxFloat64, 3.0, math.Inf(-1), false},

		{"MultiplicationWithNaN", math.NaN(), 2.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. Error: %v\nStack trace:\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			r, w, _ := os.Pipe()
			defer r.Close()
			defer w.Close()

			oldStdout := os.Stdout
			os.Stdout = w
			defer func() { os.Stdout = oldStdout }()

			actualResult := Multiply(tt.num1, tt.num2)

			if tt.expectNaN {
				if !math.IsNaN(actualResult) {
					t.Errorf("Test failed for %s. Expected NaN but got %v.", tt.name, actualResult)
				} else {
					t.Logf("Test passed for %s. Result is NaN as expected.", tt.name)
				}
			} else {

				if actualResult != tt.expected {
					t.Errorf("Test failed for %s. Expected %v but got %v.", tt.name, tt.expected, actualResult)
				} else {
					t.Logf("Test passed for %s. Expected %v and got %v.", tt.name, tt.expected, actualResult)
				}
			}

			output := ""
			fmt.Fscanf(r, "%s", &output)
			t.Logf("Captured output for test %s: %s", tt.name, output)
		})
	}
}


/*
ROOST_METHOD_HASH=SinCosTan_c242c1aa6d
ROOST_METHOD_SIG_HASH=SinCosTan_0f509380d6

FUNCTION_DEF=func SinCosTan(angle float64) (sin, cos, tan float64) // Trigonometric functions (Sin, Cos, Tan)


*/
func TestSinCosTan(t *testing.T) {

	tests := []struct {
		name   string
		angle  float64
		expect struct {
			sin float64
			cos float64
			tan float64
		}
	}{
		{
			name:  "Common Angle π/4",
			angle: math.Pi / 4,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: math.Sqrt(2) / 2, cos: math.Sqrt(2) / 2, tan: 1},
		},
		{
			name:  "Zero Angle",
			angle: 0,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: 0, cos: 1, tan: 0},
		},
		{
			name:  "Right Angle π/2",
			angle: math.Pi / 2,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: 1, cos: 0, tan: math.Inf(1)},
		},
		{
			name:  "Negative Angle -π/4",
			angle: -math.Pi / 4,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: -math.Sqrt(2) / 2, cos: math.Sqrt(2) / 2, tan: -1},
		},
		{
			name:  "Large Angle 4π",
			angle: 4 * math.Pi,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: 0, cos: 1, tan: 0},
		},
		{
			name:  "Small Angle 0.0001",
			angle: 0.0001,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: 0.0001, cos: 1, tan: 0.0001},
		},
		{
			name:  "Angle Close to π (math.Pi - 0.0001)",
			angle: math.Pi - 0.0001,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: math.Sin(math.Pi - 0.0001), cos: -math.Cos(math.Pi), tan: -math.Tan(math.Pi)},
		},
		{
			name:  "Very Large Negative Angle -10π",
			angle: -10 * math.Pi,
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: 0, cos: 1, tan: 0},
		},
		{
			name:  "Irrational Multiple of π (π/√2)",
			angle: math.Pi / math.Sqrt(2),
			expect: struct {
				sin float64
				cos float64
				tan float64
			}{sin: math.Sin(math.Pi / math.Sqrt(2)), cos: math.Cos(math.Pi / math.Sqrt(2)), tan: math.Tan(math.Pi / math.Sqrt(2))},
		},
	}

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			r, w, _ := os.Pipe()
			os.Stdout = w

			sin, cos, tan := SinCosTan(tc.angle)

			w.Close()
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(r); err != nil {
				t.Errorf("Error reading captured stdout: %v", err)
			}

			if buf.String() != "" {
				t.Logf("Captured Stdout: %s", buf.String())
			}

			if math.Abs(sin-tc.expect.sin) > 1e-6 {
				t.Errorf("Failed %s - Expected sin: %v, Got: %v", tc.name, tc.expect.sin, sin)
			}
			if math.Abs(cos-tc.expect.cos) > 1e-6 {
				t.Errorf("Failed %s - Expected cos: %v, Got: %v", tc.name, tc.expect.cos, cos)
			}
			if math.IsNaN(tan) || math.IsInf(tan, 0) || math.Abs(tan-tc.expect.tan) > 1e-6 {
				t.Errorf("Failed %s - Expected tan: %v, Got: %v", tc.name, tc.expect.tan, tan)
			}

			t.Logf("Success %s - sin: %v, cos: %v, tan: %v", tc.name, sin, cos, tan)

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
		name        string
		input       float64
		expect      float64
		shouldPanic bool
	}{
		{
			name:        "Square root of positive integer",
			input:       4.0,
			expect:      2.0,
			shouldPanic: false,
		},
		{
			name:        "Square root of zero",
			input:       0.0,
			expect:      0.0,
			shouldPanic: false,
		},
		{
			name:        "Square root of negative number",
			input:       -4.0,
			expect:      0.0,
			shouldPanic: true,
		},
		{
			name:        "Square root of a large positive number",
			input:       1e10,
			expect:      1e5,
			shouldPanic: false,
		},
		{
			name:        "Square root of a small positive number",
			input:       1e-10,
			expect:      1e-5,
			shouldPanic: false,
		},
		{
			name:        "Square root of a non-integer positive number",
			input:       2.25,
			expect:      1.5,
			shouldPanic: false,
		},
		{
			name:        "Square root of a number close to machine limit",
			input:       math.MaxFloat64,
			expect:      math.Sqrt(math.MaxFloat64),
			shouldPanic: false,
		},
		{
			name:        "Square root of a denormalized floating-point number",
			input:       5e-324,
			expect:      math.Sqrt(5e-324),
			shouldPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.shouldPanic {
						t.Logf("Panic successfully handled: %v", r)
					} else {
						t.Errorf("Unexpected panic occurred: %v", r)
					}
				}
			}()

			var result float64
			if !tc.shouldPanic {
				result = SquareRoot(tc.input)
			}

			if !tc.shouldPanic {

				if math.Abs(result-tc.expect) > 1e-12 {
					t.Errorf("Failed %s: expected %v, got %v", tc.name, tc.expect, result)
				} else {
					t.Logf("Passed %s: expected %v, got %v", tc.name, tc.expect, result)
				}
			}
		})
	}

	t.Run("Concurrent execution of square root", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic encountered during concurrent execution: %v", r)
			}
		}()

		inputs := []float64{4.0, 9.0, 16.0, 25.0}
		expectedOutputs := []float64{2.0, 3.0, 4.0, 5.0}
		results := make([]float64, len(inputs))
		var wg sync.WaitGroup

		for i := range inputs {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				results[idx] = SquareRoot(inputs[idx])
			}(i)
		}
		wg.Wait()

		for i := range inputs {
			if math.Abs(results[i]-expectedOutputs[i]) > 1e-12 {
				t.Errorf("Concurrent test failed for input %.2f: expected %.2f, got %.2f", inputs[i], expectedOutputs[i], results[i])
			} else {
				t.Logf("Concurrent test passed for input %.2f: expected %.2f, got %.2f", inputs[i], expectedOutputs[i], results[i])
			}
		}
	})
}

