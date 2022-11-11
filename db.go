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

// this does all the database stuff
func Decide(i int, id int, item string, amount int) {
	if i == 1 {
		insert(item, amount)
	} else if i == 2 {
		del(id)
	} else if i == 3 {
		updt(id, item, amount)
	} else {
		fmt.Println("Error")
	}

}

// This checks for errors
func CheckError(err error) {
	if err != nil {
		log.Fatal()
	}
}

// This inserts a new entry into the database
func insert(item string, amount int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	_, err = db.Exec("INSERT INTO shopping (item, amount) VALUES ($1, $2)", item, amount)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}

// This deletes any entry at given ID
func del(id int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Exec("DELETE FROM shopping WHERE id = $1;", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Success")
	appendToJson()
	return nil
}

// This updates any entry at given ID
func updt(id int, item string, amount int) error {

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Query("UPDATE shopping SET item = $2, amount = $3 WHERE id = $1", id, item, amount)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}
