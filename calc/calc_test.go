package calc

import (
	os "os"
	debug "runtime/debug"
	testing "testing"
)

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
