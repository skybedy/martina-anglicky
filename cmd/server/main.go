package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/skybedy/martina-anglicky/internal/web"
)

func main() {
	homeTemplate, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	app := web.New(homeTemplate)

	log.Println("Spouštím server na http://127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", app.Routes()); err != nil {
		log.Fatal(err)
	}
}
