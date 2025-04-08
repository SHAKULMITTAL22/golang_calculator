package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
)

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
