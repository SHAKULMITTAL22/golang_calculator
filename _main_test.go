package golang_calculator

import (
	fmt "fmt"
	math "math"
	debug "runtime/debug"
	strconv "strconv"
	testing "testing"
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
		expectedResult float64
		expectError    bool
	}{

		{
			name:           "Scenario 1: Valid Positive Float",
			input:          "123.456",
			expectedResult: 123.456,
			expectError:    false,
		},

		{
			name:           "Scenario 2: Valid Negative Float",
			input:          "-98.7",
			expectedResult: -98.7,
			expectError:    false,
		},

		{
			name:           "Scenario 3: Valid Integer String",
			input:          "250",
			expectedResult: 250.0,
			expectError:    false,
		},

		{
			name:           "Scenario 4: String Representing Zero",
			input:          "0",
			expectedResult: 0.0,
			expectError:    false,
		},

		{
			name:           "Scenario 5: String with Whitespace",
			input:          "  75.5  ",
			expectedResult: 75.5,
			expectError:    false,
		},

		{
			name:           "Scenario 6: Scientific Notation Positive Exp",
			input:          "1.234e+3",
			expectedResult: 1234.0,
			expectError:    false,
		},

		{
			name:           "Scenario 7: Scientific Notation Negative Exp",
			input:          "5.67E-2",
			expectedResult: 0.0567,
			expectError:    false,
		},

		{
			name:        "Scenario 8: Invalid Non-Numeric String",
			input:       "not-a-number",
			expectError: true,
		},

		{
			name:        "Scenario 9: Empty String",
			input:       "",
			expectError: true,
		},

		{
			name:        "Scenario 10: Invalid Format (Multiple Decimals)",
			input:       "1.2.3",
			expectError: true,
		},

		{
			name:           "Edge Case: Positive Infinity",
			input:          "Inf",
			expectedResult: math.Inf(1),
			expectError:    false,
		},
		{
			name:           "Edge Case: Negative Infinity",
			input:          "-Inf",
			expectedResult: math.Inf(-1),
			expectError:    false,
		},
		{
			name:           "Edge Case: Max Float64",
			input:          fmt.Sprintf("%g", math.MaxFloat64),
			expectedResult: math.MaxFloat64,
			expectError:    false,
		},
		{
			name:           "Edge Case: Smallest Non-zero Positive Float64",
			input:          fmt.Sprintf("%g", math.SmallestNonzeroFloat64),
			expectedResult: math.SmallestNonzeroFloat64,
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered during test execution: %v\n%s", r, string(debug.Stack()))
					t.Errorf("FAIL: Test panicked unexpectedly for input %q", tc.input)
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input string: %q", tc.input)

			if tc.expectError {

				t.Log("Expecting an error condition (which leads to os.Exit in the original function)")

				_, err := strconv.ParseFloat(tc.input, 64)

				if err == nil {

					t.Errorf("FAIL: Input %q: Expected strconv.ParseFloat to return an error, but it succeeded.", tc.input)
				} else {

					t.Logf("PASS: Input %q: Correctly identified as invalid by strconv.ParseFloat, returning error: %v (expected behavior before os.Exit)", tc.input, err)
				}

			} else {

				t.Logf("Expecting result: %f", tc.expectedResult)

				actualResult := stringToFloat64(tc.input)

				if !almostEqual(actualResult, tc.expectedResult) {

					t.Errorf("FAIL: Input %q: Expected %f, but got %f", tc.input, tc.expectedResult, actualResult)
				} else {

					t.Logf("PASS: Input %q: Correctly converted to %f", tc.input, actualResult)
				}

				if math.IsInf(tc.expectedResult, 0) && !math.IsInf(actualResult, 0) {
					t.Errorf("FAIL: Input %q: Expected infinity (%f), but got a finite number %f", tc.input, tc.expectedResult, actualResult)
				}
				if !math.IsInf(tc.expectedResult, 0) && math.IsInf(actualResult, 0) {
					t.Errorf("FAIL: Input %q: Expected a finite number (%f), but got infinity %f", tc.input, tc.expectedResult, actualResult)
				}

				if math.IsNaN(tc.expectedResult) {
					if !math.IsNaN(actualResult) {
						t.Errorf("FAIL: Input %q: Expected NaN, but got %f", tc.input, actualResult)
					} else {
						t.Logf("PASS: Input %q: Correctly resulted in NaN", tc.input)
					}
				} else if math.IsNaN(actualResult) {
					t.Errorf("FAIL: Input %q: Expected %f, but got NaN", tc.input, tc.expectedResult)
				}
			}
		})
	}
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expectedInt int
		expectExit  bool
	}{

		{
			name:        "Scenario 1: Valid Positive Integer String Conversion",
			input:       "123",
			expectedInt: 123,
			expectExit:  false,
		},
		{
			name:        "Scenario 2: Valid Negative Integer String Conversion",
			input:       "-45",
			expectedInt: -45,
			expectExit:  false,
		},
		{
			name:        "Scenario 3: Valid Zero Integer String Conversion",
			input:       "0",
			expectedInt: 0,
			expectExit:  false,
		},

		{
			name:       "Scenario 4: Invalid Non-Numeric String Input",
			input:      "abc",
			expectExit: true,
		},
		{
			name:       "Scenario 5: Invalid Empty String Input",
			input:      "",
			expectExit: true,
		},
		{
			name:       "Scenario 6: Invalid String with Leading/Trailing Spaces",
			input:      " 123 ",
			expectExit: true,
		},
		{
			name:       "Scenario 6b: Invalid String with Leading Space",
			input:      " 456",
			expectExit: true,
		},
		{
			name:       "Scenario 6c: Invalid String with Trailing Space",
			input:      "789 ",
			expectExit: true,
		},
		{
			name:       "Scenario 7: Invalid Floating-Point String Input",
			input:      "1.23",
			expectExit: true,
		},
		{

			name:       "Scenario 8: String Representing Number Outside Integer Range (Overflow)",
			input:      "99999999999999999999",
			expectExit: true,
		},
		{
			name:       "Scenario 8: String Representing Number Outside Integer Range (Underflow)",
			input:      "-99999999999999999999",
			expectExit: true,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test case '%s' panicked unexpectedly.", tc.name)
				}
			}()

			t.Logf("Running test case: %s", tc.name)
			t.Logf("Input string: %q", tc.input)

			if tc.expectExit {

				t.Log("Expecting function to call os.Exit(2). Test runner may terminate.")

				_ = stringToInt(tc.input)

				t.Errorf("FAIL: Test case '%s' expected os.Exit(2) but the function returned.", tc.name)
				t.Logf("This indicates stringToInt did not exit as expected for input %q.", tc.input)

			} else {

				t.Logf("Expecting integer result: %d", tc.expectedInt)

				actualInt := stringToInt(tc.input)

				if actualInt != tc.expectedInt {
					t.Errorf("FAIL: Test case '%s' failed.", tc.name)
					t.Logf("Input: %q", tc.input)
					t.Logf("Expected integer: %d", tc.expectedInt)
					t.Logf("Actual integer: %d", actualInt)
				} else {
					t.Logf("PASS: Test case '%s' passed.", tc.name)
					t.Logf("Input %q correctly converted to %d.", tc.input, actualInt)
				}
			}
		})
	}
}

