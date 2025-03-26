// ********RoostGPT********
/*

roost_feedback [3/26/2025, 8:17:19 AM]:Please identify the appropriate errors and make this file fully compilable with zero errors even if something needs to be commented or remove.\n\n
*/

// ********RoostGPT********

package golang_calculator

import (
	"fmt"
	"os"
	"strings"
	"testing"

	calc "github.com/SHAKULMITTAL22/golang_calculator/calc"
)

var osExit = os.Exit

// Updated TestStringToFloat64
func TestStringToFloat64(t *testing.T) {

	testCases := []struct {
		name          string
		input         string
		expected      float64
		expectPanic   bool
		expectedError string
	}{
		{
			name:        "Valid Float String Conversion",
			input:       "123.456",
			expected:    123.456,
			expectPanic: false,
		},
		{
			name:        "Conversion of Integer String as Float",
			input:       "100",
			expected:    100.0,
			expectPanic: false,
		},
		{
			name:          "Input of a Non-Numeric String",
			input:         "abc123",
			expectPanic:   true,
			expectedError: "strconv.ParseFloat: parsing 'abc123': invalid syntax",
		},
		{
			name:          "Handling Empty String Input",
			input:         "",
			expectPanic:   true,
			expectedError: "strconv.ParseFloat: parsing '': invalid syntax",
		},
		{
			name:        "Input with Leading Whitespaces",
			input:       "   456.789",
			expected:    456.789,
			expectPanic: false,
		},
		{
			name:        "Input with Trailing Whitespaces",
			input:       "789.321  ",
			expected:    789.321,
			expectPanic: false,
		},
		{
			name:        "Very Large Float String Input",
			input:       "1.7976931348623157e+308",
			expected:    1.7976931348623157e+308,
			expectPanic: false,
		},
		{
			name:        "Very Small Float String Input",
			input:       "5e-324",
			expected:    5e-324,
			expectPanic: false,
		},
		{
			name:          "Input with Invalid Characters Around a Number",
			input:         "123abc",
			expectPanic:   true,
			expectedError: "strconv.ParseFloat: parsing '123abc': invalid syntax",
		},
		{
			name:        "Scientific Notation String Conversion",
			input:       "1.23e4",
			expected:    12300,
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Fail()
					}
					if tc.expectedError != "" && !strings.Contains(fmt.Sprintf("%v", r), tc.expectedError) {
						t.Errorf("Error mismatch: got %v, expected %v", r, tc.expectedError)
					}
				}
			}()

			if !tc.expectPanic {
				actual := calc.StringToFloat64(tc.input)
				if actual != tc.expected {
					t.Errorf("Actual value mismatch: got %v, expected %v", actual, tc.expected)
				}
			}
		})
	}
}

// Updated TestStringToInt
func TestStringToInt(t *testing.T) {

	type testCase struct {
		name           string
		input          string
		expectedResult int
		expectError    bool
		exitCode       int
	}

	testCases := []testCase{
		{
			name:           "Valid Integer String Conversion",
			input:          "123",
			expectedResult: 123,
			expectError:    false,
		},
		{
			name:        "Empty String Conversion",
			input:       "",
			expectError: true,
			exitCode:    2,
		},
		{
			name:        "Non-Numeric String Conversion",
			input:       "abcd",
			expectError: true,
			exitCode:    2,
		},
		{
			name:           "String with Leading Zeros",
			input:          "007",
			expectedResult: 7,
			expectError:    false,
		},
		{
			name:           "Negative Integer String",
			input:          "-45",
			expectedResult: -45,
			expectError:    false,
		},
		{
			name:           "Large Integer String",
			input:          "2147483647",
			expectedResult: 2147483647,
			expectError:    false,
		},
		{
			name:        "Overflow Integer String",
			input:       "9223372036854775808",
			expectError: true,
			exitCode:    2,
		},
		{
			name:           "String with Whitespace Characters",
			input:          "   456   ",
			expectedResult: 456,
			expectError:    false,
		},
		{
			name:           "Zero String",
			input:          "0",
			expectedResult: 0,
			expectError:    false,
		},
		{
			name:        "Mixed Alphanumeric String Conversion",
			input:       "12a45",
			expectError: true,
			exitCode:    2,
		},
		{
			name:        "Multi-Line String Input",
			input:       "123\n456",
			expectError: true,
			exitCode:    2,
		},
		{
			name:        "Unicode Characters in String",
			input:       "123漢字",
			expectError: true,
			exitCode:    2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					if !tc.expectError {
						t.Errorf("Unexpected panic during test: %v", r)
					}
				}
			}()

			if !tc.expectError {
				result := calc.StringToInt(tc.input)
				if result != tc.expectedResult {
					t.Errorf("Expected result %d but got %d for test case: %v", tc.expectedResult, result, tc.name)
				}
			}
		})
	}
}
