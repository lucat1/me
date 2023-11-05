package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/lucat1/me/config"
	router "github.com/lucat1/me/http"
	"github.com/lucat1/me/templates"
)

func main() {
	config.ParseConfig("./config.yaml")
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

	if err := templates.Load(os.DirFS(conf.Builtin.TemplatesDir)); err != nil {
		slog.With("err", err).Error("Could not load templates")
	}

	addr := fmt.Sprintf("%s:%d", conf.WebConfig.Ip, conf.WebConfig.Port)
	slog.With("addr", addr).Info("Listening")
	http.ListenAndServe(addr, router.Router())
}
