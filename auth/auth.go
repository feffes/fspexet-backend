package auth

import (
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"encoding/json"
)

//User struct, might be placed in model in future? as AuthUser, different from normal User. Maybe same idk lol gotta think about it
type User struct {
	Username   string `form:"userid" json:"userid" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
// JwtToken is holds token string that goes back to the client
type JwtToken struct {
	Token string `json:"token"`
}
// Exception holds a message for use in json response during auth
type Exception struct {
	Message string `json:"message"`
}


// CreateToken creates new token for users on login
func CreateToken(w http.ResponseWriter, r *http.Request) {
	var u User
	_ = json.NewDecoder(r.Body).Decode(&u)
	// never store password in token!
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
	})
	// secret argument should be a special and unique string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

// VerifyToken runs as middleware before "next" to auth the user
func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	// return a HandlerFunc that performs authourisation
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token as an authorisation header from request header
		authorisationHeader := r.Header.Get("authorization")
		// make sure authorisation header is not empty
		if authorisationHeader != "" {
			// We only need to check that authorisation header is not empty
			token, err := jwt.Parse(authorisationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				json.NewEncoder(w).Encode(Exception{Message: err.Error()})
				return
			}
			if token.Valid {
				context.Set(r, "decoded", token.Claims)
				next(w, r)
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}