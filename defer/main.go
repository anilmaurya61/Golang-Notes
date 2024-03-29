package main

import "fmt"

func main(){
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("four")
	myDeferFun()
}

// First in last out

// "one", "two", "three"

func myDeferFun(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
