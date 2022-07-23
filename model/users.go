package model

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuelsih/gosckt/config"
	"golang.org/x/crypto/bcrypt"
)

// User is the model for users table
type User struct {
	db       *pgxpool.Pool `json:"-"`
	ID       uuid.UUID     `json:"id"`
	Email    string        `json:"email"`
	Name     string        `json:"name"`
	Password []byte        `json:"-"`
}

// NewUser return struct User with db given
func NewUser(db *pgxpool.Pool) *User {
	return &User{db: db}
}

// Insert will insert users' email, name, and hashed password
func (u *User) Insert() bool {
	ctx, cancel := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancel()

	var id uuid.UUID

	query := `INSERT INTO users(email, name, password) values($1, $2, $3) returning id;`

	err := u.db.QueryRow(ctx, query, u.Email, u.Name, u.Password).Scan(&id)

	if err != nil {
		return false
	}

	if id != uuid.Nil {
		u.ID = id
		return true
	}

	return false
}

// FindByEmail will get user by its email
// Return error if user not found or internal error
func (u *User) FindByEmail(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancel()

	query := `SELECT id, email, name, password FROM USERS WHERE email = $1`

	var (
		userUUID uuid.UUID
		userEmail, userName, userPassword []byte
	)

	err := u.db.QueryRow(ctx, query, email).Scan(&userUUID, &userEmail, &userName, &userPassword)

	if err != nil {
		println("error slur:", err.Error())
		return false
	}

	u.ID = userUUID
	u.Email = string(userEmail)
	u.Name = string(userName)
	u.Password = userPassword

	return u.ID != uuid.Nil
}

// HasEmail will find user by its email
// This method is different with FindByEmail, this will return true if user is found. Otherwise false
func (u *User) HasEmail(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancel()

	var userEmail []byte

	query := `SELECT email FROM USERS WHERE email = $1`

	err := u.db.QueryRow(ctx, query, email).Scan(&userEmail)

	if err != nil {
		return false
	}

	return userEmail != nil
}

// SetPassword will encrypted the password with salt
func (u *User) SetPassword(password string) {
	saltedBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
		return
	}

	u.Password = hashedBytes
}

// PasswordMatch will check if incoming password is equal to users' password
func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		println(`User.PasswordMatch:`, err.Error())
		return false
	}

	return true
}

// Clean will set the db connection to nil
func (u *User) Clean() *User {
	u.db = nil
	return u
}
