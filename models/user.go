package models

import (
	"time"
)

// User does stuff i guess lol
type User struct{
	ID string
	Name string
	Password string
	LastLogin time.Time
	CreationTime time.Time
}

func (db *DB) UserID(id string) (User, error){
	row := db.QueryRow("SELECT * FROM users WHERE ID=$1", id )
	usr := User{}
	err := row.Scan(&usr.ID, &usr.Name, &usr.Password, &usr.LastLogin, &usr.CreationTime)
	if err != nil {
		return usr, err
	}
	return usr, nil
}