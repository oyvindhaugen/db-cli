package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// this is the handler for cli.html and it gives all the parameters needed for database queries
func cliHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	id := r.FormValue("id")
	Id, _ = strconv.Atoi(id)
	Item = r.FormValue("item")
	amount := r.FormValue("amount")
	Amount, _ = strconv.Atoi(amount)
	go decide()
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
