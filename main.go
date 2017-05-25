package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Getting the port on which the server will listen to
	port := os.Getenv("PORT")

	// If no port exists (to run on localhost) use port 1337
	if port == "" {
		port = "1337"
	}

	http.HandleFunc("/hello", AppHandle.HelloHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
