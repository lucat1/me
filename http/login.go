package http

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/lucat1/me/client"
	"github.com/lucat1/me/config"
)

const (
	HTMX_REDIRECT = "HX-Redirect"
)

type LoginFormData struct {
	Username      string
	UsernameError string
	Password      string
	PasswordError string
	Error         string
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r)
	user := GetUser(r, false)
	logger.Info("Rendering login index")
	if user != nil {
		http.Redirect(w, r, INDEX_ROUTE, http.StatusTemporaryRedirect)
		return
	}

	if err := RenderPage[LoginFormData](w, r, "login", "Login", LoginFormData{}); err != nil {
		logger.With("err", err).Error("Could not render page")
	}
}

type LoginDoData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginDo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	logger := GetLogger(r)
	var loginData LoginDoData

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		Error(w, r, http.StatusBadRequest, true, nil, "Invalid login data")
		return
	}

	logger = logger.With("username", loginData.Username)
	logger.Info("Login attempt")

	// Basic returnData to persist form values
	returnData := LoginFormData{
		Username: loginData.Username,
		Password: loginData.Password,
	}

	if loginData.Username == "" {
		returnData.UsernameError = "The username cannot be empty"
	} else if loginData.Password == "" {
		returnData.PasswordError = "The password cannot be empty"
	} else {
		authCfg := config.Get().Auth
		filter := strings.ReplaceAll(authCfg.Filter, "{username}", loginData.Username)
		c, err := client.StartRoot(logger)
		if err != nil {
			Error(w, r, http.StatusInternalServerError, true, err, "Could not start the LDAP client")
			return
		}
		logger.With("base", authCfg.BaseDN, "filter", filter).Debug("Searching for user")
		results, err := c.Search(ldap.NewSearchRequest(
			authCfg.BaseDN,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			filter,
			[]string{"dn"},
			nil,
		))
		if err != nil {
			Error(w, r, http.StatusInternalServerError, true, err, "LDAP Search Error")
			return
		}
		if err := c.Close(); err != nil {
			Error(w, r, http.StatusInternalServerError, true, err, "LDAP Close Error")
			return
		}

		for _, entry := range results.Entries {
			logger.With("entry", entry.DN).Info("Found user")
		}
		if len(results.Entries) < 1 || len(results.Entries) > 1 {
			returnData.Error = "Invalid credentials"
			return
		} else {
			userDN := results.Entries[0].DN

			c, err := client.Start(logger)
			if err != nil {
				Error(w, r, http.StatusInternalServerError, true, err, "Could not start the LDAP client")
				return
			}
			if err := c.Bind(userDN, loginData.Password); err != nil {
				returnData.Error = "Invalid credentials"
			} else {
				logger.With("username", loginData.Username, "dn", userDN).Info("Logged in")
				if err := Authenticate(w, User{DN: userDN}); err != nil {
					Error(w, r, http.StatusInternalServerError, true, err, "Could not store generate the authentication cookie")
					return
				}
				w.Header().Add(HTMX_REDIRECT, "/")
			}
			if err := c.Close(); err != nil {
				Error(w, r, http.StatusInternalServerError, true, err, "LDAP Close Error")
				return
			}
		}
	}

	if err := RenderBlockPage[LoginFormData](w, r, "login.form", returnData); err != nil {
		logger.With("err", err).Error("Could not render page")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     AUTH_COOKIE_NAME,
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
	http.Redirect(w, r, INDEX_ROUTE, http.StatusTemporaryRedirect)
	return
}
