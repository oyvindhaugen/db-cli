package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// this is the handler for cli.html and it gives all the parameters needed for database queries
func cliHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	id := r.FormValue("Id")
	Id, _ = strconv.Atoi(id)
	Item = r.FormValue("Item")
	amount := r.FormValue("Amount")
	Amount, _ = strconv.Atoi(amount)
	Decide()
}

// this tells the db.go what database action is gonna be performed
func formTestHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	http.Redirect(w, r, "/cli.html", http.StatusFound)
	fmt.Fprintf(w, "POST request successful\n")
	action := r.FormValue("action")
	ActionInt, _ = strconv.Atoi(action)
}
func handle() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/index", formTestHandler)
	http.HandleFunc("/cli", cliHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

const pass = "iktfag"

func appendToJson(j *toJson) {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = postgres password = %s  dbname = postgres sslmode=disable", pass)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal()
	}
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
			fmt.Println(err)
		}
		items := &toJson{Id: id, Item: item, Amount: amount}
		file, _ := json.MarshalIndent(items, "", " ")

		_ = os.WriteFile("selectQuery.json", file, 0666)
	}
}

type toJson struct {
	Id     int
	Item   string
	Amount int
}
