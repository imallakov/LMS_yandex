package main

import (
	"errors"
	"strconv"
)

func SumTwoIntegers(a, b string) (int ,error) {
	num1,err1 := strconv.Atoi(a)
	num2,err2 := strconv.Atoi(b)
	if err1 != nil || err2!=nil{
		return 0,errors.New("invalid input, please provide two integers")
	}
	return num1+num2,nil
}