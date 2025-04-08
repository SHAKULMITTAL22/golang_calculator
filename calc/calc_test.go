package calc

import (
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

	type testCase struct {
		name     string
		input    float64
		expected float64
		isNaN    bool
	}

	tests := []testCase{
		{
			name:     "Absolute Value of a Positive Float",
			input:    3.14,
			expected: 3.14,
			isNaN:    false,
		},
		{
			name:     "Absolute Value of a Negative Float",
			input:    -3.14,
			expected: 3.14,
			isNaN:    false,
		},
		{
			name:     "Absolute Value of Zero",
			input:    0.0,
			expected: 0.0,
			isNaN:    false,
		},
		{
			name:     "Absolute Value of a Large Positive Float",
			input:    math.MaxFloat64,
			expected: math.MaxFloat64,
			isNaN:    false,
		},
		{
			name:     "Absolute Value of a Large Negative Float",
			input:    -math.MaxFloat64,
			expected: math.MaxFloat64,
			isNaN:    false,
		},
		{
			name:     "Absolute Value of a Small Fractional Positive Float",
			input:    0.0001,
			expected: 0.0001,
			isNaN:    false,
		},
		{
			name:     "Absolute Value of a Small Fractional Negative Float",
			input:    -0.0001,
			expected: 0.0001,
			isNaN:    false,
		},
		{
			name:  "Absolute Value of NaN",
			input: math.NaN(),
			isNaN: true,
		},
		{
			name:     "Absolute Value of Positive Infinity",
			input:    math.Inf(1),
			expected: math.Inf(1),
			isNaN:    false,
		},
		{
			name:     "Absolute Value of Negative Infinity",
			input:    math.Inf(-1),
			expected: math.Inf(1),
			isNaN:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := Absolute(test.input)

			if test.isNaN {
				if !math.IsNaN(result) {
					t.Errorf("%s: expected NaN, but got %v", test.name, result)
				} else {
					t.Logf("%s: Passed, result is NaN as expected.", test.name)
				}
			} else {
				if result != test.expected {
					t.Errorf("%s: expected %v, but got %v", test.name, test.expected, result)
				} else {
					t.Logf("%s: Passed, result is %v as expected.", test.name, result)
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
	t.Parallel()

	tests := []struct {
		name        string
		num1        int
		num2        int
		expected    int
		expectPanic bool
	}{
		{
			name:     "Add Two Small Positive Numbers",
			num1:     3,
			num2:     5,
			expected: 8,
		},
		{
			name:     "Add Two Zero Numbers",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Add a Positive Number and Zero",
			num1:     7,
			num2:     0,
			expected: 7,
		},
		{
			name:     "Add a Negative Number and Zero",
			num1:     -4,
			num2:     0,
			expected: -4,
		},
		{
			name:     "Add Two Negative Numbers",
			num1:     -3,
			num2:     -6,
			expected: -9,
		},
		{
			name:     "Add a Positive and a Negative Number",
			num1:     10,
			num2:     -3,
			expected: 7,
		},
		{
			name:     "Add Maximum Integer Values (Overflow Check)",
			num1:     math.MaxInt,
			num2:     0,
			expected: math.MaxInt,
		},
		{
			name:     "Add a Large Positive and a Large Negative Number",
			num1:     math.MaxInt,
			num2:     -math.MaxInt,
			expected: 0,
		},
		{
			name:        "Add Two Minimum Integer Values",
			num1:        math.MinInt,
			num2:        0,
			expected:    math.MinInt,
			expectPanic: false,
		},
		{
			name:     "Commutativity Property of Addition",
			num1:     12,
			num2:     5,
			expected: 17,
		},
		{
			name:     "Associativity Property of Addition",
			num1:     3,
			num2:     4,
			expected: 7,
		},

		{
			name:     "Add Extremely Large Numbers Beyond Data Type Limits",
			num1:     math.MaxInt - 10,
			num2:     10,
			expected: math.MaxInt,
		},
		{
			name:     "Add Numbers Close to Zero",
			num1:     1,
			num2:     -1,
			expected: 0,
		},
		{
			name:     "Add Large Numbers Close to Integer Limit",
			num1:     math.MaxInt - 1,
			num2:     0,
			expected: math.MaxInt - 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("Test %s panicked unexpectedly: %v", tt.name, r)
					}
				}
			}()

			result := Add(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.num1, tt.num2, result, tt.expected)
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
	t.Run("Test scenarios for Divide function", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test: %v\n", r)
				t.Fail()
			}
		}()

		tests := []struct {
			name          string
			num1          float64
			num2          float64
			expected      float64
			expectPanic   bool
			expectedError string
		}{
			{
				name:        "Division of two positive numbers",
				num1:        10,
				num2:        2,
				expected:    5,
				expectPanic: false,
			},
			{
				name:        "Division involving zero numerator",
				num1:        0,
				num2:        5,
				expected:    0,
				expectPanic: false,
			},
			{
				name:          "Division by zero raises a panic",
				num1:          10,
				num2:          0,
				expectPanic:   true,
				expectedError: "division by zero is not allowed",
			},
			{
				name:        "Division with negative numbers",
				num1:        -10,
				num2:        2,
				expected:    -5,
				expectPanic: false,
			},
			{
				name:        "Division of two negative numbers",
				num1:        -10,
				num2:        -2,
				expected:    5,
				expectPanic: false,
			},
			{
				name:        "Floating-point precision handling",
				num1:        1.123,
				num2:        2.456,
				expected:    0.457364,
				expectPanic: false,
			},
			{
				name:        "Large number inputs handling",
				num1:        math.MaxFloat64,
				num2:        2.0,
				expected:    math.MaxFloat64 / 2.0,
				expectPanic: false,
			},
			{
				name:        "Small denominator resulting in large result",
				num1:        1000.0,
				num2:        0.0001,
				expected:    10000000.0,
				expectPanic: false,
			},
			{
				name:        "Division with equal numerator and denominator",
				num1:        42.0,
				num2:        42.0,
				expected:    1,
				expectPanic: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				defer func() {
					if r := recover(); r != nil {
						if tt.expectPanic {
							t.Logf("Expected panic occurred: %v", r)
							if r != tt.expectedError {
								t.Errorf("Incorrect panic message: got %v, expected %v", r, tt.expectedError)
							}
						} else {
							t.Errorf("Unexpected panic occurred: %v", r)
							t.Fail()
						}
					}
				}()

				if tt.expectPanic {
					Divide(tt.num1, tt.num2)

					t.Errorf("Expected panic, but none occurred")
				} else {
					result := Divide(tt.num1, tt.num2)
					if math.Abs(result-tt.expected) > 1e-6 {
						t.Errorf("Incorrect result: got %v, expected %v", result, tt.expected)
					} else {
						t.Logf("Test succeeded with result: %v", result)
					}
				}
			})
		}
	})
}

/*
ROOST_METHOD_HASH=GCD_6cf0472095
ROOST_METHOD_SIG_HASH=GCD_ab1c91475d

FUNCTION_DEF=func GCD(a, b int) int
*/
func TestGcd(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
			t.Fail()
		}
	}()

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Compute GCD for Two Positive Integers",
			a:        48,
			b:        18,
			expected: 6,
		},
		{
			name:     "Compute GCD When One Argument is Zero",
			a:        25,
			b:        0,
			expected: 25,
		},
		{
			name:     "Compute GCD When Both Arguments are Zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "Compute GCD When One Argument is Negative",
			a:        -36,
			b:        24,
			expected: 12,
		},
		{
			name:     "Compute GCD When Both Arguments are Negative",
			a:        -48,
			b:        -18,
			expected: 6,
		},
		{
			name:     "Compute GCD for Prime Numbers",
			a:        13,
			b:        17,
			expected: 1,
		},
		{
			name:     "Compute GCD for One Being a Multiple of Another",
			a:        36,
			b:        12,
			expected: 12,
		},
		{
			name:     "Compute GCD for Large Integers",
			a:        109739369,
			b:        120,
			expected: 1,
		},
		{
			name:     "Compute GCD for Floating-Point Conversions to Integers",
			a:        int(math.Floor(float64(20.5))),
			b:        int(math.Floor(float64(5.3))),
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test %s. %v\n%s", tt.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			result := GCD(tt.a, tt.b)

			if result != tt.expected {
				t.Errorf("Test '%s' failed: GCD(%d, %d) = %d; expected %d", tt.name, tt.a, tt.b, result, tt.expected)
			} else {
				t.Logf("Test '%s' passed: GCD(%d, %d) = %d; expected %d", tt.name, tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
