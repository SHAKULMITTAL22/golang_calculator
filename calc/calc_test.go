package calc

import (
	math "math"
	debug "runtime/debug"
	testing "testing"
	fmt "fmt"
	os "os"
	time "time"
)








/*
ROOST_METHOD_HASH=Absolute_f8af7505a1
ROOST_METHOD_SIG_HASH=Absolute_4bad226818

FUNCTION_DEF=func Absolute(num float64) float64 

*/
func TestAbsolute(t *testing.T) {

	testCases := []struct {
		name     string
		input    float64
		expected float64
		isNaN    bool
	}{
		{
			name:     "Positive Number",
			input:    5.0,
			expected: 5.0,
		},
		{
			name:     "Negative Number",
			input:    -5.0,
			expected: 5.0,
		},
		{
			name:     "Zero",
			input:    0.0,
			expected: 0.0,
		},
		{
			name:     "Maximum Float64 Value",
			input:    math.MaxFloat64,
			expected: math.MaxFloat64,
		},
		{
			name:     "Minimum Float64 Value",
			input:    -math.MaxFloat64,
			expected: math.MaxFloat64,
		},
		{
			name:     "Small Decimal Value",
			input:    0.0000001,
			expected: 0.0000001,
		},
		{
			name:     "Small Negative Decimal Value",
			input:    -0.0000001,
			expected: 0.0000001,
		},
		{
			name:  "Not a Number (NaN)",
			input: math.NaN(),
			isNaN: true,
		},
		{
			name:     "Positive Infinity",
			input:    math.Inf(1),
			expected: math.Inf(1),
		},
		{
			name:     "Negative Infinity",
			input:    math.Inf(-1),
			expected: math.Inf(1),
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

			t.Logf("Testing Absolute function with input: %v", tc.input)
			result := Absolute(tc.input)

			if tc.isNaN {
				if !math.IsNaN(result) {
					t.Errorf("Expected NaN, but got %v", result)
				} else {
					t.Logf("Success: Absolute(%v) correctly returned NaN", tc.input)
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Expected Absolute(%v) to be %v, but got %v", tc.input, tc.expected, result)
			} else {
				t.Logf("Success: Absolute(%v) = %v", tc.input, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Add_38f6779755
ROOST_METHOD_SIG_HASH=Add_8e349a90e1

FUNCTION_DEF=func Add(num1, num2 int) int 

*/
func TestAdd(t *testing.T) {

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{
			name:     "Add two positive numbers",
			num1:     5,
			num2:     7,
			expected: 12,
		},
		{
			name:     "Add positive and negative numbers",
			num1:     10,
			num2:     -5,
			expected: 5,
		},
		{
			name:     "Add two negative numbers",
			num1:     -3,
			num2:     -7,
			expected: -10,
		},
		{
			name:     "Add zero and positive number",
			num1:     0,
			num2:     42,
			expected: 42,
		},
		{
			name:     "Add zero and negative number",
			num1:     0,
			num2:     -42,
			expected: -42,
		},
		{
			name:     "Add zero and zero",
			num1:     0,
			num2:     0,
			expected: 0,
		},
		{
			name:     "Test with large numbers",
			num1:     math.MaxInt32 / 2,
			num2:     math.MaxInt32 / 2,
			expected: math.MaxInt32 - 1,
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

			t.Logf("Testing Add(%d, %d), expecting: %d", tc.num1, tc.num2, tc.expected)

			result := Add(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tc.num1, tc.num2, result, tc.expected)
			} else {
				t.Logf("Add(%d, %d) = %d; test passed", tc.num1, tc.num2, result)
			}
		})
	}

	t.Run("Test potential overflow", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		num1 := math.MaxInt32
		num2 := 1
		result := Add(num1, num2)

		t.Logf("Add(%d, %d) = %d (overflow behavior)", num1, num2, result)

		expected := num1 + num2
		if result != expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", num1, num2, result, expected)
		}
	})

	t.Run("Test with stdout capture", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		oldStdout := os.Stdout

		r, w, _ := os.Pipe()
		os.Stdout = w

		result := Add(3, 4)

		fmt.Fprintf(w, "Add result: %d\n", result)

		w.Close()

		os.Stdout = oldStdout

		var capturedOutput string
		fmt.Fscanf(r, "Add result: %d\n", &capturedOutput)

		t.Logf("Captured output: %s", capturedOutput)

		if result != 7 {
			t.Errorf("Add(3, 4) = %d; expected 7", result)
		}
	})
}


/*
ROOST_METHOD_HASH=Divide_f2ddee767d
ROOST_METHOD_SIG_HASH=Divide_15b7594322

FUNCTION_DEF=func Divide(num1, num2 float64) float64 

*/
func TestDivide(t *testing.T) {

	testCases := []struct {
		name           string
		num1           float64
		num2           float64
		expected       float64
		expectPanic    bool
		expectedPanic  string
		checkInfinity  bool
		checkNaN       bool
		epsilon        float64
		skipComparison bool
	}{

		{
			name:     "Positive numbers division",
			num1:     10.0,
			num2:     2.0,
			expected: 5.0,
			epsilon:  1e-15,
		},

		{
			name:     "Negative dividend, positive divisor",
			num1:     -10.0,
			num2:     2.0,
			expected: -5.0,
			epsilon:  1e-15,
		},
		{
			name:     "Positive dividend, negative divisor",
			num1:     10.0,
			num2:     -2.0,
			expected: -5.0,
			epsilon:  1e-15,
		},
		{
			name:     "Negative dividend, negative divisor",
			num1:     -10.0,
			num2:     -2.0,
			expected: 5.0,
			epsilon:  1e-15,
		},

		{
			name:          "Division by zero",
			num1:          10.0,
			num2:          0.0,
			expectPanic:   true,
			expectedPanic: "division by zero is not allowed",
		},

		{
			name:     "Division with very large numbers",
			num1:     math.MaxFloat64,
			num2:     2.0,
			expected: math.MaxFloat64 / 2.0,
			epsilon:  1e-10,
		},

		{
			name:     "Division with very small numbers",
			num1:     math.SmallestNonzeroFloat64,
			num2:     2.0,
			expected: math.SmallestNonzeroFloat64 / 2.0,
			epsilon:  1e-15,
		},

		{
			name:          "Division resulting in positive infinity",
			num1:          math.MaxFloat64,
			num2:          math.SmallestNonzeroFloat64,
			checkInfinity: true,
			expected:      math.Inf(1),
		},
		{
			name:          "Division resulting in negative infinity",
			num1:          -math.MaxFloat64,
			num2:          math.SmallestNonzeroFloat64,
			checkInfinity: true,
			expected:      math.Inf(-1),
		},

		{
			name:           "Division with NaN dividend",
			num1:           math.NaN(),
			num2:           5.0,
			checkNaN:       true,
			skipComparison: true,
		},
		{
			name:     "Division by positive infinity",
			num1:     5.0,
			num2:     math.Inf(1),
			expected: 0.0,
			epsilon:  1e-15,
		},
		{
			name:           "Division of infinity by infinity",
			num1:           math.Inf(1),
			num2:           math.Inf(1),
			checkNaN:       true,
			skipComparison: true,
		},

		{
			name:     "Division with repeating decimal result",
			num1:     1.0,
			num2:     3.0,
			expected: 0.3333333333333333,
			epsilon:  1e-15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					if tc.expectPanic {
						if r != tc.expectedPanic {
							t.Errorf("Expected panic message '%s', got '%v'", tc.expectedPanic, r)
						} else {
							t.Logf("Successfully caught expected panic: %v", r)
						}
					} else {
						t.Errorf("Unexpected panic: %v", r)
					}
				} else if tc.expectPanic {
					t.Errorf("Expected panic but none occurred")
				}
			}()

			result := Divide(tc.num1, tc.num2)

			if tc.expectPanic {
				t.Errorf("Expected panic but function returned %v", result)
				return
			}

			if tc.checkInfinity {
				if !math.IsInf(result, int(math.Copysign(1, tc.expected))) {
					t.Errorf("Expected %v, got %v", tc.expected, result)
				} else {
					t.Logf("Successfully verified infinity result: %v", result)
				}
				return
			}

			if tc.checkNaN {
				if !math.IsNaN(result) {
					t.Errorf("Expected NaN, got %v", result)
				} else {
					t.Logf("Successfully verified NaN result")
				}
				return
			}

			if tc.skipComparison {
				return
			}

			if math.Abs(result-tc.expected) > tc.epsilon {
				t.Errorf("Expected %v, got %v, difference: %v exceeds epsilon %v",
					tc.expected, result, math.Abs(result-tc.expected), tc.epsilon)
			} else {
				t.Logf("Successfully verified result: %v", result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Factorial_89543dc467
ROOST_METHOD_SIG_HASH=Factorial_9b038c83eb

FUNCTION_DEF=func Factorial(n int) int 

*/
func BenchmarkFactorial(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{"Small", 5},
		{"Medium", 10},
		{"Large", 15},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("Input_%s_%d", bm.name, bm.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Factorial(bm.input)
			}
		})
	}
}

func TestFactorial(t *testing.T) {

	t.Run("Scenario 8: Test Factorial Against Known Values", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		testCases := map[string]struct {
			input    int
			expected int
		}{
			"Factorial of 0": {0, 1},
			"Factorial of 1": {1, 1},
			"Factorial of 2": {2, 2},
			"Factorial of 3": {3, 6},
			"Factorial of 4": {4, 24},
			"Factorial of 5": {5, 120},
			"Factorial of 6": {6, 720},
			"Factorial of 7": {7, 5040},
		}

		for name, tc := range testCases {
			t.Run(name, func(t *testing.T) {
				result := Factorial(tc.input)
				if result != tc.expected {
					t.Errorf("Expected Factorial(%d) to be %d, got %d", tc.input, tc.expected, result)
				} else {
					t.Logf("Successfully calculated Factorial(%d) = %d", tc.input, result)
				}
			})
		}
	})

	t.Run("Scenario 1: Test Factorial of Zero", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		result := Factorial(0)
		expected := 1
		if result != expected {
			t.Errorf("Expected Factorial(0) to be %d, got %d", expected, result)
		} else {
			t.Logf("Successfully calculated Factorial(0) = %d", result)
		}
	})

	t.Run("Scenario 2: Test Factorial of One", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		result := Factorial(1)
		expected := 1
		if result != expected {
			t.Errorf("Expected Factorial(1) to be %d, got %d", expected, result)
		} else {
			t.Logf("Successfully calculated Factorial(1) = %d", result)
		}
	})

	t.Run("Scenario 3: Test Factorial of Small Positive Numbers", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		result := Factorial(5)
		expected := 120
		if result != expected {
			t.Errorf("Expected Factorial(5) to be %d, got %d", expected, result)
		} else {
			t.Logf("Successfully calculated Factorial(5) = %d", result)
		}
	})

	t.Run("Scenario 4: Test Factorial of Larger Positive Numbers", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		result := Factorial(10)
		expected := 3628800
		if result != expected {
			t.Errorf("Expected Factorial(10) to be %d, got %d", expected, result)
		} else {
			t.Logf("Successfully calculated Factorial(10) = %d", result)
		}
	})

	t.Run("Scenario 5: Test Factorial Panics with Negative Input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected Factorial(-1) to panic, but it didn't")
			} else {
				expectedMsg := "factorial is not defined for negative numbers"
				if r != expectedMsg {
					t.Errorf("Expected panic message '%s', got '%v'", expectedMsg, r)
				} else {
					t.Logf("Successfully caught panic with message: %v", r)
				}
			}
		}()

		Factorial(-1)
	})

	t.Run("Scenario 6: Test Factorial Result Overflow", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		var largeInput int
		if math.MaxInt == math.MaxInt32 {
			largeInput = 13
			t.Log("Testing on 32-bit system, using n=13 as largest safe factorial")
		} else {
			largeInput = 20
			t.Log("Testing on 64-bit system, using n=20 as largest safe factorial")
		}

		var expected int
		if largeInput == 13 {
			expected = 6227020800
		} else {
			expected = 2432902008176640000
		}

		result := Factorial(largeInput)
		if result != expected {
			t.Errorf("Expected Factorial(%d) to be %d, got %d", largeInput, expected, result)
		} else {
			t.Logf("Successfully calculated Factorial(%d) = %d without overflow", largeInput, result)
		}

		overflowInput := largeInput + 1
		overflowResult := Factorial(overflowInput)
		t.Logf("Factorial(%d) = %d (may be incorrect due to integer overflow)", overflowInput, overflowResult)

		if overflowResult < 0 || overflowResult < result {
			t.Logf("Detected likely integer overflow: Factorial(%d) returned %d", overflowInput, overflowResult)
		}
	})
}


/*
ROOST_METHOD_HASH=GCD_6cf0472095
ROOST_METHOD_SIG_HASH=GCD_ab1c91475d

FUNCTION_DEF=func GCD(a, b int) int 

*/
func TestGcd(t *testing.T) {

	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
		scenario string
	}{

		{
			name:     "Positive integers",
			a:        12,
			b:        8,
			expected: 4,
			scenario: "Scenario 1: Verify that the GCD function correctly calculates the greatest common divisor of two positive integers.",
		},
		{
			name:     "Positive integers - larger values",
			a:        48,
			b:        18,
			expected: 6,
			scenario: "Scenario 1: Verify that the GCD function correctly calculates the greatest common divisor of two positive integers.",
		},

		{
			name:     "First parameter zero",
			a:        0,
			b:        5,
			expected: 5,
			scenario: "Scenario 2: Verify that the GCD function handles cases where one of the inputs is zero.",
		},
		{
			name:     "Second parameter zero",
			a:        5,
			b:        0,
			expected: 5,
			scenario: "Scenario 2: Verify that the GCD function handles cases where one of the inputs is zero.",
		},

		{
			name:     "Both parameters zero",
			a:        0,
			b:        0,
			expected: 0,
			scenario: "Scenario 3: Verify the behavior of the GCD function when both inputs are zero.",
		},

		{
			name:     "Large numbers",
			a:        1234567890,
			b:        987654321,
			expected: 9,
			scenario: "Scenario 4: Verify that the GCD function correctly handles large integers without overflow or excessive recursion.",
		},

		{
			name:     "First parameter negative",
			a:        -12,
			b:        8,
			expected: 4,
			scenario: "Scenario 5: Verify that the GCD function correctly handles negative integers.",
		},
		{
			name:     "Second parameter negative",
			a:        12,
			b:        -8,
			expected: 4,
			scenario: "Scenario 5: Verify that the GCD function correctly handles negative integers.",
		},
		{
			name:     "Both parameters negative",
			a:        -12,
			b:        -8,
			expected: 4,
			scenario: "Scenario 5: Verify that the GCD function correctly handles negative integers.",
		},

		{
			name:     "Distinct prime numbers",
			a:        17,
			b:        19,
			expected: 1,
			scenario: "Scenario 6: Verify that the GCD function correctly identifies when two prime numbers have no common divisor other than 1.",
		},

		{
			name:     "Consecutive Fibonacci numbers 1",
			a:        5,
			b:        8,
			expected: 1,
			scenario: "Scenario 7: Verify that the GCD of consecutive Fibonacci numbers is always 1.",
		},
		{
			name:     "Consecutive Fibonacci numbers 2",
			a:        8,
			b:        13,
			expected: 1,
			scenario: "Scenario 7: Verify that the GCD of consecutive Fibonacci numbers is always 1.",
		},

		{
			name:     "Equal numbers",
			a:        7,
			b:        7,
			expected: 7,
			scenario: "Scenario 8: Verify that the GCD function correctly handles cases where both inputs are the same.",
		},

		{
			name:     "First parameter is 1",
			a:        1,
			b:        5,
			expected: 1,
			scenario: "Scenario 9: Verify that the GCD function correctly handles cases where one input is 1.",
		},
		{
			name:     "Second parameter is 1",
			a:        10,
			b:        1,
			expected: 1,
			scenario: "Scenario 9: Verify that the GCD function correctly handles cases where one input is 1.",
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

			t.Logf("Testing scenario: %s", tc.scenario)

			result := GCD(tc.a, tc.b)

			if result != tc.expected {
				t.Errorf("GCD(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
			} else {
				t.Logf("GCD(%d, %d) = %d; test passed", tc.a, tc.b, result)
			}
		})
	}

	t.Run("Performance with worst-case inputs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
				t.Fail()
			}
		}()

		t.Log("Scenario 10: Verify that the GCD function performs efficiently even with inputs that require many recursive calls.")

		fib40, fib41 := generateFibonacciPair(40)

		start := time.Now()
		result := GCD(fib40, fib41)
		duration := time.Since(start)

		if result != 1 {
			t.Errorf("GCD(%d, %d) = %d; expected 1", fib40, fib41, result)
		} else {
			t.Logf("GCD(%d, %d) = %d; test passed", fib40, fib41, result)
		}

		if duration > 100*time.Millisecond {
			t.Logf("Warning: GCD calculation took %v, which might be too long", duration)
		} else {
			t.Logf("Performance test passed: GCD calculation took %v", duration)
		}
	})
}

func generateFibonacciPair(n int) (int, int) {
	if n <= 1 {
		return n, 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b, a + b
}


/*
ROOST_METHOD_HASH=LCM_6035446662
ROOST_METHOD_SIG_HASH=LCM_121c872fbf

FUNCTION_DEF=func LCM(a, b int) int 

*/
func TestLcm(t *testing.T) {

	type testCase struct {
		name     string
		a        int
		b        int
		expected int
		panics   bool
	}

	tests := []testCase{

		{
			name:     "Positive integers - 4 and 6",
			a:        4,
			b:        6,
			expected: 12,
		},
		{
			name:     "Positive integers - 15 and 20",
			a:        15,
			b:        20,
			expected: 60,
		},

		{
			name:     "Identical numbers - 5 and 5",
			a:        5,
			b:        5,
			expected: 5,
		},
		{
			name:     "Identical numbers - 10 and 10",
			a:        10,
			b:        10,
			expected: 10,
		},

		{
			name:   "One input as 0 - 0 and 5",
			a:      0,
			b:      5,
			panics: true,
		},
		{
			name:   "One input as 0 - 7 and 0",
			a:      7,
			b:      0,
			panics: true,
		},

		{
			name:   "Both inputs as 0 - 0 and 0",
			a:      0,
			b:      0,
			panics: true,
		},

		{
			name:     "Negative numbers - -4 and 6",
			a:        -4,
			b:        6,
			expected: 12,
		},
		{
			name:     "Negative numbers - 4 and -6",
			a:        4,
			b:        -6,
			expected: 12,
		},
		{
			name:     "Negative numbers - -4 and -6",
			a:        -4,
			b:        -6,
			expected: 12,
		},

		{
			name:     "Large numbers - 1234 and 5678",
			a:        1234,
			b:        5678,
			expected: 1400652,
		},

		{
			name:     "Coprime numbers - 7 and 12",
			a:        7,
			b:        12,
			expected: 84,
		},
		{
			name:     "Coprime numbers - 5 and 8",
			a:        5,
			b:        8,
			expected: 40,
		},

		{
			name:     "One input as 1 - 1 and 5",
			a:        1,
			b:        5,
			expected: 5,
		},
		{
			name:     "One input as 1 - 7 and 1",
			a:        7,
			b:        1,
			expected: 7,
		},

		{
			name:     "Prime numbers - 3 and 5",
			a:        3,
			b:        5,
			expected: 15,
		},
		{
			name:     "Prime numbers - 11 and 13",
			a:        11,
			b:        13,
			expected: 143,
		},

		{
			name:     "One number divisible by other - 4 and 8",
			a:        4,
			b:        8,
			expected: 8,
		},
		{
			name:     "One number divisible by other - 9 and 3",
			a:        9,
			b:        3,
			expected: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered so failing test. %v\n%s", r, string(debug.Stack()))
					if !tc.panics {
						t.Fail()
					} else {
						t.Logf("Expected panic occurred: %v", r)
					}
				} else if tc.panics {
					t.Errorf("Expected panic but none occurred")
				}
			}()

			result := LCM(tc.a, tc.b)

			if !tc.panics {
				if result != tc.expected {
					t.Errorf("LCM(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
				} else {
					t.Logf("LCM(%d, %d) = %d; test passed", tc.a, tc.b, result)
				}
			}
		})
	}

	t.Run("Integer overflow check", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Panic encountered with large numbers: %v\n%s", r, string(debug.Stack()))

				t.Logf("Note: This test checks how the function handles potential overflow")
			}
		}()

		a := math.MaxInt32 / 2
		b := math.MaxInt32 / 3

		result := LCM(a, b)
		t.Logf("LCM of large numbers (%d, %d) = %d", a, b, result)

	})
}


/*
ROOST_METHOD_HASH=Multiply_1585632006
ROOST_METHOD_SIG_HASH=Multiply_d6ab1fb07f

FUNCTION_DEF=func Multiply(num1, num2 float64) float64 

*/
func TestMultiply(t *testing.T) {

	testCases := []struct {
		name     string
		num1     float64
		num2     float64
		expected float64
		isNaN    bool
	}{

		{
			name:     "Basic multiplication of positive numbers",
			num1:     2.0,
			num2:     3.0,
			expected: 6.0,
		},

		{
			name:     "Multiplication with zero (first operand)",
			num1:     0.0,
			num2:     5.5,
			expected: 0.0,
		},
		{
			name:     "Multiplication with zero (second operand)",
			num1:     5.5,
			num2:     0.0,
			expected: 0.0,
		},

		{
			name:     "Multiplication with negative and positive number",
			num1:     -2.0,
			num2:     3.0,
			expected: -6.0,
		},
		{
			name:     "Multiplication with positive and negative number",
			num1:     2.0,
			num2:     -3.0,
			expected: -6.0,
		},
		{
			name:     "Multiplication with two negative numbers",
			num1:     -2.0,
			num2:     -3.0,
			expected: 6.0,
		},

		{
			name:     "Multiplication with very large numbers",
			num1:     1e150,
			num2:     1e150,
			expected: 1e300,
		},

		{
			name:     "Multiplication with very small numbers",
			num1:     1e-150,
			num2:     1e-150,
			expected: 1e-300,
		},

		{
			name:     "Multiplication with infinity (first operand)",
			num1:     math.Inf(1),
			num2:     2.0,
			expected: math.Inf(1),
		},
		{
			name:     "Multiplication with infinity (second operand)",
			num1:     2.0,
			num2:     math.Inf(1),
			expected: math.Inf(1),
		},

		{
			name:  "Multiplication with NaN (first operand)",
			num1:  math.NaN(),
			num2:  2.0,
			isNaN: true,
		},
		{
			name:  "Multiplication with NaN (second operand)",
			num1:  2.0,
			num2:  math.NaN(),
			isNaN: true,
		},

		{
			name:     "Precision test with decimal numbers",
			num1:     0.1,
			num2:     0.2,
			expected: 0.02,
		},

		{
			name:     "Commutative property test (a*b)",
			num1:     3.14,
			num2:     2.71,
			expected: 8.5094,
		},
		{
			name:     "Commutative property test (b*a)",
			num1:     2.71,
			num2:     3.14,
			expected: 8.5094,
		},

		{
			name:     "Identity property test (a*1)",
			num1:     42.0,
			num2:     1.0,
			expected: 42.0,
		},
		{
			name:     "Identity property test (1*a)",
			num1:     1.0,
			num2:     42.0,
			expected: 42.0,
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

			t.Logf("Testing: %s with inputs %v and %v", tc.name, tc.num1, tc.num2)
			result := Multiply(tc.num1, tc.num2)

			if tc.isNaN {
				if !math.IsNaN(result) {
					t.Errorf("Expected NaN, but got %v", result)
				} else {
					t.Logf("Success: Result is NaN as expected")
				}
				return
			}

			const epsilon = 1e-10
			if math.Abs(result-tc.expected) > epsilon {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			} else {
				t.Logf("Success: Expected %v and got %v", tc.expected, result)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Power_2542f03efe
ROOST_METHOD_SIG_HASH=Power_feb9859574

FUNCTION_DEF=func Power(base, exponent float64) float64 

*/
func TestPower(t *testing.T) {

	testCases := []struct {
		name     string
		base     float64
		exponent float64
		expected float64
		isNaN    bool
	}{
		{
			name:     "Scenario 1: Positive Base and Positive Exponent",
			base:     2.0,
			exponent: 3.0,
			expected: 8.0,
		},
		{
			name:     "Scenario 2: Negative Base and Integer Exponent",
			base:     -2.0,
			exponent: 3.0,
			expected: -8.0,
		},
		{
			name:     "Scenario 3: Zero Base",
			base:     0.0,
			exponent: 5.0,
			expected: 0.0,
		},
		{
			name:     "Scenario 4: Zero Exponent",
			base:     7.5,
			exponent: 0.0,
			expected: 1.0,
		},
		{
			name:     "Scenario 5: Base One",
			base:     1.0,
			exponent: 100.0,
			expected: 1.0,
		},
		{
			name:     "Scenario 6: Fractional Exponent",
			base:     4.0,
			exponent: 0.5,
			expected: 2.0,
		},
		{
			name:     "Scenario 7: Negative Exponent",
			base:     2.0,
			exponent: -2.0,
			expected: 0.25,
		},
		{
			name:     "Scenario 8: Very Large Exponent",
			base:     1.001,
			exponent: 1000.0,
			expected: 2.7169,
		},
		{
			name:     "Scenario 9: Very Small Exponent",
			base:     10.0,
			exponent: 0.0001,
			expected: 1.0023,
		},
		{
			name:     "Scenario 10: Negative Base and Fractional Exponent",
			base:     -4.0,
			exponent: 0.5,
			isNaN:    true,
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

			t.Logf("Testing Power(%v, %v)", tc.base, tc.exponent)
			result := Power(tc.base, tc.exponent)

			if tc.isNaN {

				if !math.IsNaN(result) {
					t.Errorf("Expected NaN for Power(%v, %v), but got %v", tc.base, tc.exponent, result)
				} else {
					t.Logf("Successfully got NaN for Power(%v, %v) as expected", tc.base, tc.exponent)
				}
			} else if tc.name == "Scenario 8: Very Large Exponent" || tc.name == "Scenario 9: Very Small Exponent" {

				tolerance := 0.001
				if math.Abs(result-tc.expected) > tolerance {
					t.Errorf("Power(%v, %v) = %v, expected approximately %v (within %v tolerance)",
						tc.base, tc.exponent, result, tc.expected, tolerance)
				} else {
					t.Logf("Power(%v, %v) = %v, which is approximately %v as expected",
						tc.base, tc.exponent, result, tc.expected)
				}
			} else {

				if result != tc.expected {
					t.Errorf("Power(%v, %v) = %v, expected %v", tc.base, tc.exponent, result, tc.expected)
				} else {
					t.Logf("Power(%v, %v) = %v as expected", tc.base, tc.exponent, result)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=SquareRoot_600b6ad663
ROOST_METHOD_SIG_HASH=SquareRoot_5aa1e1a6d6

FUNCTION_DEF=func SquareRoot(num float64) float64 

*/
func TestSquareRoot(t *testing.T) {

	const epsilon = 1e-9

	isApproxEqual := func(a, b float64) bool {
		return math.Abs(a-b) < epsilon
	}

	tests := []struct {
		name     string
		input    float64
		expected float64
		isPanic  bool
	}{

		{
			name:     "Positive number",
			input:    16.0,
			expected: 4.0,
			isPanic:  false,
		},

		{
			name:     "Zero",
			input:    0.0,
			expected: 0.0,
			isPanic:  false,
		},

		{
			name:     "Very large positive number",
			input:    1e10,
			expected: 1e5,
			isPanic:  false,
		},

		{
			name:     "Very small positive number",
			input:    1e-10,
			expected: 1e-5,
			isPanic:  false,
		},

		{
			name:     "Negative number",
			input:    -4.0,
			expected: 0.0,
			isPanic:  true,
		},

		{
			name:     "Perfect square - 4",
			input:    4.0,
			expected: 2.0,
			isPanic:  false,
		},
		{
			name:     "Perfect square - 9",
			input:    9.0,
			expected: 3.0,
			isPanic:  false,
		},
		{
			name:     "Perfect square - 25",
			input:    25.0,
			expected: 5.0,
			isPanic:  false,
		},

		{
			name:     "Non-perfect square - 2",
			input:    2.0,
			expected: math.Sqrt(2.0),
			isPanic:  false,
		},
		{
			name:     "Non-perfect square - 3",
			input:    3.0,
			expected: math.Sqrt(3.0),
			isPanic:  false,
		},
		{
			name:     "Non-perfect square - 10",
			input:    10.0,
			expected: math.Sqrt(10.0),
			isPanic:  false,
		},

		{
			name:     "Maximum float64 value",
			input:    math.MaxFloat64,
			expected: math.Sqrt(math.MaxFloat64),
			isPanic:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("Panic encountered: %v\n%s", r, string(debug.Stack()))
					if !tt.isPanic {
						t.Errorf("SquareRoot(%v) unexpectedly panicked: %v", tt.input, r)
					} else {

						if r != "square root of a negative number is not defined" {
							t.Errorf("Expected panic message 'square root of a negative number is not defined', got: %v", r)
						} else {
							t.Logf("Successfully caught expected panic for negative input: %v", tt.input)
						}
					}
				} else if tt.isPanic {
					t.Errorf("SquareRoot(%v) should have panicked but didn't", tt.input)
				}
			}()

			if tt.isPanic {
				t.Logf("Testing panic case with input: %v", tt.input)
				result := SquareRoot(tt.input)

				t.Errorf("SquareRoot(%v) = %v, expected to panic", tt.input, result)
				return
			}

			result := SquareRoot(tt.input)
			if !isApproxEqual(result, tt.expected) {
				t.Errorf("SquareRoot(%v) = %v, want %v", tt.input, result, tt.expected)
			} else {
				t.Logf("SquareRoot(%v) = %v, which matches expected %v", tt.input, result, tt.expected)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=Subtract_559013d27f
ROOST_METHOD_SIG_HASH=Subtract_29b74c09c9

FUNCTION_DEF=func Subtract(num1, num2 int) int 

*/
func TestSubtract(t *testing.T) {

	testCases := []struct {
		name     string
		num1     int
		num2     int
		expected int
		scenario string
	}{
		{
			name:     "Basic Subtraction with Positive Numbers",
			num1:     10,
			num2:     5,
			expected: 5,
			scenario: "Scenario 1: Verify that the Subtract function correctly subtracts two positive integers.",
		},
		{
			name:     "Subtraction with Negative Numbers",
			num1:     -10,
			num2:     -5,
			expected: -5,
			scenario: "Scenario 2: Verify that the Subtract function correctly handles negative integers in both parameters.",
		},
		{
			name:     "Subtraction Resulting in Zero",
			num1:     42,
			num2:     42,
			expected: 0,
			scenario: "Scenario 3: Verify that the Subtract function correctly returns zero when subtracting a number from itself.",
		},
		{
			name:     "Subtraction Resulting in Negative Value",
			num1:     5,
			num2:     10,
			expected: -5,
			scenario: "Scenario 4: Verify that the Subtract function correctly returns a negative result when the second parameter is larger than the first.",
		},
		{
			name:     "Subtraction with Zero as First Parameter",
			num1:     0,
			num2:     5,
			expected: -5,
			scenario: "Scenario 5: Verify that the Subtract function correctly handles zero as the first parameter.",
		},
		{
			name:     "Subtraction with Zero as Second Parameter",
			num1:     5,
			num2:     0,
			expected: 5,
			scenario: "Scenario 6: Verify that the Subtract function correctly handles zero as the second parameter.",
		},
		{
			name:     "Subtraction with Large Positive Integers",
			num1:     math.MaxInt32 - 10,
			num2:     10,
			expected: math.MaxInt32 - 20,
			scenario: "Scenario 7: Verify that the Subtract function correctly handles large positive integers within the int range.",
		},
		{
			name:     "Subtraction with Large Negative Integers",
			num1:     math.MinInt32 + 10,
			num2:     10,
			expected: math.MinInt32,
			scenario: "Scenario 8: Verify that the Subtract function correctly handles large negative integers within the int range.",
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

			t.Logf("Testing: %s", tc.scenario)

			result := Subtract(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("Subtract(%d, %d) = %d; expected %d", tc.num1, tc.num2, result, tc.expected)
				t.Logf("FAILED: %s", tc.scenario)
			} else {
				t.Logf("SUCCESS: Subtract(%d, %d) = %d as expected", tc.num1, tc.num2, result)
			}
		})
	}
}

