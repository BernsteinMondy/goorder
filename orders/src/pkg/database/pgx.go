package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewSQL(c Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", connString(c))
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	return db, nil
}

func connString(c Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}
