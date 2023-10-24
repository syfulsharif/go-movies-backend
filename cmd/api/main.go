package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type applicatoin struct {
	Domain string
}

func main() {
	//set application config
	var app applicatoin
	//read from command line

	//connect to the database
	app.Domain = "example.com"
	log.Println("Starting Application on port", port)
	http.HandleFunc("/", Hello)
	//start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
