package controller

import (
	"list/model"
	"list/service"
	"net/http"
	"text/template"
)

var DataUser string

type NoteController struct {
	service service.CommonNoteService
}

func (u *NoteController) roots(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/notes.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := u.service.GetAllData()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (u *NoteController) Add(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/formAdd.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "POST" {
		http.Redirect(w, r, "http://localhost:8080/notes", http.StatusSeeOther)
	}

	input := model.UserInput{
		Header: r.FormValue("judul"),
		Body:   r.FormValue("isi"),
	}

	DataUser = input.Header

	if input.Header != "" {

		if err := u.service.Insert(input.Header, input.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (u *NoteController) Delete(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("header")

	http.Redirect(w, r, "http://localhost:8080/notes", http.StatusSeeOther)

	if err := u.service.Delete(param); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func NewNoteController(service service.CommonNoteService) *NoteController {
	return &NoteController{
		service: service,
	}
}

func (u *NoteController) RouteCommonNote() {
	http.HandleFunc("/notes", u.roots)
	http.HandleFunc("/notes/add", u.Add)
	http.HandleFunc("/delete", u.Delete)
}
