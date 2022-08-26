package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	//handle()
	appendToJson(&toJson{0, "", 0})
}
func appendToJson(j *toJson) {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = first_db sslmode=disable", pass2)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	var (
		id     int
		item   string
		amount int
	)
	res, errs := db.Query("select * from shopping;")
	if errs != nil {
		return
	}
	defer res.Close()
	f, err := os.OpenFile("./selectQuery.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for res.Next() {
		err := res.Scan(&id, &item, &amount)
		if err != nil {
			log.Fatal(err)
		}
		varsForJson := &toJson{id: id, item: item, amount: amount}
		byteArray, err := json.Marshal(varsForJson)
		if err != nil {
			fmt.Println(err)
		}
		n, err := f.Write(byteArray)
		if err != nil {
			fmt.Println(n, err)
		}
		if n, err = f.WriteString("\n"); err != nil {
			fmt.Println(n, err)
		}
	}
}

type toJson struct {
	id     int    `json: id`
	item   string `json: item`
	amount int    `json: amount`
}
