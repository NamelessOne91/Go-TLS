package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	caBytes, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	ca := x509.NewCertPool()
	if !ca.AppendCertsFromPEM(caBytes) {
		log.Fatal("ca.cert not valid")
	}

	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      ca,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	res, err := client.Get("https://go-demo.localtest.me/common-name")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Printf("Body (status %d): %s\n", res.StatusCode, body)
}
