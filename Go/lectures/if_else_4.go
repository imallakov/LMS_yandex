package main

import "fmt"

func main() {
	var a,b,positive,null int
	fmt.Scan(&a,&b)
	if a>0{
		positive++
	}
	if b>0 {
		positive++
	}
	if a==0 {
		null++
	}
	if b==0 {
		null++
	}
	if positive==2 {
		fmt.Println("Оба числа положительные")
	} else {
		if null>0 {
			fmt.Println("Одно из чисел равно нулю")
		} else {
			if positive>0 {
				fmt.Println("Одно число положительное, а другое отрицательное")
			} else {
				fmt.Println("Оба числа отрицательные")
			}
		}
	}
}