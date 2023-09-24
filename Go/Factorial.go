package main

import "errors"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	var ans int = 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans, nil
}
