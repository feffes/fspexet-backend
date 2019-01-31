package models

import (
	"time"
)

type News struct {
    ID   	string
    Title  	string
    Author 	string
	Content	string
	Time  	time.Time
}

func (db *DB) AllNews() ([]*News, error) {
    rows, err := db.Query("SELECT * FROM news")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    bks := make([]*News, 0)
    for rows.Next() {
        bk := new(News)
        err := rows.Scan(&bk.ID, &bk.Title, &bk.Author, &bk.Content, &bk.Time)
        if err != nil {
            return nil, err
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return bks, nil
}

func (db *DB) NewsID(id string) (News, error) {
	row := db.QueryRow("SELECT * FROM news WHERE ID=$1", id)
	nws := News{}
	err := row.Scan(&nws.ID, &nws.Title, &nws.Author, &nws.Content, &nws.Time)
	if err != nil {
		return nws, err
	}
	return nws, nil
}