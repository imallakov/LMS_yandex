package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	var ch1,ch2 string
	var count1,count2 int
	fmt.Scanln(&ch1)
	fmt.Scanln(&ch2)
	reader := bufio.NewReader(os.Stdin)
    line, _ := reader.ReadString('\n')
	var length int = len([]rune(line))
	for i:=0;i<length;i++{
		if []rune(line)[i]==[]rune(ch1)[0] {
			count1++
		}
		if []rune(line)[i]==[]rune(ch2)[0] {
			count2++
		}
	}
	if count1>count2 {
		fmt.Println(ch1," ",count1)
		fmt.Println(ch2," ",count2)
	}else {
		fmt.Println(ch2," ",count2)
		fmt.Println(ch1," ",count1)
	}

}
