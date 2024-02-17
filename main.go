package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var database []int

func form(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/list.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	input := r.FormValue("input")

	angkaIn, err := strconv.Atoi(input)

	if err != nil {
		http.ServeFile(w, r, "views/warning.html")
		return
	}

	database = append(database, angkaIn)
	fmt.Println("GINIH", database)

	datas := map[string]interface{}{
		"data":  database,
		"input": input,
	}

	if err := tmpl.Execute(w, datas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", form)
	http.ListenAndServe(":8080", nil)
}
