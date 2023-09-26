package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var positive bool
	if(n>0) {
		positive=true
	}
	var even int;
	even=n%2
	if(positive) {
		fmt.Print("Число положительное и ")
	} else {
		fmt.Print("Число отрицательное и ")
	}
	if(even==0) {
		fmt.Println("четное")
	} else {
		fmt.Println("нечетное")
	}

}