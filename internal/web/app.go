package web

import "html/template"

type App struct {
	homeTemplate    *template.Template
	martinaPortrait []byte
}

func New(homeTemplate *template.Template, martinaPortrait []byte) *App {
	return &App{
		homeTemplate:    homeTemplate,
		martinaPortrait: martinaPortrait,
	}
}
