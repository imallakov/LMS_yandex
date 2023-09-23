package main

func ConcatStringsAndInt(str1, str2 string, num int) string {
	var ans string = str1+str2
	var temp string
	for num>0 {
		var n int = num%10
		temp+=string(n+48)
		num/=10
	}
	for i:=len(temp)-1;i>=0;i--{
		ans+=string(temp[i])
	}
	return ans
}