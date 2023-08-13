package handlers

import "github.com/feynmaz/fiberg/models"

type UsersStorage interface {
	Get() []*models.User
	Insert(user *models.User)
}
