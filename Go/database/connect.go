package database

import (
	"database/sql"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func ConnectDb() (*sql.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
