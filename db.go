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
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	_, err = db.Exec("insert into shopping (item, amount) values ($1, $2)", item, amount)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}

// This deletes any entry at given ID
func del(id int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Exec("delete from shopping where id = $1;", id)
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

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Query("update shopping set item = $2, amount = $3 where id = $1", id, item, amount)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}
