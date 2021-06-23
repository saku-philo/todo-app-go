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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

func GetUser(id int) (user User, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	user = User{}
	cmd :=
		`SELECT id, uuid, name, email, password, created_at, updated_at
	FROM users
	WHERE id =$1`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Fatalln(err)
	}
	return user, err
}

func (u *User) UpdateUser() (err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `UPDATE users
		SET name =$1, email =$2, updated_at =$3
		WHERE id =$4`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.UpdatedAt, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `UPDATE users
		SET is_deleted = TRUE, deleted_at =$1
		WHERE id =$2`

	_, err = Db.Exec(cmd, time.Now(), u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at, updated_at
		FROM users
		WHERE email =$1`

	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}
