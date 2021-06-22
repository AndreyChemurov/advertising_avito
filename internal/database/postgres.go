package database

import (
	"advertising_avito/internal/types"
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

func (p *postgres) Create(ctx context.Context, name string, desc string, links []string, price float64) (id int, err error) {
	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return
	}

	// Откат будет игнориться, если был коммит
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("error on create: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Подготовка к добавлению в таблицу "advertisement"
	stmt, err := tx.Prepare("INSERT INTO advertisement VALUES (DEFAULT, $1, $2, $3) RETURNING id;")

	if err != nil {
		return
	}

	defer stmt.Close()

	// Добавление в таблицу "advertisement" с возвращением ID
	if err = stmt.QueryRow(name, desc, price).Scan(&id); err != nil {
		return
	}

	// Подготовка к добавлению в таблицу "photos"
	stmt, err = tx.Prepare("INSERT INTO photos VALUES (DEFAULT, $1, $2);")

	if err != nil {
		return
	}

	// Добавление в таблицу "photos"
	for _, link := range links {
		_, err = stmt.Exec(id, link)

		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("error on get-one: %v. error on rollback: %v", err, rerr)
			}
			return
		}
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return
	}

	return id, nil
}

func (p *postgres) GetOne(ctx context.Context, id int, fields bool) (name string, price float64, desc string, allLinks []string, err error) {
	var (
		stmt *sql.Stmt
		rows *sql.Rows
	)

	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return
	}

	// Откат будет игнориться, если был коммит
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("error on get-one: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Подготовка к селекту
	stmt, err = tx.Prepare("SELECT name, link, price, description FROM advertisement INNER JOIN photos ON (advertisement.id=$1 and adv_id=$1);")

	if err != nil {
		return
	}

	if rows, err = stmt.Query(id); err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var link string
		if err = rows.Scan(&name, &link, &price, &desc); err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("error on get-one: %v. error on rollback: %v", err, rerr)
			}
			return
		}

		allLinks = append(allLinks, link)
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (p *postgres) GetAll(ctx context.Context, page int, sort string) (advs []types.Advertisement, err error) {
	var (
		rows *sql.Rows
		stmt *sql.Stmt
	)

	// Начало транзакции
	tx, err := p.db.Begin()

	if err != nil {
		return
	}

	// Откат будет игнориться, если был коммит
	defer func() {
		if err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("error on get-all: %v. error on rollback: %v", err, rerr)
			}
		}
	}()

	// Подготовка к селекту
	SQLString := fmt.Sprintf(`
	SELECT name, link, price FROM (
		SELECT DISTINCT ON (p.adv_id) name, link, price, created_at 
		FROM advertisement a INNER JOIN photos p on (a.id=p.adv_id)) subquery
	ORDER BY %s OFFSET %d - 1 LIMIT 10;
	`, sort, page)

	stmt, err = tx.Prepare(SQLString)

	if err != nil {
		return
	}

	// Получение из выборки
	if rows, err = stmt.Query(); err != nil {
		return
	}

	for rows.Next() {
		var (
			name  string
			link  string
			price float64

			adv types.Advertisement
		)

		if err = rows.Scan(&name, &link, &price); err != nil {
			rerr := tx.Rollback()
			if rerr != nil {
				err = fmt.Errorf("error on get-all: %v. error on rollback: %v", err, rerr)
			}
			return
		}

		adv = types.Advertisement{
			Name:     name,
			MainLink: link,
			Price:    price,
		}

		advs = append(advs, adv)
	}

	if err = rows.Err(); err != nil {
		return
	}

	// Коммит
	if err = tx.Commit(); err != nil {
		return
	}

	return
}
