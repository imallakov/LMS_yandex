package main

import "fmt"

func main(){
	var n, sum int
	fmt.Scan(&n)
	for n>0 {
		if n%3!=0 && n%5!=0 {
			sum+=n
		}
		n--
	}
	fmt.Println(sum)
}