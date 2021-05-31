package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `INSERT INTO todos (
		content,
		user_id,
		created_at,
		updated_at)
		VALUES ($1, $2, $3, $4)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now(), time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
