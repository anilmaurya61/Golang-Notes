package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){

	fmt.Println("Welcome to Files in golang")

	content := "What is Lorem Ipsum?"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("File length is : ", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))

}