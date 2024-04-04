// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )

// const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

// func main(){

// 	fmt.Println("LCO web request")

// 	responce, err := http.Get(myUrl)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("responce is 1", responce)

// 	defer responce.Body.Close()

// 	databytes, err := ioutil.ReadAll(responce.Body)

// 	if err != nil {
// 		panic(err)
// 	}

// 	content := string(databytes)

// 	fmt.Println("responce is ", content)

// 	// parsing
// 	result, _ := url.Parse(myUrl)

// 	fmt.Println("Scheme", result.Scheme)
// 	fmt.Println("Host", result.Host)
// 	fmt.Println("Path", result.Path)
// 	fmt.Println("Port", result.Port())
// 	fmt.Println("RawQuery", result.RawQuery)
// }


package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a route and its handler function
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

