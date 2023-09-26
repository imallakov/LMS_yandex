package main

import "fmt"

func main() {
	var n,sum int
	fmt.Scan(&n)
	if n<0 {
		fmt.Println("Некорректный ввод")
		return
	}
	for n>0 {
		if n%2==1 {
			sum+=n
		}
		n--;
	}
	fmt.Println(sum)
}