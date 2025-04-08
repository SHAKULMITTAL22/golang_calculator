package main

import (
	bytes "bytes"
	os "os"
	debug "runtime/debug"
	strings "strings"
	testing "testing"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {

	type testCase struct {
		name        string
		input       string
		expected    float64
		shouldError bool
	}

	testCases := []testCase{
		{"ValidFloatString", "123.45", 123.45, false},
		{"IntegerString", "42", 42.0, false},
		{"NegativeFloatString", "-123.45", -123.45, false},
		{"ZeroString", "0", 0.0, false},
		{"EmptyString", "", 0, true},
		{"NonNumericString", "abc", 0, true},
		{"MixedString", "123abc", 0, true},
		{"ExponentString", "1.23e4", 12300.0, false},
		{"WhitespaceString", "   123.45   ", 123.45, false},
		{"VeryLargeNumber", "1e308", 1e308, false},
		{"VerySmallNumber", "1e-308", 1e-308, false},
		{"NaNString", "NaN", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {

				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var stdout, stderr bytes.Buffer
			os.Stdout = &stdout
			os.Stderr = &stderr
			exitCode := 0

			oldExit := os.Exit
			os.Exit = func(code int) {
				exitCode = code
			}
			defer func() { os.Exit = oldExit }()

			result := 0.0
			hasExited := false
			func() {
				defer func() {
					if r := recover(); r != nil {
						hasExited = true
					}
				}()
				result = stringToFloat64(tc.input)
			}()

			os.Stdout = os.Stdout
			os.Stderr = os.Stderr

			if tc.shouldError {
				if !hasExited {
					t.Errorf("Expected an error or program exit for input '%s'", tc.input)
				}
				if !strings.Contains(stderr.String(), "parsing") {
					t.Errorf("Expected an error message in stderr for input '%s'", tc.input)
				}
			} else {
				if hasExited {
					t.Errorf("Unexpected exit for input '%s'", tc.input)
				}
				if result != tc.expected {
					t.Errorf("For input '%s': expected %f, got %f", tc.input, tc.expected, result)
				}
			}

			t.Logf("Test %s passed for input '%s'", tc.name, tc.input)
		})
	}

}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {

	testCases := []struct {
		name           string
		input          string
		expectedOutput int
		shouldPanic    bool
	}{
		{"Valid integer string", "123", 123, false},
		{"Negative integer string", "-456", -456, false},
		{"String with leading zeros", "007", 7, false},
		{"Empty string", "", 0, true},
		{"Non-numeric string", "abc", 0, true},
		{"Whitespace-only string", "   ", 0, true},
		{"Very large integer", "9223372036854775807", 9223372036854775807, false},
		{"Out-of-range integer", "9223372036854775808", 0, true},
		{"Decimal string input", "3.14", 0, true},
		{"Special character input", "@#$%", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var buf bytes.Buffer
			stdout := os.Stdout
			os.Stdout = &buf
			defer func() {
				os.Stdout = stdout
			}()

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. Panic: %v\n%s", r, string(debug.Stack()))
					if !tc.shouldPanic {
						t.Fatalf("Test case failed unexpectedly; did not expect panic for input: '%s'", tc.input)
					}
				}
			}()

			if !tc.shouldPanic {
				output := stringToInt(tc.input)
				if output != tc.expectedOutput {
					t.Errorf("Failed: expected %d, got %d, input: '%s'", tc.expectedOutput, output, tc.input)
				} else {
					t.Logf("Passed: expected %d, got %d, input: '%s'", tc.expectedOutput, output, tc.input)
				}
			} else {

				func() { _ = stringToInt(tc.input) }()
				t.Logf("Passed (expected panic): input: '%s'", tc.input)
			}

			if tc.shouldPanic {
				if buf.Len() == 0 {
					t.Errorf("Expected an error log for input '%s', but none was recorded", tc.input)
				}
			}
		})
	}
}

