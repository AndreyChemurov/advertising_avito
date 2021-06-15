package database

import "database/sql"

type postgres struct {
	db *sql.DB
}

func newPostgres(db *sql.DB) *postgres {
	return &postgres{
		db: db,
	}
}

func (p *postgres) Create(name string, desc string, links []string, price float64) (int, error) {
	//
	return 0, nil
}

func (p *postgres) GetOne(id int, fields bool) (string, float64, string, string, []string) {
	//
	return "", 0, "", "", []string{}
}

func (p *postgres) GetAll(page int, sort string) {

}
