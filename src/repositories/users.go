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
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}
