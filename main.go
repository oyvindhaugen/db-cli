package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type columns struct {
	id         int
	age        int
	first_name string
	last_name  string
	email      string
}

var pass string = "iktfag"

func insert(c *columns) error {
	psqlConnect := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlConnect)
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	_, err = db.Exec("delete from users where id = $1;", c.id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func del(c *columns) error {
	psqlConnect := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlConnect)
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	_, err = db.Exec("delete from users where id = $1;", c.id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func updt(c *columns) error {
	psqlConnect := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlConnect)
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	_, err = db.Query("update users set age = $2, first_name = $3, last_name = $4, email = $5 where id = $1", c.id, c.age, c.first_name, c.last_name, c.email)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func slct(c *columns) error {
	psqlConnect := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlConnect)
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	var (
		id        int
		age       int
		firstname string
		lastname  string
		email     string
	)

	res, err := db.Query("select * from users where id = $1;", c.id)
	if err != nil {
		return err
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(&id, &age, &firstname, &lastname, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d\n age: %d\n first_name: %s\n last_name: %s\n email: %s\n", id, age, firstname, lastname, email)
	}
	return nil
}
