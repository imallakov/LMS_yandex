package main

import (
    "fmt"
)

func Add(a,b float64) float64 {
	var sum float64
	sum=a+b
	return sum
}

func Multiply(a,b float64) float64 {
	var ans float64
	ans=a*b
	return ans
}

func PrintNumbersAscending(n int) {
	for i:=1;i<=n;i++ {
		fmt.Print(i," ")
	}
}

func main() {
	var a,b float64
	fmt.Scan(&a,&b)
	Add(a,b)
	Multiply(a,b)
	PrintNumbersAscending(15)
}