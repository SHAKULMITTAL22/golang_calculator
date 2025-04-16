package calc

import (
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=GCD_6cf0472095
ROOST_METHOD_SIG_HASH=GCD_ab1c91475d

FUNCTION_DEF=func GCD(a, b int) int
*/
func TestGcd(t *testing.T) {

	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		shouldPanic bool
	}{
		{
			name:        "Normal case with two positive integers",
			a:           12,
			b:           8,
			expected:    4,
			shouldPanic: false,
		},
		{
			name:        "One input is zero",
			a:           0,
			b:           5,
			expected:    5,
			shouldPanic: false,
		},
		{
			name:        "Both inputs are zero",
			a:           0,
			b:           0,
			expected:    0,
			shouldPanic: true,
		},
		{
			name:        "Both inputs are prime numbers",
			a:           13,
			b:           17,
			expected:    1,
			shouldPanic: false,
		},
		{
			name:        "Negative integers as input",
			a:           -16,
			b:           -24,
			expected:    8,
			shouldPanic: false,
		},
		{
			name:        "Large numbers as input",
			a:           123456789,
			b:           987654321,
			expected:    9,
			shouldPanic: false,
		},
		{
			name:        "One input is negative, the other positive",
			a:           -27,
			b:           18,
			expected:    9,
			shouldPanic: false,
		},
		{
			name:        "Inputs where one number is much larger than the other",
			a:           1000,
			b:           3,
			expected:    1,
			shouldPanic: false,
		},
		{
			name:        "Inputs are equal",
			a:           25,
			b:           25,
			expected:    25,
			shouldPanic: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {

				if r := recover(); r != nil {
					t.Logf("Panic encountered during '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					if tc.shouldPanic {
						t.Log("Test passed. Panic was expected.")
						return
					}
					t.Fail()
				}
			}()

			result := GCD(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("Scenario '%s' failed. Expected %d, got %d.", tc.name, tc.expected, result)
			} else {
				t.Logf("Scenario '%s' succeeded. GCD(%d, %d) returned %d as expected.", tc.name, tc.a, tc.b, result)
			}
		})
	}
}
