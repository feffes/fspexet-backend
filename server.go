package main

import (
	"git.f-spexet.se/feffe/fspexet-backend/auth"
	"encoding/json"
    "log"
    "net/http"
	"git.f-spexet.se/feffe/fspexet-backend/models"
	"github.com/gorilla/mux"
)

type Env struct {
    db models.Datastore
}

func main() {
	//TODO, handle args and flags
    db, err := models.NewDB("postgres://fspex:godtyckligt@postgres:5432/fspex?sslmode=disable")
    if err != nil {
        log.Panic(err)
    }

	env := &Env{db}
	
	m := mux.NewRouter()
	m.HandleFunc("/news", env.newsIndex)
	m.HandleFunc("/auth", auth.VerifyToken(env.authIndex))
	m.HandleFunc("/token", auth.CreateToken)
	log.Println("starting on :5000")
	http.Handle("/", m)
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
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
	js, err := json.Marshal(bks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

		//fmt.Fprintf(w, "%s, %s, %s, %s, %s \n", bk.ID, bk.Title, bk.Author, bk.Content, bk.Time)
}

func (env *Env) authIndex(w http.ResponseWriter, req *http.Request) {
	log.Println("authed :)")
}
