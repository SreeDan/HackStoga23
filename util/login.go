package util

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/HackStoga-2023-server/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
)

func Authenticate(ctx context.Context, userId int64, username string, password string, db *sql.DB) error {
	bytePass := []byte(password)

	passwdDb, err := getHashedPasswordFromDatabase(ctx, userId, username, db)
	if err != nil {
		return err
	}

	if comparePasswords(passwdDb, bytePass) {
		return nil
	} else {
		return fmt.Errorf("usernmae/password not found")
	}
}

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) error { //  Don't have time for JWT token authentication
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	//  Puts the body into a struct
	var loginStruct models.Login
	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	if err != nil {
		panic(err)
	}

	bytePass := []byte(loginStruct.Password)

	passwdDb, err := getHashedPasswordFromDatabase(ctx, loginStruct.UserId, loginStruct.Username, db)
	if err != nil {
		return err
	}

	if comparePasswords(passwdDb, bytePass) {
		var response = models.TodoItemJsonResponse{Type: "success"}
		json.NewEncoder(w).Encode(response)
		return nil
	} else {
		return fmt.Errorf("usernmae/password not found")
	}
}

func hashAndSalt(pwd []byte) string {
	// adds hashed and salts for security
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// verifies the passowrd
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func getHashedPasswordFromDatabase(ctx context.Context, userId int64, username string, db *sql.DB) (string, error) {
	var pwd string
	sqlStatement := `select password from credentials where user_id=$1 and username = $2`
	row := db.QueryRowContext(ctx, sqlStatement, userId, username)

	switch err := row.Scan(&pwd); err {
	case sql.ErrNoRows:
		return "", fmt.Errorf("username/password not found")
	case nil:
		return pwd, nil
	default:
		return "", err
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	//  Puts the body into a struct
	var loginStruct models.Login
	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	if err != nil {
		panic(err)
	}

	userId := rand.Intn(99999999999)

	bytePass := []byte(loginStruct.Password)
	hash := hashAndSalt(bytePass)

	insertStmt := `insert into credentials(user_id, username, password) VALUES($1, $2, $3)`
	_, err = db.ExecContext(ctx, insertStmt, userId, loginStruct.Username, hash)
	if err != nil {
		return err
	}

	var response = models.TodoItemJsonResponse{Type: "success"}
	json.NewEncoder(w).Encode(response)

	return nil
}
