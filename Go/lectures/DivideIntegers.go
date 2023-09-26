package main

import "errors"

func DivideIntegers(a,b int) (float64, error) {
	if b==0 {
		return 0, errors.New("division by zero not allowed")
	}

	return float64(a/b),nil
}