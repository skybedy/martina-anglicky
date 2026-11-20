package web_test

import (
	"bytes"
	"html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/skybedy/martina-anglicky/internal/web"
)

func loadHomeTemplate(t *testing.T) *template.Template {
	t.Helper()

	tmpl, err := template.ParseFiles(filepath.Join("..", "..", "templates", "home.html"))
	if err != nil {
		t.Fatalf("load home template: %v", err)
	}
	return tmpl
}

func TestRoutes(t *testing.T) {
	handler := web.New(loadHomeTemplate(t)).Routes()

	tests := []struct {
		name        string
		method      string
		target      string
		wantStatus  int
		wantHTML    bool
		wantContent bool
		wantAllow   string
	}{
		{
			name: "home", method: http.MethodGet, target: "/",
			wantStatus: http.StatusOK, wantHTML: true, wantContent: true,
		},
		{
			name: "home with query", method: http.MethodGet, target: "/?source=test",
			wantStatus: http.StatusOK, wantHTML: true, wantContent: true,
		},
		{
			name: "head", method: http.MethodHead, target: "/",
			wantStatus: http.StatusOK, wantHTML: true,
		},
		{
			name: "unknown path", method: http.MethodGet, target: "/rezervace",
			wantStatus: http.StatusNotFound,
		},
		{
			name: "nested path", method: http.MethodGet, target: "/students/1",
			wantStatus: http.StatusNotFound,
		},
		{
			name: "post", method: http.MethodPost, target: "/",
			wantStatus: http.StatusMethodNotAllowed, wantAllow: "GET, HEAD",
		},
		{
			name: "put", method: http.MethodPut, target: "/",
			wantStatus: http.StatusMethodNotAllowed, wantAllow: "GET, HEAD",
		},
		{
			name: "delete", method: http.MethodDelete, target: "/",
			wantStatus: http.StatusMethodNotAllowed, wantAllow: "GET, HEAD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(tt.method, tt.target, nil)
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d; body: %s", recorder.Code, tt.wantStatus, recorder.Body.String())
			}
			if tt.wantHTML {
				if got := recorder.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
					t.Errorf("Content-Type = %q, want HTML with UTF-8", got)
				}
			}
			if got := recorder.Header().Get("Allow"); got != tt.wantAllow {
				t.Errorf("Allow = %q, want %q", got, tt.wantAllow)
			}
			if tt.wantContent {
				for _, fragment := range []string{
					"<!doctype html>",
					`<html lang="cs">`,
					"<title>Angličtina s Martinou</title>",
					"<h1>Angličtina s Martinou</h1>",
					"<p>Individuální online lekce angličtiny s Martinou.</p>",
				} {
					if !strings.Contains(recorder.Body.String(), fragment) {
						t.Errorf("response is missing %q", fragment)
					}
				}
			}
		})
	}
}

func TestHomeTemplateError(t *testing.T) {
	// Execution writes this prefix before the missing field causes an error.
	// The handler must discard the partial page and return only a public error.
	tmpl, err := template.New("home").Parse("private partial page{{.MissingField}}")
	if err != nil {
		t.Fatalf("parse error fixture: %v", err)
	}
	handler := web.New(tmpl).Routes()
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", nil))

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", recorder.Code)
	}
	if got := recorder.Header().Get("Content-Type"); got != "text/plain; charset=utf-8" {
		t.Errorf("Content-Type = %q, want plain text with UTF-8", got)
	}
	if got, want := recorder.Body.String(), "Stránku se nepodařilo načíst.\n"; got != want {
		t.Errorf("body = %q, want %q without partial HTML or internal error details", got, want)
	}
}

func TestHomeTemplateData(t *testing.T) {
	tmpl := loadHomeTemplate(t)
	tests := []struct {
		name        string
		title       string
		description string
	}{
		{
			name:        "HTML is escaped",
			title:       `<script>alert("x")</script>`,
			description: `<img src=x onerror="alert(1)"> & další lekce`,
		},
		{name: "empty fields"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data := struct {
				Title       string
				Description string
			}{Title: tt.title, Description: tt.description}
			var output bytes.Buffer
			if err := tmpl.Execute(&output, data); err != nil {
				t.Fatalf("render template: %v", err)
			}

			for _, fragment := range []string{
				"<title>" + html.EscapeString(tt.title) + "</title>",
				"<h1>" + html.EscapeString(tt.title) + "</h1>",
				"<p>" + html.EscapeString(tt.description) + "</p>",
			} {
				if !strings.Contains(output.String(), fragment) {
					t.Errorf("rendered HTML is missing %q", fragment)
				}
			}
		})
	}
}
