package main

import (
	bytes "bytes"
	fmt "fmt"
	os "os"
	strconv "strconv"
	testing "testing"
	debug "runtime/debug"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {

	type testCase struct {
		name           string
		input          string
		expectedResult float64
		expectError    bool
		exitCode       int
	}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {

	captureOutput := func(f func()) string {
		var buf bytes.Buffer
		stdout := os.Stdout
		defer func() { os.Stdout = stdout }()
		os.Stdout = &buf
		f()
		return buf.String()
	}

	type testCase struct {
		name           string
		input          string
		expectedOutput int
		expectExit     bool
	}

	testCases := []testCase{
		{
			name:           "Scenario 1: Valid Integer Conversion",
			input:          "123",
			expectedOutput: 123,
			expectExit:     false,
		},
		{
			name:           "Scenario 2: Zero in String Format",
			input:          "0",
			expectedOutput: 0,
			expectExit:     false,
		},
		{
			name:           "Scenario 3: Negative Integer Conversion",
			input:          "-123",
			expectedOutput: -123,
			expectExit:     false,
		},
		{
			name:           "Scenario 4: Non-Numeric String",
			input:          "abc",
			expectedOutput: 0,
			expectExit:     true,
		},
		{
			name:           "Scenario 5: Empty String",
			input:          "",
			expectedOutput: 0,
			expectExit:     true,
		},
		{
			name:           "Scenario 6: String with Whitespace",
			input:          " 123 ",
			expectedOutput: 0,
			expectExit:     true,
		},
		{
			name:           "Scenario 7: Very Large Integer String",
			input:          "9223372036854775807",
			expectedOutput: 9223372036854775807,
			expectExit:     false,
		},
		{
			name:           "Scenario 8: Out-of-Bounds Integer String",
			input:          "99999999999999999999",
			expectedOutput: 0,
			expectExit:     true,
		},
		{
			name:           "Scenario 9: String with Special Characters",
			input:          "1#2",
			expectedOutput: 0,
			expectExit:     true,
		},
		{
			name:           "Scenario 10: System Graceful Termination Check",
			input:          "xyz",
			expectedOutput: 0,
			expectExit:     true,
		},
		{
			name:           "Scenario 11: Floating Point Numeric String",
			input:          "12.34",
			expectedOutput: 0,
			expectExit:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			if tc.expectExit {
				output := captureOutput(func() {
					defer func() {
						r := recover()
						if r != nil {
							if code, ok := r.(int); !ok || code != 2 {
								t.Errorf("Unexpected exit code: got %v, expected 2.", code)
							}
						} else {
							t.Errorf("Expected os.Exit but did not occur.")
						}
					}()
					stringToInt(tc.input)
				})

				if len(output) == 0 {
					t.Errorf("Expected error output, but none received.")
				}
				return
			}

			actualOutput := stringToInt(tc.input)

			if actualOutput != tc.expectedOutput {
				t.Errorf("Test '%s' failed: expected %d, got %d.", tc.name, tc.expectedOutput, actualOutput)
			} else {
				t.Logf("Test '%s' passed: input '%s' successfully converted to %d.", tc.name, tc.input, actualOutput)
			}
		})
	}
}

