package main

import (
	"bytes"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/kataras/muxie"
	"github.com/lucat1/me/config"
)

var (
	templates *template.Template
	counter   = Counter{0}
)

func main() {
	config.ParseConfig("./config/config.yaml")
	conf := config.Get()

	//Set log level
	var logLevel slog.Level

	switch conf.LogLevel {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelWarn
	}
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(handler))

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
