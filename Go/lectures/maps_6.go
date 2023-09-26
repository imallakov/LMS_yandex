package main

import (
	"strings"
)

func IsPalindrome(input string) bool{
	var str1,str2 string
	for _,ch:=range input {
		if(ch!=' '){
			str1=str1 + string(ch)
			str2=string(ch)+str2
		}
	}
	str1=strings.ToLower(str1)
	str2=strings.ToLower(str2)
	return str1==str2

}