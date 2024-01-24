package models

import (
	"errors"

	"example.com/test/db"
	"example.com/test/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	query := "INSERT INTO users(email, password) VALUES($1, $2)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil

}
