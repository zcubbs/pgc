package pgc

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateDatabase(db *sql.DB, dbName, owner string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s WITH OWNER %s;", dbName, owner))
	if err != nil {
		if err.Error() == fmt.Sprintf("pq: database \"%s\" already exists", dbName) {
			log.Printf("database %s already exists, skipping...", dbName)
			return nil
		}
		return fmt.Errorf("error creating database: %w", err)
	}
	return nil
}

func CreateUser(db *sql.DB, username, password string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE USER %s WITH ENCRYPTED PASSWORD '%s';", username, password))
	if err != nil {
		if err.Error() == fmt.Sprintf("pq: role \"%s\" already exists", username) {
			log.Printf("user %s already exists, skipping...", username)
			return nil
		}
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func GrantPrivileges(db *sql.DB, dbName, username, privileges string) error {
	_, err := db.Exec(fmt.Sprintf("GRANT %s ON DATABASE %s TO %s;", privileges, dbName, username))
	if err != nil {
		if err.Error() == fmt.Sprintf("pq: role \"%s\" does not exist", username) {
			log.Printf("user %s does not exist, skipping...", username)
			return nil
		}
		return fmt.Errorf("error granting privileges: %w", err)
	}
	return nil
}
