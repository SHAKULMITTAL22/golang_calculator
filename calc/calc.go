package calc

import (
	"math"
)

// Add two integers
func Add(num1, num2 int) int {
	return num1 + num2
}

// Subtract two integers
func Subtract(num1, num2 int) int {
	return num1 - num2
}

// Multiply two floating-point numbers
func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

// Divide two floating-point numbers (with error handling)
func Divide(num1, num2 float64) float64 {
	if num2 == 0 {
		panic("division by zero is not allowed")
	}
	return num1 / num2
}

// Modulo operation
func Modulo(num1, num2 int) int {
	return num1 % num2
}

// Power function
func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Absolute value
func Absolute(num float64) float64 {
	return math.Abs(num)
}

// Square root (with error handling)
func SquareRoot(num float64) float64 {
	if num < 0 {
		panic("square root of a negative number is not defined")
	}
	return math.Sqrt(num)
}

// Factorial (Recursive)
func Factorial(n int) int {
	if n < 0 {
		panic("factorial is not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Greatest Common Divisor (GCD) using Euclidean algorithm
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// Least Common Multiple (LCM) using GCD
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

// Logarithm function (log_base of num)
func Logarithm(num, base float64) float64 {
	if num <= 0 || base <= 0 || base == 1 {
		panic("logarithm is not defined for these values")
	}
	return math.Log(num) / math.Log(base)
}

// Trigonometric functions (Sin, Cos, Tan)
func SinCosTan(angle float64) (sin, cos, tan float64) {
	sin = math.Sin(angle)
	cos = math.Cos(angle)
	tan = math.Tan(angle)
	return
}
