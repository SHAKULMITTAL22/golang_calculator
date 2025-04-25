package golang_calculator

import (
	fmt "fmt"
	os "os"
	debug "runtime/debug"
	strings "strings"
	testing "testing"
	calc "github.com/SHAKULMITTAL22/golang_calculator/calc"
	bytes "bytes"
)



var osExit = os.Exit




/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
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
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					if !tc.expectPanic {
						t.Fail()
					}
					if tc.expectedError != "" && !strings.Contains(fmt.Sprintf("%v", r), tc.expectedError) {
						t.Errorf("Error mismatch: got %v, expected %v", r, tc.expectedError)
					}
				}
			}()

			rPipe, wPipe, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe for stdout redirection: %v", err)
			}
			oldStdout := os.Stdout
			os.Stdout = wPipe

			defer func() {
				os.Stdout = oldStdout
				wPipe.Close()
				rPipe.Close()
			}()

			var actual float64
			if !tc.expectPanic {
				actual = calc.StringToFloat64(tc.input)
			}

			wPipe.Close()
			output := &strings.Builder{}
			_, _ = fmt.Fscanf(rPipe, "%s", output)

			if !tc.expectPanic {
				if actual != tc.expected {
					t.Errorf("Actual value mismatch: got %v, expected %v", actual, tc.expected)
				}
			} else {
				if len(output.String()) == 0 {
					t.Errorf("Expected error not logged for input %v", tc.input)
				}
			}

			t.Logf("Test case '%s' passed successfully.", tc.name)
		})
	}
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
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
					t.Logf("Panic encountered during test - failing test: %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var buf bytes.Buffer
			stdout := os.Stdout
			os.Stdout = &buf
			defer func() { os.Stdout = stdout }()

			exitCode := -1
			osExit = func(code int) {
				exitCode = code
				panic("os.Exit called")
			}
			defer func() { osExit = os.Exit }()

			var result int
			var err error
			if tc.expectError {
				defer func() {

					if r := recover(); r == nil {
						t.Errorf("Expected os.Exit to be called but it wasn't for test case: %v", tc.name)
					} else if exitCode != tc.exitCode {
						t.Errorf("Expected exit code %d but got %d for test case: %v", tc.exitCode, exitCode, tc.name)
					}
				}()
			}
			result = stringToInt(tc.input)

			if !tc.expectError {
				if result != tc.expectedResult {
					t.Errorf("Expected result %d but got %d for test case: %v", tc.expectedResult, result, tc.name)
				}
			} else {

				output := buf.String()
				t.Logf("Captured function output: %v", output)
				if exitCode != tc.exitCode {
					t.Errorf("Expected exit code %d but got %d for test case: %v", tc.exitCode, exitCode, tc.name)
				}
			}
		})
	}
}

