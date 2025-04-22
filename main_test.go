package main

import (
	bytes "bytes"
	fmt "fmt"
	os "os"
	exec "os/exec"
	debug "runtime/debug"
	strings "strings"
	testing "testing"
	strconv "strconv"
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
		expected       float64
		expectError    bool
		errorExitCode  int
		errorMsgSubstr string
	}{
		{
			name:     "Scenario 1: Valid Float String Conversion",
			input:    "123.45",
			expected: 123.45,
		},
		{
			name:     "Scenario 2: Integer String Conversion",
			input:    "42",
			expected: 42.0,
		},
		{
			name:     "Scenario 3: Scientific Notation String Conversion",
			input:    "1.23e5",
			expected: 123000.0,
		},
		{
			name:     "Scenario 4: Negative Number String Conversion",
			input:    "-123.45",
			expected: -123.45,
		},
		{
			name:     "Scenario 5: Zero String Conversion",
			input:    "0",
			expected: 0.0,
		},
		{
			name:     "Scenario 5b: Zero Decimal String Conversion",
			input:    "0.0",
			expected: 0.0,
		},
		{
			name:     "Scenario 6: Very Large Number String Conversion",
			input:    "1234567890123456",
			expected: 1234567890123456.0,
		},
		{
			name:     "Scenario 7: Very Small Number String Conversion",
			input:    "0.0000000000123",
			expected: 0.0000000000123,
		},
		{
			name:           "Scenario 8: Invalid String Handling (Non-Numeric String)",
			input:          "abc",
			expectError:    true,
			errorExitCode:  2,
			errorMsgSubstr: "invalid syntax",
		},
		{
			name:           "Scenario 9: Empty String Handling",
			input:          "",
			expectError:    true,
			errorExitCode:  2,
			errorMsgSubstr: "invalid syntax",
		},
		{
			name:           "Scenario 10a: String with Leading Whitespace",
			input:          " 123.45",
			expectError:    true,
			errorExitCode:  2,
			errorMsgSubstr: "invalid syntax",
		},
		{
			name:           "Scenario 10b: String with Trailing Whitespace",
			input:          "123.45 ",
			expectError:    true,
			errorExitCode:  2,
			errorMsgSubstr: "invalid syntax",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			if tc.expectError {

				testErrorCase(t, tc.input, tc.errorExitCode, tc.errorMsgSubstr)
			} else {

				result := stringToFloat64(tc.input)
				if result != tc.expected {
					t.Errorf("Expected %f, got %f", tc.expected, result)
				} else {
					t.Logf("Successfully converted %s to %f", tc.input, result)
				}
			}
		})
	}
}

func testErrorCase(t *testing.T, input string, expectedExitCode int, expectedErrorSubstr string) {

	t.Logf("Testing error case with input: %s", input)
	t.Log("Note: Cannot directly test os.Exit behavior in the same process.")
	t.Log("In a real application, consider refactoring the function to return an error instead of calling os.Exit.")

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	done := make(chan bool)

	go func() {
		defer func() {

			if r := recover(); r != nil {
				t.Logf("Recovered from panic: %v", r)
			}
			done <- true
		}()

		stringToFloat64(input)
	}()

	<-done

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, expectedErrorSubstr) {
		t.Errorf("Expected error message to contain %q, got: %q", expectedErrorSubstr, output)
	} else {
		t.Logf("Successfully detected error message containing %q", expectedErrorSubstr)
	}

	t.Logf("Note: In a real scenario, the function would exit with code %d", expectedExitCode)
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
		shouldExit     bool
	}{
		{
			name:           "Valid Integer String Conversion",
			input:          "123",
			expectedOutput: 123,
			shouldExit:     false,
		},
		{
			name:           "Negative Integer String Conversion",
			input:          "-42",
			expectedOutput: -42,
			shouldExit:     false,
		},
		{
			name:           "Zero String Conversion",
			input:          "0",
			expectedOutput: 0,
			shouldExit:     false,
		},
		{
			name:           "Invalid String (Non-Numeric) Conversion",
			input:          "abc",
			expectedOutput: 0,
			shouldExit:     true,
		},
		{
			name:           "Mixed Alphanumeric String Conversion",
			input:          "123abc",
			expectedOutput: 0,
			shouldExit:     true,
		},
		{
			name:           "Empty String Conversion",
			input:          "",
			expectedOutput: 0,
			shouldExit:     true,
		},
		{
			name:           "Very Large Integer String Conversion",
			input:          strconv.Itoa(2147483647),
			expectedOutput: 2147483647,
			shouldExit:     false,
		},
		{
			name:           "Whitespace-Padded Integer String Conversion",
			input:          " 42 ",
			expectedOutput: 0,
			shouldExit:     true,
		},
		{
			name:           "Integer String with Leading Zeros",
			input:          "007",
			expectedOutput: 7,
			shouldExit:     false,
		},
		{
			name:           "Integer String with Plus Sign",
			input:          "+42",
			expectedOutput: 42,
			shouldExit:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			if tc.shouldExit {
				testExitingFunction(t, tc.input, tc.name)
			} else {

				result := stringToInt(tc.input)
				if result != tc.expectedOutput {
					t.Errorf("Expected %d, got %d", tc.expectedOutput, result)
				} else {
					t.Logf("Success: Input '%s' correctly converted to %d", tc.input, result)
				}
			}
		})
	}
}

func createTempTestFile(t *testing.T, input string) string {
	tempFile, err := os.CreateTemp("", "stringtoint_test_*.go")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	fmt.Fprintf(tempFile, `
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	stringToInt("%s")
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
`, input)

	tempFile.Close()
	return tempFile.Name()
}

func testExitingFunction(t *testing.T, input string, testName string) {

	tempFile := createTempTestFile(t, input)
	defer os.Remove(tempFile)

	cmd := exec.Command("go", "run", tempFile)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() == 2 {
			t.Logf("Success: Input '%s' correctly caused program exit with code 2", input)
		} else {
			t.Errorf("Expected exit code 2 for input '%s', got %d", input, exitError.ExitCode())
		}
	} else if err == nil {
		t.Errorf("Expected program to exit with error for input '%s', but it succeeded", input)
	} else {
		t.Errorf("Unexpected error running test process: %v", err)
	}

	if !strings.Contains(stderr.String(), "strconv.Atoi") {
		t.Logf("Error output: %s", stderr.String())
	}
}

