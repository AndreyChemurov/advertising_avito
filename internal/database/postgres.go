package database

import (
	"advertising_avito/internal/types"
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

func (p *postgres) Create(name string, desc string, links []string, price float64) (int, error) {
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

func (p *postgres) GetOne(id int, fields bool) (string, float64, string, []string, error) {
	var (
		stmt *sql.Stmt
		rows *sql.Rows
	)

	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return "", 0, "", []string{}, err
	}

	// Откат будет игнориться, если был коммит
	// TODO: correct rollback with error handling
	defer tx.Rollback()

	// Подготовка к селекту
	stmt, err = tx.Prepare("SELECT name, link, price, description FROM advertisement INNER JOIN photos ON (advertisement.id=$1 and adv_id=$1);")

	if err != nil {
		return "", 0, "", []string{}, err
	}

	if rows, err = stmt.Query(id); err != nil {
		return "", 0, "", []string{}, err
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
			return "", 0, "", []string{}, err
		}

		allLinks = append(allLinks, link)
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return "", 0, "", []string{}, err
	}

	return name, price, description, allLinks, nil
}

func (p *postgres) GetAll(page int, sort string) ([]types.Advertisement, error) {
	var (
		rows *sql.Rows
		stmt *sql.Stmt
	)

	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return nil, err
	}

	// Откат будет игнориться, если был коммит
	// TODO: correct rollback with error handling
	defer tx.Rollback()

	// Подготовка к селекту
	SQLString := fmt.Sprintf(`
	SELECT name, link, price FROM (
		SELECT DISTINCT ON (p.adv_id) name, link, price, created_at 
		FROM advertisement a INNER JOIN photos p on (a.id=p.adv_id)) subquery
	ORDER BY %s OFFSET %d - 1 LIMIT 10;
	`, sort, page)

	stmt, err = tx.Prepare(SQLString)

	if err != nil {
		return nil, err
	}

	// Получение из выборки
	if rows, err = stmt.Query(); err != nil {
		return nil, err
	}

	var advertisements = make([]types.Advertisement, 0)

	for rows.Next() {
		var (
			name  string
			link  string
			price float64

			adv types.Advertisement
		)

		if err = rows.Scan(&name, &link, &price); err != nil {
			// TODO: correct rollback with error handling
			_ = tx.Rollback()
			return nil, err
		}

		adv = types.Advertisement{
			Name:     name,
			MainLink: link,
			Price:    price,
		}

		advertisements = append(advertisements, adv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return advertisements, nil
}
