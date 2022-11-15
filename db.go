package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Id     int
	Item   string
	Amount int
)

// This checks for errors
func CheckError(err error) {
	if err != nil {
		log.Fatal()
	}
}

// This inserts a new entry into the database
func Insert(item string, amount int, userId int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	fmt.Println(userId)
	defer db.Close()
	_, err = db.Exec("INSERT INTO shopping (item, amount, owner) VALUES ($1, $2, $3)", item, amount, userId)
	if err != nil {
		return err
	}
	return nil
}

// This deletes any entry at given ID
func Del(id int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Exec("DELETE FROM shopping WHERE id = $1;", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// This updates any entry at given ID
func Updt(id int, item string, amount int) error {

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Query("UPDATE shopping SET item = $2, amount = $3 WHERE id = $1", id, item, amount)
	if err != nil {
		return err
	}
	return nil
}
