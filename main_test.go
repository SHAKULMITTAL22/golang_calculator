package main

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	os "os"
	exec "os/exec"
	debug "runtime/debug"
	strconv "strconv"
	strings "strings"
	testing "testing"
)

/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64
*/
func TestStringToFloat64(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expected    float64
		expectNaN   bool
		expectInf   int
		expectError bool
		description string
	}{

		{
			name:        "Scenario 1: Valid Positive Integer",
			input:       "123",
			expected:    123.0,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of a simple positive integer string.",
		},

		{
			name:        "Scenario 2: Valid Negative Integer",
			input:       "-45",
			expected:    -45.0,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of a simple negative integer string.",
		},

		{
			name:        "Scenario 3: Valid Positive Float",
			input:       "3.14159",
			expected:    3.14159,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of a positive float string.",
		},

		{
			name:        "Scenario 4: Valid Negative Float",
			input:       "-0.005",
			expected:    -0.005,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of a negative float string.",
		},

		{
			name:        "Scenario 5: Valid Zero String",
			input:       "0",
			expected:    0.0,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of '0'.",
		},
		{
			name:        "Scenario 5: Valid Zero Float String",
			input:       "0.0",
			expected:    0.0,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of '0.0'.",
		},

		{
			name:        "Scenario 6: Valid Sci Notation (Pos Exp)",
			input:       "1.23e4",
			expected:    12300.0,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of scientific notation with positive exponent.",
		},
		{
			name:        "Scenario 6: Valid Sci Notation (Pos Exp Uppercase)",
			input:       "1.23E4",
			expected:    12300.0,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of scientific notation with uppercase E.",
		},

		{
			name:        "Scenario 7: Valid Sci Notation (Neg Exp)",
			input:       "5.67e-3",
			expected:    0.00567,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of scientific notation with negative exponent.",
		},
		{
			name:        "Scenario 7: Valid Sci Notation (Neg Exp Uppercase)",
			input:       "5.67E-3",
			expected:    0.00567,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct conversion of scientific notation with uppercase E and negative exponent.",
		},

		{
			name:        "Scenario 8: Leading/Trailing Whitespace",
			input:       "  78.9  ",
			expected:    78.9,
			expectNaN:   false,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct handling of leading/trailing whitespace.",
		},

		{
			name:        "Scenario 9: Invalid Non-Numeric String",
			input:       "not-a-number",
			expected:    0,
			expectNaN:   false,
			expectInf:   0,
			expectError: true,
			description: "Verifies error condition for non-numeric input (expects exit).",
		},

		{
			name:        "Scenario 10: Invalid Empty String",
			input:       "",
			expected:    0,
			expectNaN:   false,
			expectInf:   0,
			expectError: true,
			description: "Verifies error condition for empty input (expects exit).",
		},

		{
			name:        "Scenario 11: Invalid Multiple Decimals",
			input:       "1.2.3",
			expected:    0,
			expectNaN:   false,
			expectInf:   0,
			expectError: true,
			description: "Verifies error condition for multiple decimal points (expects exit).",
		},

		{
			name:        "Scenario 12: Valid Positive Infinity",
			input:       "Inf",
			expected:    0,
			expectNaN:   false,
			expectInf:   1,
			expectError: false,
			description: "Verifies correct parsing of 'Inf'.",
		},
		{
			name:        "Scenario 12: Valid Positive Infinity (+Inf)",
			input:       "+Inf",
			expected:    0,
			expectNaN:   false,
			expectInf:   1,
			expectError: false,
			description: "Verifies correct parsing of '+Inf'.",
		},

		{
			name:        "Scenario 13: Valid Negative Infinity",
			input:       "-Inf",
			expected:    0,
			expectNaN:   false,
			expectInf:   -1,
			expectError: false,
			description: "Verifies correct parsing of '-Inf'.",
		},

		{
			name:        "Scenario 14: Valid NaN",
			input:       "NaN",
			expected:    0,
			expectNaN:   true,
			expectInf:   0,
			expectError: false,
			description: "Verifies correct parsing of 'NaN'.",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {

					t.Logf("Panic encountered in test case '%s': %v\n%s", tc.name, r, string(debug.Stack()))
					t.Errorf("Test case '%s' panicked unexpectedly", tc.name)
				}
			}()

			t.Logf("Running Test Case: %s", tc.name)
			t.Logf("Description: %s", tc.description)
			t.Logf("Input: \"%s\"", tc.input)

			if tc.expectError {

				_, err := strconv.ParseFloat(tc.input, 64)
				if err == nil {
					t.Errorf("Expected strconv.ParseFloat to return an error for input '%s', but it did not. The function stringToFloat64 would not have exited as expected.", tc.input)
				} else {

					t.Logf("Success: strconv.ParseFloat correctly returned an error ('%v') for input '%s', which would trigger os.Exit in stringToFloat64.", err, tc.input)
				}
				return
			}

			actual := stringToFloat64(tc.input)

			if tc.expectNaN {
				if !math.IsNaN(actual) {
					t.Errorf("Expected NaN for input '%s', but got %v", tc.input, actual)
				} else {
					t.Logf("Success: Correctly parsed '%s' as NaN.", tc.input)
				}
			} else if tc.expectInf != 0 {
				if !math.IsInf(actual, tc.expectInf) {
					expectedInfStr := "+Inf"
					if tc.expectInf < 0 {
						expectedInfStr = "-Inf"
					}
					t.Errorf("Expected %s for input '%s', but got %v", expectedInfStr, tc.input, actual)
				} else {
					expectedInfStr := "+Inf"
					if tc.expectInf < 0 {
						expectedInfStr = "-Inf"
					}
					t.Logf("Success: Correctly parsed '%s' as %s.", tc.input, expectedInfStr)
				}
			} else {

				if actual != tc.expected {
					t.Errorf("Expected %v for input '%s', but got %v", tc.expected, tc.input, actual)
				} else {
					t.Logf("Success: Correctly parsed '%s' as %v.", tc.input, actual)
				}
			}
		})
	}
}

/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int
*/
func TestStringToInt(t *testing.T) {

	t.Run("SuccessfulConversions", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		testCases := []struct {
			name     string
			input    string
			expected int
			scenario string
		}{
			{
				name:     "Scenario 1: Positive Integer",
				input:    "123",
				expected: 123,
				scenario: "This test verifies the correct conversion of a standard positive integer string.",
			},
			{
				name:     "Scenario 2: Zero",
				input:    "0",
				expected: 0,
				scenario: "This test verifies the correct conversion of the string \"0\".",
			},
			{
				name:     "Scenario 3: Negative Integer",
				input:    "-45",
				expected: -45,
				scenario: "This test verifies the correct conversion of a valid negative integer string.",
			},
		}

		for _, tc := range testCases {

			t.Run(tc.name, func(t *testing.T) {
				t.Logf("Running Test Case: %s", tc.name)
				t.Logf("Scenario: %s", tc.scenario)
				t.Logf("Arrange: Input string = %q", tc.input)

				result := stringToInt(tc.input)
				t.Logf("Act: Called stringToInt(%q), got %d", tc.input, result)

				if result != tc.expected {
					t.Errorf("Assert: FAILED: For input %q, expected %d, but got %d", tc.input, tc.expected, result)
				} else {
					t.Logf("Assert: PASSED: Result %d matches expected value %d", result, tc.expected)
				}

			})
		}
	})

	t.Run("ErrorHandlingAndExit", func(t *testing.T) {

		errorTestCases := []struct {
			name                 string
			input                string
			expectedExitCode     int
			expectErrorOutput    bool
			expectedErrorSubstr  string
			scenario             string
			validationLogic      string
			validationImportance string
		}{
			{
				name:                 "Scenario 4: Non-Numeric String",
				input:                "abc",
				expectedExitCode:     2,
				expectErrorOutput:    true,
				expectedErrorSubstr:  "invalid syntax",
				scenario:             "Tests behavior with non-numeric input.",
				validationLogic:      "strconv.Atoi(\"abc\") fails, triggering the error path and os.Exit(2). Assertion checks the captured exit code and stderr.",
				validationImportance: "Crucial for validating robustness against invalid input and predictable termination.",
			},
			{
				name:                 "Scenario 5: Empty String",
				input:                "",
				expectedExitCode:     2,
				expectErrorOutput:    true,
				expectedErrorSubstr:  "invalid syntax",
				scenario:             "Tests behavior with an empty input string.",
				validationLogic:      "strconv.Atoi(\"\") fails, triggering the error path and os.Exit(2). Assertion checks the captured exit code and stderr.",
				validationImportance: "Tests handling of a common edge case (empty input).",
			},
			{
				name:                 "Scenario 6: String with Leading/Trailing Spaces",
				input:                " 123 ",
				expectedExitCode:     2,
				expectErrorOutput:    true,
				expectedErrorSubstr:  "invalid syntax",
				scenario:             "Tests behavior with leading/trailing whitespace.",
				validationLogic:      "strconv.Atoi fails if whitespace is present. Error path and os.Exit(2) should execute. Assertion checks exit code and stderr.",
				validationImportance: "Validates strict adherence to strconv.Atoi's expected format (no implicit trimming).",
			},
			{
				name:                 "Scenario 7: Floating-Point String",
				input:                "12.34",
				expectedExitCode:     2,
				expectErrorOutput:    true,
				expectedErrorSubstr:  "invalid syntax",
				scenario:             "Tests behavior with a floating-point string.",
				validationLogic:      "strconv.Atoi cannot parse decimal points. Error path and os.Exit(2) should execute. Assertion checks exit code and stderr.",
				validationImportance: "Ensures the function correctly rejects non-integer numeric formats.",
			},
			{
				name:                 "Scenario 8: Out of Range String (Overflow)",
				input:                "9223372036854775808",
				expectedExitCode:     2,
				expectErrorOutput:    true,
				expectedErrorSubstr:  "value out of range",
				scenario:             "Tests behavior with a number larger than max int64.",
				validationLogic:      "strconv.Atoi detects range errors. Error path and os.Exit(2) should execute. Assertion checks exit code and stderr.",
				validationImportance: "Validates handling of numerical limits, preventing potential issues if the error wasn't caught.",
			},
		}

		for _, tc := range errorTestCases {

			t.Run(tc.name, func(t *testing.T) {
				defer func() {
					if r := recover(); r != nil {

						t.Logf("Panic encountered in test setup/assertion. %v\n%s", r, string(debug.Stack()))
						t.FailNow()
					}
				}()

				t.Logf("Running Test Case: %s", tc.name)
				t.Logf("Scenario: %s", tc.scenario)
				t.Logf("Arrange: Input string = %q", tc.input)
				t.Logf("Arrange: Expecting exit code = %d", tc.expectedExitCode)
				if tc.expectErrorOutput {
					t.Logf("Arrange: Expecting error message containing = %q", tc.expectedErrorSubstr)
				}

				if os.Getenv("GO_TEST_SUBPROCESS_EXIT") == "1" {

					stringToInt(os.Getenv("SUBPROCESS_INPUT"))
					return
				}

				cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=^%s$", t.Name()))

				cmd.Env = append(os.Environ(),
					"GO_TEST_SUBPROCESS_EXIT=1",
					"SUBPROCESS_INPUT="+tc.input,
				)

				var stderr bytes.Buffer
				cmd.Stderr = &stderr
				var stdout bytes.Buffer
				cmd.Stdout = &stdout

				t.Logf("Act: Executing test binary as subprocess: %s with input %q", os.Args[0], tc.input)

				err := cmd.Run()
				stderrStr := stderr.String()
				stdoutStr := stdout.String()
				t.Logf("Act: Subprocess completed. Stderr:\n%s", stderrStr)
				t.Logf("Act: Subprocess completed. Stdout:\n%s", stdoutStr)

				exitErr, ok := err.(*exec.ExitError)
				if !ok {
					t.Errorf("Assert: FAILED: Expected command to exit with an error, but got %v", err)
					return
				}

				actualExitCode := exitErr.ExitCode()
				if actualExitCode != tc.expectedExitCode {
					t.Errorf("Assert: FAILED: Expected exit code %d, but got %d", tc.expectedExitCode, actualExitCode)
				} else {
					t.Logf("Assert: PASSED: Exit code %d matches expected value %d (Scenario 9 implicitly verified)", actualExitCode, tc.expectedExitCode)
				}

				outputToCheck := stdoutStr + stderrStr
				if tc.expectErrorOutput {
					if !strings.Contains(outputToCheck, tc.expectedErrorSubstr) {
						t.Errorf("Assert: FAILED: Expected output to contain %q, but it didn't. Full Output:\n%s", tc.expectedErrorSubstr, outputToCheck)
					} else {
						t.Logf("Assert: PASSED: Output contains expected error substring %q", tc.expectedErrorSubstr)
					}
				}

				t.Logf("Validation Logic: %s", tc.validationLogic)
				t.Logf("Validation Importance: %s", tc.validationImportance)
			})
		}
	})

}
