package util

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

const CODE_HASH = "E^dx&e$?]ot5#qMQq~YP6dG&W{5t"

// Secret envionment variable not found in code and not included in git repo
var SYS_HASH = os.Getenv("SYSTEM_HASH")
var PEPPER = SYS_HASH + CODE_HASH


func GenID(length int) string {
	randBytes := make([]byte, 32)
	io.ReadFull(rand.Reader, randBytes)
	// The output of this is actually like 42 chars long, so shrink it.
	return base64.URLEncoding.EncodeToString(randBytes)[:length]
}

func HashPass(password string) string {
	// Increse this to like 12 once we have better hardware
	b, _ := bcrypt.GenerateFromPassword([]byte(password+PEPPER), 8)
	return string(b)
}

// Checks if the hash is bcrypt for the password
func AuthPass(hash, password string) bool {
	resultErr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+PEPPER))
	return resultErr == nil
}
