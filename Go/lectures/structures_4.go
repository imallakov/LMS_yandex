package main

import "fmt"

type Dog struct {}

type Cat struct {}

type Animal interface {
	MakeSound()
}

func (Dog) MakeSound(){
	fmt.Println("Гав!")
}

func (Cat) MakeSound(){
	fmt.Println("Мяу!")
}
