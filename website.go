package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"unicode/utf8"

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
	//Decide()
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
	ActionInt, _ := strconv.Atoi(action)
	Decide(ActionInt)
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
	appendToJson(&toJson{0, "", 0})
}

const pass = "iktfag"

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
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
	res, err := db.Query("select * from shopping;")
	if err != nil {
		return
	}
	defer res.Close()
	var toJsonString string
	for res.Next() {
		err := res.Scan(&id, &item, &amount)
		if err != nil {
			fmt.Println(err)
		}
		items := &toJson{Id: id, Item: item, Amount: amount}
		file, _ := json.MarshalIndent(items, "", " ")
		toJsonString = toJsonString + string(file) + ","
	}

	toJsonString = trimLastChar(toJsonString)
	toJsonString = "[" + toJsonString + "]"
	_ = os.WriteFile("selectQuery.json", []byte(toJsonString), 0666)
}

type toJson struct {
	Id     int
	Item   string
	Amount int
}
