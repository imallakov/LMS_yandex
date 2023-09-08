package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	if number>0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}
