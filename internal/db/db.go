package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zhetkerbaevan/statistics_collection/internal/config"
)

func NewPostgreSQLStorage(cfg config.Config) (*sql.DB, error) {
	//Open connection to database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword))
	if err != nil {
		return nil, err
	}

	//Verify connection to database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}