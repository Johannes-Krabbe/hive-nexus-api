package auth

import (
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/scrypt"
)


func hash(p string, s string) (hp string, err error) {
	dk, err := scrypt.Key([]byte(p), []byte(s), 32768, 8, 1, 32)
	if (err != nil){
		panic(err)
	}
	hp = hex.EncodeToString(dk)
	return
}



func getSalt() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
    rand.Seed(time.Now().UnixNano())

    b := make([]byte, 32)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }

    return string(b)
}