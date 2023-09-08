package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHellowWorld)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/new-endpoint", handleNewEndpoint)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHellowWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(
			writer,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
	}

	writeResponse(writer, "Hello World")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(
			writer,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
	}

	writeResponse(writer, "OK!")
}

func handleNewEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(
			writer,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
	}

	writeResponse(writer, "OK!")
}

func writeResponse(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	fmt.Println(response)
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
