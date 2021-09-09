package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// get data

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}

	//read data

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//print response in terminal

	sb := string(body)
	log.Printf(sb)
}
