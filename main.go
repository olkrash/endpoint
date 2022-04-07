package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func users(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	content, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	_, err = writer.Write(content)
	if err != nil {
		log.Fatal("Error writing output: ", err)
	}
}

func main() {
	http.HandleFunc("/users", users)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Error setting up the route: ", err)
	}
}
