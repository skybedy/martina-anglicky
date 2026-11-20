package web

import (
	"bytes"
	"log"
	"net/http"
)

type homePageData struct {
	Title       string
	Description string
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	data := homePageData{
		Title:       "Angličtina s Martinou",
		Description: "Individuální online lekce angličtiny s Martinou.",
	}

	var output bytes.Buffer
	if err := app.homeTemplate.Execute(&output, data); err != nil {
		log.Printf("vykreslení úvodní stránky: %v", err)
		http.Error(w, "Stránku se nepodařilo načíst.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := output.WriteTo(w); err != nil {
		log.Printf("odeslání odpovědi: %v", err)
	}
}
