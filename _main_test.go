package golang_calculator

import (
	fmt "fmt"
	testing "testing"
	require "github.com/stretchr/testify/require"
	strconv "strconv"
)








/*
ROOST_METHOD_HASH=stringToFloat64_d38659cd50
ROOST_METHOD_SIG_HASH=stringToFloat64_44e80853e6

FUNCTION_DEF=func stringToFloat64(str string) float64 

*/
func TestStringToFloat64(t *testing.T) {

	testCases := []struct {
		name          string
		inputStr      string
		expectedFloat float64
		shouldError   bool
	}{
		{
			name:          "Scenario 1: Valid String to Float Conversion",
			inputStr:      "42.00",
			expectedFloat: 42.00,
			shouldError:   false,
		},
		{
			name:        "Scenario 2: Invalid String to Float Conversion",
			inputStr:    "abc",
			shouldError: true,
		},
		{
			name:          "Scenario 3: Conversion of String Containing Large Float to Float64",
			inputStr:      "1.0e50",
			expectedFloat: 1.0e50,
			shouldError:   false,
		},
		{
			name:          "Scenario 4: Conversion of Empty String to Float64",
			inputStr:      "",
			expectedFloat: 0,
			shouldError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Recovered from panic: %v", r)
					t.Fail()
				}
			}()

			res := stringToFloat64(tc.inputStr)

			if tc.shouldError {
				if res != nil {
					t.Logf("Failed %s: expected nil, got %v", tc.name, res)
					t.Fail()
				}
				return
			}

			require.Equal(t, tc.expectedFloat, res, fmt.Sprintf("Failed Test %s: expected %v, got %v", tc.name, tc.expectedFloat, res))

		})
	}
}


/*
ROOST_METHOD_HASH=stringToInt_73b9cbccee
ROOST_METHOD_SIG_HASH=stringToInt_e7cc66ec50

FUNCTION_DEF=func stringToInt(str string) int 

*/
func TestStringToInt(t *testing.T) {
	var tests = []struct {
		name, input string
		want        int
		errExpected bool
	}{
		{"Valid conversion", "123", 123, false},
		{"Invalid conversion", "hello", 0, true},
		{"Boundary value conversion - lower", "-2147483648", -2147483648, false},
		{"Boundary value conversion - upper", "2147483647", 2147483647, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Test case panicked: %v", tt.name)
					t.Fail()
				}
			}()

			r, err := strconv.Atoi(tt.input)

			if (err != nil) != tt.errExpected {
				t.Fatalf("Error difference: got %v, want %v", err, tt.errExpected)
			}
			if r != tt.want {
				t.Errorf("Conversion result difference: got %v, want %v", r, tt.want)
			}
		})
	}
}

