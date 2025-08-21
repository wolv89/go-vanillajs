package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"frontendmasters.com/reelingit/data"
	"frontendmasters.com/reelingit/handlers"
	"frontendmasters.com/reelingit/logger"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
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

	if err := godotenv.Load(); err != nil {
		slog.Error("tried to boot without .env file", err)
		log.Fatalf("unable to find .env file %v", err)
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		slog.Error("DATABASE_URL not set in .env", nil)
		log.Fatalf("DATABASE_URL not set in .env")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		slog.Error("could not connect to database:", err)
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	movieRepo, _ := data.NewMovieRepository(db, slog)

	mh := handlers.MovieHandler{
		Storage: movieRepo,
		Logger:  slog,
	}
	http.HandleFunc("/api/movies/top", mh.GetTopMovies)
	http.HandleFunc("/api/movies/random", mh.GetRandomMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8017"

	fmt.Println("starting server with address:", addr)
	err = http.ListenAndServe(addr, nil)

	if err != nil {
		slog.Error("server failed", err)
		log.Fatalf("server failed: %v", err)
	}

}
