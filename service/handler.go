package server

import (
	"html/template"
	"log"
	"net/http"
)

var Tmpl *template.Template
var err error

func init() {
	Tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	page := Ascii{}
	err = Tmpl.ExecuteTemplate(w, "index.html", page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	page := Ascii{
		Input: r.FormValue("input"),
		Fs:    r.FormValue("fs"),
	}
	page.Output, err = GetArt(page.Input, page.Fs)
	if err != nil {
		if err.Error() == "non valid symbols" {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = Tmpl.ExecuteTemplate(w, "index.html", page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
