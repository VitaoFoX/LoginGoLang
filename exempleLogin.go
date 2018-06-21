package main

import (
	"html/template"
	"log"
	"net/http"
)

//Login Struct
type Login struct {
	Name string
	Pass string
}

//Login Example - Hello Guys - By Vitor Brunoo

var tpl *template.Template

func login(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {

		login := req.FormValue("login")
		pass := req.FormValue("pass")

		if login == "admin" && pass == "admin" {
			var l Login

			l.Name = req.FormValue("login")
			l.Pass = req.FormValue("pass")

			err := tpl.ExecuteTemplate(w, "login.gohtml", l)
			if err != nil {
				http.Error(w, err.Error(), 500)
				log.Fatalln(err)
			}
			return
		}

	}
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

}

func init() {
	tpl = template.Must(template.ParseFiles("login.gohtml"))
}

func main() {
	http.HandleFunc("/", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
