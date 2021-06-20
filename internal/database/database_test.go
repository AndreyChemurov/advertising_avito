package database

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq" //
)

func TestDatabaseGetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	query := "SELECT name, link, price, description FROM advertisement INNER JOIN photos ON (advertisement.id=$1 and adv_id=$1);"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	p := &postgres{
		db: db,
	}

	if _, _, _, _, err = p.GetOne(context.Background(), 1, true); err != nil {
		t.Errorf("error was not expected while get one: %v", err)
	}
}

func TestDatabaseCreateFail(t *testing.T) {
	_, err := GetDatabase("error-database")
	if err == nil {
		t.Errorf("database creation return wrong response: %w", err)
	}
}

func TestDatabaseOpenConnectionFail(t *testing.T) {
	_, err := GetDatabase("postgres")
	if err != nil {
		t.Errorf("database open connection return wrong response: %w", err)
	}
}
