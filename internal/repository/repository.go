package repository

import (
	"database/sql"

	"github.com/syfulsharif/go-movies-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
