package http

import (
	"github.com/kataras/muxie"
)

func Router() *muxie.Mux {
	router := muxie.NewMux()
	router.Use(LogMiddleware, AuthMiddleware)
	return router
}
