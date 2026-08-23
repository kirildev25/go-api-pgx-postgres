package store

import (
	"database/sql"
	"time"

	_ "golang.org/x/crypto/bcrypt"
)

type password struct {
	plainText *string
	hash      []byte
}

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Bio          string    `json:"bio"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserStore interface {
	CreateUser(*User) error
	GetUserByUsername(username string) (*User, error)
	UpdateUser(*User) error
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

func (s *PostgresUserStore) CreateUser(user *User) error                      { return nil }
func (s *PostgresUserStore) GetUserByUsername(username string) (*User, error) { return nil, nil }
func (s *PostgresUserStore) UpdateUser(*User) error                           { return nil }
