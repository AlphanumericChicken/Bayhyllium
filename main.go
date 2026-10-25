package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http_GET_print("http://roblox.com/")
}

func http_GET_print(URL string) {
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
