package web

import "net/http"

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /assets/martina-koncept.png", app.martinaPortraitImage)
	return mux
}
