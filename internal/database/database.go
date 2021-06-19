package database

import (
	"context"
	"errors"
)

// Database - интефейс базы данных.
// Описывает методы:
//	Create - добавить объявление;
//	GetOne - получить информацию по конкретному объявлению;
//	GetAll - получить информацию по всем объявлениям.
type Database interface {
	Create(context.Context, string, string, []string, float64) (int, error)
	GetOne(context.Context, int, bool) (string, float64, string, string, []string, error)
	GetAll(context.Context, int, string) (string, float64, string, string, []string, error)
}

var databases = map[string]Database{
	"postgres": newPostgres(),
}

// GetDatabase - фабличный метод (здесь только для postgres)
// получения объекта базы данных.
func GetDatabase(driver string) (Database, error) {
	database, found := databases[driver]
	if !found {
		return nil, errors.New("no such database driver")
	}
	return database, nil
}
