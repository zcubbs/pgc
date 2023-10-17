package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/zcubbs/pgc/cmd/cli/pgc"
	"log"
)

var (
	Version = "0.0.0"
	Commit  = "none"
	Date    = "unknown"
)
var (
	configPath = flag.String("c", "", "Path to config file")
)

func main() {
	flag.Parse()

	//Print out version information
	fmt.Printf("pgc %s %s %s\n", Version, Commit, Date)

	if *configPath == "" {
		log.Fatalf("config file path not specified.\nUsage: pgc -c /path/to/config.yaml")
	}

	log.Printf("Loading configuration: %s", *configPath)
	config, err := pgc.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	log.Println("Connecting to the database...")
	db, err := pgc.Connect(config)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("could not close db: %v", err)
		}
	}(db)

	// First, create users
	log.Println("Creating users...")
	for _, user := range config.Users {
		log.Printf("Creating user: %s", user.Name)
		err := pgc.CreateUser(db, user.Name, user.Password)
		if err != nil {
			log.Fatalf("could not create user %s: %v", user.Name, err)
		}
	}

	// Then, create databases and assign owners
	log.Println("Creating databases and assigning owners...")
	for _, database := range config.Databases {
		log.Printf("Creating database: %s with owner: %s", database.Name, database.Owner)
		err := pgc.CreateDatabase(db, database.Name, database.Owner)
		if err != nil {
			log.Fatalf("could not create db %s: %v", database.Name, err)
		}
	}

	// Then, grant privileges
	log.Println("Granting privileges...")
	for _, privilege := range config.Privileges {
		log.Printf("Granting privileges on database: %s to user: %s", privilege.Database, privilege.User)
		err := pgc.GrantPrivileges(db, privilege.Database, privilege.User, privilege.Privileges)
		if err != nil {
			log.Fatalf("could not grant privileges: %v", err)
		}
	}

	log.Println("Updating pg_hba.conf...")
	err = pgc.UpdatePgHba(config.PgHba, config.PgHbaConfPath)
	if err != nil {
		log.Printf("could not update pg_hba: %v", err)
	}

	log.Println("Restarting PostgreSQL...")
	err = pgc.RestartPostgres(config.RestartCommand)
	if err != nil {
		log.Printf("could not restart postgres: %v", err)
	}

	log.Println("Operation completed successfully.")
}
