package dynamic

// Fibo returns num's fibonnacci number.
// Reference: https://en.wikipedia.org/wiki/Fibonacci_number
func Fibo(num int) int {
	memo := []int{0, 1}
	return fibo(num, memo)
}

func fibo(num int, memo []int) int {
	if num < len(memo) {
		return memo[num]
	}
	fib := fibo(num-1, memo) + fibo(num-2, memo)
	memo = append(memo, fib)
	return fib
}
