package data

import "frontendmasters.com/reelingit/models"

type MovieStorage interface {
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	GetMovieById(id int) (models.Movie, error)
	SearchMoviesByName(name string) (models.Movie, error)
	GetAllGenres() ([]models.Genre, error)
}
