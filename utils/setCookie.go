package utils

import "net/http"

func SetCookie(w http.ResponseWriter, name, value string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // production only
		SameSite: http.SameSiteStrictMode,
		MaxAge:   maxAge,
	})
}
