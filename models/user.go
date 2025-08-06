package models

import (
	"github.com/slimreaper35/konflux-test/database"
	"github.com/slimreaper35/konflux-test/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Create() error {
	var query = `
	INSERT INTO users(email, password)
	VALUES (?, ?);
	`
	hashedPassword, err := utils.Hash(u.Password)
	if err != nil {
		panic(err)
	}

	result, err := database.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetUserBy(email string) (*User, error) {
	var query = `
	SELECT *
	FROM users
	WHERE email = ?;
	`
	var row = database.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
