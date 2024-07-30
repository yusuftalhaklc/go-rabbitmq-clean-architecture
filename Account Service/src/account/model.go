package account

import "time"

type Account struct {
	Id            string
	FullName      string
	PasswordHash  string
	Email         string
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
