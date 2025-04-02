package golang_calculator

import (
	fmt "fmt"
	os "os"
	reflect "reflect"
	debug "runtime/debug"
	strconv "strconv"
	testing "testing"
	strings "strings"
	bytes "bytes"
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
		expected      interface{}
		expectError   bool
		errorContains string
		expectedLog   string
	}

	testCases := []TestCase{
		{name: "Valid Numeric String", input: "123.45", expected: float64(123.45), expectError: false},
		{name: "Negative Numeric String", input: "-456.78", expected: float64(-456.78), expectError: false},
		{name: "Zero Numeric String", input: "0", expected: float64(0.0), expectError: false},
		{name: "Non-Numeric String", input: "hello", expected: "exit", expectError: true, errorContains: "invalid syntax"},
		{name: "Empty String", input: "", expected: "exit", expectError: true, errorContains: "invalid syntax"},
		{name: "Large Numeric String", input: "1e+308", expected: float64(1e+308), expectError: false},
		{name: "Special Float Value NaN", input: "NaN", expected: float64(NaN()), expectError: false},
		{name: "Special Float Value Infinity", input: "Infinity", expected: float64(Inf(1)), expectError: false},
		{name: "String with Extra Whitespace", input: "  123.456  ", expected: float64(123.456), expectError: false},
		{name: "Scientific Notation String", input: "6.022e23", expected: float64(6.022e23), expectError: false},
		{name: "High Precision Float String", input: "0.123456789012345678", expected: float64(0.12345678901234568), expectError: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered in test case \"%s\": %v\nStack: %s\n", tc.name, r, string(debug.Stack()))
					t.Fail()
				}
			}()

			var res interface{}
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			if !tc.expectError {

				res = stringToFloat64(tc.input)
				w.Close()
				os.Stdout = oldStdout
				if !reflect.DeepEqual(res, tc.expected) {
					t.Errorf("Test case \"%s\" failed. Expected: %v, Got: %v", tc.name, tc.expected, res)
				} else {
					t.Logf("Test case \"%s\" passed. Expected: %v, Got: %v", tc.name, tc.expected, res)
				}
			} else {

				defer func() { os.Stdout = oldStdout }()
				defer func() {
					if r := recover(); r != nil {
						res = "exit"
						w.Close()
						os.Stdout = oldStdout
						if !reflect.DeepEqual(res, tc.expected) {
							t.Errorf("Test case \"%s\" failed due to unexpected behavior. Expected exit behavior but observed panic recovery.", tc.name)
						} else {
							t.Logf("Test case \"%s\" passed. Expected behavior matches.", tc.name)
						}
					}
				}()
				stringToFloat64(tc.input)
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

	type TestCase struct {
		name   string
		input  string
		expect int
		wantErr bool
		exitCode int
	}

