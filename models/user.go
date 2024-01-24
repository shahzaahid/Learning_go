package models

import (
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
