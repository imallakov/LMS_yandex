package main

import (
	"errors"
	"strconv"
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "",errors.New("negative numbers are not allowed")
	}
	var str string
	for num>0 {
		var temp int = num%2
		str += strconv.Itoa(temp)
		num/=2;
	}
	str=reverse(str)
	return str,nil
}

func reverse(s string) string {
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
    return string(rns)
}