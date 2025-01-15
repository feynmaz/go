package database

import (
	"sync"

	"github.com/feynmaz/fiberg/models"
	"go.uber.org/zap"
)

type Database struct {
	storage []*models.User
	mu      sync.Mutex
}

func Connect() *Database {
	s := make([]*models.User, 0)
	zap.L().Info("Connected with Database")
	return &Database{storage: s}
}

func (db *Database) Insert(user *models.User) {
	db.mu.Lock()
	db.storage = append(db.storage, user)
	db.mu.Unlock()
}

func (db *Database) Get() []*models.User {
	return db.storage
}
