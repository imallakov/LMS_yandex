package main

import "fmt"

func main() {
	var a,b,c, count int
	fmt.Scanln(&a,&b,&c)
	if a<0 || b<0 || c<0 {
		fmt.Println("Некорректный ввод")
		return
	}
	if(a==b) {
		count++
	}
	if b==c {
		count++
	}
	if a==c {
		count++
	}
	if(count>=2){
		fmt.Println("Все числа равны")
	} else {
		if count==1 {
			fmt.Println("Два числа равны")
		} else {
			fmt.Println("Все числа разные")
		}
	}
}