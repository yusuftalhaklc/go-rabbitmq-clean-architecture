package infra

import (
	"account"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository() *AccountRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "123456"
	dbName := "AccountDB"
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("[X] Database Connection Error")
	}
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Create(account account.Account) error {
	query := `
	INSERT INTO users 
	(id, full_name, email, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query, account.Id, account.FullName, account.Email, account.PasswordHash, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) Update(email string, verified bool) error {
	query := `
	UPDATE users
	SET email_verified = $2 ,updated_at = $3
	WHERE email = $1
	`
	_, err := r.db.Exec(query, email, verified, time.Now())
	if err != nil {
		return err
	}

	return nil
}
