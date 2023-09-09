package main

import "fmt"

func main() {
	var n, ans int = 0,1
	fmt.Scan(&n)
	for(n>0){
		ans*=n
		n--
	}
	fmt.Println(ans)
}