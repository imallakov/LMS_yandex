package main

import "fmt"

func main() {
	var n int
	var a int = 3
	fmt.Scan(&n)
	for a<=n {
		if is_prime(a){
			fmt.Print("хоп ")
		} else {
			fmt.Print(a," ")	
		}
		a+=5
	}
}

func is_prime(n int) bool {
	for i:=2;i<n;i++{
		if n%i==0 {
			return false
		}
	}
	return true
}