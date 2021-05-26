package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
