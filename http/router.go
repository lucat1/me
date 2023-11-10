package http

import (
	"net/http"

	"github.com/kataras/muxie"
)

const (
	INDEX_ROUTE    = "/"
	LOGIN_ROUTE    = "/login"
	LOGIN_DO_ROUTE = "/login/do"
	LOGOUT_ROUTE   = "/logout"
)

func Router() *muxie.Mux {
	router := muxie.NewMux()
	router.Use(IdMiddleware, LogMiddleware, AuthMiddleware, ErrorMiddleware)
	router.Handle(INDEX_ROUTE, AuthMust(http.HandlerFunc(Index)))
	router.HandleFunc(LOGIN_ROUTE, Login)
	router.HandleFunc(LOGIN_DO_ROUTE, LoginDo)
	router.HandleFunc(LOGOUT_ROUTE, Logout)
	return router
}
