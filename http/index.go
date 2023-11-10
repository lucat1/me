package http

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r)
	user := GetUser(r, false)
	logger.With("user", user).Info("Recieved index request")
	if user == nil {
		http.Redirect(w, r, LOGIN_ROUTE, http.StatusTemporaryRedirect)
		return
	}

	if err := RenderPage[interface{}](w, r, "index", "Index", nil); err != nil {
		logger.With("err", err).Error("Could not render page")
	}
}
