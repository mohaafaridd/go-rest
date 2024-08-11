package models

import (
	"errors"

	"example.com/events/db"
	"example.com/events/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashPassword)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string

	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !isValidPassword {
		return errors.New("invalid credentials")
	}

	return nil
}
