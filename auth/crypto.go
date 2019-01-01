package auth

import (
	"os"
	"crypto/rand"
	"crypto/rsa"
)

var keyDir os.File
var oldKeyDir os.File
//GenRS4096KeyPair returns privatekey, publickey
func GenRS4096KeyPair() error{
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	err = key.Validate()
	if err != nil {
		return err
	}
	PrivateKey = key
	PublicKey = &key.PublicKey

	return nil
}
