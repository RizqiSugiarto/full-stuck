package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type UserInput struct {
	Header string
	Body   string
}

var UserData []UserInput

func roots(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/notes.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(len(UserData), "GINIj")

	if err := tmpl.Execute(w, UserData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("views", "formAdd.html")
	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "POST" {
		http.Redirect(w, r, "http://localhost:8080/notes", http.StatusSeeOther)
	}

	input := UserInput{
		Header: r.FormValue("judul"),
		Body:   r.FormValue("isi"),
	}

	if input.Header != "" {
		UserData = append(UserData, input)
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("header")

	http.Redirect(w, r, "http://localhost:8080/notes", http.StatusSeeOther)

	for i := range UserData {
		if i == 0 || i%2 == 1 && UserData[i].Header == param {
			UserData = append(UserData[:i], UserData[i+1:]...)
			return
		}
	}

}

func main() {
	http.HandleFunc("/notes", roots)
	http.HandleFunc("/notes/add", Add)
	http.HandleFunc("/delete", Delete)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
