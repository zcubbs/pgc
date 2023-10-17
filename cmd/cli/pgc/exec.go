package pgc

import (
	"database/sql"
	"fmt"
)

func CreateDatabase(db *sql.DB, dbName, owner string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s WITH OWNER %s;", dbName, owner))
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}
	return nil
}

func CreateUser(db *sql.DB, username, password string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE USER %s WITH ENCRYPTED PASSWORD '%s';", username, password))
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func GrantPrivileges(db *sql.DB, dbName, username, privileges string) error {
	_, err := db.Exec(fmt.Sprintf("GRANT %s ON DATABASE %s TO %s;", privileges, dbName, username))
	if err != nil {
		return fmt.Errorf("error granting privileges: %w", err)
	}
	return nil
}
