package main

import (
	"fmt"
	"math/rand"

	"github.com/YuriiTseplitskyi/puppy"
	//"studyingGO/mypackage"
)

func From13() {
	fmt.Println("I am from version 1.3.0")
}

func puppytest() {
	s1 := puppy.Barks()
	fmt.Println(s1)
}

func okidiom(){
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not present")
	}
	arv := m["dsfsd"]
	fmt.Println(arv)
	fmt.Println(m)

	counter := 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("found 3")
			counter++
		}

	}
	fmt.Println(counter)
}

func main() {

	Arrays2()
}
