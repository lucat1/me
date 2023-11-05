package http

import "net/http"

type LoginData struct {
	Message string
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := GetLogger(r)
	logger.Info("Received login request")

	if err := RenderPage[LoginData](w, r, "me.login", "Login", LoginData{
		Message: "test",
	}); err != nil {
		logger.With("err", err).Error("Could not render page")
	}
}
