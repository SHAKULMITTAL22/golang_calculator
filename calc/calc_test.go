package calc

import (
	debug "runtime/debug"
	testing "testing"
)

/*
ROOST_METHOD_HASH=Absolute_d231f0ab10
ROOST_METHOD_SIG_HASH=Absolute_ec3c06e5a3

FUNCTION_DEF=func Absolute(num float64) float64 // Absolute value
*/
func TestAbsolute(t *testing.T) {

	testCases := []struct {
		name           string
		input          float64
		expectedOutput float64
	}{
		{
			name:           "Testing Absolute Function With a Positive Number",
			input:          45.76,
			expectedOutput: 45.76,
		},
		{
			name:           "Testing Absolute Function With a Negative Number",
			input:          -56.87,
			expectedOutput: 56.87,
		},
		{
			name:           "Testing Absolute Function With Zero",
			input:          0,
			expectedOutput: 0,
		},
		{
			name:           "Testing Absolute Function With Extremely Large Values",
			input:          1.7976931348623157e+308,
			expectedOutput: 1.7976931348623157e+308,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in %v. %v\n%s", tc.name, r, debug.Stack())
					t.Fail()
				}
			}()

			result := Absolute(tc.input)
			if result != tc.expectedOutput {
				t.Errorf("Failed in test %v. Expected: %v but returned: %v",
					tc.name, tc.expectedOutput, result)
			}
			t.Logf("Success in test: %v", tc.name)
		})
	}
}
