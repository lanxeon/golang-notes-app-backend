package config

import (
	"log"
	"os"

	pg "github.com/go-pg/pg/v9"

	controllers "notes_app/controllers"
)

// Connecting to db
func Connect() *pg.DB {

	opts := &pg.Options{
		User:     "postgres",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "notes",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")
	controllers.CreateNotesTable(db) // create table
	controllers.InitiateDB(db)       // pass the DB instance reference

	return db
}
