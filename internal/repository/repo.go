package repository

import (
	"database/sql"
	"greenlight/internal/data"
)

type Repository struct {
	MovieRepo data.MovieRepo
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		MovieRepo: data.MovieRepo{DB: db},
	}
}
