package fibonacci

//*求斐波那契数列，返回第n个值
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
