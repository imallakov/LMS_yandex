package main

import (
    "fmt"
    "math"
)

func SqRoots() {
    var a, b, c float64
    fmt.Scanln(&a, &b, &c)
	D:=math.Pow(b,2)-4*a*c;
	if D<0 {
		fmt.Println(0," ",0)
		return
	}
	var x_1, x_2 float64
	x_1=(-b+math.Sqrt(D))/(2*a)
	x_2=(-b-math.Sqrt(D))/(2*a)
	if x_1>x_2 {
		temp:=x_1
		x_1=x_2
		x_2=temp
	}
	fmt.Printf("%f",x_1)
	if x_1!=x_2 {
		fmt.Printf(" %f\n",x_2)
	} else {
		fmt.Print("\n")
	}
}

func main() {
	SqRoots()
}