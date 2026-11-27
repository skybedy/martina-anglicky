package web_test

import (
	"bytes"
	"html"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
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

func loadMartinaPortrait(t *testing.T) []byte {
	t.Helper()

	portrait, err := os.ReadFile(filepath.Join("..", "..", "prototype", "martina-koncept.png"))
	if err != nil {
		t.Fatalf("load Martina portrait: %v", err)
	}
	return portrait
}

func TestRoutes(t *testing.T) {
	handler := web.New(loadHomeTemplate(t), loadMartinaPortrait(t)).Routes()

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
					"<title>Martina — angličtina po svém | vizuální návrh</title>",
					"<h1>Angličtina, kterou si",
					`src="/assets/martina-koncept.png"`,
					"Pojďme zjistit, jestli si sedneme.",
				} {
					if !strings.Contains(recorder.Body.String(), fragment) {
						t.Errorf("response is missing %q", fragment)
					}
				}
			}
		})
	}
}

func TestMartinaPortraitImage(t *testing.T) {
	portrait := loadMartinaPortrait(t)
	handler := web.New(loadHomeTemplate(t), portrait).Routes()
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(
		recorder,
		httptest.NewRequest(http.MethodGet, "/assets/martina-koncept.png", nil),
	)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body: %s", recorder.Code, recorder.Body.String())
	}
	if got := recorder.Header().Get("Content-Type"); got != "image/png" {
		t.Errorf("Content-Type = %q, want image/png", got)
	}
	if got := recorder.Header().Get("Cache-Control"); got != "public, max-age=3600" {
		t.Errorf("Cache-Control = %q, want public cache header", got)
	}
	if !bytes.Equal(recorder.Body.Bytes(), portrait) {
		t.Error("response body does not match Martina portrait")
	}
}

func TestHomeTemplateError(t *testing.T) {
	// Execution writes this prefix before the missing field causes an error.
	// The handler must discard the partial page and return only a public error.
	tmpl, err := template.New("home").Parse("private partial page{{.MissingField}}")
	if err != nil {
		t.Fatalf("parse error fixture: %v", err)
	}
	handler := web.New(tmpl, nil).Routes()
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
				`<meta name="description" content="` + html.EscapeString(tt.description) + `">`,
			} {
				if !strings.Contains(output.String(), fragment) {
					t.Errorf("rendered HTML is missing %q", fragment)
				}
			}
		})
	}
}
