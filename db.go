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

type columns struct {
	id     int
	item   string
	amount int
}

// this does all the database stuff
func Decide(i int, id int, item string, amount int) {
	if i == 1 {
		insert(&columns{Id, Item, Amount}, item, amount)
	} else if i == 2 {
		del(&columns{Id, Item, Amount}, id)
	} else if i == 3 {
		updt(&columns{Id, Item, Amount}, id, item, amount)
	} else if i == 4 {
		slct(&columns{Id, Item, Amount}, id)
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

func insert(c *columns, item string, amount int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass3)
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
func del(c *columns, id int) error {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass3)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Exec("delete from shopping where id = $1;", id)
	if err != nil {
		fmt.Println("mo")
		fmt.Println(err)
		return err
	}
	fmt.Println("Success")
	appendToJson()
	return nil
}
func updt(c *columns, id int, item string, amount int) error {

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass3)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	_, err = db.Query("update shopping set item = $2, amount = $3 where id = $1", c.id, c.item, c.amount)
	if err != nil {
		fmt.Println("mo")
		return err
	}
	fmt.Println("Success")
	return nil
}
func slct(c *columns, ids int) error {

	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass3)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	var (
		id     int
		item   string
		amount int
	)
	res, err := db.Query("select * from shopping where id = $1;", c.id)
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
