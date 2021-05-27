package models

import (
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at,
		updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now(),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
