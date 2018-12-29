package main

import (
    "fmt"
    "log"
    "net/http"
    "git.f-spexet.se/feffe/fspexet-backend/models"
)

type Env struct {
    db models.Datastore
}

func main() {
    db, err := models.NewDB("postgres://fspex:godtyckligt@postgres:5432/fspex?sslmode=disable")
    if err != nil {
        log.Panic(err)
    }

    env := &Env{db}
	http.HandleFunc("/news", env.newsIndex)
    http.ListenAndServe(":5000", nil)
}

func (env *Env) newsIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), 405)
        return
    }
    bks, err := env.db.AllNews()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, %s, %s \n", bk.ID, bk.Title, bk.Author, bk.Content, bk.Time)
    }
}