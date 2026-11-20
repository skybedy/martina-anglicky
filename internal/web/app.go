package web

import "html/template"

type App struct {
	homeTemplate *template.Template
}

func New(homeTemplate *template.Template) *App {
	return &App{homeTemplate: homeTemplate}
}
