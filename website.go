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

func insertHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	http.Redirect(w, r, "/", http.StatusFound)
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	item := r.FormValue("newItem")
	amount := r.FormValue("newAmount")
	amountInt, _ := strconv.Atoi(amount)
	fmt.Printf("id: %v\n item: %s\n amount: %v", idInt, item, amountInt)
}
func updatehandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	http.Redirect(w, r, "/", http.StatusFound)
	id := r.FormValue("id")
	idInt, _ := strconv.Atoi(id)
	item := r.FormValue("newItem")
	amount := r.FormValue("newAmount")
	amountInt, _ := strconv.Atoi(amount)
	fmt.Printf("id: %v\n item: %s\n amount: %v", idInt, item, amountInt)
}
func deleteRow(w http.ResponseWriter, r *http.Request) {
	var data deletedRow

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		var resData deletedRowRes
		resData.Id = 0
		resData.Result = "There was an error with json data"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resData)
		return
	}
	Decide(2, data.Id, "", 0)
	var resData deletedRowRes
	resData.Id = data.Id
	resData.Result = "Successfully deleted"
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resData)

	appendToJson()
}
func handle() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/delete_row", deleteRow)
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/update", updatehandler)
	appendToJson()
	fmt.Printf("Starting server at port 127.0.0.1:5500\n")
	if err := http.ListenAndServe("127.0.0.1:5500", nil); err != nil {
		log.Fatal(err)
	}
}

const pass = "iktfag"

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}
func appendToJson() {
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
	_ = os.WriteFile("./static/selectQuery.json", []byte(toJsonString), 0666)
}

type toJson struct {
	Id     int
	Item   string
	Amount int
}
type deletedRow struct {
	Id int
}
type deletedRowRes struct {
	Result string
	Id     int
}
