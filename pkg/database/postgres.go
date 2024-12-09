package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewPostgresConnection(cfg Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InitDatabase(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS buildings (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        city VARCHAR(100) NOT NULL,
        year INTEGER NOT NULL,
        floor_count INTEGER NOT NULL
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return err
	}

	return nil
}
