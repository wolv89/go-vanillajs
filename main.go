package main

import (
	"fmt"
	"log"
	"net/http"

	"frontendmasters.com/reelingit/logger"
)

func initialiseLogger() *logger.Logger {

	l, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("unable to set up logger: %v", err)
	}
	return l

}

func main() {

	// Server Logger
	slog := initialiseLogger()
	defer slog.Close()

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8017"

	fmt.Println("starting server with address:", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		slog.Error("server failed", err)
		log.Fatalf("server failed: %v", err)
	}

}
