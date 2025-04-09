package main

import (
	bytes "bytes"
	fmt "fmt"
	os "os"
	strconv "strconv"
	strings "strings"
	testing "testing"
	sync "sync"
	debug "runtime/debug"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedValue float64
		expectedError string
		expectedExit  bool
	}

	tests := []testCase{
		{name: "Valid numeric string", input: "12.34", expectedValue: 12.34},
		{name: "Invalid numeric string", input: "abc", expectedError: strconv.ErrSyntax.Error(), expectedExit: true},
		{name: "Integer-like numeric string", input: "42", expectedValue: 42.0},
		{name: "Fractional number with high precision", input: "0.123456789", expectedValue: 0.123456789},
		{name: "Leading/trailing whitespace string", input: "   23.5   ", expectedValue: 23.5},
		{name: "Very large numeric string", input: "1e20", expectedValue: 1e20},
		{name: "Zero numeric string", input: "0", expectedValue: 0.0},
		{name: "Negative numeric string", input: "-45.67", expectedValue: -45.67},
		{name: "Empty string", input: "", expectedError: strconv.ErrSyntax.Error(), expectedExit: true},
		{name: "Special character string", input: "_@%$", expectedError: strconv.ErrSyntax.Error(), expectedExit: true},
	}

	originalExit := os.Exit
	defer func() { os.Exit = originalExit }()

	var wg sync.WaitGroup
	for _, tc := range tests {
		wg.Add(1)
		t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var stdout bytes.Buffer
			fmt.Fprintf(&stdout, "")

			exitCalled := false
			exitCode := -1
			os.Exit = func(code int) {
				exitCalled = true
				exitCode = code
			}

			var result float64
			var caughtError string
			func() {
				defer func() {
					if r := recover(); r != nil {
						if exitCalled && tc.expectedExit {
							t.Logf("Exit confirmed with code %d", exitCode)
						} else {
							t.Logf("Unexpected panic occurred: %v", r)
							t.Fail()
						}
					}
				}()

				result = stringToFloat64(tc.input)
				caughtError = stdout.String()
			}()

			if tc.expectedExit {
				if !exitCalled {
					t.Errorf("Expected exit but no Exit call occurred")
				}
				if !strings.Contains(caughtError, tc.expectedError) {
					t.Errorf("Expected error '%v' but got '%v'", tc.expectedError, caughtError)
				}
			} else {
				if exitCalled {
					t.Errorf("Unexpected exit occurred for input '%v'", tc.input)
				}
				if result != tc.expectedValue {
					t.Errorf("Expected value %v but got %v for input '%v'", tc.expectedValue, result, tc.input)
				}
				t.Logf("Test '%v' successful with value: %v", tc.name, result)
			}
		})
	}

	wg.Wait()
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {

	var outputBuffer bytes.Buffer
	os.Stdout = &outputBuffer

	tests := []struct {
		name          string
		input         string
		expected      int
		shouldExit    bool
		expectedError string
	}{
		{"Valid Integer String", "123", 123, false, ""},
		{"Zero in String Format", "0", 0, false, ""},
		{"Negative Integer String", "-123", -123, false, ""},
		{"Non-Numeric String", "abc", 0, true, "strconv.Atoi: parsing \"abc\": invalid syntax"},
		{"Empty String Input", "", 0, true, "strconv.Atoi: parsing \"\": invalid syntax"},
		{"Whitespaced String", " 123 ", 0, true, "strconv.Atoi: parsing \" 123 \": invalid syntax"},
		{"Very Large Integer", "9223372036854775807", 9223372036854775807, false, ""},
		{"Out-of-Bounds Integer", "99999999999999999999", 0, true, "strconv.Atoi: parsing \"99999999999999999999\": value out of range"},
		{"String with Special Characters", "1#3", 0, true, "strconv.Atoi: parsing \"1#3\": invalid syntax"},
		{"Graceful Termination Check", "xyz", 0, true, "strconv.Atoi: parsing \"xyz\": invalid syntax"},
		{"Floating Point String", "12.34", 0, true, "strconv.Atoi: parsing \"12.34\": invalid syntax"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					t.Fail()
				}
			}()

			outputBuffer.Reset()
			expectedExit := tt.shouldExit
			exitCode := 0
			mockExit := func(code int) {
				exitCode = code
				panic(fmt.Sprintf("os.Exit(%d)", code))
			}
			osExitBackup := os.Exit
			os.Exit = mockExit
			defer func() { os.Exit = osExitBackup }()

			var result int
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("%v", r)
					}
				}()
				result = stringToInt(tt.input)
			}()

			if expectedExit {
				if err == nil {
					t.Errorf("Expected os.Exit but function completed execution for test case: %s", tt.name)
				} else {

					if outputBuffer.String() != tt.expectedError+"\n" {
						t.Errorf("Expected error output '%s' but got '%s'", tt.expectedError, outputBuffer.String())
					}

					if exitCode != 2 {
						t.Errorf("Expected exit code 2 but got %d for test case: %s", exitCode, tt.name)
					}

					t.Logf("Test passed with error handling: %s", tt.name)
				}
			} else {
				if result != tt.expected {
					t.Errorf("Expected output %d but got %d for input: %s", tt.expected, result, tt.input)
				}

				t.Logf("Test passed for successful integer conversion: %s", tt.name)
			}
		})
	}
}

