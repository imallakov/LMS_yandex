package main

import "fmt"

func CalculateDigitalRoot(n int) int {
	for (n>10) {
		m:=n
		var sum int
		for(m>0) {
			sum+=m%10
			m/=10
		}
		n=sum
	}
	return n
}

func main() {
	fmt.Println(CalculateDigitalRoot(15))
}