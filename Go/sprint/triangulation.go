package main

import "fmt"

func main() {
	var a,b int
	fmt.Scan(&a,&b)
	if a==0 {
		b++
	}
	for a>0 {
		for i:=1;i<=a;i++ {
			fmt.Print("*")
		}
		fmt.Println()
		a--;
	}
	for b>1 {
		fmt.Println("*")
		b--;
	}
}