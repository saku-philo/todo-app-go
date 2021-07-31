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
	Todos     []Todo
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

func (u *User) CreateSession() (session Session, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()
	session = Session{}
	cmd1 := `INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at,
		updated_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `SELECT id, uuid, email, user_id, created_at, updated_at
		FROM sessions
		WHERE user_id =$1
		AND email =$2
		AND is_deleted = FALSE`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `SELECT id, uuid, email, user_id, created_at, updated_at
		FROM sessions
		WHERE uuid =$1
		AND is_deleted = FALSE`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
		&sess.UpdatedAt,
	)

	if err != nil {
		valid = false
		return
	}

	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	cmd := `UPDATE sessions
		SET is_deleted = TRUE, deleted_at =$1
		WHERE uuid =$2`
	_, err = Db.Exec(cmd, time.Now(), sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	// Connent to DB
	Db, err = connectDB()
	defer Db.Close()

	user = User{}
	cmd := `SELECT id, uuid, name, email, created_at, updated_at
		FROM users
		WHERE id =$1
		AND is_deleted = FALSE`

	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
