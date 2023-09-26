package main

import "fmt"

func main() {
	var n,a,b,c,stop,count int = 0,0,1,1,0,0
	fmt.Scan(&n)

	if n<=0 {
		fmt.Println(0)
		count++
	}
	if n<=1 {
		fmt.Println(1)
		count++
	}
	for stop==0 {
		c=a+b
		a=b
		b=c
		if c>=n {
			fmt.Println(c)
			count++
		}
		if count==10 {
			stop=1
		}
	}
}