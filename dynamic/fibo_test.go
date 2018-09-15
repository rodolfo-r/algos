package dynamic_test

import (
	"testing"

	"github.com/rodolfo-r/algos/dynamic"
)

func TestMemo(t *testing.T) {
	fibs := []int{0, 1, 1, 2, 3, 5, 8,
		13, 21, 34, 55, 89}

	for i, n := 0, len(fibs); i < n; i++ {
		fib := dynamic.Fibo(i)
		if fib != fibs[i] {
			t.Fatalf("wrong fibonacci num for %v. have %v, want %v", i, fib, fibs[i])
		}
	}
}
