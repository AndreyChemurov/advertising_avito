package database

import (
	"testing"
)

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
