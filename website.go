package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/index", queryHandler)
}
