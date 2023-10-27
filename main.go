package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/kataras/muxie"
)

var (
	templates *template.Template
	counter   = Counter{0}
)

func main() {
	fs := os.DirFS(".")
	var err error
	templates, err = template.ParseFS(fs, "templates/*.html")
	if err != nil {
		panic(err)
	}

	router := muxie.NewMux()
	router.HandleFunc("/", index)
	router.HandleFunc("/inc", mod(1))
	router.HandleFunc("/dec", mod(-1))
	http.ListenAndServe(":3001", router)
}

type Render struct {
	Page string
	Data template.HTML
}

type Counter struct {
	Count int
}

func renderCounter() string {
	var bytes bytes.Buffer
	err := templates.ExecuteTemplate(&bytes, "counter", counter)
	if err != nil {
		panic(err)
	}
	return bytes.String()
}

func index(w http.ResponseWriter, r *http.Request) {
	render := Render{Page: "counter", Data: template.HTML(renderCounter())}
	err := templates.ExecuteTemplate(w, "root", render)
	if err != nil {
		panic(err)
	}
}

func mod(diff int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counter.Count += diff
		w.Write([]byte(renderCounter()))
	}
}
