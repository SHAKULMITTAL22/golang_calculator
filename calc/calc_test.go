package calc

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	os "os"
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Absolute_f8af7505a1
ROOST_METHOD_SIG_HASH=Absolute_4bad226818

FUNCTION_DEF=func Absolute(num float64) float64
*/
func TestAbsolute(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		output float64
	}{

		{name: "Positive Input", input: 5.7, output: 5.7},

		{name: "Negative Input", input: -3.14, output: 3.14},

		{name: "Zero Input", input: 0.0, output: 0.0},

		{name: "Large Positive Number", input: 1e10, output: 1e10},

		{name: "Large Negative Number", input: -1e10, output: 1e10},

		{name: "Tiny Positive Number", input: 1e-10, output: 1e-10},

		{name: "Tiny Negative Number", input: -1e-10, output: 1e-10},

		{name: "Positive Infinity", input: math.Inf(1), output: math.Inf(1)},
		{name: "Negative Infinity", input: math.Inf(-1), output: math.Inf(1)},

		{name: "Not-a-Number (NaN)", input: math.NaN(), output: math.NaN()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution. %v \n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var buf bytes.Buffer
			stdout := &buf
			fmt.Fprint(stdout, tt.name)

			result := Absolute(tt.input)

			if math.IsNaN(result) && math.IsNaN(tt.output) {

				t.Log("Both output and expected value are NaN. Test case passed.")
			} else if result != tt.output {
				t.Errorf("Test '%s' failed | Input: %v | Expected: %v | Got: %v", tt.name, tt.input, tt.output, result)
				return
			}

			t.Logf("Test case '%s' passed | Input: %v | Output matched: %v", tt.name, tt.input, result)
		})
	}

}

/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int
*/
func TestAdd(t *testing.T) {

	type testCase struct {
		name        string
		num1        int
		num2        int
		expected    int
		panicExpect bool
	}

	testCases := []testCase{
		{
			name:     "Adding Two Positive Integers",
			num1:     3,
			num2:     5,
			expected: 8,
		},
		{
			name:     "Adding a Positive and a Negative Integer",
			num1:     10,
			num2:     -5,
			expected: 5,
		},
		{
			name:     "Adding Two Negative Integers",
			num1:     -4,
			num2:     -7,
			expected: -11,
		},
		{
			name:     "Adding Zero to a Number",
			num1:     42,
			num2:     0,
			expected: 42,
		},
		{
			name:     "Adding Zero to Itself",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Adding Large Positive Integers",
			num1:     math.MaxInt32 - 1,
			num2:     1,
			expected: math.MaxInt32,
		},
		{
			name:        "Adding Large and Small Integers (Overflow Check)",
			num1:        math.MaxInt32,
			num2:        1,
			expected:    math.MaxInt32 + 1,
			panicExpect: true,
		},
		{
			name:     "Adding Small and Large Negative Integers",
			num1:     math.MinInt32 + 1,
			num2:     -1,
			expected: math.MinInt32,
		},
		{
			name:     "Adding Two Equal Negative Numbers",
			num1:     -10,
			num2:     -10,
			expected: -20,
		},
		{
			name:     "Adding Numbers with Opposite Large Magnitudes",
			num1:     math.MaxInt32,
			num2:     -math.MaxInt32,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered. %v\n%s", r, string(debug.Stack()))
					if tc.panicExpect {
						t.Logf("Test passed due to expected panic.")
					} else {
						t.Fail()
					}
				}
			}()

			result := Add(tc.num1, tc.num2)

			if !tc.panicExpect {
				if result != tc.expected {
					t.Errorf("Expected %d, got %d for inputs (%d + %d)", tc.expected, result, tc.num1, tc.num2)
				} else {
					t.Logf("Success: %d + %d = %d", tc.num1, tc.num2, result)
				}
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

	tests := []struct {
		name        string
		input       int
		expected    interface{}
		expectPanic bool
	}{
		{name: "Positive Integer Input (n=5)", input: 5, expected: 120, expectPanic: false},
		{name: "Edge Case (n=0)", input: 0, expected: 1, expectPanic: false},
		{name: "Edge Case (n=1)", input: 1, expected: 1, expectPanic: false},
		{name: "Large Positive Integer Input (n=15)", input: 15, expected: 1307674368000, expectPanic: false},
		{name: "Invalid Input (n=-1)", input: -1, expected: "factorial is not defined for negative numbers", expectPanic: true},
		{name: "Boundary Input (math.MaxInt)", input: math.MaxInt, expected: "stack overflow or integer overflow", expectPanic: true},
		{name: "Intermediate Recursive Steps (n=3)", input: 3, expected: 6, expectPanic: false},
		{name: "Zero with No Recursion (n=0)", input: 0, expected: 1, expectPanic: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectPanic {
						t.Logf("Panic expected and observed. Message: %v\nStack: %s", r, string(debug.Stack()))
						if msg, ok := r.(string); ok {
							if msg != tt.expected {
								t.Errorf("Unexpected panic message: got '%v', want '%v'", msg, tt.expected)
							} else {
								t.Log("Panic message validation passed.")
							}
						} else {
							t.Errorf("Panic message type is unexpected: got '%T', want string", r)
						}
					} else {
						t.Errorf("Unexpected panic encountered: %v", r)
					}
					t.FailNow()
				}
			}()

			if tt.expectPanic {
				_ = Factorial(tt.input)
			} else {
				result := Factorial(tt.input)

				if res, ok := tt.expected.(int); ok && result != res {
					t.Errorf("Failed '%s': got '%v', want '%v'", tt.name, result, res)
				} else {
					t.Logf("Passed '%s': got expected result '%v'", tt.name, result)
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
	}

	testCases := []testCase{
		{name: "Two Positive Numbers", a: 48, b: 18, expected: 6},
		{name: "One Number Zero", a: 45, b: 0, expected: 45},
		{name: "Two Identical Numbers", a: 25, b: 25, expected: 25},
		{name: "Larger Number Divisible by Smaller Number", a: 56, b: 14, expected: 14},
		{name: "Two Prime Numbers", a: 13, b: 7, expected: 1},
		{name: "One Negative, One Positive Number", a: -36, b: 60, expected: 12},
		{name: "Both Numbers Negative", a: -42, b: -56, expected: 14},
		{name: "Two Very Large Numbers", a: 987654321, b: 123456789, expected: 9},
		{name: "Both Numbers Zero", a: 0, b: 0, expected: 0},
		{name: "Recursive Stopping Condition", a: 252, b: 105, expected: 21},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			outputBuffer := os.Stdout
			defer func() { os.Stdout = outputBuffer }()

			tempOutput, _ := os.Create("temporary_output.log")
			defer tempOutput.Close()
			os.Stdout = tempOutput

			result := GCD(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("Failed test %q! GCD(%d, %d) = %d; expected %d", tc.name, tc.a, tc.b, result, tc.expected)
			} else {
				t.Logf("Passed test %q! GCD(%d, %d) = %d as expected", tc.name, tc.a, tc.b, result)
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

	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{
			name:    "Scenario 1: Verify LCM of Two Positive Integers",
			a:       6,
			b:       8,
			want:    24,
			wantErr: false,
		},
		{
			name:    "Scenario 2: Verify LCM of Two Equal Integers",
			a:       7,
			b:       7,
			want:    7,
			wantErr: false,
		},
		{
			name:    "Scenario 3: Verify LCM of Zero and a Positive Integer",
			a:       0,
			b:       5,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Scenario 4: Verify LCM of Zero and Zero",
			a:       0,
			b:       0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Scenario 5: Verify LCM of Two Large Integers",
			a:       123456,
			b:       789012,
			want:    8117355456,
			wantErr: false,
		},
		{
			name:    "Scenario 6: Verify LCM with Negative Integers",
			a:       -6,
			b:       -8,
			want:    24,
			wantErr: false,
		},
		{
			name:    "Scenario 7: Verify LCM of a Positive and a Negative Integer",
			a:       6,
			b:       -8,
			want:    24,
			wantErr: false,
		},
		{
			name:    "Scenario 8: Verify LCM with Prime Numbers",
			a:       11,
			b:       13,
			want:    143,
			wantErr: false,
		},
		{
			name:    "Scenario 9: Verify LCM of Co-Primes",
			a:       9,
			b:       10,
			want:    90,
			wantErr: false,
		},
		{
			name:    "Scenario 10: Verify LCM with One Input as Maximum Integer Value",
			a:       math.MaxInt64,
			b:       1,
			want:    math.MaxInt64,
			wantErr: false,
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

			t.Logf("Running test: %s | Input: a=%d, b=%d | Expected Output: %d", tt.name, tt.a, tt.b, tt.want)

			got := LCM(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Test Failed [%s]: Expected %d, Got %d", tt.name, tt.want, got)
			} else {
				t.Logf("Test Passed [%s]: Got the expected result %d", tt.name, tt.want)
			}

			if tt.wantErr {
				t.Errorf("Test Expected an error but none was observed for [%s]: Inputs a=%d, b=%d", tt.name, tt.a, tt.b)
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
		name      string
		num       float64
		base      float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "Scenario 1 - Valid numbers and base greater than 1",
			num:       100,
			base:      10,
			expected:  2,
			expectErr: false,
		},
		{
			name:      "Scenario 2 - Logarithm base equal to Euler's number",
			num:       math.E,
			base:      math.E,
			expected:  1,
			expectErr: false,
		},
		{
			name:      "Scenario 3 - Number and base are both 1",
			num:       1,
			base:      1,
			expectErr: true,
		},
		{
			name:      "Scenario 4 - Negative number as num",
			num:       -5,
			base:      10,
			expectErr: true,
		},
		{
			name:      "Scenario 5 - Zero as num",
			num:       0,
			base:      10,
			expectErr: true,
		},
		{
			name:      "Scenario 6 - Zero as base",
			num:       10,
			base:      0,
			expectErr: true,
		},
		{
			name:      "Scenario 7 - Negative base",
			num:       10,
			base:      -2,
			expectErr: true,
		},
		{
			name:      "Scenario 8 - Non-integer base (e.g., base = 2.5)",
			num:       32,
			base:      2.5,
			expected:  3.161158,
			expectErr: false,
		},
		{
			name:      "Scenario 9 - Large numbers",
			num:       1e10,
			base:      2,
			expected:  33.219280948873624,
			expectErr: false,
		},
		{
			name:      "Scenario 10 - Small fractional numbers (num < 1)",
			num:       0.5,
			base:      10,
			expected:  -0.30103,
			expectErr: false,
		},
		{
			name:      "Scenario 11 - Very small base (base = 1.01)",
			num:       10,
			base:      1.01,
			expected:  230.258509,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.expectErr {
						t.Logf("Expected panic for test '%s': %v\n%s", tt.name, r, string(debug.Stack()))
					} else {
						t.Errorf("Unexpected panic for test '%s': %v\n%s", tt.name, r, string(debug.Stack()))
						t.FailNow()
					}
				}
			}()

			if tt.expectErr {

				Logarithm(tt.num, tt.base)
				t.Errorf("Expected error for test '%s', but got none", tt.name)
			} else {
				result := Logarithm(tt.num, tt.base)

				if math.Abs(result-tt.expected) > 1e-6 {
					t.Errorf("Test '%s' failed: Expected %.6f, got %.6f", tt.name, tt.expected, result)
				} else {
					t.Logf("Test '%s' succeeded: Expected %.6f, got %.6f", tt.name, tt.expected, result)
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
		expectInf bool
	}{
		{
			name:     "Scenario 1: Multiplication of Two Positive Numbers",
			num1:     2.5,
			num2:     4.0,
			expected: 10.0,
		},
		{
			name:     "Scenario 2: Multiplication with One Operand as Zero",
			num1:     0.0,
			num2:     5.0,
			expected: 0.0,
		},
		{
			name:     "Scenario 3: Multiplication of Two Negative Numbers",
			num1:     -3.0,
			num2:     -2.0,
			expected: 6.0,
		},
		{
			name:     "Scenario 4: Multiplication of Positive and Negative Numbers",
			num1:     3.0,
			num2:     -2.0,
			expected: -6.0,
		},
		{
			name:     "Scenario 5: Multiplication of a Number by One",
			num1:     7.5,
			num2:     1.0,
			expected: 7.5,
		},
		{
			name:     "Scenario 6: Multiplication Using Large Numbers",
			num1:     1e10,
			num2:     2e10,
			expected: 2e20,
		},
		{
			name:     "Scenario 7: Multiplication Resulting in Precision Loss",
			num1:     0.1,
			num2:     0.2,
			expected: 0.02,
		},
		{
			name:     "Scenario 8: Multiplication Using Very Small Numbers",
			num1:     1e-10,
			num2:     2e-10,
			expected: 2e-20,
		},
		{
			name:      "Scenario 9: Multiplication Involving Infinity",
			num1:      math.Inf(1),
			num2:      2.0,
			expectInf: true,
		},
		{
			name:      "Scenario 10: Multiplication Resulting in NaN",
			num1:      math.NaN(),
			num2:      3.0,
			expectNaN: true,
		},
		{
			name:     "Scenario 11: Multiplication of Extremely Large and Small Numbers",
			num1:     1e10,
			num2:     1e-10,
			expected: 1.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\nStack trace: %s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			actual := Multiply(test.num1, test.num2)

			if test.expectNaN {
				if !math.IsNaN(actual) {
					t.Fatalf("Expected NaN, but got: %v", actual)
				}
				t.Logf("Success: %s - Result correctly identified as NaN", test.name)
			} else if test.expectInf {
				if !math.IsInf(actual, 1) {
					t.Fatalf("Expected Infinity, but got: %v", actual)
				}
				t.Logf("Success: %s - Result correctly identified as Infinity", test.name)
			} else {
				const tolerance = 1e-9
				if math.Abs(actual-test.expected) > tolerance {
					t.Fatalf("Failure: %s - Expected: %v, but got: %v", test.name, test.expected, actual)
				}
				t.Logf("Success: %s - Expected result matched", test.name)
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
		name     string
		base     float64
		exponent float64
		expected float64
		explain  string
	}{
		{"Positive_Base_Positive_Exponent", 2, 3, 8.0, "Mathematical correctness verified."},
		{"Positive_Base_Zero_Exponent", 5, 0, 1.0, "Respects zero exponentiation property."},
		{"Positive_Base_Negative_Exponent", 2, -3, 0.125, "Handles reciprocal calculation correctly."},
		{"Zero_Base_Positive_Exponent", 0, 5, 0.0, "Zero base to positive exponent returns zero."},
		{"Zero_Base_Zero_Exponent", 0, 0, 1.0, "Go conventions for zero base and zero exponent yield 1."},
		{"Negative_Base_Positive_Odd_Exponent", -2, 3, -8.0, "Negative base with odd exponent yields negative value."},
		{"Negative_Base_Negative_Odd_Exponent", -2, -3, -0.125, "Reciprocal negative calculation verified for odd exponents."},
		{"Fractional_Base_Positive_Exponent", 0.5, 2, 0.25, "Floating-point precision for positive exponent verified."},
		{"Fractional_Base_Negative_Exponent", 0.5, -2, 4.0, "Handles fractional negative exponents properly."},
		{"Large_Base_Large_Exponent", 1e6, 10, math.Pow(1e6, 10), "Tests overflow behavior for large values."},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			oldStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe for stdout: %v", err)
			}
			os.Stdout = w

			result := Power(test.base, test.exponent)

			w.Close()
			os.Stdout = oldStdout

			if _, err = fmt.Fscanf(r, ""); err != nil {
				r.Close()
				t.Fatalf("Failed to read stdout: %v", err)
			}
			r.Close()

			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("For %s: Expected %.5f but got %.5f. Explanation: %s", test.name, test.expected, result, test.explain)
			} else {
				t.Logf("SUCCESS: %s passed! Result is %.5f as expected. Explanation: %s", test.name, result, test.explain)
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

	testCases := []struct {
		name        string
		angle       float64
		expectedSin float64
		expectedCos float64
		expectedTan float64
		expectPanic bool
	}{
		{
			name:        "Zero Angle",
			angle:       0,
			expectedSin: 0,
			expectedCos: 1,
			expectedTan: 0,
		},
		{
			name:        "Acute Angle (π/4)",
			angle:       math.Pi / 4,
			expectedSin: math.Sqrt(2) / 2,
			expectedCos: math.Sqrt(2) / 2,
			expectedTan: 1,
		},
		{
			name:        "Right Angle (π/2)",
			angle:       math.Pi / 2,
			expectedSin: 1,
			expectedCos: 0,
			expectedTan: math.Inf(1),
		},
		{
			name:        "Negative Angle (-π/4)",
			angle:       -math.Pi / 4,
			expectedSin: -math.Sqrt(2) / 2,
			expectedCos: math.Sqrt(2) / 2,
			expectedTan: -1,
		},
		{
			name:        "Full Circle (2π)",
			angle:       2 * math.Pi,
			expectedSin: 0,
			expectedCos: 1,
			expectedTan: 0,
		},
		{
			name:        "Large Positive Angle (1000π)",
			angle:       1000 * math.Pi,
			expectedSin: 0,
			expectedCos: 1,
			expectedTan: 0,
		},
		{
			name:        "Large Negative Angle (-1000π)",
			angle:       -1000 * math.Pi,
			expectedSin: 0,
			expectedCos: 1,
			expectedTan: 0,
		},
		{
			name:        "Arbitrary Common Angle (π/3)",
			angle:       math.Pi / 3,
			expectedSin: math.Sqrt(3) / 2,
			expectedCos: 0.5,
			expectedTan: math.Sqrt(3),
		},
		{
			name:        "Small Angle Approximation (1e-9)",
			angle:       1e-9,
			expectedSin: 1e-9,
			expectedCos: 1,
			expectedTan: 1e-9,
		},
		{
			name:        "Undefined Tangent Scenario (π/2)",
			angle:       math.Pi / 2,
			expectedSin: 1,
			expectedCos: 0,
			expectedTan: math.Inf(1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Logf("Panic encountered so failing test for case %s: %v\n%s", tc.name, r, string(debug.Stack()))
						t.Fail()
					} else {
						t.Logf("Expected panic in test case '%s', and it happened successfully.", tc.name)
					}
				}
			}()

			sin, cos, tan := SinCosTan(tc.angle)

			if !tc.expectPanic {
				if math.Abs(sin-tc.expectedSin) > 1e-9 {
					t.Errorf("Test '%s' failed: Expected sin %.8f but got %.8f", tc.name, tc.expectedSin, sin)
				} else {
					t.Logf("Test '%s' passed for sin", tc.name)
				}

				if math.Abs(cos-tc.expectedCos) > 1e-9 {
					t.Errorf("Test '%s' failed: Expected cos %.8f but got %.8f", tc.name, tc.expectedCos, cos)
				} else {
					t.Logf("Test '%s' passed for cos", tc.name)
				}

				if math.IsInf(tan, 1) && math.IsInf(tc.expectedTan, 1) {
					t.Logf("Test '%s' passed for tan (Infinity handled correctly)", tc.name)
				} else if math.Abs(tan-tc.expectedTan) > 1e-9 {
					t.Errorf("Test '%s' failed: Expected tan %.8f but got %.8f", tc.name, tc.expectedTan, tan)
				} else {
					t.Logf("Test '%s' passed for tan", tc.name)
				}
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
	type testCase struct {
		input          float64
		expectedOutput float64
		expectPanic    bool
		description    string
	}

	tests := []testCase{
		{input: 4, expectedOutput: 2, expectPanic: false, description: "Square root of 4 should be 2."},
		{input: 16, expectedOutput: 4, expectPanic: false, description: "Square root of 16 should be 4."},
		{input: 25, expectedOutput: 5, expectPanic: false, description: "Square root of 25 should be 5."},
		{input: 0, expectedOutput: 0, expectPanic: false, description: "Square root of 0 should be 0."},
		{input: -1, expectedOutput: 0, expectPanic: true, description: "Square root of -1 should cause panic."},
		{input: -25, expectedOutput: 0, expectPanic: true, description: "Square root of -25 should cause panic."},
		{input: 0.25, expectedOutput: math.Sqrt(0.25), expectPanic: false, description: "Square root of 0.25."},
		{input: 2.25, expectedOutput: math.Sqrt(2.25), expectPanic: false, description: "Square root of 2.25."},
		{input: 1e10, expectedOutput: math.Sqrt(1e10), expectPanic: false, description: "Square root of 1e10."},
		{input: 1e16, expectedOutput: math.Sqrt(1e16), expectPanic: false, description: "Square root of 1e16."},
		{input: 1e-10, expectedOutput: math.Sqrt(1e-10), expectPanic: false, description: "Square root of 1e-10."},
		{input: 1e-15, expectedOutput: math.Sqrt(1e-15), expectPanic: false, description: "Square root of 1e-15."},
		{input: 123.456, expectedOutput: math.Sqrt(123.456), expectPanic: false, description: "Square root of 123.456."},
		{input: 789.123, expectedOutput: math.Sqrt(789.123), expectPanic: false, description: "Square root of 789.123."},
		{input: math.Inf(1), expectedOutput: math.Inf(1), expectPanic: false, description: "Square root of positive infinity."},
		{input: math.NaN(), expectedOutput: math.NaN(), expectPanic: false, description: "Square root of NaN."},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("Panic correctly encountered for input: %f. %s", tc.input, string(debug.Stack()))
					} else {
						t.Errorf("Unexpected panic for input: %f. %v", tc.input, r)
					}
				}
			}()

			actualOutput := 0.0
			if !tc.expectPanic {
				actualOutput = SquareRoot(tc.input)

				if math.IsNaN(tc.expectedOutput) && !math.IsNaN(actualOutput) ||
					!math.IsNaN(tc.expectedOutput) && math.Abs(actualOutput-tc.expectedOutput) > 1e-9 {
					t.Errorf("Test failed for input %f. Expected %f, got %f", tc.input, tc.expectedOutput, actualOutput)
				} else {
					t.Logf("Test passed for input %f. Output matches expected value %f", tc.input, tc.expectedOutput)
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
	t.Helper()

	tests := []struct {
		description string
		num1        int
		num2        int
		expected    int
		expectPanic bool
	}{
		{
			description: "Scenario 1: Subtraction of Two Positive Integers",
			num1:        10,
			num2:        5,
			expected:    5,
			expectPanic: false,
		},
		{
			description: "Scenario 2: Subtraction with Zero as One Operand",
			num1:        0,
			num2:        5,
			expected:    -5,
			expectPanic: false,
		},
		{
			description: "Scenario 3: Subtraction Resulting in Zero",
			num1:        7,
			num2:        7,
			expected:    0,
			expectPanic: false,
		},
		{
			description: "Scenario 4: Subtraction of Negative Numbers",
			num1:        -8,
			num2:        -3,
			expected:    -5,
			expectPanic: false,
		},
		{
			description: "Scenario 5: Mixed Signs Leading to Positive Result",
			num1:        -5,
			num2:        -10,
			expected:    5,
			expectPanic: false,
		},
		{
			description: "Scenario 6: Resulting in MinInt Value",
			num1:        math.MinInt,
			num2:        0,
			expected:    math.MinInt,
			expectPanic: false,
		},
		{
			description: "Scenario 7: Resulting in MaxInt Value",
			num1:        math.MaxInt - 1,
			num2:        -1,
			expected:    math.MaxInt,
			expectPanic: false,
		},
		{
			description: "Scenario 8: Rounded Float Numbers Subtracted",
			num1:        int(math.Round(3.7)),
			num2:        int(math.Round(1.2)),
			expected:    3,
			expectPanic: false,
		},
		{
			description: "Scenario 9: Large Difference in Operands",
			num1:        1000000,
			num2:        3,
			expected:    999997,
			expectPanic: false,
		},
		{
			description: "Scenario 10: Identical Extreme Values Resulting in Zero",
			num1:        math.MaxInt,
			num2:        math.MaxInt,
			expected:    0,
			expectPanic: false,
		},
		{
			description: "Scenario 11: Negative Extreme Values Resulting in Zero",
			num1:        math.MinInt,
			num2:        math.MinInt,
			expected:    0,
			expectPanic: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.description, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered. %v\n%s", r, string(debug.Stack()))
					if tc.expectPanic {
						t.Log("Expected panic; test succeeded.")
					} else {
						t.Fail()
					}
				}
			}()

			result := Subtract(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Test '%s' failed. Expected %d, got %d", tc.description, tc.expected, result)
			} else {
				t.Logf("Test '%s' succeeded. Expected result matches actual.", tc.description)
			}

			if result == tc.expected {
				t.Logf("Passed: %s | Inputs: %d, %d | Output: %d", tc.description, tc.num1, tc.num2, result)
			}
		})
	}

}
