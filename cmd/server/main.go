package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Working...")
}

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Error starting TLS enabled server: %v", err)
	}
}
