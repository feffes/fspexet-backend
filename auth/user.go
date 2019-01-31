package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fspexet/fspexet-backend/models"
)

type AuthUser struct {
	ID       string
	Password string
}

// CheckCredentials checks if the given credentials for an authenticating user are correct. SUPER BAD IMPLEMENTATION DO NOT PUSH I REPEAT DO NOT PUSH
func checkCredentials(ausr AuthUser, usr models.User) bool {
	return (ausr.ID == usr.ID && ausr.Password == usr.Password)
}

func LoginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var ausr AuthUser
		err := decoder.Decode(&ausr)
		if err != nil {
			// ????????
		}
		log.Printf("%v", ausr)
		usr, err := models.DataBase.UserID(ausr.ID)
		if err != nil {
			log.Println(usr)
			w.WriteHeader(500)
		} else {
			if checkCredentials(ausr, usr) {
				log.Println(checkCredentials(ausr, usr))
				next(w, r)
			} else {
				w.WriteHeader(403)
			}
		}

	})
}
