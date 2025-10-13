package database

import (
	"database/sql"
	"fmt"
	"log"
	"proyecto-bd-final/config"

	_ "github.com/godror/godror"
)

var DB *sql.DB

func ConnectDatabase(cfg *config.Config) error {
	connStr := fmt.Sprintf(
		"%s/%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBService,
	)

	var err error
	DB, err = sql.Open("godror", connStr)
	if err != nil {
		return fmt.Errorf("error opening db: %w", err)
	}

	err = DB.Ping()

	if err != nil {
		return fmt.Errorf("error connecting to db: %w", err)
	}

	log.Println("Successfully connected to Oracle Database")
	return nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
