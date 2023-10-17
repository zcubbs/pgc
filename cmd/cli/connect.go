package main

import (
	"database/sql"
	"fmt"
)

func connect(config Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		config.Postgresql.Host, config.Postgresql.Port, config.Postgresql.User, config.Postgresql.Password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}
	return db, nil
}
