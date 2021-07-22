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

func GetTodo(id int) (todo Todo, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `SELECT id, content, user_id, created_at, updated_at
		FROM todos
		WHERE id =$1`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `SELECT id, content, user_id, created_at, updated_at
		FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	// Fetch todos data specify user
	cmd := `SELECT id, content, user_id, created_at, updated_at
		FROM todos
		WHERE user_id =$1`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodo() error {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `UPDATE todos set content =$1, user_id =$2
		WHERE id =$3`

	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `UPDATE todos
		SET is_deleted = TRUE, deleted_at =$1
		WHERE id =$2`

	_, err = Db.Exec(cmd, time.Now(), t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
