package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos"

func main(){

	fmt.Println("LCO web request")

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("responce is 1", responce)

	defer responce.Body.Close()

	databytes, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("responce is ", content)

}