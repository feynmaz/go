package handlers

import "github.com/feynmaz/fiberg/database"

type Handler struct {
	db UsersStorage
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{
		db: db,
	}
}
