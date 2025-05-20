package calc

import (
	math "math"
	debug "runtime/debug"
	strconv "strconv"
	testing "testing"
)

type squareRootTestCase struct {
	input    float64
	expected float64
	isPanic  bool
}

func TestAbsolute(t *testing.T) {
	scenarios := []struct {
		desc            string
		num             float64
		expected_result float64
	}{
		{
			desc:            "Positive Number Test",
			num:             12.34,
			expected_result: 12.34,
		},
		{
			desc:            "Negative Number Test",
			num:             -12.34,
			expected_result: 12.34,
		},
		{
			desc:            "Zero Number Test",
			num:             0,
			expected_result: 0,
		},
		{
			desc:            "Large Number Test",
			num:             math.MaxFloat64,
			expected_result: math.MaxFloat64,
		},
		{
			desc:            "Very Small Negative Number Test",
			num:             -0.0000001,
			expected_result: 0.0000001,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution. %v\n%s", r, debug.Stack())
					t.Fail()
				}
			}()

			actual_result := Absolute(s.num)
			if actual_result != s.expected_result {
				t.Errorf("absolute value of %.8f should be %.8f, got %.8f", s.num, s.expected_result, actual_result)
			}

			t.Logf("Success: Expected absolute value of %.8f to be %.8f and got %.8f", s.num, s.expected_result, actual_result)
		})

	}

}

func TestSquareRoot(t *testing.T) {

	testcases := []squareRootTestCase{
		{input: 16, expected: 4},
		{input: 0, expected: 0},
		{input: -4, isPanic: true},
		{input: math.MaxFloat64, expected: math.Sqrt(math.MaxFloat64)},
		{input: math.NaN(), expected: math.NaN()},
	}

	for i, test := range testcases {

		t.Run("TestCaseNumber"+strconv.Itoa(i), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					if !test.isPanic {
						t.Fail()
					}
				}
			}()

			result := SquareRoot(test.input)

			if !math.IsNaN(result) && !math.IsNaN(test.expected) && result != test.expected {
				t.Logf("Fail: Expected %v but got %v", test.expected, result)
				t.Fail()
			} else if math.IsNaN(result) != math.IsNaN(test.expected) {
				t.Logf("Fail: Expected %v but got %v", test.expected, result)
				t.Fail()
			} else {
				t.Logf("Success: Expected %v and got %v", test.expected, result)
			}
		})
	}
}
