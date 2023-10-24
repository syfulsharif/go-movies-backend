package repository

import "github.com/syfulsharif/go-movies-backend/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}
