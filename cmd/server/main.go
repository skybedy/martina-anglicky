package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/skybedy/martina-anglicky/internal/web"
)

const defaultAddress = "127.0.0.1:8080"

func main() {
	homeTemplate, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	martinaPortrait, err := os.ReadFile("prototype/martina-koncept.png")
	if err != nil {
		log.Fatal(err)
	}

	app := web.New(homeTemplate, martinaPortrait)

	address := serverAddress()
	log.Printf("Spouštím server na http://%s", address)
	if err := http.ListenAndServe(address, app.Routes()); err != nil {
		log.Fatal(err)
	}
}

func serverAddress() string {
	if address := os.Getenv("ADDR"); address != "" {
		return address
	}
	return defaultAddress
}
