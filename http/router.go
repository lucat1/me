package http

import (
	"github.com/kataras/muxie"
)

const (
	INDEX_ROUTE = "/"
	LOGIN_ROUTE = "/login"
)

func Router() *muxie.Mux {
	router := muxie.NewMux()
	router.Use(IdMiddleware, LogMiddleware, AuthMiddleware)
	router.HandleFunc(INDEX_ROUTE, Index)
	router.HandleFunc(LOGIN_ROUTE, Login)
	return router
}
