package store

import (
	"database/sql"
	"time"

	"github.com/kirildev25/go-api-pgx-postgres/internal/tokens"
)

type PostgresTokenStore struct {
	db *sql.DB
}

func NewPostgresTokenStore(db *sql.DB) *PostgresTokenStore {
	return &PostgresTokenStore{
		db: db,
	}
}

type TokenStore interface {
	Insert(token *tokens.Token) error
	CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error)
	DeleteAllTokensForUser(userID int, scope string) error
}

func (t *PostgresTokenStore) Insert(token *tokens.Token) error {
	return nil
}
func (t *PostgresTokenStore) CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error) {
	return nil, nil
}
func (t *PostgresTokenStore) DeleteAllTokensForUser(userID int, scope string) error {
	return nil
}
