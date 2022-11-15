package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(username, password string) (bool, int) {
	hash, err := HashPassword(password)
	fmt.Println(password)
	var dbPassword string
	var id int
	if err != nil {
		log.Fatal("Could not hash password: ", err)
	}

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable") //implement .env
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Could not secure connection to database: ", err)
	}
	defer db.Close()
	fmt.Printf("balls %s\n", username)
	row := db.QueryRow("SELECT password, id FROM users WHERE username = $1", username)
	switch err := row.Scan(&dbPassword, &id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
	case nil:
		if !CheckPasswordHash(password, hash) {
			fmt.Println("Not matching")
			return false, 0
		}
		fmt.Println("Matching")
		return true, id
	default:
		log.Fatal("Error")
	}
	return false, 0
}

func NewUser(username, password string) bool {
	password, err := HashPassword(password)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable") //implement .env
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}
