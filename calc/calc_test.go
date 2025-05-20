package calc

import testing "testing"

type testCase struct {
	name                 string
	input                float64
	expectedOutput       float64
	expectedPanicMessage string
}

/*
ROOST_METHOD_HASH=SquareRoot_17095d9165
ROOST_METHOD_SIG_HASH=SquareRoot_232943a56a

FUNCTION_DEF=func SquareRoot(num float64) float64 // Square root (with error handling)
*/
func TestSquareRoot(t *testing.T) {

	cases := []testCase{
		{
			name:           "Square root of a positive number",
			input:          25,
			expectedOutput: 5.0,
		},
		{
			name:           "Square root of zero",
			input:          0,
			expectedOutput: 0.0,
		},
		{
			name:                 "Square root of a negative number",
			input:                -4,
			expectedPanicMessage: "square root of a negative number is not defined",
		},
		{
			name:           "Square root of a non-integer number (float)",
			input:          2.25,
			expectedOutput: 1.5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					expected := c.expectedPanicMessage
					if r != expected {
						t.Errorf("Expected panic: %v, got: %v", expected, r)
					}
				}
			}()

			got := SquareRoot(c.input)
			expected := c.expectedOutput
			if got != expected {
				t.Errorf("SquareRoot(%f) == %f, want %f", c.input, got, expected)
			} else {
				t.Logf("Success: SquareRoot(%f) == %f", c.input, got)
			}
		})
	}
}
