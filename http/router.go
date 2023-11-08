package http

import (
	"github.com/kataras/muxie"
)

const (
	INDEX_ROUTE    = "/"
	LOGIN_ROUTE    = "/login"
	LOGIN_DO_ROUTE = "/login/do"
)

func Router() *muxie.Mux {
	router := muxie.NewMux()
	router.Use(IdMiddleware, LogMiddleware, AuthMiddleware, ErrorMiddleware)
	router.HandleFunc(INDEX_ROUTE, Index)
	router.HandleFunc(LOGIN_ROUTE, Login)
	router.HandleFunc(LOGIN_DO_ROUTE, LoginDo)
	return router
}
