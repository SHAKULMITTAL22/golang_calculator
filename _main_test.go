package golang_calculator

import (
	fmt "fmt"
	os "os"
	strconv "strconv"
	testing "testing"
	bytes "bytes"
	debug "runtime/debug"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput float64
		expectPanic    bool
	}{
		{
			name:           "Valid Numeric String",
			input:          "123.45",
			expectedOutput: 123.45,
			expectPanic:    false,
		},
		{
			name:           "Negative Numeric String",
			input:          "-456.78",
			expectedOutput: -456.78,
			expectPanic:    false,
		},
		{
			name:           "Zero Numeric String",
			input:          "0",
			expectedOutput: 0.0,
			expectPanic:    false,
		},
		{
			name:           "Non-Numeric String",
			input:          "hello",
			expectedOutput: 0,
			expectPanic:    true,
		},
		{
			name:           "Empty String",
			input:          "",
			expectedOutput: 0,
			expectPanic:    true,
		},
		{
			name:           "Large Numeric String",
			input:          "1e+308",
			expectedOutput: 1e+308,
			expectPanic:    false,
		},
		{
			name:           "Special Float Value NaN",
			input:          "NaN",
			expectedOutput: 0,
			expectPanic:    true,
		},
		{
			name:           "String with Extra Whitespace",
			input:          "  123.456  ",
			expectedOutput: 123.456,
			expectPanic:    false,
		},
		{
			name:           "String in Scientific Notation",
			input:          "6.022e23",
			expectedOutput: 6.022e23,
			expectPanic:    false,
		},
		{
			name:           "String with Extreme Precision",
			input:          "0.123456789012345678",
			expectedOutput: 0.123456789012345678,
			expectPanic:    false,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if tc.expectPanic {
						t.Logf("[SUCCESS] Expected panic encountered with input '%s': %v\n%s", tc.input, r, string(debug.Stack()))
					} else {
						t.Logf("[FAILURE] Unexpected panic occurred with input '%s': %v\n%s", tc.input, r, string(debug.Stack()))
						t.Fail()
					}
				} else if tc.expectPanic {
					t.Logf("[FAILURE] Expected panic was NOT encountered with input '%s'", tc.input)
					t.Fail()
				}
			}()

			var stdoutBuffer bytes.Buffer
			tempStdout := os.Stdout
			defer func() { os.Stdout = tempStdout }()
			os.Stdout = &stdoutBuffer

			output := stringToFloat64(tc.input)

			if !tc.expectPanic {
				if output != tc.expectedOutput {
					t.Errorf("[FAILURE] Test case '%s' failed. Expected: %v, Got: %v", tc.name, tc.expectedOutput, output)
				} else {
					t.Logf("[SUCCESS] Test case '%s' passed. Output matched: %v", tc.name, output)
				}
			} else {
				t.Logf("[INFO] Test case '%s': Panic expected and handled", tc.name)
			}
		})
	}
}

