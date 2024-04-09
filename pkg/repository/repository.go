package repository

import "github.com/jmoiron/sqlx"

type Banner interface {
}

type Tag interface {
}

type Feature interface {
}

type Repository struct {
	Banner
	Tag
	Feature
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
