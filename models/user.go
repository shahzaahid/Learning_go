package models

import "example.com/test/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES($1, $2)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}
