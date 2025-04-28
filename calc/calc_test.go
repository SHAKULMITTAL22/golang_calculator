package calc

import (
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=GCD_1da681d86b
ROOST_METHOD_SIG_HASH=GCD_39a1228f3a

FUNCTION_DEF=func GCD(a, b int) int // Greatest Common Divisor (GCD) using Euclidean algorithm
*/
func TestGcd(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic encountered, test failed. Detail: %v\n%s", r, string(debug.Stack()))
			t.Fail()
		}
	}()

	testCases := []struct {
		name       string
		a, b       int
		expected   int
		shouldFail bool
	}{

		{"Positive Integers", 56, 98, 14, false},

		{"Zero as First Number", 0, 45, 45, false},
		{"Zero as Second Number", 45, 0, 45, false},

		{"Equal Integers", 25, 25, 25, false},

		{"Co-Prime Integers", 17, 13, 1, false},

		{"Negative Positive Combo", -48, 18, 6, false},
		{"Positive Negative Combo", 18, -48, 6, false},

		{"Both Negative", -36, -60, 12, false},

		{"Both Zero (Undefined)", 0, 0, 0, true},

		{"Very Large Integers", 1000000000, 500000000, 500000000, false},

		{"Prime and Its Multiple", 7, 35, 7, false},

		{"No Common Divisors", 9, 28, 1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic in test `%s`: %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			actual := GCD(tc.a, tc.b)

			if tc.shouldFail {

				t.Logf("Scenario '%s' expected failure due to undefined behavior. Observed result: %d", tc.name, actual)
			} else if actual != tc.expected {
				t.Errorf(
					"Scenario `%s` failed. Expected (%d), but got (%d). a = %d, b = %d",
					tc.name, tc.expected, actual, tc.a, tc.b,
				)
			} else {
				t.Logf("Scenario `%s` passed. Correct GCD: %d", tc.name, actual)
			}
		})
	}
}
