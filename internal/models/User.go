package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
}

func (u *User) PasswordMatches(plaintText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(plaintText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) :
			// invalid password
			return false, nil
		default:
			return false, err

		}
	}
	return true,nil


}