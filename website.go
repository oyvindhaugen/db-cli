package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"unicode/utf8"

	_ "github.com/lib/pq"
)

// This tells db.go to insert a new entry, giving it the Item and Amount
func insertRow(w http.ResponseWriter, r *http.Request) {
	var data insertedRow
	var resData insertedRowRes
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())

		resData.Item = data.Item
		resData.Result = "There was an error with json data"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resData)
		return
	}
	fmt.Println(data.Item, data.Amount)
	Decide(1, 0, data.Item, data.Amount)
	resData.Item = data.Item
	resData.Result = "Successfully inserted"
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resData)

	appendToJson()
}

// This tells db.go to update an entry at given ID, giving it the new Item and Amount
func updateRow(w http.ResponseWriter, r *http.Request) {
	var data updatedRow
	var resData updatedRowRes
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())

		resData.Id = data.Id
		resData.Result = "There was an error with json data"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resData)
		return
	}
	fmt.Println(data.Id, data.Item, data.Amount)
	Decide(3, data.Id, data.Item, data.Amount)

	resData.Id = data.Id
	resData.Result = "Successfully updated"
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resData)

	appendToJson()
}

// This tells db.go to delete an entry at given ID
//
//	func deleteRow(w http.ResponseWriter, r *http.Request) {
//		var data deletedRow
func deleteRow(w http.ResponseWriter, r *http.Request) {
	var data deletedRow
	var resData deletedRowRes
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())

		resData.Id = data.Id
		resData.Result = "There was an error with the json data"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resData)
		return
	}
	Decide(2, data.Id, "", 0)

	resData.Id = data.Id
	resData.Result = "Successfully deleted"
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resData)

	appendToJson()
}

// This handles all the websites, giving them functions
func handle() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/delete_row", deleteRow)
	http.HandleFunc("/insert_row", insertRow)
	http.HandleFunc("/update_row", updateRow)
	appendToJson()
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

// This trims the last character of a string
func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

// This selects everything from the database, then adds it into a JSON file for the frontend to use.
func appendToJson() {
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
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
	res, err := db.Query("SELECT * FROM shopping;")
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
	_ = os.WriteFile("./static/data.json", []byte(toJsonString), 0666)
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
type insertedRow struct {
	Item   string
	Amount int
}
type insertedRowRes struct {
	Result string
	Item   string
}
type updatedRow struct {
	Id     int
	Item   string
	Amount int
}
type updatedRowRes struct {
	Result string
	Id     int
}
