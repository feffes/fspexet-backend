package auth

import (
	"os"
	"time"
	"io/ioutil"
	"github.com/dgrijalva/jwt-go"
	"log"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
)

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Bytes:   privDER,
	}
	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)
	return privatePEM
}
func encodePublicKeyToPEM(publicKey *rsa.PublicKey) []byte {
	pubDER := x509.MarshalPKCS1PublicKey(publicKey)
	pubBlock := pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: pubDER,
	}
	pubPEM := pem.EncodeToMemory(&pubBlock)
	return pubPEM

}

// LoadKeys loads the keys into variables
func loadKeys(private []byte, public []byte) {
	PrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(private)
	PublicKey, _ = jwt.ParseRSAPublicKeyFromPEM(public)
}

func writeKeys(private []byte, public []byte){
	ioutil.WriteFile(keyDir.Name()+"rsa.key", private, 0770)
	ioutil.WriteFile(keyDir.Name()+"rsa.pub", public, 0770)

}

func readKeys() ([]byte, []byte){
	priv, err := ioutil.ReadFile(keyDir.Name()+"rsa.key")
	pub, err := ioutil.ReadFile(keyDir.Name()+"rsa.pub")
	if err != nil {
		log.Println(err)
	}
	return priv, pub
}

// CopyCurrentKey stores a copy of the current rsa.pub and rsa.key to old/rsa.pub.Time.Now() and old/rsa.key.Time.Now() respectively  
func CopyCurrentKey(){
	now := time.Now()
	priv, pub := readKeys()
	ioutil.WriteFile(oldKeyDir.Name()+"rsa.key."+now.String(), priv, 0770)
	ioutil.WriteFile(oldKeyDir.Name()+"rsa.pub."+now.String(), pub, 0770)

}
// InitKeyDir sets the path to the RSA key directory
func InitKeyDir(path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err){
			os.MkdirAll(path, 0770) //0770 corresponds to rwx for user and group and no permission for others
		}
	}
	tmp, err := os.Open(path)
	if err != nil{
		log.Println(err)
	}
	defer tmp.Close()
	keyDir = *tmp
	if _, err := os.Stat(keyDir.Name()+"old/"); err != nil {
		if os.IsNotExist(err){
			os.MkdirAll(path+"old/", 0770) //0770 corresponds to rwx for user and group and no permission for others
		}
	}
	oldtmp, err := os.Open(path+"old/")
	if err != nil{
		log.Println(err)
	}
	oldKeyDir = *oldtmp
	log.Printf("key dir set to : %s and old keys are stored in: %s  \n", keyDir.Name(), oldKeyDir.Name())

}