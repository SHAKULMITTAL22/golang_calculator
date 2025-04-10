package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
)








/*
ROOST_METHOD_HASH=Multiply_1585632006
ROOST_METHOD_SIG_HASH=Multiply_d6ab1fb07f

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 

*/
func TestMultiply(t *testing.T) {

	const epsilon = 1e-9

	testCases := []struct {
		name            string
		num1            float64
		num2            float64
		expected        float64
		expectNaN       bool
		expectInf       bool
		expectedInfSign int
	}{

		{
			name:     "Scenario 1: Positive Integers",
			num1:     5.0,
			num2:     4.0,
			expected: 20.0,
		},

		{
			name:     "Scenario 2: Positive Floats",
			num1:     2.5,
			num2:     1.5,
			expected: 3.75,
		},

		{
			name:     "Scenario 3: Positive and Negative",
			num1:     10.0,
			num2:     -3.5,
			expected: -35.0,
		},

		{
			name:     "Scenario 4: Two Negatives",
			num1:     -4.0,
			num2:     -2.5,
			expected: 10.0,
		},

		{
			name:     "Scenario 5: Zero First Operand",
			num1:     0.0,
			num2:     123.45,
			expected: 0.0,
		},

		{
			name:     "Scenario 6: Zero Second Operand",
			num1:     -987.65,
			num2:     0.0,
			expected: 0.0,
		},

		{
			name:     "Scenario 7: Zero by Zero",
			num1:     0.0,
			num2:     0.0,
			expected: 0.0,
		},

		{
			name:            "Scenario 8: Positive Infinity First Operand",
			num1:            math.Inf(1),
			num2:            10.0,
			expectInf:       true,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 9: Positive Infinity Second Operand",
			num1:            100.0,
			num2:            math.Inf(1),
			expectInf:       true,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 10: Negative Infinity and Positive",
			num1:            5.0,
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: -1,
		},

		{
			name:            "Scenario 11: Negative Infinity and Negative",
			num1:            -10.0,
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 12: Positive Infinity and Negative Infinity",
			num1:            math.Inf(1),
			num2:            math.Inf(-1),
			expectInf:       true,
			expectedInfSign: -1,
		},

		{
			name:      "Scenario 13: NaN First Operand",
			num1:      math.NaN(),
			num2:      10.0,
			expectNaN: true,
		},

		{
			name:      "Scenario 14: NaN Second Operand",
			num1:      -5.0,
			num2:      math.NaN(),
			expectNaN: true,
		},

		{
			name:      "Scenario 15: Infinity by Zero",
			num1:      math.Inf(1),
			num2:      0.0,
			expectNaN: true,
		},

		{
			name:      "Scenario 16: Zero by Infinity",
			num1:      0.0,
			num2:      math.Inf(-1),
			expectNaN: true,
		},

		{
			name:     "Scenario 17: Near MaxFloat64",
			num1:     math.MaxFloat64 / 2.0,
			num2:     1.9,
			expected: (math.MaxFloat64 / 2.0) * 1.9,
		},

		{
			name:            "Scenario 18: Overflow to Positive Infinity",
			num1:            math.MaxFloat64,
			num2:            2.0,
			expectInf:       true,
			expectedInfSign: 1,
		},

		{
			name:            "Scenario 19: Overflow to Negative Infinity",
			num1:            math.MaxFloat64,
			num2:            -2.0,
			expectInf:       true,
			expectedInfSign: -1,
		},

		{
			name:     "Scenario 20: Underflow to Zero",
			num1:     1e-200,
			num2:     1e-200,
			expected: 0.0,
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
			t.Logf("  Inputs: num1 = %v, num2 = %v", tc.num1, tc.num2)

			actual := Multiply(tc.num1, tc.num2)

			if tc.expectNaN {

				if !math.IsNaN(actual) {
					t.Errorf("Multiply(%v, %v) = %v; expected NaN, but got a number", tc.num1, tc.num2, actual)
					t.Logf("  Failure Reason: Expected NaN result for inputs %v, %v, but received %v.", tc.num1, tc.num2, actual)
				} else {
					t.Logf("  Success: Correctly resulted in NaN for inputs %v, %v.", tc.num1, tc.num2)
				}
			} else if tc.expectInf {

				expectedSignStr := "Positive"
				if tc.expectedInfSign == -1 {
					expectedSignStr = "Negative"
				}
				if !math.IsInf(actual, tc.expectedInfSign) {
					t.Errorf("Multiply(%v, %v) = %v; expected %s Infinity, but got %v", tc.num1, tc.num2, actual, expectedSignStr, actual)
					t.Logf("  Failure Reason: Expected %s Infinity for inputs %v, %v, but received %v.", expectedSignStr, tc.num1, tc.num2, actual)
				} else {
					t.Logf("  Success: Correctly resulted in %s Infinity for inputs %v, %v.", expectedSignStr, tc.num1, tc.num2)
				}
			} else {

				if math.IsNaN(actual) {

					t.Errorf("Multiply(%v, %v) = NaN; expected %v, but got NaN", tc.num1, tc.num2, tc.expected)
					t.Logf("  Failure Reason: Expected a numeric result (%v) for inputs %v, %v, but received NaN.", tc.expected, tc.num1, tc.num2)
				} else if math.IsInf(actual, 0) {

					infSign := 1
					if math.IsInf(actual, -1) {
						infSign = -1
					}
					infSignStr := "Positive"
					if infSign == -1 {
						infSignStr = "Negative"
					}
					t.Errorf("Multiply(%v, %v) = %s Infinity; expected %v, but got Infinity", tc.num1, tc.num2, infSignStr, tc.expected)
					t.Logf("  Failure Reason: Expected a numeric result (%v) for inputs %v, %v, but received %s Infinity.", tc.expected, tc.num1, tc.num2, infSignStr)
				} else if diff := math.Abs(actual - tc.expected); diff > epsilon {

					if !(tc.expected == 0.0 && actual == 0.0) {
						t.Errorf("Multiply(%v, %v) = %v; expected %v (difference %v exceeds epsilon %v)", tc.num1, tc.num2, actual, tc.expected, diff, epsilon)
						t.Logf("  Failure Reason: Result %v is not within epsilon %v of expected %v for inputs %v, %v. Difference: %v", actual, epsilon, tc.expected, tc.num1, tc.num2, diff)
					} else {

						t.Logf("  Success: Result %v is within epsilon %v of expected %v for inputs %v, %v.", actual, epsilon, tc.expected, tc.num1, tc.num2)
					}

				} else {

					t.Logf("  Success: Result %v is within epsilon %v of expected %v for inputs %v, %v.", actual, epsilon, tc.expected, tc.num1, tc.num2)
				}
			}
		})
	}
}

