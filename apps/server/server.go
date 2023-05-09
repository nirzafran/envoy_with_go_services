package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// response hello world
	w.WriteHeader(http.StatusOK)
	currentTime := time.Now().Format(time.RFC850)
	w.Write([]byte("hello world " + currentTime))
}

func main() {
	const port = 8080
	fmt.Printf("Starting server on port %d", port)
	http.HandleFunc("/", rootHandler)
	// convert port to string
	// start the server on port 8080
	serverPort := ":" + strconv.Itoa(port)
	http.ListenAndServe(serverPort, nil)
}
