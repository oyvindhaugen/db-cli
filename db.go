package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var ActionInt int
var Id int
var Item string
var Amount int

type columns struct {
	id     int
	item   string
	amount int
}

// this does all the database stuff
func decide() {
	fmt.Println("here")
	fmt.Println(ActionInt)
	if ActionInt == 1 {
		insert(&columns{Id, Item, Amount})
	} else if ActionInt == 2 {
		del(&columns{Id, Item, Amount})
	} else if ActionInt == 3 {
		updt(&columns{Id, Item, Amount})
	} else if ActionInt == 4 {
		slct(&columns{Id, Item, Amount})
	} else {
		fmt.Println("Error")
	}

}
func CheckError(err error) {
	if err != nil {
		log.Fatal()
	}
}

const pass3 = "iktfag"

func insert(c *columns) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass3)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Exec("insert into shopping (Item, Amount) values ($1, $2)", c.item, c.amount)
	if err != nil {
		fmt.Println("mo")
		return err
	}
	fmt.Println("Success")
	return nil
}
func del(c *columns) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Exec("delete from shopping where Id = $1;", c.id)
	if err != nil {
		fmt.Println("mo")
		return err
	}
	fmt.Println("Success")
	return nil
}
func updt(c *columns) error {

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Query("update shopping set Item = $2, Amount = $3 where Id = $1", c.id, c.item, c.amount)
	if err != nil {
		fmt.Println("mo")
		return err
	}
	fmt.Println("Success")
	return nil
}
func slct(c *columns) error {

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = first_db sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	var (
		id     int
		item   string
		amount int
	)
	res, err := db.Query("select * from shopping where Id = $1;", c.id)
	if err != nil {
		fmt.Println("mo")
		return err
	}
	defer res.Close()
	for res.Next() {
		err := res.Scan(&id, &item, &amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %d\n Item: %s\n Amount: %d\n", id, item, amount)
	}

	fmt.Println("Success")
	return nil
}
