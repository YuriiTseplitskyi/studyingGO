package main

import (
	"fmt"
	"github.com/YuriiTseplitskyi/puppy"
)

func From13() {
	fmt.Println("I am from version 1.3.0")
}

func main(){
	fmt.Println("Hello World2")
	s1:=puppy.Barks()
	fmt.Println(s1)
}