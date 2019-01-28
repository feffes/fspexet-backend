package auth

import (
	"errors"
	"regexp"
	"os"
	"crypto/rand"
	"crypto/rsa"
)
const argonRegex = `^(?:\$([a-z0-9-]+))(?:\$v=([a-z0-9]+))?(?:\$((?:[a-z0-9-]+=[a-zA-Z0-9/+.-]+,?)+))?(?:\$([a-zA-Z0-9/+.-]+))?(?:\$([a-zA-Z0-9/+.-]+))?$`

var keyDir os.File
var oldKeyDir os.File

type argon struct{
	id 		string
	ver 	string
	conf	string
	salt 	string
	hash	string
}

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
// CompareArgon2id compares an input string to a given argon2id hash


func parseArgon(hash string) (argon, error){
	rex := regexp.MustCompile(argonRegex)
	out := rex.FindAllStringSubmatch(hash, -1)[0]
	if hash != out[0] || len(out) != 6{
		return argon{}, errors.New("invalid argon hash:" + hash)
	}
	return argon{
		id: out[1],
		ver: out[2],
		conf: out[3],
		salt: out[4],
		hash: out[5],
	}, nil
}