package calc

import (
	fmt "fmt"
	math "math"
	testing "testing"
	os "os"
)

const Tolerance = 1e-9

var testTable = []struct {
	a        int
	b        int
	expected int
}{
	{6, 8, 24},
	{0, 4, 0},
	{4, 4, 4},
	{-3, -6, 6},
}

type TestCase struct {
	Input    float64
	Expected float64
	Scenario string
}
type scenario struct {
		description string
		num1        float64
		num2        float64
		expect      float64
	}
type testCase struct {
	input          float64
	expectedResult float64
	expectPanic    bool
	desc           string
}


/*
ROOST_METHOD_HASH=Absolute_d231f0ab10
ROOST_METHOD_SIG_HASH=Absolute_ec3c06e5a3

FUNCTION_DEF=func Absolute(num float64) float64 // Absolute value


*/
func TestAbsolute(t *testing.T) {

	testCases := []TestCase{
		{Input: 20, Expected: 20, Scenario: "Testing the Absolute function with a positive integer"},
		{Input: -30, Expected: 30, Scenario: "Testing the Absolute function with a negative integer"},
		{Input: 3.5, Expected: 3.5, Scenario: "Testing the Absolute function with a decimal number"},
		{Input: -2.5, Expected: 2.5, Scenario: "Testing the Absolute function with a negative decimal number"},
		{Input: 0, Expected: 0, Scenario: "Testing the Absolute function with zero"},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			t.Logf("Executing: %s", testCase.Scenario)
			result := Absolute(testCase.Input)

			if math.Abs(result-testCase.Expected) > 1e-9 {
				t.Errorf("For: %.2f, Expected: %.2f, but got: %.2f", testCase.Input, testCase.Expected, result)
			} else {
				t.Logf("For: %.2f, Expected: %.2f, Success", testCase.Input, testCase.Expected)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Add_61b83cbd8d
ROOST_METHOD_SIG_HASH=Add_9b37ae6611

FUNCTION_DEF=func Add(num1, num2 int) int // Add two integers


*/
func TestAdd(t *testing.T) {

	testCases := []struct {
		num1        int
		num2        int
		expected    int
		description string
	}{
		{2, 3, 5, "Testing Addition of Positive Numbers"},
		{-1, -3, -4, "Testing Addition of Negative Numbers"},
		{0, 5, 5, "Testing Addition of Zero"},
		{2147483647, 2147483647, 4294967294, "Testing Addition of Large Numbers"},
		{math.MaxInt64, 1, math.MinInt64, "Testing Addition with Maximum Integer Limit"},
	}

	for _, tc := range testCases {

		t.Run(tc.description, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered during test execution: %v ", r)
					t.Fail()
				}
			}()

			got := Add(tc.num1, tc.num2)

			if got != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.num1, tc.num2, got, tc.expected)
			} else {
				t.Logf("Success: %s. For inputs (%d, %d) function returned output as expected: %d", tc.description, tc.num1, tc.num2, tc.expected)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Divide_6fe509f399
ROOST_METHOD_SIG_HASH=Divide_d926fccfc9

FUNCTION_DEF=func Divide(num1, num2 float64) float64 // Divide two floating-point numbers (with error handling)


*/
func TestDivide(t *testing.T) {

	type test struct {
		num1     float64
		num2     float64
		expected float64
		name     string
		hasPanic bool
	}

	var tests = []test{
		{4.0, 2.0, 2.0, "Accurate Division When Both The Numerator And Denominator Are Positive", false},
		{-4.0, -2.0, 2.0, "Accurate Division When Both The Numerator And Denominator Are Negative", false},
		{4.0, 0.0, 0.0, "Error Handling When The Denominator Is Zero", true},
		{2.0, 3.0, 0.6666666666666666, "Precision Of The Result", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. the reason of panic is %v\n", r)
					t.Fail()
				}
			}()

			res := Divide(tt.num1, tt.num2)

			if tt.hasPanic {
				t.Errorf("Divide() does not panic")
				return
			}

			diff := math.Abs(res - tt.expected)
			if diff > 1e-9 {
				t.Errorf("Divide() did not return the correct value, got: %v, want: %v, with difference of %v", res, tt.expected, diff)
			} else {
				t.Logf("Divide() test passed for %v, got: %v, expected: %v, with difference of %v", tt.name, res, tt.expected, diff)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Logarithm_4092f1cba7
ROOST_METHOD_SIG_HASH=Logarithm_0780d00fe8

FUNCTION_DEF=func Logarithm(num, base float64) float64 // Logarithm function (log_base of num)


*/
func TestLogarithm(t *testing.T) {

	tests := []struct {
		name           string
		num, base      float64
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "Valid parameters",
			num:            100,
			base:           10,
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "Zero parameters",
			num:            0,
			base:           0,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "Base parameter equal to 1",
			num:            50,
			base:           1,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "Negative parameters",
			num:            -20,
			base:           -4,
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, test := range tests {

		t.Logf("Executing %s test case:\nnum=%f\nbase=%f\n", test.name, test.num, test.base)

		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered on %s test. Validate if error is actually expected and not false positive: %v", test.name, r)
					if test.expectError == false {
						t.Fail()
					}
				}
			}()

			result := Logarithm(test.num, test.base)

			if (math.Abs(result-test.expectedResult) > Tolerance) && test.expectError == false {
				t.Fatalf("Expected result was %f, but received %f", test.expectedResult, result)
			}
			t.Log("Test case passed successfully!")
		})
	}
}


/*
ROOST_METHOD_HASH=Modulo_7e9e651e69
ROOST_METHOD_SIG_HASH=Modulo_502e1458a3

FUNCTION_DEF=func Modulo(num1, num2 int) int // Modulo operation


*/
func TestModulo(t *testing.T) {

	tests := []struct {
		name        string
		num1        int
		num2        int
		want        int
		shouldPanic bool
	}{
		{
			name:        "Positive Dividend, Positive Divisor",
			num1:        10,
			num2:        3,
			want:        1,
			shouldPanic: false,
		},
		{
			name:        "Positive Dividend, Negative Divisor",
			num1:        10,
			num2:        -3,
			want:        -1,
			shouldPanic: false,
		},
		{
			name:        "Negative Dividend, Positive Divisor",
			num1:        -10,
			num2:        3,
			want:        2,
			shouldPanic: false,
		},
		{
			name:        "Dividend is Zero",
			num1:        0,
			num2:        3,
			want:        0,
			shouldPanic: false,
		},
		{
			name:        "Divisor is Zero",
			num1:        10,
			num2:        0,
			want:        0,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.shouldPanic {
					t.Errorf("Modulo() = panic occurred unexpectedly: %v", r)
				}
			}()

			got := Modulo(tt.num1, tt.num2)

			if !tt.shouldPanic && got != tt.want {
				t.Errorf("Modulo() = %v, want %v", got, tt.want)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Multiply_7a2824e2c7
ROOST_METHOD_SIG_HASH=Multiply_0911ef76c1

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 // Multiply two floating-point numbers


*/
func TestMultiply(t *testing.T) {
	type scenario struct {
		description string
		num1        float64
		num2        float64
		expect      float64
	}

	scenarios := []scenario{
		{"Regular Positive Inputs", 1.2, 2.3, 2.76},
		{"Zero Inputs", 0.0, 3.2, 0.0},
		{"Negative Inputs", -3.0, -2.0, 6.0},
		{"Combination of Positive and Negative Inputs", -5.5, 1.0, -5.5},

		{"Large Inputs", 1e308, 1e308, math.Inf(1)},
	}

	for _, s := range scenarios {
		t.Run(s.description, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Panic encountered in %s. %v", s.description, r)
					t.Fail()
				}
			}()
			got := Multiply(s.num1, s.num2)
			if got != s.expect {
				t.Errorf("Test failed for %s: got %v, expected %v.", s.description, got, s.expect)
			} else {
				t.Logf("Test passed for %s, got expected value: %v", s.description, got)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Power_1c67a5d8b5
ROOST_METHOD_SIG_HASH=Power_c74b8edd76

FUNCTION_DEF=func Power(base, exponent float64) float64 // Power function


*/
func TestPower(t *testing.T) {

	testScenarios := []struct {
		desc     string
		base     float64
		exponent float64
		expected float64
	}{
		{
			desc:     "Testing Power Function with Positive Integer values",
			base:     2,
			exponent: 3,
			expected: 8,
		},
		{
			desc:     "Testing Power Function with Negative Integer values",
			base:     -2,
			exponent: -3,
			expected: 0.125,
		},
		{
			desc:     "Testing Power Function with Zero values",
			base:     0,
			exponent: 0,
			expected: 1,
		},
	}

	for i := range testScenarios {
		scenario := testScenarios[i]

		t.Run(fmt.Sprintf("Scenario %d: %s", i+1, scenario.desc), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			actual := Power(scenario.base, scenario.exponent)
			if actual != scenario.expected {
				t.Errorf("Expected: '%v', got: '%v'", scenario.expected, actual)
			} else {
				t.Logf("Success: Scenario '%v' passed!", scenario.desc)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=SinCosTan_c242c1aa6d
ROOST_METHOD_SIG_HASH=SinCosTan_0f509380d6

FUNCTION_DEF=func SinCosTan(angle float64) (sin, cos, tan float64) // Trigonometric functions (Sin, Cos, Tan)


*/
func TestSinCosTan(t *testing.T) {
	scenarios := []scenario{
		{
			desc:      "Testing Trigonometric Functions of Zero Degree",
			input:     0,
			wantSin:   0,
			wantCos:   1,
			wantTan:   0,
			tolerance: 1e-10,
		},
		{
			desc:      "Testing Trigonometric Functions of 90 Degrees",
			input:     math.Pi / 2,
			wantSin:   1,
			wantCos:   0,
			wantTan:   "Infinity",
			tolerance: 1e-10,
		},
		{
			desc:      "Testing Trigonometric Functions of Negative Values",
			input:     -math.Pi / 2,
			wantSin:   -1,
			wantCos:   0,
			wantTan:   "-Infinity",
			tolerance: 1e-10,
		},
		{
			desc:      "Testing The Precision of the Trigonometric Functions",
			input:     1.23456789,
			wantSin:   math.Sin(1.23456789),
			wantCos:   math.Cos(1.23456789),
			wantTan:   math.Tan(1.23456789),
			tolerance: 1e-10,
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Panic encountered during test '%s'. %v", s.desc, r)
				}
			}()

			gotSin, gotCos, gotTan := SinCosTan(s.input)

			if diff := math.Abs(s.wantSin - gotSin); diff > s.tolerance {
				t.Errorf("Test '%s' failed, Sin value: got %f, want %f", s.desc, gotSin, s.wantSin)
			}

			if diff := math.Abs(s.wantCos - gotCos); diff > s.tolerance {
				t.Errorf("Test '%s' failed, Cos value: got %f, want %f", s.desc, gotCos, s.wantCos)
			}

			switch expected := s.wantTan.(type) {
			case float64:
				if diff := math.Abs(expected - gotTan); diff > s.tolerance {
					t.Errorf("Test '%s' failed, Tan value: got %f, want %f", s.desc, gotTan, expected)
				}
			case string:

				if (expected == "Infinity" && !math.IsInf(gotTan, 1)) || (expected == "-Infinity" && !math.IsInf(gotTan, -1)) {
					t.Errorf("Test '%s' failed expected: %s but got Tan value: %f", s.desc, expected, gotTan)
				}
			default:
				t.Errorf("Invalid type for wantTan in scenario: %s", s.desc)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=SquareRoot_17095d9165
ROOST_METHOD_SIG_HASH=SquareRoot_232943a56a

FUNCTION_DEF=func SquareRoot(num float64) float64 // Square root (with error handling)


*/
func TestSquareRoot(t *testing.T) {

	testCases := []testCase{
		{input: 100, expectedResult: 10, expectPanic: false, desc: "Valid Positive Number"},
		{input: 0, expectedResult: 0, expectPanic: false, desc: "Zero"},
		{input: -1, expectedResult: 0, expectPanic: true, desc: "Negative Input Handling"},
		{input: 1e14, expectedResult: math.Sqrt(1e14), expectPanic: false, desc: "Large Number Input"},
		{input: 1e-14, expectedResult: math.Sqrt(1e-14), expectPanic: false, desc: "Small Number Input"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if testCase.expectPanic {
						t.Logf("Panic occurred as expected: %v", r)
					} else {
						t.Errorf("Panic occurred but not expected: %v", r)
					}
				}
			}()

			result := SquareRoot(testCase.input)

			if !testCase.expectPanic {
				if result != testCase.expectedResult {
					t.Errorf("SquareRoot(%v): Expected %v, got %v", testCase.input, testCase.expectedResult, result)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Subtract_58eac52f91
ROOST_METHOD_SIG_HASH=Subtract_b1211baa34

FUNCTION_DEF=func Subtract(num1, num2 int) int // Subtract two integers


*/
func TestSubtract(t *testing.T) {

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Subtracting two positive numbers",
			num1:     10,
			num2:     5,
			expected: 5,
		},
		{
			name:     "Subtracting two negative numbers",
			num1:     -7,
			num2:     -3,
			expected: -4,
		},
		{
			name:     "Subtracting a positive from a negative number",
			num1:     -7,
			num2:     3,
			expected: -10,
		},
		{
			name:     "Subtracting a large number from a smaller one",
			num1:     3,
			num2:     7,
			expected: -4,
		},
		{
			name:     "Subtracting zero from a number",
			num1:     3,
			num2:     0,
			expected: 3,
		},
		{
			name:     "Subtracting maximum possible integers",
			num1:     math.MaxInt64,
			num2:     math.MaxInt64,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			result := Subtract(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Subtract(%d, %d) = %d; expected %d", tc.num1, tc.num2, result, tc.expected)
			} else {
				t.Logf("Subtract(%d, %d) = %d; passed", tc.num1, tc.num2, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Factorial_68fe6fb960
ROOST_METHOD_SIG_HASH=Factorial_3d037eec72

FUNCTION_DEF=func Factorial(n int) int // Factorial (Recursive)


*/
func TestFactorial(t *testing.T) {

	type test struct {
		input       int
		output      int
		shouldPanic bool
	}

	tests := []test{
		{0, 1, false},
		{5, 120, false},
		{-5, 0, true},
		{20, 2432902008176640000, false},
	}

	for _, tc := range tests {

		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered, test failed. %v", r)
				t.Fail()
			}
		}()

		t.Run("", func(t *testing.T) {

			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("Function did not panic when it should have.")
					}
				}()

				Factorial(tc.input)
			} else {

				if output := Factorial(tc.input); output != tc.output {
					t.Errorf("For input %v expected output %v, but got output %v", tc.input, tc.output, output)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=GCD_1da681d86b
ROOST_METHOD_SIG_HASH=GCD_39a1228f3a

FUNCTION_DEF=func GCD(a, b int) int // Greatest Common Divisor (GCD) using Euclidean algorithm


*/
func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantPanic bool
	}{

		{
			name:      "Testing with Positive Integers",
			args:      args{a: 48, b: 18},
			want:      6,
			wantPanic: false,
		},
		{
			name:      "Zero Divisor Testing",
			args:      args{a: 0, b: 20},
			want:      20,
			wantPanic: false,
		},
		{
			name:      "Negative Integers Test",
			args:      args{a: -21, b: 7},
			want:      7,
			wantPanic: false,
		},
		{
			name:      "Testing with Prime Numbers",
			args:      args{a: 11, b: 13},
			want:      1,
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Unexpected Panic occured. %v", r)
					t.Fail()
				}
			}()

			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=LCM_85c2702b86
ROOST_METHOD_SIG_HASH=LCM_fb713f0b10

FUNCTION_DEF=func LCM(a, b int) int // Least Common Multiple (LCM) using GCD


*/
func TestLcm(t *testing.T) {

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	for _, tt := range testTable {
		t.Run(fmt.Sprintf("LCM of %d and %d", tt.a, tt.b), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v", r)
					t.Fail()
				}
			}()

			result := LCM(tt.a, tt.b)

			if result == tt.expected {
				t.Logf("LCM(%d, %d) = %d; Expected = %d. Test Passed!", tt.a, tt.b, result, tt.expected)
			} else {
				t.Errorf("LCM(%d, %d) = %d; Expected = %d. Test Failed!", tt.a, tt.b, result, tt.expected)
			}

			if tt.a == 0 || tt.b == 0 {
				if int(math.Abs(float64(tt.b))) != GCD(tt.a, tt.b) {
					t.Errorf("Expected GCD(%d, %d) = %d. But got GCD = %d. External Dependency assumption validation Failed!", tt.a, tt.b, int(math.Abs(float64(tt.b))), GCD(tt.a, tt.b))
				}
			}
		})
	}

	_ = w.Close()
	out := make([]byte, 1024)
	for {
		n, _ := r.Read(out)
		if n == 0 {
			break
		}
	}
	os.Stdout = old
}

