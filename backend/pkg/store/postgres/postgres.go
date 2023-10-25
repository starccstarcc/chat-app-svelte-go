package postgres

import "database/sql"

const (
	PostgresDNS = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

type postgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
