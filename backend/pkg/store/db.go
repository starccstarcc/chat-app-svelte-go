package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB(driver string, dns string) (*sql.DB, error) {
	db, err := sql.Open(driver, dns)

	if err != nil {
		return nil, err
	}

	return db, nil
}
