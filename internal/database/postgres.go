package database

import (
	"database/sql"
	"os"
)

// postgres - реализация интерфейса Database для postgres.
type postgres struct {
	db *sql.DB
}

var instance *postgres

// Констуктор для postgres.
func newPostgres(driver string) *postgres {
	if instance != nil {
		return instance
	}

	// TODO:
	//	1. add error handling
	//	2. close connection
	db, _ := sql.Open(driver, os.Getenv("DATABASE_URL"))
	instance = &postgres{
		db: db,
	}

	return instance
}

func (p *postgres) Create(name string, desc string, links []string, price float64) (int, error) {
	//
	return 1000, nil
}

func (p *postgres) GetOne(id int, fields bool) (string, float64, string, string, []string, error) {
	//
	return "", 0, "", "", []string{}, nil
}

func (p *postgres) GetAll(page int, sort string) (string, float64, string, string, []string, error) {
	//
	return "", 0, "", "", []string{}, nil
}
