package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func connectDB() (db *sql.DB, err error) {
	sqlDriver := fmt.Sprintf("%v", config.Config.SQLDriver)
	connInfo := fmt.Sprintf(
		"user=%v dbname=%v password=%v sslmode=disable",
		config.Config.DbUser,
		config.Config.DbName,
		config.Config.DbPassword,
	)

	Db, err = sql.Open(sqlDriver, connInfo)
	// Error handling
	if err != nil {
		log.Panicln(err)
		fmt.Println("error")
	}
	return Db, err
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
