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
		Title:       "Martina — angličtina po svém | vizuální návrh",
		Description: "Individuální online lekce angličtiny pro dospělé.",
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

func (app *App) martinaPortraitImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", http.DetectContentType(app.martinaPortrait))
	w.Header().Set("Cache-Control", "public, max-age=3600")
	if _, err := w.Write(app.martinaPortrait); err != nil {
		log.Printf("odeslání portrétu: %v", err)
	}
}
