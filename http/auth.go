package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucat1/me/config"
)

const (
	AUTH_CONTEXT_KEY = "auth"
	AUTH_COOKIE_NAME = "meAuthToken"
)

type User struct {
	DN string `json:"string"`
}

type Claims struct {
	User
	jwt.RegisteredClaims
}

func Authenticate(w http.ResponseWriter, user User) (err error) {
	authConfig := config.Get().Auth
	claims := Claims{User: user}
	expiry := time.Now().Add(time.Duration(authConfig.Duration) * time.Minute)
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiry)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(authConfig.Secret))
	if err != nil {
		err = fmt.Errorf("Could not sign JWT: %v", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    AUTH_COOKIE_NAME,
		Value:   tokenString,
		Expires: expiry,
		Path:    "/",
	})
	return
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := GetLogger(r)
		cookie, err := r.Cookie(AUTH_COOKIE_NAME)
		if err != nil {
			logger.With("err", err, "missing", err == http.ErrNoCookie).Debug("Could not get auth cookie")
			// TODO: We may consider returning a 400 for requests which give an error different
			// from ErrNoCookie, as done here: https://www.sohamkamani.com/golang/jwt-authentication/
			h.ServeHTTP(w, r)
			return
		}

		claims := Claims{}
		tkn, err := jwt.ParseWithClaims(cookie.Value, &claims, func(token *jwt.Token) (any, error) {
			return []byte(config.Get().Auth.Secret), nil
		})
		if err != nil || !tkn.Valid {
			logger.With("err", err, "valid", tkn.Valid).Debug("Invalid token")
		} else {
			ctx := context.WithValue(r.Context(), AUTH_CONTEXT_KEY, claims.User)
			r = r.WithContext(ctx)
		}
		h.ServeHTTP(w, r)
	})
}

func AuthMust(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := GetLogger(r)
		user, ok := r.Context().Value(AUTH_CONTEXT_KEY).(*User)

		if !ok {
			http.Redirect(w, r, LOGIN_ROUTE, http.StatusTemporaryRedirect)
			return
		}

		logger.With("user", user).Debug("Granted authenticated request")
		h.ServeHTTP(w, r)
	})
}

// if must is set, GetUser can only be called in a http.Handler wrapped in a
// AuthMust middleware, but the nil checking on the returned variable can be avoided.
// Otherwise, the user should be always checked for existance (nil if unlogged)
func GetUser(r *http.Request, must bool) (user *User) {
	var ok bool
	user, ok = r.Context().Value(AUTH_CONTEXT_KEY).(*User)
	if !ok {
		user = nil
		if must {
			panic("Could not get request user")
		}
		return
	}

	return
}
