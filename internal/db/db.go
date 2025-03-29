package db

import "github.com/jmoiron/sqlx"

func NewDB() (*sqlx.DB, error) {
	connStr := "user=postgres password=password dbname=postgres sslmode=disable host=localhost port=5432"
	return sqlx.Connect("postgres", connStr)
}
