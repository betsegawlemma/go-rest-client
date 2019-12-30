package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/betsegawlemma/restclient/data"
)

var tmpl = template.Must(template.ParseGlob("ui/templates/*.html"))

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/user", singleUser)
	mux.HandleFunc("/users", allUsers)

	http.ListenAndServe(":8181", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}

func singleUser(w http.ResponseWriter, r *http.Request) {
	idraw := r.FormValue("id")
	id, err := strconv.Atoi(idraw)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	user, err := data.FetchUser(id)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		tmpl.ExecuteTemplate(w, "error.layout", nil)
	}
	tmpl.ExecuteTemplate(w, "user.layout", user)

}

func allUsers(w http.ResponseWriter, r *http.Request) {
	pageraw := r.FormValue("page")
	page, err := strconv.Atoi(pageraw)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	users, err := data.FetchUsers(page)

	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		tmpl.ExecuteTemplate(w, "error.layout", nil)
	}

	tmpl.ExecuteTemplate(w, "users.layout", users)
}
