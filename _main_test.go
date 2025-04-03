package main

import (
	bytes "bytes"
	os "os"
	debug "runtime/debug"
	testing "testing"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {

	type TestCase struct {
		name          string
		input         string
		expectedValue float64
		expectedExit  bool
		expectedError string
	}

	originalStdout := os.Stdout

	captureBuffer := &bytes.Buffer{}
	t.Cleanup(func() {
		os.Stdout = originalStdout
	})

	os.Stdout = captureBuffer

	testCases := []TestCase{
		{
			name:          "Valid Float String Conversion",
			input:         "123.45",
			expectedValue: 123.45,
			expectedExit:  false,
		},
		{
			name:          "Invalid Float String",
			input:         "abc",
			expectedExit:  true,
			expectedError: "strconv.ParseFloat: parsing \"abc\": invalid syntax",
		},
		{
			name:          "Empty String as Input",
			input:         "",
			expectedExit:  true,
			expectedError: "strconv.ParseFloat: parsing \"\": invalid syntax",
		},
		{
			name:          "Large Float String",
			input:         "1.7976931348623157e+308",
			expectedValue: 1.7976931348623157e+308,
			expectedExit:  false,
		},
		{
			name:          "Very Small Float String",
			input:         "2.2250738585072014e-308",
			expectedValue: 2.2250738585072014e-308,
			expectedExit:  false,
		},
		{
			name:          "Negative Float String",
			input:         "-1234.56",
			expectedValue: -1234.56,
			expectedExit:  false,
		},
		{
			name:          "String with Whitespace",
			input:         "   123.45   ",
			expectedValue: 123.45,
			expectedExit:  false,
		},
		{
			name:          "Special Float Values: NaN",
			input:         "NaN",
			expectedValue: 0,
			expectedExit:  false,
		},
		{
			name:          "Special Float Values: Infinity",
			input:         "Infinity",
			expectedValue: 0,
			expectedExit:  false,
		},
		{
			name:          "Non-Numeric Characters in String",
			input:         "123abc",
			expectedExit:  true,
			expectedError: "strconv.ParseFloat: parsing \"123abc\": invalid syntax",
		},
		{
			name:         "Exceedingly Long Input String",
			input:        string(make([]byte, 10000000)),
			expectedExit: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					if tc.expectedExit {
						t.Logf("Exit behavior is expected. Test case succeeded.")
					} else {
						t.FailNow()
					}
				}
			}()

			var result float64
			func() {
				result = stringToFloat64(tc.input)
			}()

			if tc.expectedExit {

				output := captureBuffer.String()
				if tc.expectedError != "" && !bytes.Contains([]byte(output), []byte(tc.expectedError)) {
					t.Errorf("Expected error '%s' not found in output '%s'", tc.expectedError, output)
				}
			} else {

				if result != tc.expectedValue {
					t.Errorf("Expected %v but got %v", tc.expectedValue, result)
				}
			}

			t.Logf("Test case '%s' succeeded with input '%s'", tc.name, tc.input)
		})
	}
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {
	t.Log("Starting TestStringToInt...")

	tests := []struct {
		name        string
		input       string
		expected    int
		expectError bool
	}{

		{"ValidNumericString", "123", 123, false},

		{"LeadingAndTrailingSpaces", "   456   ", 456, false},

		{"NonNumericString", "abc123", 0, true},

		{"NegativeIntegerString", "-987", -987, false},

		{"EmptyStringInput", "", 0, true},

		{"ZeroStringInput", "0", 0, false},

		{"LargeValidNumericString", "2147483647", 2147483647, false},

		{"OutOfRangeStringInput", "9223372036854775808", 0, true},

		{"SpecialCharsAndWhitespaces", "12! #45", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered, failing test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var buf bytes.Buffer
			stdOut := os.Stdout
			os.Stdout = &buf
			defer func() { os.Stdout = stdOut }()

			var actual int
			if tt.expectError {
				defer func() {
					if r := recover(); r != nil {

						t.Logf("Expected error behavior for input '%s'. Error correctly handled.", tt.input)
					} else {
						t.Errorf("Expected panic for input '%s', but test did not panic.", tt.input)
					}
				}()
			}
			actual = stringToInt(tt.input)

			if !tt.expectError {
				if actual != tt.expected {
					t.Errorf("Test '%s' failed. Expected %d, got %d", tt.name, tt.expected, actual)
				} else {
					t.Logf("Test '%s' passed successfully. Expected and actual match: %d", tt.name, actual)
				}
			}
		})
	}

	t.Log("All scenarios tested for TestStringToInt.")
}

