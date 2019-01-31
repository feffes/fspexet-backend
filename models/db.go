package models

import (
    _ "github.com/lib/pq" // postgres driver
    "database/sql"
)

var DataBase Datastore

type Datastore interface {
	AllNews() ([]*News, error)
	UserID(id string) (User, error)
}

type DB struct {
    *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}