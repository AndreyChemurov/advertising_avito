package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" //
)

// postgres - реализация интерфейса Database для postgres.
type postgres struct {
	db *sql.DB
}

var instance *postgres

// Констуктор для postgres.
func newPostgres() *postgres {
	if instance != nil {
		return instance
	}

	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	instance = &postgres{
		db: db,
	}

	return instance
}

func (p *postgres) Create(ctx context.Context, name string, desc string, links []string, price float64) (int, error) {
	var id int

	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return 0, err
	}

	// Откат будет игнориться, если был коммит
	// TODO: correct rollback with error handling
	defer tx.Rollback()

	// Подготовка к добавлению в таблицу "advertisement"
	stmt, err := tx.Prepare("INSERT INTO advertisement VALUES (DEFAULT, $1, $2, $3) RETURNING id;")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	// Добавление в таблицу "advertisement" с возвращением ID
	if err = stmt.QueryRow(name, desc, price).Scan(&id); err != nil {
		return 0, err
	}

	// Подготовка к добавлению в таблицу "photos"
	stmt, err = tx.Prepare("INSERT INTO photos VALUES (DEFAULT, $1, $2);")

	if err != nil {
		return 0, err
	}

	// Добавление в таблицу "photos"
	for _, link := range links {
		_, err = stmt.Exec(id, link)

		if err != nil {
			// TODO: correct rollback with error handling
			_ = tx.Rollback()
			return 0, err
		}
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *postgres) GetOne(ctx context.Context, id int, fields bool) (string, float64, string, string, []string, error) {
	var (
		stmt *sql.Stmt
		rows *sql.Rows
	)

	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return "", 0, "", "", []string{}, err
	}

	// Откат будет игнориться, если был коммит
	// TODO: correct rollback with error handling
	defer tx.Rollback()

	// Подготовка к селекту
	stmt, err = tx.Prepare("SELECT name, link, price, description FROM advertisement INNER JOIN photos ON (advertisement.id=$1 and adv_id=$1);")

	if err != nil {
		return "", 0, "", "", []string{}, err
	}

	if rows, err = stmt.Query(id); err != nil {
		return "", 0, "", "", []string{}, err
	}

	defer rows.Close()

	var (
		name        string
		description string
		link        string
		price       float64
		allLinks    = make([]string, 0, 3)
	)

	for rows.Next() {
		if err = rows.Scan(&name, &link, &price, &description); err != nil {
			// TODO: correct rollback with error handling
			_ = tx.Rollback()
			return "", 0, "", "", []string{}, err
		}

		allLinks = append(allLinks, link)
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return "", 0, "", "", []string{}, err
	}

	return name, price, allLinks[0], description, allLinks, nil
}

func (p *postgres) GetAll(ctx context.Context, page int, sort string) (string, float64, string, string, []string, error) {
	//
	return "", 0, "", "", []string{}, nil
}
