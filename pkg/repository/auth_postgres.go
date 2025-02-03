package repository

import (
	"eduProject"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user eduProject.User) (int, error) {
	logrus.Error("CreateUser called without a user")
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, passwordhash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (eduProject.User, error) {
	var user eduProject.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND passwordhash = $2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
