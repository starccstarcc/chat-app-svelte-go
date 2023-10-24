package postgres

import "database/sql"

type postgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
