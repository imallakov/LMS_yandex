package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n<0 {
		fmt.Println("Некорректный ввод")
		return
	}
	if n<10 {
		fmt.Println("Число меньше 10")
	}else {
		if n<100 {
			fmt.Println("Число меньше 100")
		} else {
			if n<1000 {
				fmt.Println("Число меньше 1000")
			}else {
				fmt.Println("Число больше или равно 1000")
			}
		}
	}
}