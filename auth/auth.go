package auth

import (
	"crypto/rsa"
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"net/http"
	"encoding/json"
)
var(
	//PublicKey decryption key
	PublicKey *rsa.PublicKey//[]byte
	//PrivateKey encryption key
	PrivateKey *rsa.PrivateKey//[]byte
)

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
	var u AuthUser
	log.Println(string(encodePrivateKeyToPEM(PrivateKey)[:]))
	log.Println(string(encodePublicKeyToPEM(PublicKey)[:]))
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil{
		// HANDLE IT
	}
	//Handle auth here? probably in another method in the chain though
	// AKA Compare with database > this method > next and if it fucks up anywhere in the chain return 403
	
	signer:= jwt.New(jwt.GetSigningMethod("RS512"))
	//set claims
	claims := make(jwt.MapClaims)
	claims["id"]  = u.ID
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	signer.Claims = claims

	log.Println(signer)

	tokenString, err := signer.SignedString(PrivateKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error signing token: %v\n", err)
	}else {
		response := JwtToken{tokenString}
		JSONResponse(response, w)
	}

}

// VerifyToken runs as middleware before "next" to auth the user
func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	// return a HandlerFunc that performs authourisation
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//validate token
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return PublicKey, nil
			})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorised access to this resource")
			
		} else {
			if token.Valid {
				next(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid")
			}
		}
	})

}
// JSONResponse takes a struct for example and marshalls it into json and writes it the the response.
func JSONResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	
}