package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Working...")
}

func main() {
	http.HandleFunc("/", index)

	err := http.Serve(autocert.NewListener("my-host.address"), nil)
	if err != nil {
		log.Fatalf("Error starting TLS enabled server: %v", err)
	}
}
