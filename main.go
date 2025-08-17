package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8017"
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
