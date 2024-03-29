package main

import "fmt"

func main(){
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("four")
}

// First in last out

// "one", "two", "three"
