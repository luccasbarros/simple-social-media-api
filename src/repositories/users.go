package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUsersRepositories create users repositories
func NewUsersRepositories(db *sql.DB) *users {
	return &users{db}
}

// Create insert user on database
func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}
