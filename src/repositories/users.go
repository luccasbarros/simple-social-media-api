package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// List get all users with name or nick filter
func (repository users) List(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %string%

	rows, erro := repository.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if erro = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, erro
}

func (repository users) GetById(ID uint64) (models.User, error) {
	rows, erro := repository.db.Query("select id, name, nick, email, created_at from users where id = ?", ID)

	if erro != nil {
		return models.User{}, erro
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if erro = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository users) Update(ID uint64, user models.User) error {
	statement, erro := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repository users) Delete(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
