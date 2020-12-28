package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"os"
)

func main() {

	// json sample backend
	// https://jsonplaceholder.typicode.com/users
	// https://jsonplaceholder.typicode.com/todos
	// https://jsonplaceholder.typicode.com/posts
	// https://jsonplaceholder.typicode.com/comments

	// get command line argument
	userID := os.Args[1]

	// make a sample HTTP GET request
	res, err := http.Get("http://jsonplaceholder.typicode.com/users/" + string(userID))

	// check for response error
	if err != nil {
		log.Fatal( err )
	}

	// read all response body
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()

	// print `data` as string
	fmt.Printf("%s\n",data)

}