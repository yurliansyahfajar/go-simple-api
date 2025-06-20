package models

import (
	"fmt"

	"github.com/yurliansyahfajar/go-simple-api/db"
	"github.com/yurliansyahfajar/go-simple-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}

	defer stmt.Close()

	hashedPassowrd, err := utils.HashPassword(u.Password)

	if err != nil {
		return fmt.Errorf("hashing password error: %w", err)
	}

	result, err := stmt.Exec(u.Email, hashedPassowrd)

	if err != nil {
		return fmt.Errorf("error exec insert: %w", err)
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("error call last insert ID: %w", err)
	}

	u.ID = userId
	return nil

}
