package main

import (
	"fmt"
	"os"
	"github.com/fspexet/fspexet-backend/auth"
	"encoding/json"
    "log"
    "net/http"
	"github.com/fspexet/fspexet-backend/models"
	"github.com/gorilla/mux"
)
//Env holds the datastore
type Env struct {
    db models.Datastore
}

func main() {
	//TODO, handle args and flags
	dbHost 	:= os.Getenv("POSTGRES_HOST")
	dbUser 	:= os.Getenv("POSTGRES_USER")
	dbPass 	:= os.Getenv("POSTGRES_PASSWORD")
	dbName 	:= os.Getenv("POSTGRES_DB")
	dbSsl 	:= os.Getenv("POSTGRES_SSL")
	//keyDir := os.Getenv("FSPEXET_BACKEND_KEYSDIR")

	options := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUser, dbPass, dbName, dbSsl)

    //db, err := models.NewDB("postgres://fspex:godtyckligt@postgres:5432/fspex?sslmode=disable")
    db, err := models.NewDB(options)
	if err != nil {
        log.Panic(err)
    }

	err = auth.GenRS4096KeyPair()
	if err != nil {
		log.Println(err)
	}

	env := &Env{db}
	models.DataBase = env.db

	m := mux.NewRouter()
	m.HandleFunc("/news", env.newsIndex)
	m.HandleFunc("/auth", auth.VerifyToken(env.authIndex)).Methods("GET")
	m.HandleFunc("/token", auth.LoginMiddleware(auth.CreateToken)).Methods("POST")
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
}

func (env *Env) authIndex(w http.ResponseWriter, req *http.Request) {
	log.Println("authed :)")
}

